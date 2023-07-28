package models

import "time"

type Disease struct {
	DiseaseId string    `json:"id" valid:"uuid"`
	Name      string    `json:"name" valid:"required~name is blank"`
	Medicine  Medicine  `json:"medicine" valid:"-"`
	CreatedAt time.Time `json:"created_at" valid:"-"`
	UpdatedAt time.Time `json:"updated_at" valid:"-"`
}

type Diseases []Disease
