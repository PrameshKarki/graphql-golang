package services

import (
	"fmt"

	"github.com/PrameshKarki/event-management-golang/configs"
	"github.com/PrameshKarki/event-management-golang/graph/model"
	"github.com/doug-martin/goqu/v9"
)

func CreateExpense(eventID string, body model.ExpenseInput) (int, error) {
	db := configs.GetDatabaseConnection()
	ds := configs.GetDialect().Insert(configs.TABLE_NAME["EXPENSE"]).
		Cols("item_name", "cost", "description", "category", "event_id").
		Vals(
			goqu.Vals{body.ItemName, body.Cost, body.Description, body.Category, eventID},
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

func GetExpensesOfEvent(eventID string) ([]*model.Expense, error) {
	db := configs.GetDatabaseConnection()
	ds := configs.GetDialect().Select("ID", "item_name", "cost", "description", "category").From(configs.TABLE_NAME["EXPENSE"]).Where(goqu.Ex{"event_id": eventID})
	sql, _, _ := ds.ToSQL()
	fmt.Println("SQL", sql)
	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}
	var expenses []*model.Expense
	for rows.Next() {
		var expense model.Expense
		err := rows.Scan(&expense.ID, &expense.ItemName, &expense.Cost, &expense.Description, &expense.Category)
		if err != nil {
			panic(err)
		}
		expenses = append(expenses, &expense)
	}
	return expenses, nil
}

func RemoveExpense(expenseID string) (int, error) {
	db := configs.GetDatabaseConnection()
	ds := configs.GetDialect().Delete(configs.TABLE_NAME["EXPENSE"]).Where(goqu.Ex{"id": expenseID})
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

func GetExpense(expenseID string) (*model.ExpenseWithEvent, error) {
	db := configs.GetDatabaseConnection()
	ds := configs.GetDialect().Select("expenses.ID", "expenses.item_name", "expenses.cost", "expenses.description", "expenses.category", "events.id", "events.name", "events.start_date", "events.end_date", "events.location", "events.description").From(configs.TABLE_NAME["EXPENSE"]).Join(goqu.T(configs.TABLE_NAME["EVENT"]), goqu.On(goqu.Ex{"expenses.event_id": goqu.I("events.id")})).Where(goqu.Ex{"expenses.id": expenseID})
	sql, _, _ := ds.ToSQL()
	fmt.Println("SQL", sql)
	row := db.QueryRow(sql)
	var expense model.ExpenseWithEvent
	var event model.Event
	err := row.Scan(&expense.ID, &expense.ItemName, &expense.Cost, &expense.Description, &expense.Category, &event.ID, &event.Name, &event.StartDate, &event.EndDate, &event.Location, &event.Description)
	if err != nil {
		panic(err)
	}
	expense.Event = &event
	return &expense, nil
}

func GetEventID(expenseID string) (string, error) {
	db := configs.GetDatabaseConnection()
	ds := configs.GetDialect().Select("event_id").From(configs.TABLE_NAME["EXPENSE"]).Where(goqu.Ex{"id": expenseID})
	sql, _, _ := ds.ToSQL()
	fmt.Println("SQL", sql)
	row := db.QueryRow(sql)
	var eventID string
	err := row.Scan(&eventID)
	if err != nil {
		panic(err)
	}
	return eventID, nil
}

func SumOfExpenseByCategory(eventID string) ([]*model.ExpensesByCategory, error) {
	db := configs.GetDatabaseConnection()
	ds := configs.GetDialect().Select("category", goqu.SUM("cost").As("total_cost")).From(configs.TABLE_NAME["EXPENSE"]).Where(goqu.Ex{"event_id": eventID}).GroupBy("category")
	sql, _, _ := ds.ToSQL()
	fmt.Println("SQL", sql)
	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}
	var expenses []*model.ExpensesByCategory
	for rows.Next() {
		var expense model.ExpensesByCategory
		err := rows.Scan(&expense.Category, &expense.Total)
		if err != nil {
			panic(err)
		}
		expenses = append(expenses, &expense)
	}
	return expenses, nil
}
