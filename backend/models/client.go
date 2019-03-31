package models

import "time"

//Client ...
type Client struct {
	ID          int32
	Type        string
	DateStart   time.Time
	DateEnd     time.Time
	Sum         float32
	IsAccedents bool
	CompanyID   int32
	Name        string
	Contacts    string
	Areas       int32
	AreasIot    int32
	Squares     int32
}

//Client ...
type NiceClient struct {
	ID          int32
	Type        string
	DateStart   string
	DateEnd     string
	Sum         float32
	IsAccedents bool
	CompanyID   int32
	Name        string
	Contacts    string
	Areas       int32
	AreasIot    int32
	Squares     int32
}
