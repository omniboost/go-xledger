package xledger

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-xledger/utils"
)

func NewGraphQLRequest[T any](c *Client) GraphQLRequest[T] {
	r := GraphQLRequest[T]{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type GraphQLRequest[T any] struct {
	client      *Client
	queryParams *GraphQLRequestQueryParams
	pathParams  *GraphQLRequestPathParams
	method      string
	headers     http.Header
	requestBody GraphQLRequestBody
}

func (r GraphQLRequest[T]) NewQueryParams() *GraphQLRequestQueryParams {
	return &GraphQLRequestQueryParams{}
}

type GraphQLRequestQueryParams struct {
}

func (p GraphQLRequestQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	encoder.RegisterEncoder(DateTime{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GraphQLRequest[T]) QueryParams() *GraphQLRequestQueryParams {
	return r.queryParams
}

func (r GraphQLRequest[T]) NewPathParams() *GraphQLRequestPathParams {
	return &GraphQLRequestPathParams{}
}

type GraphQLRequestPathParams struct {
}

func (p *GraphQLRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GraphQLRequest[T]) PathParams() *GraphQLRequestPathParams {
	return r.pathParams
}

func (r *GraphQLRequest[T]) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GraphQLRequest[T]) SetMethod(method string) {
	r.method = method
}

func (r *GraphQLRequest[T]) Method() string {
	return r.method
}

func (r GraphQLRequest[T]) NewRequestBody() GraphQLRequestBody {
	return GraphQLRequestBody{}
}

type GraphQLRequestBody struct {
	Query         string                 `json:"query"`
	Variables     map[string]interface{} `json:"variables,omitempty"`
	OperationName *string                `json:"operationName,omitempty"`
}

func (r *GraphQLRequest[T]) RequestBody() *GraphQLRequestBody {
	return &r.requestBody
}

func (r *GraphQLRequest[T]) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *GraphQLRequest[T]) SetRequestBody(body GraphQLRequestBody) {
	r.requestBody = body
}

func (r *GraphQLRequest[T]) NewResponseBody() *T {
	return new(T)
}

func (r *GraphQLRequest[T]) URL() *url.URL {
	u := r.client.GetEndpointURL("", r.PathParams())
	return &u
}

func (r *GraphQLRequest[T]) Do(ctx context.Context) (T, error) {
	// Create http request
	req, err := r.client.NewRequest(ctx, r)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Process query parameters
	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, true)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}
