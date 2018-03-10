.PHONY: bundle

bundle: prepare_output_dir service1 service2
	mv web-service-1/web-service-1 build/web-service-1/
	mv web-service-2/web-service-2 build/web-service-2/
	cp -r proxy Dockerrun.aws.json build/
	cd build; zip -r bundle.zip .

prepare_output_dir:
	rm -rf build
	mkdir -p build/web-service-1
	mkdir -p build/web-service-2

service1: prepare_output_dir
	cd web-service-1; go build

service2: prepare_output_dir
	cd web-service-2; go build
