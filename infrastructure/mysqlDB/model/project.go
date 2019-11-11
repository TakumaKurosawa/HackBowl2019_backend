package model

import "database/sql"

type Project struct {
	Id        int
	Name      string
	Level     int
	StartDate sql.NullString
}
