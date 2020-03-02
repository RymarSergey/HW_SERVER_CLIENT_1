package serverGin

import (
	"awesomeProject1/server/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var hb []model.Human

func isExist(firstName string, hb []model.Human) (int, bool) {
	for i, value := range hb {
		if firstName != value.FirstName {
			continue
		} else {
			return i, true
		}
	}
	return 0, false
}
func StartServer() {
	router := gin.Default()
	router.GET("/", getting)
	router.POST("/", posting)
	router.PUT("/", putting)
	router.DELETE("/", deleting)
	router.Run()
}
func posting(c *gin.Context) {
	humanTemp := model.Human{}
	err := c.Bind(&humanTemp)
	if err != nil {
		c.String(http.StatusBadRequest, "Bind Error")
		return
	}
	if len(hb) != 0 {
		_, ok := isExist(humanTemp.FirstName, hb)
		if ok {
			c.String(http.StatusAccepted, " Human exist")
			return
		}
	}
	hb = append(hb, humanTemp)
	c.String(http.StatusOK, " Human saved")
}

func getting(c *gin.Context) {
	if len(hb) == 0 {
		return
	}
	i, ok := isExist(c.Query("first_name"), hb)
	if ok {
		c.JSON(http.StatusOK, &hb[i])
	} else {
		c.String(http.StatusNoContent, "Nor exist")
	}
}
func putting(c *gin.Context) {
	if len(hb) == 0 {
		return
	}
	humanTemp := model.Human{}
	err := c.Bind(&humanTemp)
	if err != nil {
		c.String(http.StatusBadRequest, "Bind Error")
		return
	}
	i, ok := isExist(humanTemp.FirstName, hb)
	if ok {
		hb[i] = humanTemp
		fmt.Println(hb[i])
		c.String(http.StatusOK, "human updated! ")
	} else {
		c.String(http.StatusNoContent, "Nor exist")
	}
}
func deleting(c *gin.Context) {
	if len(hb) == 0 {
		return
	}
	i, ok := isExist(c.Query("first_name"), hb)
	if ok {
		hb = append(hb[:i], hb[i+1:]...)
		c.String(http.StatusOK, "Human was deleted ")
	} else {
		c.String(http.StatusNoContent, "Nor exist")
	}
}
