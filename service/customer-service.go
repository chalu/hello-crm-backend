package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/chalu/hello-crm-backend/entity"
	"github.com/google/uuid"
)

type CustomerService interface {
	Find(string) entity.Customer
	FindAll() []entity.Customer
	Add(entity.Customer) string
	Update(entity.Customer) string
	Delete(string) string
}

type CustomerServiceImpl struct {
	customers map[string]entity.Customer
}

func New() CustomerService {
	jsonFile, err := os.Open("./customers.json")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	dataBytes, _ := ioutil.ReadAll(jsonFile)
	customerData := map[string]entity.Customer{}
	json.Unmarshal(dataBytes, &customerData)

	return &CustomerServiceImpl{
		customers: customerData,
	}
}

func (service *CustomerServiceImpl) Find(id string) entity.Customer {
	if cust, exists := service.customers[id]; exists {
		return cust
	}
	return entity.Customer{}
}

func (service *CustomerServiceImpl) FindAll() []entity.Customer {
	customers := []entity.Customer{}
	for _, c := range service.customers {
		customers = append(customers, c)
	}
	return customers
}

func (service *CustomerServiceImpl) Add(c entity.Customer) string {
	c.ID = uuid.NewString()
	service.customers[c.ID] = c

	defer persist(service)
	return c.ID
}

func (service *CustomerServiceImpl) Update(c entity.Customer) string {
	if _, exists := service.customers[c.ID]; exists {
		service.customers[c.ID] = c
		defer persist(service)
		return c.ID
	}
	return ""
}

func (service *CustomerServiceImpl) Delete(id string) string {
	if _, exists := service.customers[id]; exists {
		delete(service.customers, id)
		defer persist(service)
		return id
	}
	return ""
}

func persist(service *CustomerServiceImpl) {
	jsonStr, err := json.MarshalIndent(service.customers, "", "\t")
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	} else {
		fmt.Println(string(jsonStr))
	}
	ioutil.WriteFile("./customers.json", jsonStr, 0644)
}
