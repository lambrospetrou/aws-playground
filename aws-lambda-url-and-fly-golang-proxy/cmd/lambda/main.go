package lib

import (
	"com.lambrospetrou/aws-playground/aws-lambda-url-and-fly-golang-proxy/core"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
)

func main() {
	mutex := core.NewMutex()
	lambda.Start(httpadapter.NewV2(mutex).ProxyWithContext)
}
