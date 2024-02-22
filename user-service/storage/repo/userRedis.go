package repo

import (
	pbu "registration/user-service/genproto/user"
)

type UserStorageRedisI interface {
	Sign(user *pbu.UserDetail) (*pbu.ResponseMessage, error)
	Verification(req *pbu.VerificationUserRequest) (*pbu.User, error)
}
