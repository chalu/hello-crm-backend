package controller

import (
	"github.com/chalu/hello-crm-backend/entity"
	"github.com/chalu/hello-crm-backend/service"
	"github.com/gin-gonic/gin"
)

type CustomerController interface {
	Find(string) entity.Customer
	FindAll() []entity.Customer
	AddOrUpdate(ctx *gin.Context) string
	Remove(id string) string
}

type Controller struct {
	service service.CustomerService
}

func New(s service.CustomerService) Controller {
	return Controller{
		service: s,
	}
}

func (c Controller) Find(id string) entity.Customer {
	cust := c.service.Find(id)
	return cust
}

func (c Controller) FindAll() []entity.Customer {
	return c.service.FindAll()
}

func (c Controller) AddOrUpdate(ctx *gin.Context) string {
	var customer entity.Customer
	ctx.BindJSON(&customer)

	custId := ""
	switch m := ctx.Request.Method; m {
	case "POST":
		custId = c.service.Add(customer)
	case "PUT":
		customer.ID = ctx.Param("id")
		custId = c.service.Update(customer)
	}

	return custId
}

func (c Controller) Remove(id string) string {
	return c.service.Delete(id)
}
