package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/Mrinal-Agrawal21/student-api/internal/types"
	"github.com/Mrinal-Agrawal21/student-api/internal/utils/response"
)
func NewStudentHandler() http.HandlerFunc{
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

		response.WriteJson(w,http.StatusCreated,map[string]string {"success":"OK"})
	}
}