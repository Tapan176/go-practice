package constants

type ContextKey string

const (
	DBContextKey       ContextKey = "dbClient"
	UserIDContextKey   ContextKey = "userId"
	UserRoleContextKey ContextKey = "userRole"
	ValidatedBodyKey   ContextKey = "validatedBody"
)

const (
	DefaultPageSize = 10
	DefaultPage     = 1
	MaxPageSize     = 100
)

const (
	RoleAdmin  = "admin"
	RoleUser   = "user"
	RoleAuthor = "author"
)
