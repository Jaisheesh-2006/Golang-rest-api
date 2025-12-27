package student

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/Jaisheesh-2006/Golang-rest-api/internal/types"
	"github.com/Jaisheesh-2006/Golang-rest-api/internal/utils/responses"
	"github.com/go-playground/validator/v10"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student)

		if errors.Is(err, io.EOF) || err != nil {
			responses.WriteJson(w, http.StatusBadRequest, responses.GeneralError(err))
			return
		}
		slog.Info("Creating a student")

		//* request validation
		if err := validator.New().Struct(student); err != nil {
			validationErrors := err.(validator.ValidationErrors)
			responses.WriteJson(w, http.StatusBadRequest, responses.ValidationError(validationErrors))
			return
		}

		responses.WriteJson(w, http.StatusCreated, map[string]string{
			"message": "Student created successfully",
		})

	}
}
