package repo

import (
	"context"
	"time"

	"pbdoc/util/db"
)

type Project struct {
	ID         int32
	Name       string
	CreateUser string
	CreateAt   time.Time
	ModifyOn   time.Time
}

// GetProjectByID
func GetProjectByID(ctx context.Context, id int32) (project Project, err error) {
	conn := db.Get(ctx, "")
	sql := "select id, name, create_user, create_at, modify_on from t_project where id = ?"
	if err = conn.QueryRowContext(ctx, sql, id).Scan(
		&project.ID,
		&project.Name,
		&project.CreateUser,
		&project.CreateAt,
		&project.ModifyOn,
	); err != nil {
		if db.IsNoRow(err) {
			err = nil
		}

		return
	}

	return
}

// GetProjectByName
func GetProjectByName(ctx context.Context, name string) (project Project, err error) {
	conn := db.Get(ctx, "")
	sql := "select id, name, create_user, create_at, modify_on from t_project where name = ?"
	if err = conn.QueryRowContext(ctx, sql, name).Scan(
		&project.ID,
		&project.Name,
		&project.CreateUser,
		&project.CreateAt,
		&project.ModifyOn,
	); err != nil {
		if db.IsNoRow(err) {
			err = nil
		}

		return
	}
	return
}

// AddProject
func AddProject(ctx context.Context, project Project) (id int32, err error) {
	return
}

func DeleteProject(ctx context.Context, id int32) (err error) {
	return
}

func UpdateProject(ctx context.Context, project Project) (err error) {
	return
}

func ListAllProject(ctx context.Context) (list []Project, err error) {
	conn := db.Get(ctx, "")
	sql := "select id, name, create_user, create_at, modify_on from t_project"
	rows, err := conn.QueryContext(ctx, sql)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var p Project
		if err = rows.Scan(
			&p.ID,
			&p.Name,
			&p.CreateUser,
			&p.CreateAt,
			&p.ModifyOn,
		); err != nil {
			return
		}
		list = append(list, p)
	}

	return
}
