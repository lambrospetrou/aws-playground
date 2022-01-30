.PHONY: bundle

bundle: prepare_output_dir app
	mv app/app build/bin/
	cp build-tools/build-hello.sh build/bin/build-hello.sh
	cp -rf .platform .ebextensions build-tools/Buildfile build-tools/Procfile build/
	cd build; zip -r bundle.zip .

prepare_output_dir:
	rm -rf build
	mkdir -p build/bin

app: prepare_output_dir
	cd app; GOARCH=amd64 GOOS=linux go build -o app
