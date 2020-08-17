package service

// Authenticate generates a JWT token based on given data
func Authenticate(email string, name string) (string, error) {
	token, err := GenerateJWT(email, name)

	return token, err
}
