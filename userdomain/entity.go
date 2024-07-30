package userdomain

import "time"

type Entity struct {
	ID             string
	Name           string
	Email          string
	DateOfRegister time.Time
	Role           string
}
