package controllers

import "github.com/gin-gonic/gin"

//Health .
func Health(c *gin.Context) {
	c.String(200, "ok")
}
