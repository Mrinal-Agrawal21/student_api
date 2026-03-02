package student

import "net/http"


func NewStudentHandler() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, new student handler!"))
	}
}