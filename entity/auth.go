package entity

type AuthKey struct {
	ForeignID uint64 `json:"-"`
}

type Auth struct {
	AuthKey
	Column string `json:"column,omitempty"`
}

type AuthLogin struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
