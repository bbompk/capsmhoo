package apigateway

type StudentClient struct {
}

type StudentClientRest interface {
}

func ProvideStudentClientRest() *StudentClient {
	return &StudentClient{}
}
