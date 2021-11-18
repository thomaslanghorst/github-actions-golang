package meals

import (
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type MealHandler struct {
	service MealsServiceInterface
}

func NewMealHandler(s MealsServiceInterface) *MealHandler {
	return &MealHandler{
		service: s,
	}
}

func (h *MealHandler) ListMeals(w http.ResponseWriter, r *http.Request) {
	meals, err := h.service.ListMeals()
	if err != nil {
		log.Warnf(err.Error())
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	err = json.NewEncoder(w).Encode(meals)
	if err != nil {
		log.Warnf("error encoding meals. Error: %s", err.Error())
		sendError(w, http.StatusInternalServerError, "error encoding meals")
		return
	}
}

func (h *MealHandler) CreateMeal(w http.ResponseWriter, r *http.Request) {
	var m Meal

	err := json.NewDecoder(r.Body).Decode(&m)
	defer r.Body.Close()

	if err != nil {
		log.Warnf(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error while decoding meal"))
		return
	}

	id, err := h.service.CreateMeal(m)
	if err != nil {
		log.Warnf(err.Error())
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte(fmt.Sprintf("{\"id\":\"%d\"}", id)))
	return
}

func sendError(w http.ResponseWriter, statusCode int, msg string) {

	body := []byte(fmt.Sprintf("{\"error_message\":\"%s\"}", msg))

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(body)
}
