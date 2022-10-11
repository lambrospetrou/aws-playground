package core

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"

	"github.com/go-redis/redis/v8"
	"github.com/segmentio/ksuid"

	"com.lambrospetrou/aws-playground/aws-lambda-upstash-redis-article/redisdb"
)

// contextKey is a value for use with context.WithValue. It's used as
// a pointer so it fits in an interface{} without allocation. This technique
// for defining context keys was copied from Go 1.7's new use of context in net/http.
type contextKey struct {
	name string
}

const (
	COOKIE_AUTH_NAME = "xxx_session_id"
)

var (
	CTX_USER_ID = &contextKey{"LoggedInUserId"}
	redisDb     = redisdb.NewNative()
)

func NewMux() *chi.Mux {
	// https://github.com/go-chi/chi/blob/b6a2c5a909f66db8b2166b69628fff095ed51adc/_examples/rest/main.go
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)

	r.Get("/login", login)
	r.Post("/login", login)
	r.Group(func(r chi.Router) {
		r.Use(UsersWithSessionOnly)
		r.Get("/lessons/completed", listLessonsCompleted)
		r.Post("/lessons/{lessonSlug}/mark-complete", markLessonComplete)
	})

	return r
}

func login(w http.ResponseWriter, r *http.Request) {
	// Check credentials and update redis session and return Set-Cookie
	// WARNING: You should do an actual validation in production for credentials!
	// ...
	// For now we always assume correctness and automatically create a session token
	// by saving it to Redis, and also setting it as a cookie.

	userId := strings.TrimSpace(r.FormValue("userId"))
	if userId == "" {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, struct{ Message string }{Message: "Missing required userId"})
		return
	}

	sessionId := ksuid.New()
	redisDb.Set(r.Context(), "session:"+sessionId.String(), userId, time.Hour*1)
	http.SetCookie(w, &http.Cookie{Name: COOKIE_AUTH_NAME, Value: sessionId.String()})

	http.Redirect(w, r, "/lessons/completed", http.StatusTemporaryRedirect)
}

func listLessonsCompleted(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId := r.Context().Value(CTX_USER_ID).(string)

	lessons, err := redisDb.HGetAll(ctx, "lessons:"+userId).Result()
	if err == redis.Nil {
		lessons = map[string]string{}
	} else if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, struct{ Message string }{Message: "We could not load your lessons..."})
		return
	}

	render.JSON(w, r, struct {
		Lessons map[string]string
	}{
		lessons,
	})
}

func markLessonComplete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	lessonSlug := chi.URLParam(r, "lessonSlug")
	userId := r.Context().Value(CTX_USER_ID).(string)
	timeNow := time.Now().Format(time.RFC3339)

	err := redisDb.HSet(ctx, "lessons:"+userId, lessonSlug, timeNow).Err()
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, struct{ Message string }{Message: "We could not save your progression..."})
		return
	}

	render.JSON(w, r, struct {
		LessonSlug    string
		LastCompleted string
	}{
		lessonSlug,
		timeNow,
	})
}

// UsersOnly middleware restricts access to just logged-in users.
// If validation passes, then the context will contain the user id (CTX_USER_ID).
func UsersWithSessionOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check the cookie token against Redis and fetch the user id
		c, err := r.Cookie(COOKIE_AUTH_NAME)
		if err != nil {
			render.Status(r, http.StatusForbidden)
			render.JSON(w, r, struct{}{})
			return
		}
		ctx := r.Context()
		userId, err := redisDb.Get(ctx, "session:"+c.Value).Result()
		if err == redis.Nil {
			render.Status(r, http.StatusForbidden)
			render.JSON(w, r, struct{}{})
			return
		} else if err != nil {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, struct{ Message string }{Message: "We could not validate the provided session ID"})
			return
		}
		// Set it for downstream middleware and handlers.
		next.ServeHTTP(w, r.WithContext(context.WithValue(ctx, CTX_USER_ID, userId)))
	})
}
