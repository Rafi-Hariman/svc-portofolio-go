package valueobject

import "svc-portofolio-golang/entity"

type Auth struct {
	entity.Auth
	entity.StandardKey
	entity.Pagination
	entity.Time
}

type AuthPayloadInsert struct {
	Data []Auth `json:"data" binding:"required"`
	User string
}

type AuthPayloadUpdate struct {
	Data []AuthDataUpdate `json:"data" binding:"required"`
	User string
}

type AuthDataUpdate struct {
	Param Auth `json:"param" binding:"required"`
	Body  Auth `json:"body" binding:"required"`
}

type AuthPayloadDelete struct {
	Param []Auth `json:"param" binding:"required"`
}
