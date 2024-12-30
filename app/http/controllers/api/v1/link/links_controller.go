package v1

import (
	v1 "gohub/app/http/controllers/api/v1"
	"gohub/app/models/link"
	"gohub/pkg/response"

	"github.com/gin-gonic/gin"
)

type LinksController struct {
	v1.BaseAPIController
}

func (ctrl *LinksController) Index(c *gin.Context) {
	links := link.All()
	response.Data(c, links)
}
