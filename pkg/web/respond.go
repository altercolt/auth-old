package web

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

var (
	ErrRespondCannotMarshal = errors.New("TODO")
)

func Respond(ctx context.Context, w http.ResponseWriter, data any, statusCode int) error {
	res, err := json.Marshal(data)
	if err != nil {
		return ErrRespondCannotMarshal
	}

	w.WriteHeader(statusCode)
	_, err = w.Write(res)
	if err != nil {
		return err
	}

	return nil
}
