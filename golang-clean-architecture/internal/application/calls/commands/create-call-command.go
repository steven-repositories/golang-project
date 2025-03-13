package commands

import (
	"time"

	"golang-clean-architecture/internal/shared/interfaces"

	"github.com/google/uuid"
)

type CreateCallCommand struct {
	CallingUserId   uuid.UUID
	DateCallStarted time.Time
}

type CreateCallResult struct {
	Id uuid.UUID
}

type CreateCallCommandHandler struct {
	// dependency injection here (e.g., applicationdbcontext in .net)
}

func (handler *CreateCallCommandHandler) Handle(request CreateCallCommand) CreateCallResult {
	return CreateCallResult{
		Id: uuid.New(),
	}
}

var _ interfaces.RequestHandlerInterface[CreateCallCommand, CreateCallResult] = &CreateCallCommandHandler{}
