package models

type Plan struct {
	BaseModel
	Name      string
	RateLimit int
}
