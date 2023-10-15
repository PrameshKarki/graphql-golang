package services

import (
	"fmt"

	"github.com/PrameshKarki/event-management-golang/configs"
	"github.com/PrameshKarki/event-management-golang/graph/model"
	"github.com/doug-martin/goqu/v9"
)

func CreateSession(eventID string, body *model.SessionInput) (int, error) {
	db := configs.GetDatabaseConnection()
	ds := configs.GetDialect().Insert(configs.TABLE_NAME["SESSION"]).
		Cols("name", "start_time", "end_time", "description", "event_id").
		Vals(
			goqu.Vals{body.Name, body.StartTime, body.EndTime, body.Description, eventID},
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

func GetEventSession(eventID string) ([]*model.Session, error) {
	db := configs.GetDatabaseConnection()
	ds := configs.GetDialect().Select(
		"id", "name", "start_time", "end_time", "description",
	).From(configs.TABLE_NAME["SESSION"]).Where(goqu.Ex{
		"event_id": eventID,
	})
	sql, _, _ := ds.ToSQL()
	fmt.Println("SQL", sql)
	var sessions []*model.Session
	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var session model.Session
		err := rows.Scan(
			&session.ID,
			&session.Name,
			&session.StartTime,
			&session.EndTime,
			&session.Description,
		)
		if err != nil {
			panic(err)
		}
		sessions = append(sessions, &session)
	}

	if err != nil {
		panic(err)
	}
	return sessions, nil
}

func DeleteSession(id string) (int, error) {
	db := configs.GetDatabaseConnection()
	ds := configs.GetDialect().Delete(configs.TABLE_NAME["SESSION"]).Where(goqu.Ex{
		"id": id,
	})
	sql, _, _ := ds.ToSQL()
	fmt.Println("SQL", sql)
	res, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
	affectedRows, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	return int(affectedRows), nil
}

func GetEventIDFromSession(id string) (string, error) {
	db := configs.GetDatabaseConnection()
	ds := configs.GetDialect().Select("event_id").From(configs.TABLE_NAME["SESSION"]).Where(goqu.Ex{
		"id": id,
	})
	sql, _, _ := ds.ToSQL()
	fmt.Println("SQL", sql)
	var eventID string
	err := db.QueryRow(sql).Scan(&eventID)
	if err != nil {
		panic(err)
	}
	return eventID, nil
}

func UpdateSession(id string, body *model.SessionInput) (int, error) {
	db := configs.GetDatabaseConnection()
	ds := configs.GetDialect().Update(configs.TABLE_NAME["SESSION"]).Set(goqu.Record{
		"name":        body.Name,
		"start_time":  body.StartTime,
		"end_time":    body.EndTime,
		"description": body.Description,
	}).Where(goqu.Ex{
		"id": id,
	})
	sql, _, _ := ds.ToSQL()
	fmt.Println("SQL", sql)
	res, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
	affectedRows, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	return int(affectedRows), nil
}
