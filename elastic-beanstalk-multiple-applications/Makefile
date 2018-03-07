.PHONY: bundle

bundle: prepare_output_dir service1 service2 workerservice
	cp web-service-1/web-service-1 build/bin/
	cp web-service-2/web-service-2 build/bin/
	cp worker-service/worker-service build/bin/
	cp build-tools/hello.sh build/bin/
	cp Buildfile Procfile build/
	cd build; zip -r bundle.zip .

prepare_output_dir:
	rm -rf build
	mkdir -p build/bin

service1:
	cd web-service-1; go build

service2:
	cd web-service-2; go build

workerservice:
	cd worker-service; go build