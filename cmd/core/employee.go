package core

type EmployeeRepository interface {
	Store(employee *Employee) error
}

type Employee struct {
	ID           string        `json:"_" bson:"_id"`
	FirstName    string        `json:"firstName" bson:"firstName"`
	LastName     string        `json:"lastName" bson:"lastName"`
	MobileNumber string        `json:"mobileNumber" bson:"mobileNumber"`
	BankAccounts []BankAccount `json:"bankAccounts" bson:"bankAccounts"`
}

type BankAccount struct {
	CardNumber    string `json:"cardNumber" bson:"cardNumber"`
	AccountNumber string `json:"accountNumber" bson:"accountNumber"`
}
