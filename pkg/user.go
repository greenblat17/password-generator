package util

type UserInput struct {
	Service       string
	SecretKey     string
	SecretNumbers string
	Count         int
}

func NewUserInput() *UserInput {
	return &UserInput{
		Service:       "",
		SecretKey:     "",
		SecretNumbers: "",
		Count:         0,
	}
}
