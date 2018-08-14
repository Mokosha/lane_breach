package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"log"
)

func main() {

	dl, err := NewDataLayer()
	if err != nil {
		log.Fatal(err)
	}

	s := service{
		dataLayer: dl,
	}

	lambda.Start(s.lamdaHandler)
}
