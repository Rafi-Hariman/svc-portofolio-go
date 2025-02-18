package entity

type PortofolioKey struct {
	ForeignID uint64 `json:"-"`
}

type Portofolio struct {
	PortofolioKey
	Column string `json:"column,omitempty"`
}
