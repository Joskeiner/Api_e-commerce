package helper

import "errors"

// constants for massage in HTTP response
const (
	FORBIDDEN         = "forbidden"
	UNAUTHORIZED      = "Unauthorized"
	SUCCEEDPOSTDATA   = "Succeed to POST data"
	SUCCEEDGETDATA    = "Succeed to GET data"
	SUCCEEDPUTDATA    = "Succeed to PUT data"
	SUCCEEDDELETEDATA = "Succeed to DELETE data"
	FAILEDPOSTDATA    = "Faild to POST data"
	FAILEDGETDATA     = "Faild to GET  data"
	FAILEDPUTDATA     = "Faild to PUT data"
	FAILEDDELETEDATA  = "Faild to DELETE data"
)

// ERROR Massage
var (
	ErrUnsupportedDriver          = errors.New("unsupported database driver")
	ErrDataAlreadyExist           = errors.New("data already exists")
	ErrUnsupportedTokenType       = errors.New("unsupported token type")
	ErrExpiredToken               = errors.New("token has expired")
	ErrInvalidToken               = errors.New("token is invalid")
	ErrDataNotFound               = errors.New("data not found")
	ErrUnauthorized               = errors.New("unauthorized")
	ErrForbidden                  = errors.New("forbidden")
	ErrInsufficientPermission     = errors.New("only admin can access this resource")
	ErrEmptyAuthorizationHeader   = errors.New("empty authorization header")
	ErrInvalidAuthorizationHeader = errors.New("invalid authorization header")
	ErrInvalidAuthorizationType   = errors.New("invalid authorization type")
)
