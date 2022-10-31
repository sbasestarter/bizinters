package userinters

import (
	"context"
	"time"
)

//
//
//

type Authenticator interface {
	GetMethodName() (method string)
	Verify(ctx context.Context) (uid uint64, ok bool, err error)
}

//
//
//

type AuthForUserPolicy struct {
	UserID           uint64
	VerifiedMethods  []string
	SupportedMethods []string
}

type Policy interface {
	RequireAuthMethod(ctx context.Context, d *AuthForUserPolicy) (requiredOrMethods []string, err error)
}

type StatusController interface {
	IsTokenBanned(ctx context.Context, id uint64) (bool, error)
	BanToken(ctx context.Context, id uint64, expireAt int64) error
}

//
//
//

type LoginRequest struct {
	ContinueID        uint64
	Authenticators    []Authenticator
	TokenLiveDuration time.Duration
}

type LoginStatus int

const (
	LoginStatusSuccess LoginStatus = iota
	LoginStatusNeedMoreAuthenticator
)

type LoginResponse struct {
	Status            LoginStatus
	RequiredOrMethods []string
	UserID            uint64
	Token             string
	ContinueID        uint64
}

type UserCenter interface {
	Login(ctx context.Context, request *LoginRequest) (resp *LoginResponse, err error)
	Logout(ctx context.Context, token string) (err error)

	CheckToken(ctx context.Context, token string) (uid uint64, err error)
}
