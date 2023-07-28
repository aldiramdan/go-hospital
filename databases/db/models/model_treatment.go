package models

import "time"

type Treatment struct {
	TreatmentId string    `json:"id" valid:"uuid"`
	Patient     Patient   `json:"patient" valid:"-"`
	Disease     []Disease `json:"disease" valid:"-"`
	Doctor      Doctor    `json:"doctor" valid:"-"`
	ServiceFee  float64   `json:"service_fee" valid:"-"`
	CreatedAt   time.Time `json:"created_at" valid:"-"`
	UpdatedAt   time.Time `json:"updated_at" valid:"-"`
}

type Treatments []Treatment
