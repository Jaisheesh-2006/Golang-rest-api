package student

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/Jaisheesh-2006/Golang-rest-api/internal/storage"
	"github.com/Jaisheesh-2006/Golang-rest-api/internal/types"
	"github.com/Jaisheesh-2006/Golang-rest-api/internal/utils/responses"
	"github.com/go-playground/validator/v10"
)

func New(storage storage.Storage) http.HandlerFunc {
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
			validationErrors := err.(validator.ValidationErrors) //*type assertion
			responses.WriteJson(w, http.StatusBadRequest, responses.ValidationError(validationErrors))
			return
		}

		//* create student in the database
		id, err := storage.CreateStudent(student.Name, student.Age, student.Email)
		if err != nil {
			slog.Error("Error while creating student", slog.String("error", err.Error()))
			responses.WriteJson(w, http.StatusInternalServerError, responses.GeneralError(err))
			return
		}
		slog.Info("Student created successfully", slog.Int64("id", id))
		responses.WriteJson(w, http.StatusCreated, map[string]int64{
			"id": id,
		})

	}
}

func GetById(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		//* Log first
		slog.Info("Getting student by ID", slog.String("id", id))

		//* Lets convert id to int64
		intId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			responses.WriteJson(w, http.StatusBadRequest, responses.GeneralError(err))
			return
		}
		student, err := storage.GetStudentById(intId)
		if err != nil {
			slog.Error("Error while fetching student", slog.String("error", err.Error()))
			responses.WriteJson(w, http.StatusInternalServerError, responses.GeneralError(err))
			return
		}
		responses.WriteJson(w, http.StatusOK, student)
	}

}
func GetList(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Getting list of students")

		students, err := storage.GetStudents()
		if err != nil {
			slog.Error("Error while fetching students", slog.String("error", err.Error()))
			responses.WriteJson(w, http.StatusInternalServerError, responses.GeneralError(err))
			return
		}
		responses.WriteJson(w, http.StatusOK, students)
		slog.Info("Students fetched successfully", slog.Int("count", len(students)))

	} 

}
