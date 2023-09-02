package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Invoice struct {
	ID               primitive.ObjectID `bson:"_id"`
	Invoice_id       string             `json:"invoice_id"`
	Order_id         string             `json:"order_id"`
	Payment_method   *string            `json:"payment_method" validate:"eq=CARD|eq=CASH|eq="`
	Payment_status   *string            `json:"payment_status" validate:"required,eq=PENDING|eq=PAID"`
	Payment_due_date time.Time          `json:"Payment_due_date"`
	Created_at       time.Time          `json:"created_at"`
	Updated_at       time.Time          `json:"updated_at"`
}

///////////////////////////////////////////////////////////// percobaan 1
// package models

// import (
// 	"time"

// 	"go.mongodb.org/mongo-driver/bson/primitive"
// )

// type Invoice struct {
// 	InvoiceID      primitive.ObjectID `bson:"invoice_id,omitempty" json:"invoice_id,omitempty"`
// 	PaymentMethod  string             `bson:"payment_method" json:"payment_method"`
// 	OrderID        string             `bson:"order_id" json:"order_id"`
// 	PaymentStatus  string             `bson:"payment_status" json:"payment_status"`
// 	PaymentDue     interface{}        `bson:"payment_due" json:"payment_due"`
// 	TableNumber    interface{}        `bson:"table_number" json:"table_number"`
// 	PaymentDueDate time.Time          `bson:"payment_due_date" json:"payment_due_date"`
// 	OrderDetails   interface{}        `bson:"order_details" json:"order_details"`
// 	CreatedAt      time.Time          `bson:"created_at,omitempty" json:"created_at,omitempty"`
// 	UpdatedAt      time.Time          `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
// }

// // //////////////////////////////////////////////////////// percobaan 2
// package models

// import (
// 	"time"

// 	"go.mongodb.org/mongo-driver/bson/primitive"
// )

// type Invoice struct {
// 	ID             primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
// 	CreatedAt      time.Time          `json:"created_at" bson:"created_at"`
// 	InvoiceID      string             `json:"invoice_id" bson:"invoice_id"`
// 	OrderID        string             `json:"order_id" bson:"order_id"`
// 	PaymentDueDate time.Time          `json:"payment_due_date" bson:"payment_due_date"`
// 	PaymentMethod  string             `json:"payment_method" bson:"payment_method"`
// 	PaymentStatus  string             `json:"payment_status" bson:"payment_status"`
// 	UpdatedAt      time.Time          `json:"updated_at" bson:"updated_at"`
// 	TableNumber    string             `json:"table_number" bson:"table_number"`
// 	OrderDetails   []OrderDetail      `json:"order_details" bson:"order_details"`
// }

// type OrderDetail struct {
// 	TableNumber string `json:"table_number" bson:"table_number"`
// 	Name        string `json:"name" bson:"name"`
// }
