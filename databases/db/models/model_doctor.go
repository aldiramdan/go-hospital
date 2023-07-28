package models

import "time"

type Doctor struct {
	DoctorId        string    `json:"id" valid:"uuid"`
	Name            string    `json:"name" valid:"required~name is blank"`
	Specialization  string    `json:"specialization" valid:"required~specialization is blank"`
	ConsultationFee float64   `json:"consultation_fee" valid:"required~consultation_fee is blank"`
	CreatedAt       time.Time `json:"created_at" valid:"-"`
	UpdatedAt       time.Time `json:"updated_at" valid:"-"`
}

type Doctors []Doctor
