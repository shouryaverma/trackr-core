package mock

// Repository ...
type Repository struct {
	ReturnObject interface{}
	IsError      bool
	ErrorMessage string
}
