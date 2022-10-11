.Phony: default, build, clean

default: build

clean:
	rm -rf build/

build: build-lambda build-server

build-lambda: clean
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o build/handler cmd/lambda/main.go
	cd build/ && zip handler.zip ./handler
	cp ./build/handler.zip /mnt/c/Users/lambros/Downloads/handler.zip

build-server: clean
	CGO_ENABLED=0 go build -o build/server cmd/server/main.go

sam-invoke:
	sam local invoke -t aws-iac/sam-template.yml -e sample-event.json

sam-package: build
	sam package -t aws-iac/sam-template.yml --output-template-file "./build/sam-packaged.yml" --region eu-west-1 --s3-prefix "hello-go-redis-" --resolve-s3
	cp ./build/sam-packaged.yml /mnt/c/Users/lambros/Downloads/sam-packaged.yml

sam-deploy: build-lambda
	sam deploy -t aws-iac/sam-template.yml --stack-name "UpstashRedisGoArticleStackDemo" --region eu-west-1 --resolve-s3 --no-confirm-changeset --no-fail-on-empty-changeset --capabilities CAPABILITY_IAM CAPABILITY_NAMED_IAM CAPABILITY_AUTO_EXPAND
