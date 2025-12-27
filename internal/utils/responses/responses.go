package responses

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

const (
	StatusOK    = "OK"
	StatusError = "Error"
)

func WriteJson(w http.ResponseWriter, status int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

func GeneralError(err error) Response {
	return Response{
		Status: StatusError,
		Error:  err.Error(),
	}
}

func ValidationError(errs validator.ValidationErrors) Response {

	var errorMsg []string

	for _,err:=range errs{
		switch err.ActualTag(){
		case "required":
			errorMsg=append(errorMsg,err.Field()+" is required")
		case "min":
			errorMsg=append(errorMsg,err.Field()+" must be at least "+err.Param()+" characters long")
		case "max":
			errorMsg=append(errorMsg,err.Field()+" must be at most "+err.Param()+" characters long")
		case "email":
			errorMsg=append(errorMsg,err.Field()+" must be a valid email address")
		default:
			errorMsg=append(errorMsg,err.Field()+" is not valid")
		}

	}
	return Response{
		Status: StatusError,
		Error: strings.Join(errorMsg, ", "),
	}
	
}