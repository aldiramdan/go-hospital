package models

import "time"

type Medicine struct {
	MedicineId string    `json:"id" valid:"uuid"`
	Name       string    `json:"name" valid:"required~name is blank"`
	Price      float64   `json:"price" valid:"required~price is blank"`
	CreatedAt  time.Time `json:"created_at" valid:"-"`
	UpdatedAt  time.Time `json:"updated_at" valid:"-"`
}

type Medicines []Medicine
