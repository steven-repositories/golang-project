package interfaces

type RequestHandlerInterface[TRequest any, TResponse any] interface {
	Handle(request TRequest) TResponse
}
