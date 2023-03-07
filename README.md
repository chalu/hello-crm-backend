# hello-crm-backend

This is a simple Golang backend API for CRUD operations on Customer data. 

Users will be able to retrieve (`GET`), create (`POST`), update (`PUT`), and delete (`DELETE`) customers on the `/customers` resource endpoint of the server's deployed URL which defaults to `http://localhost:8080` when running on local

## Usage

> To **get started**, issue the `go run server.go` command in a shell once you are inside the `hello-crm-backend` folder. With the server running, go to `<SERVER_URL>/` to see a simple welcome message and hints on the available endpoints to call

---

1.  `cd` into the `hello-crm-backend` codebase and start the server by issuing the `go run server.go` command. This will display the server URL in the console. For the sake of this guide, let's call the URL `<SERVER_URL>`
2.  Using your HTTP client of choice (e.g Postman, cURL e.t.c), make a GET request to `<SERVER_URL>/customers`. This should result in a 200 response with a response payload of existing customers in the database
3.  Take a look at the `Customer` struct at `./entity/customer.go` for the data fields required to create a Customer object
4.  Making **POST/PUT/DELETE** calls to `<SERVER_URL>/customers` will return the ID of the customer created, updated or deleted.`PUT` and `DELETE` calls require using an ID path parameter (i.e `<SERVER_URL>/customers/:id`), and so does GET `<SERVER_URL>/customers/:id` for getting a specific customer
5.  Customer data is loaded from and persisted to the `customers.json` file

### Sample customer request object for POST

```json
{
    "age": 90,
    "name": "Someone Else",
    "contacted": true,
    "email": "lol@gmail.com",
    "phoneNumber": "54739534534",
    "address": "Somewhere they call homme"
}

```
