package request

type LoginRefreshTokenRequest struct {
	RefreshToken string `validate:"required" json:"refreshToken"`
}
