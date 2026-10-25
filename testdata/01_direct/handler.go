package direct

import (
	"database/sql"
	"net/http"
)

// Handler уязвим: значение из HTTP-запроса склеивается прямо в текст SQL.
func Handler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	query := "SELECT * FROM users WHERE id = " + id

	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
}
