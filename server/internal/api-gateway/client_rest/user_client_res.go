package apigateway

type UserClient struct {
}

type UserClientRest interface {
}

func ProvideUserClientRest() *UserClient {
	return &UserClient{}
}
