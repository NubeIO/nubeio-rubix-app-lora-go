package response

import (
	"github.com/NubeIO/nubeio-rubix-lib-rest-go/rest"
	"net/http"
	"reflect"
)

func Created(id string) rest.IResponse {
	return rest.Success(http.StatusCreated, rest.JSON{"id": id})
}

func Data(model interface{}) rest.IResponse {
	v := reflect.ValueOf(model)
	if v.Kind() == reflect.Slice {
		return rest.Success(http.StatusOK, rest.JSON{"count": v.Len(), "items": model})
	}
	return rest.Success(http.StatusOK, model)
}

func OK(resp interface{}) rest.IResponse {
	return rest.Success(http.StatusOK, resp)
}