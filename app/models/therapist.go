package models

type Therapist struct {
	BaseModel
	IsApproved      int
	UserID          int64
	ProfileImageKey string
	ProfileImageUrl string
}
