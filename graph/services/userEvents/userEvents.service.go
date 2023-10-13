package services

import (
	"fmt"

	"github.com/PrameshKarki/event-management-golang/configs"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
)

const TABLE_NAME = `user_events`

func CreateUserEvent(eventId string, userId string, role string) (int, error) {
	db := configs.GetDatabaseConnection()
	ds := configs.GetDialect().Insert(TABLE_NAME).
		Cols("user_id", "event_id", "role").
		Vals(
			goqu.Vals{userId, eventId, role},
		)
	sql, _, _ := ds.ToSQL()
	fmt.Println("SQL", sql)
	res, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}
	return int(id), nil

}
