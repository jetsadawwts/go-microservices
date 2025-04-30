package middlewareRepository

type (
	MiddlewareRepositoryHandler interface {}

	middlewareRepository struct {}
)


func NewMiddlewareRepository() MiddlewareRepositoryHandler {
	return &middlewareRepository{}
}