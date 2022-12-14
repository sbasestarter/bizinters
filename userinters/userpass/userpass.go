package userpass

import "context"

type UserPasswordVerifier interface {
	Verify(ctx context.Context, user, password string) (userID uint64, ok bool, err error)
}

type User struct {
	ID       uint64
	UserName string
	Password string
	CreateAt int64
	ExData   map[string]interface{}
}

type UserPasswordModel interface {
	AddUser(ctx context.Context, userName, password string) (user *User, err error)
	DeleteUser(ctx context.Context, userID uint64) error

	GetUser(ctx context.Context, userID uint64) (user *User, err error)
	GetUserByUserName(ctx context.Context, userName string) (user *User, err error)

	ListUsers(ctx context.Context) (users []*User, err error)

	UpdateUserExData(ctx context.Context, userID uint64, key string, val interface{}) (err error)
	UpdateUserAllExData(ctx context.Context, userID uint64, exData map[string]interface{}) (err error)
}
