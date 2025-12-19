// Package repository is the repository package for the Sovereign service.
package repository

type Health interface {
	Readiness() error
}
