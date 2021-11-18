package meals

import (
	"database/sql"
	"errors"
	"fmt"
)

type MealsServiceInterface interface {
	ListMeals() ([]Meal, error)
	CreateMeal(m Meal) (int, error)
}

type SqliteMealsService struct {
	db *sql.DB
}

func NewSqliteMealsService(db *sql.DB) *SqliteMealsService {
	return &SqliteMealsService{
		db: db,
	}
}

func (s *SqliteMealsService) ListMeals() ([]Meal, error) {
	meals := make([]Meal, 0)

	var rows *sql.Rows
	var err error

	rows, err = s.db.Query("SELECT * FROM meals")

	if err != nil {
		return meals, errors.New(fmt.Sprintf("unable to prepare SELECT meal statement. Error %s", err.Error()))
	}

	var mealId int
	var mealName string
	var mealCategory string

	for rows.Next() {
		err = rows.Scan(&mealId, &mealName, &mealCategory)
		if err != nil {
			return meals, errors.New(fmt.Sprintf("unable to query meals table. Error %s", err.Error()))
		}

		meals = append(meals, Meal{
			ID:       mealId,
			Name:     mealName,
			Category: mealCategory,
		})
	}

	return meals, nil
}

func (s *SqliteMealsService) CreateMeal(m Meal) (int, error) {

	stmt, err := s.db.Prepare("INSERT INTO meals(name, category) values(?,?)")
	if err != nil {
		return 0, errors.New(fmt.Sprintf("unable to prepare INSERT meal statement. Error %s", err.Error()))
	}

	res, err := stmt.Exec(m.Name, m.Category)
	if err != nil {
		return 0, errors.New(fmt.Sprintf("unable to execute create meals table statement. Error %s", err.Error()))
	}

	insertedId, err := res.LastInsertId()
	if err != nil {
		return 0, errors.New(fmt.Sprintf("unable to retrieve last inserted id. Error %s", err.Error()))
	}

	return int(insertedId), nil
}
