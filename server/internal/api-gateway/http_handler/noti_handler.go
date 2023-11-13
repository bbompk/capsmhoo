package http_handler

import (
	grpcClient "capsmhoo/internal/api-gateway/client_grpc"

	"github.com/gin-gonic/gin"
)

type NotiHandler struct {
	notiClientgRPC grpcClient.NotigRPCClient
}

type INotiHandler interface {
	GetAllNotiByUserId(c *gin.Context)
	ReadNoti(c *gin.Context)
}

func (h *NotiHandler) GetAllNotiByUserId(c *gin.Context) {
	userID := c.Param("id")
	notis, err := h.notiClientgRPC.GetAllNotiByUserId(c, userID)
	if err != nil {
		c.JSON(200, gin.H{
			"code":  "500",
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": "200",
		"data": notis,
	})
}

func (h *NotiHandler) ReadNoti(c *gin.Context) {
	userID := c.Param("id")
	err := h.notiClientgRPC.ReadNoti(c, userID)
	if err != nil {
		c.JSON(200, gin.H{
			"code":  "500",
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    "200",
		"message": "Read Notification Success",
	})
}

func ProvideNotiHandler(notiClientgRPC grpcClient.NotigRPCClient) *NotiHandler {
	return &NotiHandler{notiClientgRPC: notiClientgRPC}
}
