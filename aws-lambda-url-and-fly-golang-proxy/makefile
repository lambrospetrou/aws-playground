.Phony: default, build, clean, build-fly, sam-invoke, sam-package, sam-deploy

default: build

clean:
	rm -rf build/

build: clean
	CGO_ENABLED=0 go build -o build/handler cmd/lambda/main.go
	cd build/ && zip handler.zip ./handler
	cp ./build/handler.zip /mnt/c/Users/lambros/Downloads/handler.zip

build-fly: clean
	CGO_ENABLED=0 go build -o build/fly/app cmd/server/main.go
	cp Dockerfile build/fly/
	cp fly.toml build/fly/

sam-invoke:
	sam local invoke -t aws-iac/sam-template.yml

# Call with `make sam-package -- --region <region_name>`
sam-package: build
	sam package -t aws-iac/sam-template.yml --output-template-file "./build/sam-packaged.yml" --s3-prefix "hello-go-redis-" --resolve-s3
	cp ./build/sam-packaged.yml /mnt/c/Users/lambros/Downloads/sam-packaged.yml

sam-deploy:
	sam deploy -t aws-iac/sam-template.yml --stack-name "HelloGoRedisApiStack-${DEPLOY_STAGE:-dev}" --region "${DEPLOY_REGION:-eu-west-1}" --resolve-s3 --no-confirm-changeset --no-fail-on-empty-changeset --capabilities CAPABILITY_IAM CAPABILITY_NAMED_IAM CAPABILITY_AUTO_EXPAND
