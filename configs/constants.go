package configs

const (
	HEADER_AUTH               = "Authorization"
	ROOT_KEY_REDIS_AUTH       = "auth"
	MSG_INVALID_BODY          = "Invalid request body!"
	MSG_USER_NOT_FOUND        = "User not found!"
	MSG_INTERNAL_ERROR        = "Internal server error!"
	MSG_START_SERVER          = "Server started at http://localhost:"
	ERR_LOAD_ENV              = "Error at load env:"
	ERR_START_DB              = "Error at start database:"
	ERR_START_CACHE           = "Error at start cache:"
	ERR_START_SERVER          = "Error at start server:"
	ROOT_KEY_REDIS_RATE_LIMIT = "rate_limit"

	ROUTE_LOGIN   = "/login"
	ROUTE_ACCOUNT = "/account"
	ROUTE_HOME    = "/"
)

var route_limits = map[string]int{
	ROUTE_HOME:    10,
	ROUTE_ACCOUNT: 50,
	ROUTE_LOGIN:   10,
}

func ROUTE_LIMITS() map[string]int {
	return route_limits
}

var non_auth_routes = []string{
	ROUTE_ACCOUNT,
	ROUTE_LOGIN,
}

func NON_AUTH_ROUTES() []string {
	return non_auth_routes
}
