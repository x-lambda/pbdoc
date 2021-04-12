package db

import (
	"context"
	"database/sql"
	"pbdoc/util/conf"
	"sync"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/sync/singleflight"
)

type conn struct {
	db   *sql.DB
	name string
}

var lock sync.RWMutex
var dbs = make(map[string]*conn, 4)
var g = singleflight.Group{}

func InitTable() {
	ctx := context.TODO()
	conn := Get(ctx, "")

	s := `
    CREATE TABLE IF NOT EXISTS project(
        id          INTEGER     NOT NULL PRIMARY KEY AUTOINCREMENT,
        name        varchar(64) NOT NULL DEFAULT '' COMMENT '',
        create_user varchar(64) NOT NULL DEFAULT '' COMMENT '',
        create_at   datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
        modify_on   datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
        PRIMARY KEY (id),
		KEY ix_modify_on (modify_on)
    ) COMMENT='工程表';
    
    CREATE TABLE IF NOT EXISTS access_key(
        id          INTEGER     NOT NULL PRIMARY KEY AUTOINCREMENT,
        value       varchar(64) NOT NULL DEFAULT '' COMMENT '',
        create_user varchar(64) NOT NULL DEFAULT '' COMMENT '',
        create_at   datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
        modify_on   datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
        PRIMARY KEY (id),
		KEY ix_modify_on (modify_on) 
    ) COMMENT='';
    
    CREATE TABLE IF NOT EXISTS auth(
    	id          INTEGER     NOT NULL PRIMARY KEY AUTOINCREMENT,
        project_id  bigint(11) NOT NULL DEFAULT 0 COMMENT '',
        access_key  varchar(64) NOT NULL DEFAULT '' COMMENT '',
        create_user varchar(64) NOT NULL DEFAULT '' COMMENT '',
        create_at   datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
        modify_on   datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
        PRIMARY KEY (id),
		KEY ix_modify_on (modify_on)
    );
    
    CREATE TABLE IF NOT EXISTS branch(
        id         INTEGER     NOT NULL PRIMARY KEY AUTOINCREMENT,
        project_id INTEGER     NOT NULL DEFAULT 0  COMMENT '',
        name       varchar(64) NOT NULL DEFAULT '' COMMENT '',
        create_at   datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
        modify_on   datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
        PRIMARY KEY (id),
		KEY ix_modify_on (modify_on)
    );
    `

	conn.db.Exec(s)
}

func Get(ctx context.Context, name string) *conn {
	lock.RLock()
	db := dbs[name]
	lock.RUnlock()

	if db != nil {
		return db
	}

	v, err, _ := g.Do(name, func() (i interface{}, err error) {
		db, err := sql.Open("sqlite3", conf.DB.Name)
		if err != nil {
			return
		}

		i = &conn{
			db:   db,
			name: name,
		}
		return
	})
	if err != nil {
		panic(err)
	}

	return v.(*conn)
}

func (conn *conn) QueryRowContext(ctx context.Context, sql string, args ...interface{}) *sql.Row {
	return conn.db.QueryRowContext(ctx, sql, args)
}

func (conn *conn) QueryContext(ctx context.Context, sql string, args ...interface{}) (*sql.Rows, error) {
	return conn.db.QueryContext(ctx, sql, args)
}

func (conn *conn) Exec(ctx context.Context, sql string, args ...interface{}) (sql.Result, error) {
	return conn.db.ExecContext(ctx, sql, args)
}
