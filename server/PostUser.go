package server

import (
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"go-swagger-playground/server/gen/restapi/operations"
)

func PostUser(p operations.PostUserParams) middleware.Responder {
	fmt.Println(p)
	return operations.NewPostUserOK()
}
