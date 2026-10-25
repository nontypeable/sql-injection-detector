package safeparam

import (
	"database/sql"
	"net/http"
)

// Handler безопасен: значение из запроса передаётся как параметр.
func Handler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")

	rows, err := db.Query("SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
}
