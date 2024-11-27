// file: service.go
package main

//go:generate mockgen -source=service.go -destination=mock_service.go -package=main

// interface example for mocking
// Service defines some business logic interface.
type Service interface {
	DoSomething(param string) error
}
