package entity

type Customer struct {
	ID        string `json:"id"`
	Age       int    `json:"age"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Address   string `json:"address"`
	Contacted bool   `json:"contacted"`
	Phone     string `json:"phoneNumber"`
}
