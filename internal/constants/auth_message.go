package constants

const (
	RegistrationSuccess = "User registered successfully"
	LoginSuccess        = "User logged in successfully"
	RefreshSuccess      = "Token refreshed successfully"
	ProfileSuccess      = "User profile retrieved successfully"

	RegistrationError  = "Failed to register user"
	LoginError         = "Failed to login user"
	RefreshError       = "Failed to refresh token"
	ProfileError       = "Failed to retrieve user profile"
	InvalidCredentials = "Invalid credentials"
	InvalidToken       = "Invalid token"
	InvalidUserClaims  = "Invalid user claims"

	MissingAuthHeader    = "Missing authorization header"
	InvalidAuthHeader    = "Invalid authorization header"
	InsufficientPerms    = "Insufficient permissions"
	NotAuthenticated     = "User not authenticated"
)
