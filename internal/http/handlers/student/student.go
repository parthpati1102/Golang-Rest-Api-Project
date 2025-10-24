package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/parthpati1102/Golang-pRest-API-Project/internal/storage"
	types "github.com/parthpati1102/Golang-pRest-API-Project/internal/type"
	"github.com/parthpati1102/Golang-pRest-API-Project/internal/utils/response"
)

func New(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		slog.Info("Creating a Student")

		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student)

		if errors.Is(err, io.EOF) {
			response.Writejson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		if err != nil {
			response.Writejson(w, http.StatusBadRequest, response.GeneralError(err))
		}

		// Request Validation

		if err := validator.New().Struct(student); err != nil {

			validateErrs := err.(validator.ValidationErrors)
			response.Writejson(w, http.StatusBadRequest, response.ValidationError(validateErrs))
			return
		}

		lastId, err := storage.CreateStudent(
			student.Name,
			student.Email,
			student.Age,
		)

		slog.Info("user Created Successfully ", slog.String("userId", fmt.Sprint(lastId)))

		if err != nil {
			response.Writejson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}

		response.Writejson(w, http.StatusCreated, map[string]int64{"id": lastId})
	}
}
