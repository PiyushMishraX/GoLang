package auth

// private
func extractSession() string {
	return "LoggedIn"
}

// public method ( exposed)
func GetSession() string {
	// return "LoggedIn"
	return extractSession()
}