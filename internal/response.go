package internal

import (
	"encoding/json"
	"net/http"
)

func SendSuccessResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func SendCreatedResponse(w http.ResponseWriter, data interface{}) {
	SendSuccessResponse(w, data, http.StatusCreated)
}

func SendOKResponse(w http.ResponseWriter, data interface{}) {
	SendSuccessResponse(w, data, http.StatusOK)
}

func SendNoContentResponse(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

func HandleError(w http.ResponseWriter, err interface{}) {
	w.Header().Set("Content-Type", "application/json")

	var appErr *AppError

	switch e := err.(type) {
	case *AppError:
		appErr = e
	case string:
		appErr = GetError(e)
	case error:
		if ae, ok := e.(*AppError); ok {
			appErr = ae
		} else {
			appErr = InternalServerError
		}
	default:
		appErr = InternalServerError
	}

	w.WriteHeader(appErr.StatusCode)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"code":    appErr.Code,
		"message": appErr.Message,
	})
}
