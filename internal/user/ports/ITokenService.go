package ports

type ITokenService interface {
	GenerateAccessToken(userId, role uint) (string, error)
	GenerateRefreshToken(userId uint) (string, error)
}