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

/// new auth struct for login

type AuthLogin struct {
	entity.AuthLogin
	entity.StandardKey
	entity.Pagination
	entity.Time
}

type AuthLoginPayloadInsert struct {
	Data []AuthLogin `json:"data" binding:"required"`
	User string
}

type AuthLoginPayloadUpdate struct {
	Data []AuthLoginDataUpdate `json:"data" binding:"required"`
	User string
}

type AuthLoginDataUpdate struct {
	Param AuthLogin `json:"param" binding:"required"`
	Body  AuthLogin `json:"body" binding:"required"`
}

type AuthLoginPayloadDelete struct {
	Param []AuthLogin `json:"param" binding:"required"`
}
