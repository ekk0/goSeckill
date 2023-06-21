package server

import (
	"github.com/gin-gonic/gin"
	"goSeckill/models"
	"net/http"
)

//购买商品
func BuyGood(c *gin.Context)  {

	models.BuyGoods()

	c.JSON(http.StatusOK,gin.H{
		"code":200,
		"res":"",
	})
}


