package main

import (
	"net/http"

	"github.com/chalu/hello-crm-backend/controller"
	"github.com/chalu/hello-crm-backend/service"
	"github.com/gin-gonic/gin"
)

var (
	custService    service.CustomerService       = service.New()
	custController controller.CustomerController = controller.New(custService)
)

func main() {
	server := gin.Default()

	// setup default 404 handler
	server.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    "PAGE_NOT_FOUND",
			"message": "Page not found",
		})
	})

	server.GET("/customers", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, custController.FindAll())
	})

	server.GET("/customers/:id", func(ctx *gin.Context) {
		cust := custController.Find(ctx.Param("id"))
		status := http.StatusOK
		if cust.ID == "" {
			status = http.StatusBadRequest
			ctx.Writer.WriteHeader(status)
			return
		}
		ctx.JSON(status, cust)
	})

	server.POST("/customers", func(ctx *gin.Context) {
		custId := custController.AddOrUpdate(ctx)
		ctx.JSON(http.StatusCreated, custId)
	})

	server.PUT("/customers/:id", func(ctx *gin.Context) {
		custId := custController.AddOrUpdate(ctx)
		status := http.StatusOK
		if custId == "" {
			status = http.StatusBadRequest
		}
		ctx.JSON(status, custId)
	})

	server.DELETE("/customers/:id", func(ctx *gin.Context) {
		custId := custController.Remove(ctx.Param("id"))
		status := http.StatusOK
		if custId == "" {
			status = http.StatusBadRequest
		}
		ctx.JSON(status, custId)
	})

	server.Run(":8080")
}
