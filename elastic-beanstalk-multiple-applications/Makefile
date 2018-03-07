.PHONY: bundle

bundle: prepare_output_dir service1 service2 workerservice
	mv web-service-1/web-service-1 build/bin/
	mv web-service-2/web-service-2 build/bin/
	mv worker-service/worker-service build/bin/
	cp build-tools/hello.sh build/bin/
	cp -r .ebextensions Buildfile Procfile build/
	cd build; zip -r bundle.zip .

prepare_output_dir:
	rm -rf build
	mkdir -p build/bin

service1: prepare_output_dir
	cd web-service-1; go build

service2: prepare_output_dir
	cd web-service-2; go build

workerservice: prepare_output_dir
	cd worker-service; go build
