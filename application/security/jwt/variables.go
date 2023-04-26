package jwt

const (
	Access  = "access"
	Refresh = "refresh"
)

// TokenTypeKeyName is a map that contains the key name of the JWT in config.json
var TokenTypeKeyName = map[string]string{
	Access:  "Secure.JWTAccessSecure",
	Refresh: "Secure.JWTRefreshSecure",
}

// TokenTypeExpTime is a map that contains the expiration time of the JWT
var TokenTypeExpTime = map[string]string{
	Access:  "Secure.JWTAccessTimeMinute",
	Refresh: "Secure.JWTRefreshTimeHour",
}
