package model

import "time"

type AppConfig struct {
	Cid       string
	Config    string
	Status    int
	Memo      string
	CreatedAt time.Time
	CreatedBy string
	UpdatedAt time.Time
	UpdatedBy string
}
type AppInf struct {
	AppId      string
	AppKey     string
	AppName    string
	Department string
	Owner      string
	Status     int
	Memo       string
	CreatedAt  time.Time
	CreatedBy  string
	UpdatedAt  time.Time
	UpdatedBy  string
}
type AppProfile struct {
	AppId     string
	Cluster   string
	Profile   string
	NotifyUrl string
	Cid       string
	Status    int
	Memo      string
	CreatedAt time.Time
	CreatedBy string
	UpdatedAt time.Time
	UpdatedBy string
}
