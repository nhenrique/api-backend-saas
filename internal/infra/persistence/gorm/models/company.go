package models

import "github.com/google/uuid"

type Company struct {
	BaseModel
	Name   string
	PlanID uuid.UUID
	Plan   Plan
}
