package services

import (
	"fmt"

	"github.com/PrameshKarki/event-management-golang/configs"
	"github.com/PrameshKarki/event-management-golang/graph/model"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
)

const TABLE_NAME = `events`

type MemberInEvent struct {
	user_id  string
	role     string
	event_id string
}

func GetEvents() ([]*model.Event, error) {
	var events []*model.Event
	db := configs.GetDatabaseConnection()
	ds := configs.GetDialect().From(TABLE_NAME).Select("id", "name", "start_date", "end_date", "description")
	sql, _, _ := ds.ToSQL()
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var event model.Event
		if err := rows.Scan(&event.ID, &event.Name, &event.StartDate, &event.EndDate, &event.Description); err != nil {
			return nil, err
		}
		events = append(events, &event)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return events, nil

}
func GetEvent(id string) (*model.Event, error) {
	var event model.Event
	db := configs.GetDatabaseConnection()
	ds := configs.GetDialect().From(TABLE_NAME).Select("id", "name", "start_date", "end_date", "location", "description").Where(goqu.Ex{"id": id})
	sql, _, _ := ds.ToSQL()
	rows, err := db.Query(sql)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&event.ID, &event.Name, &event.StartDate, &event.EndDate, &event.Location, &event.Description); err != nil {
			return nil, err
		}
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return &event, nil

}

func CreateEvent(body model.EventInput, userId string) (int, error) {
	db := configs.GetDatabaseConnection()
	ds := configs.GetDialect().Insert(TABLE_NAME).
		Cols("name", "start_date", "end_date", "location", "description").
		Vals(
			goqu.Vals{body.Name, body.StartDate, body.EndDate, body.Location, body.Description},
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
func DeleteEvent(eventId string) (int, error) {
	db := configs.GetDatabaseConnection()
	ds := configs.GetDialect().Delete(TABLE_NAME).Where(goqu.Ex{"id": eventId})
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

func UpdateEventSchedule(eventId string, body model.ScheduleUpdateInput) (int, error) {
	db := configs.GetDatabaseConnection()
	ds := configs.GetDialect().Update(TABLE_NAME).
		Set(goqu.Record{
			"start_date": body.StartDate,
			"end_date":   body.EndDate,
		}).
		Where(goqu.Ex{"id": eventId})

	sql, _, _ := ds.ToSQL()
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

func UpdateEvent(eventId string, body model.EventInput) (int, error) {
	db := configs.GetDatabaseConnection()
	ds := configs.GetDialect().Update(TABLE_NAME).
		Set(goqu.Record{
			"name":        body.Name,
			"start_date":  body.StartDate,
			"end_date":    body.EndDate,
			"location":    body.Location,
			"description": body.Description,
		}).
		Where(goqu.Ex{"id": eventId})

	sql, _, _ := ds.ToSQL()
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
