package services

import (
	"fmt"

	"github.com/PrameshKarki/event-management-golang/configs"
	"github.com/PrameshKarki/event-management-golang/graph/model"
	"github.com/doug-martin/goqu/v9"
)

func UserSignUp(body model.UserInput) (int, error) {
	db := configs.GetDatabaseConnection()
	ds := configs.GetDialect().Insert(configs.TABLE_NAME["USER"]).
		Cols("name", "email", "phone_number", "password").
		Vals(
			goqu.Vals{body.Name, body.Email, body.PhoneNumber, body.Password},
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
