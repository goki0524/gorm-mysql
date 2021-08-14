package entity

type Therapist struct {
	BaseEntity
	IsApproved      int
	UserID          int64
	ProfileImageKey string
	ProfileImageUrl string
}
