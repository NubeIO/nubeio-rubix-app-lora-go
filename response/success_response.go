package response

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-rest-go/rest"
	"net/http"
	"reflect"
)

func Created(id string) rest.IResponse {
	return rest.Success(http.StatusCreated, rest.JSON{"id": id})
}

func Data(model interface{}) rest.IResponse {
	v := reflect.ValueOf(model)
	fmt.Println(5555, v, 6666,  v.Kind() == reflect.Slice)
	if v.Kind() == reflect.Slice {
		return rest.Success(http.StatusOK, rest.JSON{"count": v.Len(), "items": model})
	}
	return rest.Success(http.StatusOK, model)
}

func OK(resp interface{}) rest.IResponse {
	return rest.Success(http.StatusOK, resp)
}

func OKWithMessage(resp string) rest.IResponse {
	return rest.Success(http.StatusOK, resp)
}

