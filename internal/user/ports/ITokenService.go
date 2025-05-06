package ports

type ITokenService interface {
	GenerateAccessToken(userId, role uint) (string, error)
	GenerateRefreshToken(userId uint) (string, error)
	DecodeRefreshToken(token string) (uint, error)
	DecodeAccessToken(token string) (uint, bool, error)
}