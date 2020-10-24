release: clean compile
	zip -r bundle.zip bin/ deploy-tools/ .platform Procfile

compile:
	GOOS=linux go build -o bin/application

clean:
	rm -rf bin/ bundle.zip
