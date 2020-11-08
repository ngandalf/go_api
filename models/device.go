package models

import (
	"github.com/ngandalf/go_api/config"
	"log"
	"time"
)

type Device struct {
	Id           int       `json:"id"`
	ServerId     int       `json:"server_id"`
	Infra        int       `json:"infra"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type Devices []Device

func NewDevice(device *Device) {
	if device == nil {
		log.Fatal(c)
	}
	c.CreatedAt = time.Now()
	c.UpdatedAt = time.Now()

	err := config.Db().QueryRow("INSERT INTO devices (server_id, infra, created_at, updated_at) VALUES ($1,$2,$3,$4) RETURNING id;", device.serveur_id, device.infra, device.CreatedAt, device.UpdatedAt).Scan(&device.Id)

	if err != nil {
		log.Fatal(err)
	}
}

func FindDeviceById(id int) *Car {
	var device Device

	row := config.Db().QueryRow("SELECT * FROM devices WHERE id = $1;", id)
	err := row.Scan(&device.Id, &device.Manufacturer, &device.Design, &device.Style, &device.Doors, &device.CreatedAt, &device.UpdatedAt)

	if err != nil {
		log.Fatal(err)
	}

	return &device
}

func AllCars() *Devices {
	var devices Devices

	rows, err := config.Db().Query("SELECT * FROM devices")

	if err != nil {
		log.Fatal(err)
	}

	// Close rows after all readed
	defer rows.Close()

	for rows.Next() {
		var device Device

		err := rows.Scan(&device.Id, &device.Manufacturer, &device.Design, &device.Style, &device.Doors, &device.CreatedAt, &device.UpdatedAt)

		if err != nil {
			log.Fatal(err)
		}

		cars = append(devices, device)
	}

	return &devices
}

func UpdateDevice(device *Device) {
	device.UpdatedAt = time.Now()

	stmt, err := config.Db().Prepare("UPDATE devices SET manufacturer=$1, design=$2, style=$3, doors=$4, updated_at=$5 WHERE id=$6;")

	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(car.Manufacturer, car.Design, car.Style, car.Doors, car.UpdatedAt, car.Id)

	if err != nil {
		log.Fatal(err)
	}
}

func DeleteDeviceById(id int) error {
	stmt, err := config.Db().Prepare("DELETE FROM devices WHERE id=$1;")

	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(id)

	return err
}
