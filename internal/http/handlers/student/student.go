package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/Mrinal-Agrawal21/student-api/internal/storage"
	"github.com/Mrinal-Agrawal21/student-api/internal/types"
	"github.com/Mrinal-Agrawal21/student-api/internal/utils/response"
	"github.com/go-playground/validator"
)
func NewStudentHandler(storage storage.Storage) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {


		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err,io.EOF) {
			response.WriteJson(w,http.StatusBadRequest , response.GeneralError(fmt.Errorf("empty body")))
			return
		}

		if err != nil {
			response.WriteJson(w,http.StatusBadRequest,response.GeneralError(err))
		}

		slog.Info("creating a student")

		// Request Validation

		if err:=validator.New().Struct(student); err != nil {
			validateErrs := err.(validator.ValidationErrors)
			response.WriteJson(w,http.StatusBadRequest , response.ValidationError(validateErrs))
			return 
		}

		lastId , err :=storage.CreateStudent(
			student.Name,
			student.Email,
			student.Age,
		)
		slog.Info("user Created successfully" , slog.String("userId" , fmt.Sprint(lastId)))

		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError ,err)
		}

		response.WriteJson(w,http.StatusCreated,map[string]int64 {"id":lastId})
	}
}
func GetStudentByIdHandler(storage storage.Storage) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		slog.Info("fetching a student" , slog.String("id" , fmt.Sprint(id)))

		intId , err := strconv.ParseInt(id,10,64)
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest ,response.GeneralError(err))
		}

		student , err := storage.GetStudentById(intId)
		if err != nil {
			slog.Error("error fetching student" , slog.String("id" , fmt.Sprint(id)))
			response.WriteJson(w, http.StatusInternalServerError ,response.GeneralError(err))
			return
		}
		slog.Info("student fetched successfully" , slog.String("id" , fmt.Sprint(id)))
		response.WriteJson(w,http.StatusOK,student)
	}
}