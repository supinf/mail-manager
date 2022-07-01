package main

import (
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-xray-sdk-go/xray"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	"github.com/go-openapi/loads"
	flags "github.com/jessevdk/go-flags"
	"github.com/supinf/supinf-mail/app-api/config"
	"github.com/supinf/supinf-mail/app-api/generated/swagger/v1/restapi"
	"github.com/supinf/supinf-mail/app-api/generated/swagger/v1/restapi/operations"
	"github.com/supinf/supinf-mail/app-api/logs"
	"github.com/supinf/supinf-mail/app-api/model"
)

var httpAdapter *httpadapter.HandlerAdapter

func init() {
	logs.SetHooks()
	logs.Debug(config.ApplicationsName, nil, &logs.Map{
		"Stage":   config.AppStage,
		"Version": config.AppVersion,
	})
	if config.AwsXRay {
		err := xray.Configure(xray.Config{LogLevel: config.AwsXRayLogLevel})
		logs.Debug("AWS X-Ray configured", err, nil)
	}
	if err := model.Initialize(); err != nil {
		logs.Error("unable initialize database", err, nil)
	}

	// ----------------------------------------------------------------------------------------
	//  Copied from app-api/generated/swagger/v1/cmd/app-server/main.go
	// ----------------------------------------------------------------------------------------
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewAppAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "SUPINF MAIL API"
	parser.LongDescription = "SUPINF MAIL API 仕様\n"
	server.ConfigureFlags()
	for _, optsGroup := range api.CommandLineOptionsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			log.Fatalln(err)
		}
	}

	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}

	server.ConfigureAPI()

	//if err := server.Serve(); err != nil {
	//	log.Fatalln(err)
	//}
	httpAdapter = httpadapter.New(server.GetHandler())
	// ----------------------------------------------------------------------------------------
}

// Handler handles API requests
func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return httpAdapter.Proxy(req)
}

func main() {
	lambda.Start(Handler)
}
