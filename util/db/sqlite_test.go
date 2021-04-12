package db

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	db, err := sql.Open("sqlite3", "./foo.db")
	assert.Nil(t, err)

	s := `
    CREATE TABLE IF NOT EXISTS project(
        id          INTEGER     NOT NULL PRIMARY KEY AUTOINCREMENT,
        name        varchar(64) NOT NULL DEFAULT '',
        create_user varchar(64) NOT NULL DEFAULT '',
        create_at   datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP,
        modify_on   datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP
    );
	`
	_, err = db.Exec(s)
	assert.Nil(t, err)

	_, err = db.Exec("insert into project(name) values(?)", "a")
	assert.Nil(t, err)

	rows, err := db.Query("select id, name from project")
	assert.Nil(t, err)

	for rows.Next() {
		var id int64
		var name string
		if err = rows.Scan(&id, &name); err != nil {
			panic(err)
		}

		fmt.Println(id)
		fmt.Println(name)
	}
}
