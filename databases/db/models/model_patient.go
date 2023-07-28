package models

import "time"

type Patient struct {
	PatientId string    `json:"id" valid:"uuid"`
	Name      string    `json:"name" valid:"required~name is blank"`
	Age       int       `json:"age" valid:"required~age is blank"`
	Address   string    `json:"address" valid:"required~address is blank"`
	CreatedAt time.Time `json:"created_at" valid:"-"`
	UpdatedAt time.Time `json:"updated_at" valid:"-"`
}

type Patients []Patient
