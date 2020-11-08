package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/synbioz/go_api/models"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func DevicesIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(models.AllDevices())
}

func DevicesCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	var device models.Device

	err = json.Unmarshal(body, &device)

	if err != nil {
		log.Fatal(err)
	}

	models.NewDevice(&device)

	json.NewEncoder(w).Encode(device)
}

func DevicesShow(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		log.Fatal(err)
	}

	device := models.FindDeviceById(id)

	json.NewEncoder(w).Encode(device)
}
func DevicesUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	device := models.FindDeviceById(id)

	err = json.Unmarshal(body, &device)

	models.UpdateCar(device)

	json.NewEncoder(w).Encode(device)
}
func DevicesDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)

	// strconv.Atoi is shorthand for ParseInt
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		log.Fatal(err)
	}

	err = models.DeleteDeviceById(id)
}
