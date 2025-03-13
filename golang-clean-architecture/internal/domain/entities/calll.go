package entities

import (
	"time"

	"github.com/google/uuid"
)

type Call struct {
	Entity
	CallingUserId   uuid.UUID
	DateCallStarted time.Time
	DateDeleted     *time.Time
	User
}
