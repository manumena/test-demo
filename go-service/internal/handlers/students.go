package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"go-service/internal/students"
)

type StudentHandler struct {
	Students *students.Client
}

func NewStudentHandler(studentsClient *students.Client) *StudentHandler {
	return &StudentHandler{
		Students: studentsClient,
	}
}

// GET /students/:id
func (h *StudentHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := strings.TrimPrefix(r.URL.Path, "/students/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid student id", http.StatusBadRequest)
		return
	}

	student, err := h.Students.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	log.Printf("Fetched student: %+v\n", student)
	_ = json.NewEncoder(w).Encode(student)
}