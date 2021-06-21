package model

import "time"

//App信息表
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

//配置表
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

//App配置表
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

//已启用配置
type AppEnabledConfig  struct {
	AppId      string
	AppKey     string
	AppName    string
	Department string
	Owner      string
	Cluster   string
	Profile   string
	NotifyUrl string
	Cid       string
	Config    string
	CreatedAt  time.Time
	CreatedBy  string
}