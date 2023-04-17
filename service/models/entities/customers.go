package entities

import utils "go-clean-arch/pkg/utils/password"

//Customers ...
type Customers struct {
	CustomerId string `json:"customer_id" bson:"customer_id" redis:"customer_id"`
	Username   string `json:"username" bson:"username" redis:"username"`
	Name       string `json:"name" bson:"name" redis:"name"`
	Password   string `json:"password" bson:"password"`
	Salt       string `json:"salt" bson:"salt"`
	Email      string `json:"email" bson:"email" redis:"email"`
	Phone      string `json:"phone" json:"phone" redis:"phone"`
	Dob        string `json:"dob" json:"dob" json:"dob"`
	CreatedAt  int64  `json:"created_at" bson:"created_at" redis:"created_at"`
	UpdatedAt  int64  `json:"updated_at" bson:"updated_at" redis:"updated_at"`
}

// ComparePassword ...
func (c Customers) ComparePassword(password string) bool {
	return utils.ComparePasswordHashWithSalt(password, c.Salt, c.Password)
}

//customersBuilder ...
type customersBuilder struct {
	customer *Customers
}

func NewCustomersBuilder() *customersBuilder {
	return &customersBuilder{&Customers{}}
}

func (b *customersBuilder) WithCustomerID(customerID string) *customersBuilder {
	b.customer.CustomerId = customerID
	return b
}

func (b *customersBuilder) WithUsername(username string) *customersBuilder {
	b.customer.Username = username
	return b
}

func (b *customersBuilder) WithName(name string) *customersBuilder {
	b.customer.Name = name
	return b
}

func (b *customersBuilder) WithPassword(password string, salt string) *customersBuilder {
	hash := utils.GeneratePasswordHashWithSalt(password, salt)
	b.customer.Password = hash
	b.customer.Salt = salt
	return b
}

func (b *customersBuilder) WithEmail(email string) *customersBuilder {
	b.customer.Email = email
	return b
}

func (b *customersBuilder) WithPhone(phone string) *customersBuilder {
	b.customer.Phone = phone
	return b
}

func (b *customersBuilder) WithDOB(dob string) *customersBuilder {
	b.customer.Dob = dob
	return b
}

func (b *customersBuilder) WithCreatedAt(created int64) *customersBuilder {
	b.customer.CreatedAt = created
	return b
}

func (b *customersBuilder) WithUpdatedAt(updated int64) *customersBuilder {
	b.customer.UpdatedAt = updated
	return b
}

func (b *customersBuilder) Build() *Customers {
	return b.customer
}
