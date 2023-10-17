package model

import (
	"time"
	"github.com/asaskevitch/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Bank struct {
	ID string `json:"id" valid:"uuid"`
	CreatedAt time.Time `json:"created_at" valid:"-"`
	UpdatedAt time.Time `json:"updated_at" valid:"-"`
}
