package Repository

import (
	"time"
)

type Pull struct {
	ID int
	RemoteId int
	Url string
	Title string
	Author string
	CreatedAt time.Time
}
