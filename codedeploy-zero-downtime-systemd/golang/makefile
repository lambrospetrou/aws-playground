release: clean compile
	zip -r bundle.zip bin/ deploy-tools/ appspec.yml

compile:
	GOOS=linux go build -o bin/application

clean:
	rm -rf bin/ bundle.zip
