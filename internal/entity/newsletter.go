package entity

import "time"

//? mempresentasikan table newsletter (miri seperti model dilaravel)
type Newsletter struct {
	Id        string
	Fullname  string
	Email     string
	CreatedAt time.Time
	CreatedBy string
	UpdatedAt *time.Time
	UpdatedBy *string
	DeletedAt *time.Time
	DeletedBy *string
	IsDeleted bool
}
