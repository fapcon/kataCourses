package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"petstore/internal/models"
	"petstore/internal/service"
	"strconv"
	"strings"

	"github.com/go-pg/pg/v10"
)

type StoreCntrl struct {
	ss service.StoreServe
}

func NewStoreCntrl(db *pg.DB) *StoreCntrl {
	return &StoreCntrl{ss: service.NewStoreServe(db)}
}

// Create godoc
// @Summary Create a new order
// @Description Add a new order to the store
// @Tags order
// @Accept json
// @Produce json
// @Param order body models.Order true "Order to be created"
// @Success 201 {string} string "order created"
// @Failure 400 {string} string "request parsing error"
// @Failure 500 {string} string "order creation error"
// @Router /order [post]
func (sc StoreCntrl) Create(w http.ResponseWriter, r *http.Request) {
	order := models.Order{}
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, "request parsing error:"+err.Error(), http.StatusBadRequest)
		return
	}

	err = sc.ss.Create(context.Background(), &order)
	if err != nil {
		http.Error(w, "order creation error:"+err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, "order created")

}

// GetByID godoc
// @Summary Get an order by its ID
// @Description Retrieve an order from the store by its ID
// @Tags order
// @Produce json
// @Param id path int true "Order ID"
// @Success 200 {object} models.Order "Order found"
// @Failure 400 {string} string "id error"
// @Failure 417 {string} string "order getting error"
// @Router /order/{id} [get]
func (sc StoreCntrl) GetByID(w http.ResponseWriter, r *http.Request) {
	pathSpl := strings.Split(r.URL.Path, "/")
	id, err := strconv.Atoi(pathSpl[len(pathSpl)-1])

	if err != nil {
		http.Error(w, "id error", http.StatusBadRequest)
		return
	}

	order, err := sc.ss.GetByID(context.Background(), int64(id))
	if err != nil {
		http.Error(w, "order getting error:"+err.Error(), http.StatusExpectationFailed)
		return
	}

	fmt.Fprint(w, order)

}

// DeleteByID godoc
// @Summary Delete an order by its ID
// @Description Remove an order from the store by its ID
// @Tags order
// @Produce json
// @Param id path int true "Order ID"
// @Success 200 {string} string "order deleted"
// @Failure 400 {string} string "id error"
// @Failure 417 {string} string "order deletion error"
// @Router /order/{id} [delete]
func (sc StoreCntrl) DeleteByID(w http.ResponseWriter, r *http.Request) {
	pathSpl := strings.Split(r.URL.Path, "/")
	id, err := strconv.Atoi(pathSpl[len(pathSpl)-1])

	if err != nil {
		http.Error(w, "id error", http.StatusBadRequest)
		return
	}

	err = sc.ss.DeleteByID(context.Background(), int64(id))
	if err != nil {
		http.Error(w, "order deletion error:"+err.Error(), http.StatusExpectationFailed)
		return
	}

	fmt.Fprint(w, "order deleted")

}
