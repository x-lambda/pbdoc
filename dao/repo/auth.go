package repo

import (
	"context"
	"time"

	"pbdoc/util/db"
)

type Auth struct {
	ID         int32
	ProjectID  int32
	AccessKey  string
	CreateUser string
	CreateAt   time.Time
	ModifyOn   time.Time
}

func GetByProjectID(ctx context.Context, projectID int32) (list []Auth, err error) {
	conn := db.Get(ctx, "")
	sql := "select id, project_id, access_key, create_user, create_at, modify_on from auth where project_id = ?"
	rows, err := conn.QueryContext(ctx, sql, projectID)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var auth Auth
		if err = rows.Scan(
			&auth.ID,
			&auth.ProjectID,
			&auth.AccessKey,
			&auth.CreateUser,
			&auth.CreateAt,
			&auth.ModifyOn,
		); err != nil {
			return
		}

		list = append(list, auth)
	}
	return
}

func QueryByAccessKey(ctx context.Context, projectID int32, accessKey string) (auth Auth, err error) {
	conn := db.Get(ctx, "")
	sql := "select id, project_id, access_key, create_user, create_at, modify_on from auth where project_id = ? and access_key = ?"
	err = conn.QueryRowContext(ctx, sql, projectID, accessKey).Scan(
		&auth.ID,
		&auth.ProjectID,
		&auth.AccessKey,
		&auth.CreateUser,
		&auth.CreateAt,
		&auth.ModifyOn,
	)
	return
}

func AddAuth(ctx context.Context, projectID int32, accessKey string)
