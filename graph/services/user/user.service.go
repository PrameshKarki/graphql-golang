package services

import (
	"github.com/PrameshKarki/event-management-golang/configs"
	"github.com/PrameshKarki/event-management-golang/graph/model"
	"github.com/doug-martin/goqu/v9"
	"github.com/sirupsen/logrus"
)

func FindOneByEmail(email string) *model.User {
	db := configs.GetDatabaseConnection()
	ds := configs.GetDialect().Select("id", "email", "phone_number", "password").From(configs.TABLE_NAME["USER"]).Where(goqu.Ex{"email": email})
	sql, _, _ := ds.ToSQL()
	logrus.Info("SQL", sql)
	row := db.QueryRow(sql)
	var user model.User
	row.Scan(&user.ID, &user.Email, &user.PhoneNumber, &user.Password)
	return &user
}

func Find() ([]*model.PrivateUser, error) {
	var users []*model.PrivateUser
	db := configs.GetDatabaseConnection()
	ds := configs.GetDialect().From(configs.TABLE_NAME["USER"]).Select("id", "name")
	sql, _, _ := ds.ToSQL()
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var user model.PrivateUser
		if err := rows.Scan(&user.ID, &user.Name); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}
