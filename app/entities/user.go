package entity

import "time"

type User struct {
	BaseEntity
	IsAdmin      int       `json:"is_admin"`
	IsTherapist  int       `json:"is_therapist"`
	IsWithdrawal int       `json:"is_withdrawal"`
	OS           string    `json:"os"`
	TypeID       int       `json:"type_id"`
	LoginType    string    `json:"login_type"`
	GoogleID     int64     `json:"google_id"`
	AppleID      int64     `json:"apple_id"`
	Email        string    `json:"email"`
	LastName     string    `json:"last_name"`
	FirstName    string    `json:"first_name"`
	Gender       int       `json:"gender"`
	Birthday     time.Time `json:"birthday"`
	IPAddress    string    `json:"ip_address"`
	UserAgent    string    `json:"user_agent"`
}

func (u *User) IsLoginGoogle() bool {
	isLoginGoogle := false
	if u.LoginType == "google" {
		isLoginGoogle = true
	}
	return isLoginGoogle
}
