package handlers

import (
	"custom-server/gen/restapi/operations"
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
)
type greeting struct {

}

func Newgreeting () operations.GetGreetingHandler {
	return &greeting{}
}

func (m greeting) Handle (params operations.GetGreetingParams) middleware.Responder {
	name := swag.StringValue(params.Name)
	if name == "" {
		name = "World"
	}

	greeting := fmt.Sprintf("Hello, %s!\n", name)
	return operations.NewGetGreetingOK().WithPayload(greeting)
}
