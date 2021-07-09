package handlerfactory

import (
	"net/http"

	"github.com/HarvestStars/QuantsGateway/dtos"
	"github.com/gin-gonic/gin"
)

type BaseFactory interface {
	Create() BaseHandler
}

type BaseHandler interface {
	Get()
	Post()
}

type ImportIns struct {
	AimSite string `json:"aimsite"`
	Method  string `json:"method"`
	Url     string `json:"url"`
	Body    string `json:"body"`
}

func FactoryImport(c *gin.Context) {
	var importReq ImportIns
	err := c.Bind(&importReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": dtos.CANNOT_PARSE_POST_BODY, "msg": "Sorry", "data": err.Error()})
		return
	}

	if factory, ok := FactoryInstanceMap[importReq.AimSite]; ok {
		handler := factory.Create()
		handler.Post()
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"code": dtos.AIM_SITE_NOT_EXIST, "msg": "Sorry", "data": err.Error()})
}
