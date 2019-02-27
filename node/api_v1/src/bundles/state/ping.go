package state

import (
	"net/http"

	"../../core/common"
)

func ping(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json; charset=utf-8")
	common.WriteJSONBody(&response, http.StatusOK, Message{Msg: "pong"})
}
