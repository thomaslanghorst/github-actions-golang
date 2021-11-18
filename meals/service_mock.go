package meals

import "github.com/stretchr/testify/mock"

type MockMealsServiceInterface struct {
	mock.Mock
}

func (m *MockMealsServiceInterface) ListMeals() ([]Meal, error) {
	args := m.Called()

	var v0 []Meal
	if args.Get(0) != nil {
		v0 = args.Get(0).([]Meal)
	}
	return v0, args.Error(1)
}

func (m *MockMealsServiceInterface) CreateMeal(meal Meal) (int, error) {
	args := m.Called(meal)

	return args.Int(0), args.Error(1)
}
