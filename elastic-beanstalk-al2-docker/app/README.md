# Go + SQLite3

## How to run

```
➜  web-sqlite git:(master) ✗ GOOS=linux go build
➜  web-sqlite git:(master) ✗ docker run --rm -it --name web-sqlite --mount type=bind,source="$(pwd)",target=/app -e NUM_INSERT=0 -e SQLITE_PATH=/app/users.db alpine:3.14 "/app/web-sqlite"
2022/01/25 21:10:14 finished inserting 100 users
2022/01/25 21:10:14 Time elapsed in inserts is: 14.39ms
2022/01/25 21:10:14 finished fetching 673547 rows
2022/01/25 21:10:14 all rows: 1121910
2022/01/25 21:10:14 Time elapsed in query is: 752.3374ms
2022/01/25 21:10:14 Time elapsed in main() is: 783.4261ms
```

## BUSY TIMEOUT

At the moment the second call to the database fails with `BUSY_TIMEOUT` so we need to take care of it.
Use <https://github.com/zombiezen/go-sqlite> instead which probably allows for easier handling.
I don't care for this case though to fix it since it's just testing Docker setup.
