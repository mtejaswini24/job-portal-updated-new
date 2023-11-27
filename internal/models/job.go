package models

import (
	"gorm.io/gorm"
)

type NewCompany struct {
	CompanyName string `json:"companyname" validate:"required" `
	Location    string `json:"location" validate:"required"`
}
type Company struct {
	gorm.Model
	CompanyName string `validate:"required,unique" gorm:"unique;not null"`
	Location    string `json:"location"`
}

type NewJob struct {
	JobRole       string `json:"role" validate:"required"`
	Description   string `json:"description" validate:"required"`
	Min_Np        uint   `json:"minimum_notice_period" validate:"required"`
	Max_Np        uint   `json:"maximum_notice_period" validate:"required"`
	Budget        uint64 `json:"budget" validate:"required"`
	JobLocation   []uint `json:"locations" validate:"required"`
	Technology    []uint `json:"technology" validate:"required"`
	WorkMode      []uint `json:"workmode" validate:"required"`
	MinExp        uint64 `json:"minimum_experience" validate:"required"`
	MaxExp        uint64 `json:"maximum_experience" validate:"required"`
	Qualification []uint `json:"qualification" validate:"required"`
	Shift         []uint `json:"shift" validate:"required"`
	JobType       []uint `json:"jobtype" validate:"required"`
}
type Job struct {
	gorm.Model    `json:"-"`
	Company       Company         `json:"-" gorm:"ForeignKey:cid"`
	Cid           uint            `json:"cid"`
	JobRole       string          `json:"Role"`
	Description   string          `json:"description"`
	Min_Np        uint            `json:"minimum_notice_period"`
	Max_Np        uint            `json:"maximum_notice_period"`
	Budget        uint64          `json:"budget"`
	JobLocation   []Location      `gorm:"many2many:job_location;"`
	Technology    []Technology    `gorm:"many2many:job_technology;"`
	WorkMode      []WorkMode      `gorm:"many2many:job_workmode;"`
	MinExp        uint64          `json:"minimum_experience"`
	MaxExp        uint64          `json:"maximum_experience"`
	Qualification []Qualification `gorm:"many2many:job_qualification;"`
	Shift         []Shift         `gorm:"many2many:job_shift;"`
	JobType       []JobType       `gorm:"many2many:job_jobtype;"`
}
type Location struct {
	Id   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique; column:name"`
}
type Technology struct {
	Id   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique; column:name"`
}
type WorkMode struct {
	Id   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique; column:name"`
}
type Qualification struct {
	Id   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique; column:name"`
}
type Shift struct {
	Id   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique; column:name"`
}
type JobType struct {
	Id   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique; column:name"`
}
type RequestJob struct {
	Id            uint64 `json:"id" validate:"required"`
	JobRole       string `json:"role" validate:"required"`
	Description   string `json:"description"`
	NoticePeriod  uint64 `json:"notice_period" validate:"required"`
	Budget        uint64 `json:"budget" validate:"required"`
	JobLocation   []uint `json:"locations" validate:"required"`
	Technology    []uint `json:"technology" validate:"required"`
	WorkMode      []uint `json:"workmode" validate:"required"`
	Exp           uint64 `json:"experience" validate:"required"`
	Qualification []uint `json:"qualification" validate:"required"`
	Shift         []uint `json:"shift" validate:"required"`
	JobType       []uint `json:"jobtype" validate:"required"`
}
