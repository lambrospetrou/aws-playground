package main

import (
	"com.lambrospetrou/aws-playground/aws-lambda-upstash-redis-article/core"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
)

func main() {
	mux := core.NewMux()
	lambda.Start(httpadapter.NewV2(mux).ProxyWithContext)
}
