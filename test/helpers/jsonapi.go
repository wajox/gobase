package helpers

import (
	"bytes"
	"encoding/json"

	"github.com/google/jsonapi"
)

// JSONModelToBytesJSON - preforms model -> bytes.Reader(implements io.Reader).
// useful for passing payload to http request in tests
//
// Example:
// ```
// fakeUpdEmailReq := helpers.GetUpdateEmailRequestModel()
// payload, _ := helpers.JSONModelToBytesJSON(fakeUpdEmailReq)
// ctx.Request, _ = http.NewRequest(
// 	"POST",
// 	"/api/v1/credentials-api/update_email",
// 	payload,
// )
// ```
func JSONModelToBytesJSON(model interface{}) (*bytes.Reader, error) {
	payloadPage, err := jsonapi.Marshal(model)
	if err != nil {
		return nil, err
	}

	payloadBytes, err := json.Marshal(payloadPage)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(payloadBytes), nil
}
