.PHONY: bundle

bundle: prepare_output_dir app
	cp -rf .platform .ebextensions build/
	# cp docker-compose.yml build/ # Runs out of memory on Beanstalk
	cd build; zip -r bundle.zip .

prepare_output_dir:
	rm -rf build
	mkdir -p build/bin

app: prepare_output_dir
	cd app; GOARCH=amd64 GOOS=linux go build -o app
	mv app/app build/bin/
	cp app/Dockerfile build/Dockerfile

docker-compose: bundle
	cd build/ && docker-compose up

docker-compose-fresh: bundle
	cd build/ && docker-compose up --build --remove-orphans
