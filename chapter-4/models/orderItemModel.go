// package models

// import (
// 	"time"

// 	"go.mongodb.org/mongo-driver/bson/primitive"
// )

// type OrderItem struct {
// 	ID            primitive.ObjectID `bson:"_id"`
// 	Quantity      *string            `json:"quantity" validate:"required,eq=S|eq=M|eq=L"`
// 	Unit_price    *float64           `json:"unit_price" validate:"required"`
// 	Created_at    time.Time          `json:"created_at"`
// 	Updated_at    time.Time          `json:"updated_at"`
// 	Food_id       *string            `json:"food_id" validate:"required"`
// 	Order_item_id string             `json:"order_item_id"`
// 	Order_id      string             `json:"order_id" validate:"required"`
// }

// ////////////////////////////////////// coba
// package models

// import "time"

// type OrderItem struct {
// 	OrderItemID string    `json:"order_item_id"`
// 	OrderID     string    `json:"order_id"`
// 	FoodID      string    `json:"food_id"`
// 	Quantity    int       `json:"quantity"`
// 	CreatedAt   time.Time `json:"created_at"`
// 	UpdatedAt   time.Time `json:"updated_at"`
// }

// ///////////////////////////////////////////coba 1
// package models

// import (
// 	"time"
// )

// type OrderItem struct {
// 	OrderItemID string    `json:"order_item_id" bson:"order_item_id"`
// 	OrderID     string    `json:"order_id" bson:"order_id"`
// 	FoodID      string    `json:"food_id" bson:"food_id"`
// 	Quantity    int       `json:"quantity" bson:"quantity"`
// 	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
// 	UpdatedAt   time.Time `json:"updated_at" bson:"updated_at"`
// }

// ////////////////////////ganti 2
package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderItem struct {
	OrderItemID string             `json:"order_item_id" bson:"order_item_id"`
	OrderID     string             `json:"order_id" bson:"order_id"`
	FoodID      string             `json:"food_id" bson:"food_id"`
	Quantity    int                `json:"quantity" bson:"quantity"`
	TableID     primitive.ObjectID `json:"table_id" bson:"table_id"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
}
