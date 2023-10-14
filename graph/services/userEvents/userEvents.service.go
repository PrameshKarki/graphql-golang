package services

import (
	"fmt"

	"github.com/PrameshKarki/event-management-golang/configs"
	"github.com/PrameshKarki/event-management-golang/graph/model"
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

func AddMembersToEvent(event string, body model.AddMemberInput) (int, error) {
	db := configs.GetDatabaseConnection()

	sql := "INSERT INTO user_events(`user_id`,`event_id`,`role`) VALUES"
	for index, member := range body.Members {
		if index == len(body.Members)-1 {
			sql += fmt.Sprintf("(%s,%s,'%s')", member.ID, event, member.Role)
		} else {
			sql += fmt.Sprintf("(%s,%s,'%s')", member.ID, event, member.Role) + ","
		}
	}
	sql += ";"
	fmt.Println(sql)

	res, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
	id, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	return int(id), nil

}

func RemoveUserFromEvent(eventId string, userId string) (int, error) {
	db := configs.GetDatabaseConnection()
	ds := configs.GetDialect().Delete(TABLE_NAME).
		Where(goqu.Ex{"user_id": userId, "event_id": eventId})
	sql, _, _ := ds.ToSQL()
	fmt.Println("SQL", sql)
	res, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
	id, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	return int(id), nil
}
