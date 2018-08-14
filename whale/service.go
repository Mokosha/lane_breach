package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
)

type GatewayRequest struct {
	Path                  string            `json:"path"`
	HttpMethod            string            `json:"httpMethod"`
	Headers               map[string]string `json:"headers"`
	QueryStringParameters map[string]string `json"queryStringParameters"`
	PathParameters        map[string]string `json:"pathParameters"`
	RequestContext        interface{}       `json:"requestContext"`
}

type APIResponse struct {
	StatusCode *int64            `json:"statusCode"`
	Body       string            `json:"body"`
	Headers    map[string]string `json:"headers"`
}

type WhaleResponse struct {
	Metadata map[string]string `json:"metadata"`
	Data     interface{}       `json:"data"`
}

func (wr *WhaleResponse) toString() string {
	if d, err := json.Marshal(wr); err != nil {
		return ""
	} else {
		return string(d)
	}
}

//
// Service connects the data layer to the request handler
//
type service struct {
	dataLayer DataLayer
}

func (s *service) parseArgs(request GatewayRequest) (*GetDataArgs, error) {
	return &GetDataArgs{
		Table: request.PathParameters["table"],
		Key:   request.PathParameters["key"],
	}, nil
}

func (s *service) buildResponse(response DataResponse) (*WhaleResponse, error) {
	return &WhaleResponse{
		Data: response.Data,
	}, nil
}

func (s *service) lamdaHandler(ctx context.Context, request GatewayRequest) (*APIResponse, error) {
	args, err := s.parseArgs(request)
	if err != nil {
		return nil, err
	}

	data, err := s.dataLayer.GetData(args)
	if err != nil {
		return nil, err
	}

	response, err := s.buildResponse(data)
	if err != nil {
		return nil, err
	}
	return &APIResponse{
		StatusCode: aws.Int64(200),
		Body:       response.toString(),
		Headers: map[string]string{
			"Access-Control-Allow-Origin": "*",
		},
	}, nil
}
