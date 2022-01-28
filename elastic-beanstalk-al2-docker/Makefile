.PHONY: bundle

bundle: prepare_output_dir web1 web-sqlite
	cp -rf .platform .ebextensions build/
	cp docker-compose.prod.yml build/docker-compose.yml
	cd build; zip -r bundle.zip .

dev: prepare_output_dir web1 web-sqlite
	cp -rf .platform .ebextensions build/
	cp docker-compose.dev.yml build/docker-compose.yml
	cd build; zip -r bundle.zip .

prepare_output_dir:
	rm -rf build
	mkdir -p build/bin

web1: prepare_output_dir
	cd web-1; GOARCH=amd64 GOOS=linux go build -o web-1
	mv web-1/web-1 build/bin/
	cp web-1/Dockerfile build/Dockerfile.web1

web-sqlite: prepare_output_dir
	cd web-sqlite; GOARCH=amd64 GOOS=linux go build -o web-sqlite
	mv web-sqlite/web-sqlite build/bin/
	cp web-sqlite/Dockerfile build/Dockerfile.web-sqlite

docker-compose: bundle
	cd build/ && docker-compose up

docker-compose-fresh: bundle
	cd build/ && docker-compose up --build --remove-orphans
