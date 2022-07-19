package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"example.com/crud/pkg/models"
	"github.com/gorilla/mux"
)

func (h handler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	// Read dynamic id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Find the book by Id
	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	// Delete that book
	h.DB.Delete(&book)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Deleted")
}
