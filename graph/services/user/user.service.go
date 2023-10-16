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
