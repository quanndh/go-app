package services

import "github.com/quanndh/go-app/public/resources"

type IJwtService interface {
	Generate(payload *resources.UserResource) (string, error)
	Verify(tokenString string) (*UserClaims, error)
}
