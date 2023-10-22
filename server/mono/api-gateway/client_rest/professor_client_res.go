package apigateway

type ProfessorClient struct {
}

type ProfessorClientRest interface {
}

func ProvideProfessorClientRest() *ProfessorClient {
	return &ProfessorClient{}
}
