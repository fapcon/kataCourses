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

type PetCntrl struct {
	ps service.PetServe
}

func NewPetCntrl(db *pg.DB) *PetCntrl {
	return &PetCntrl{ps: service.NewPetServe(db)}
}

// Create godoc
// @Summary Create a new pet
// @Description Add a new pet to the store
// @Tags pets
// @Accept  json
// @Produce  json
// @Param   pet  body  models.Pet  true  "Pet object that needs to be added"
// @Success 200  {string} string "pet created"
// @Failure 400  {string} string "request parsing error: <error message>"
// @Failure 500  {string} string "pet creation error: <error message>"
// @Router /pet [post]
func (pc PetCntrl) Create(w http.ResponseWriter, r *http.Request) {
	pet := models.Pet{}
	err := json.NewDecoder(r.Body).Decode(&pet)
	if err != nil {
		http.Error(w, "request parsing error:"+err.Error(), http.StatusBadRequest)
		return
	}

	err = pc.ps.Create(context.Background(), &pet)
	if err != nil {
		http.Error(w, "pet creation error:"+err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, "pet created")

}

// Update godoc
// @Summary Update an existing pet
// @Description Update an existing pet by its ID
// @Tags pets
// @Accept  json
// @Produce  json
// @Param   pet  body  models.Pet  true  "Pet object that needs to be updated"
// @Success 200  {string} string "pet updated"
// @Failure 400  {string} string "request parsing error: <error message>"
// @Failure 500  {string} string "pet update error: <error message>"
// @Router /pet [put]
func (pc PetCntrl) Update(w http.ResponseWriter, r *http.Request) {
	pet := models.Pet{}
	err := json.NewDecoder(r.Body).Decode(&pet)
	if err != nil {
		http.Error(w, "request parsing error:"+err.Error(), http.StatusBadRequest)
		return
	}

	err = pc.ps.Update(context.Background(), &pet)
	if err != nil {
		http.Error(w, "pet update error:"+err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, "pet updated")

}

// GetByStatus godoc
// @Summary Get pets by status
// @Description Get pets by their status
// @Tags pets
// @Produce  json
// @Param   status  path  string  true  "Status of the pet"
// @Success 200  {object}  models.Pet
// @Failure 400  {string} string "empty status error"
// @Failure 417  {string} string "pet getting error: <error message>"
// @Router /pet/findByStatus/{status} [get]
func (pc PetCntrl) GetByStatus(w http.ResponseWriter, r *http.Request) {
	pathSpl := strings.Split(r.URL.Path, "/")
	status := pathSpl[len(pathSpl)-1]

	if status == "" {
		http.Error(w, "empty status error", http.StatusBadRequest)
		return
	}

	pet, err := pc.ps.GetByStatus(context.Background(), status)
	if err != nil {
		http.Error(w, "pet getting error:"+err.Error(), http.StatusExpectationFailed)
		return
	}

	fmt.Fprint(w, pet)

}

// GetById godoc
// @Summary Get pet by ID
// @Description Get a single pet by its ID
// @Tags pets
// @Produce  json
// @Param   id  path  int  true  "ID of the pet to get"
// @Success 200  {object}  models.Pet
// @Failure 400  {string} string "id error"
// @Failure 417  {string} string "pet getting error: <error message>"
// @Router /pet/{id} [get]
func (pc PetCntrl) GetById(w http.ResponseWriter, r *http.Request) {
	pathSpl := strings.Split(r.URL.Path, "/")
	id, err := strconv.Atoi(pathSpl[len(pathSpl)-1])

	if err != nil {
		http.Error(w, "id error", http.StatusBadRequest)
		return
	}

	pet, err := pc.ps.GetByID(context.Background(), int64(id))
	if err != nil {
		http.Error(w, "pet getting error:"+err.Error(), http.StatusExpectationFailed)
		return
	}

	fmt.Fprint(w, pet)

}

// UpdateById godoc
// @Summary Update pet by ID
// @Description Update a pet by its ID
// @Tags pets
// @Accept  json
// @Produce  json
// @Param   id   path     int    true  "ID of the pet to update"
// @Param   pet  body  models.Pet  true  "Pet object that needs to be updated"
// @Success 200  {string} string "pet updated"
// @Failure 400  {string} string "id error"
// @Failure 417  {string} string "pet updating error: <error message>"
// @Router /pet/{id} [put]
func (pc PetCntrl) UpdateById(w http.ResponseWriter, r *http.Request) {
	pathSpl := strings.Split(r.URL.Path, "/")
	id, err := strconv.Atoi(pathSpl[len(pathSpl)-1])

	if err != nil {
		http.Error(w, "id error", http.StatusBadRequest)
		return
	}
	pet := models.Pet{}
	err = json.NewDecoder(r.Body).Decode(&pet)
	if err != nil {
		http.Error(w, "request parsing error:"+err.Error(), http.StatusBadRequest)
		return
	}

	err = pc.ps.UpdateByID(context.Background(), int64(id), pet.Name, pet.Status)
	if err != nil {
		http.Error(w, "pet updating error:"+err.Error(), http.StatusExpectationFailed)
		return
	}

	fmt.Fprint(w, "pet updated")

}

// DeleteById godoc
// @Summary Delete pet by ID
// @Description Delete a pet by its ID
// @Tags pets
// @Produce  json
// @Param   id  path  int  true  "ID of the pet to delete"
// @Success 200  {string} string "pet deleted"
// @Failure 400  {string} string "id error"
// @Failure 417  {string} string "pet deletion error: <error message>"
// @Router /pet/{id} [delete]
func (pc PetCntrl) DeleteById(w http.ResponseWriter, r *http.Request) {
	pathSpl := strings.Split(r.URL.Path, "/")
	id, err := strconv.Atoi(pathSpl[len(pathSpl)-1])

	if err != nil {
		http.Error(w, "id error", http.StatusBadRequest)
		return
	}

	err = pc.ps.DeleteByID(context.Background(), int64(id))
	if err != nil {
		http.Error(w, "pet deletion error:"+err.Error(), http.StatusExpectationFailed)
		return
	}

	fmt.Fprint(w, "pet deleted")

}
