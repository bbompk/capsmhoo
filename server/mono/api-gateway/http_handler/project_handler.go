package http_handler

// "github.com/gin-gonic/gin"
// grpcClient "capsmhoo/mono/api-gateway/client_grpc"
// "capsmhoo/mono/api-gateway/model"

type ProjectHandler struct {
	// projectClientgRPC grpcClient.ProjectgRPCClient
}

type IProjectHandler interface {
	// TODO
}

func ProvideProjectHandler() *ProjectHandler {
	return &ProjectHandler{}
}
