package ports

type IPasswordService interface {
	Hash (password string) (string, error)
	Compare (hash, password string) bool
}