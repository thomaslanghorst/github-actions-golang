package meals

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListMealsHandler(t *testing.T) {
	// arrange
	a := assert.New(t)
	mockService := MockMealsServiceInterface{}

	r := httptest.NewRequest("GET", "/meals", nil)
	w := httptest.NewRecorder()

	mockService.
		On("ListMeals").
		Return([]Meal{}, nil)

	h := NewMealHandler(&mockService)

	// act
	h.ListMeals(w, r)

	// assert
	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)

	a.Nil(err)
	a.Equal("[]\n", string(data))
}
