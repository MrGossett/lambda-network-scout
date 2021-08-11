package main

import (
	"context"
	"log"
	"net/http/httptrace"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func main() {
	lambda.StartHandler(handler{})
}

type handler struct{}

func (h handler) Invoke(ctx context.Context, in []byte) ([]byte, error) {
	ctx = httptrace.WithClientTrace(ctx, &httptrace.ClientTrace{
		GotConn: func(connInfo httptrace.GotConnInfo) {
			log.Printf("Got Conn: %+v\n", connInfo)
		},
		DNSDone: func(dnsInfo httptrace.DNSDoneInfo) {
			log.Printf("DNS Info: %+v\n", dnsInfo)
		},
	})

	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatal(err)
	}

	svc := dynamodb.NewFromConfig(cfg)

	if _, err := svc.ListTables(ctx, &dynamodb.ListTablesInput{}); err != nil {
		return nil, err
	}

	return nil, nil
}
