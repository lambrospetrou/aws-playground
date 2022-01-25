.PHONY: bundle

bundle: prepare_output_dir service1 service2
	mv web-service-1/web-service-1 build/bin/
	mv web-service-2/web-service-2 build/bin/
	cp build-tools/hello.sh build/bin/
	cp -rf .platform .ebextensions Buildfile Procfile build/
	cd build; zip -r bundle.zip .

prepare_output_dir:
	rm -rf build
	mkdir -p build/bin

service1: prepare_output_dir
	cd web-service-1; GOARCH=amd64 GOOS=linux go build -o web-service-1

service2: prepare_output_dir
	cd web-service-2; GOARCH=amd64 GOOS=linux go build -o web-service-2
