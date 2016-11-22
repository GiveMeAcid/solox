package users

import "database/sql"
import "net/http"

var db *sql.DB

func UserIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	rows, err := db.Query("SELECT * FROM user_info")
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	w.Write([]byte("server status error"))
	defer rows.Close()
}

/**
func Insert(db *sql.DB, name, age string) (int64, error) {
	res, err := db.Exec("INSERT INTO user_info VALUES (default, $1, $2)",
		name, age)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func remove(db *sql.DB, id []string) error {
	stmt, err := db.Prepare("DELETE FROM user_info WHERE id=$1")
	if err != nil {
		return err
	}
	defer stmt.Close()
	for _, v := range id {
		_, err := stmt.Exec(v)
		if err != nil {
			return err
		}
	}
	return nil
}

func update(db *sql.DB, id, name, age string) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	_, err = tx.Exec("UPDATE user_info SET name = $1, "+
		"age = $2 WHERE id=$3",
		name, age, id)
	if err != nil {
		return err
	}
	return tx.Commit()
}

// func (u *UserInfo) addUser() {
// 	// resp, err := http.Get("http://example.com/")
// }

// func (u *UserInfo) removeUser() {

// }

// func (u *UserInfo) getUser() {

// }
**/
