package response

type LoginResponse struct {
	JwtToken     string `json:"jwtToken"`
	RefreshToken string `json:"refreshToken"`
}
