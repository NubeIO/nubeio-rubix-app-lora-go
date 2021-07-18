package response

import (
	"github.com/NubeIO/nubeio-rubix-lib-rest-go/rest"
	"net/http"
)

func BadEntity(excepted ...string) rest.IResponse {
	return rest.Failed(http.StatusUnprocessableEntity, excepted)
}

func NotFound(err string) rest.IResponse {
	return rest.Failed(http.StatusNotFound, err)
}