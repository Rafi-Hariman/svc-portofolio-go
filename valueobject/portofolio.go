package valueobject

import "svc-portofolio-golang/entity"

type Portofolio struct {
	entity.Portofolio
	entity.StandardKey
	entity.Pagination
	entity.Time
}

type PortofolioPayloadInsert struct {
	Data []Portofolio `json:"data" binding:"required"`
	User string
}

type PortofolioPayloadUpdate struct {
	Data []PortofolioDataUpdate `json:"data" binding:"required"`
	User string
}

type PortofolioDataUpdate struct {
	Param Portofolio `json:"param" binding:"required"`
	Body  Portofolio `json:"body" binding:"required"`
}

type PortofolioPayloadDelete struct {
	Param []Portofolio `json:"param" binding:"required"`
}
