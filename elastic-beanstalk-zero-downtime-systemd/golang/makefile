release: clean compile
	zip -r bundle.zip bin/application

compile:
	GOOS=linux go build -o bin/application

clean:
	rm -rf bin/ bundle.zip
