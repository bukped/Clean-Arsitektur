// package controller

// import (
// 	"context"
// 	"golang-restaurant-management/database"
// 	"golang-restaurant-management/models"
// 	"net/http"
// 	"time"

// 	"github.com/gin-gonic/gin"
// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// type OrderItemPack struct {
// 	Table_id    *string
// 	Order_items []models.OrderItem
// }

// var orderItemCollection *mongo.Collection = database.OpenCollection(database.Client, "orderItem")

// // func GetOrderItems() gin.HandlerFunc {
// // 	return func(c *gin.Context) {
// // 		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

// // 		result, err := orderItemCollection.Find(context.TODO(), bson.M{})

// // 		defer cancel()
// // 		if err != nil {
// // 			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listing ordered items"})
// // 			return
// // 		}
// // 		var allOrderItems []bson.M
// // 		if err = result.All(ctx, &allOrderItems); err != nil {
// // 			log.Fatal(err)
// // 			return
// // 		}
// // 		c.JSON(http.StatusOK, allOrderItems)
// // 	}
// // }

// // func GetOrderItemsByOrder() gin.HandlerFunc {
// // 	return func(c *gin.Context) {
// // 		orderId := c.Param("order_id")

// // 		allOrderItems, err := ItemsByOrder(orderId)

// // 		if err != nil {
// // 			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listing order items by order ID"})
// // 			return
// // 		}
// // 		c.JSON(http.StatusOK, allOrderItems)
// // 	}
// // }

// // func ItemsByOrder(id string) (OrderItems []primitive.M, err error) {
// // 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

// // 	matchStage := bson.D{{"$match", bson.D{{"order_id", id}}}}
// // 	lookupStage := bson.D{{"$lookup", bson.D{{"from", "food"}, {"localField", "food_id"}, {"foreignField", "food_id"}, {"as", "food"}}}}
// // 	unwindStage := bson.D{{"$unwind", bson.D{{"path", "$food"}, {"preserveNullAndEmptyArrays", true}}}}

// // 	lookupOrderStage := bson.D{{"$lookup", bson.D{{"from", "order"}, {"localField", "order_id"}, {"foreignField", "order_id"}, {"as", "order"}}}}
// // 	unwindOrderStage := bson.D{{"$unwind", bson.D{{"path", "$order"}, {"preserveNullAndEmptyArrays", true}}}}

// // 	lookupTableStage := bson.D{{"$lookup", bson.D{{"from", "table"}, {"localField", "order.table_id"}, {"foreignField", "table_id"}, {"as", "table"}}}}
// // 	unwindTableStage := bson.D{{"$unwind", bson.D{{"path", "$table"}, {"preserveNullAndEmptyArrays", true}}}}

// // 	projectStage := bson.D{
// // 		{"$project", bson.D{
// // 			{"id", 0},
// // 			{"amount", "$food.price"},
// // 			{"total_count", 1},
// // 			{"food_name", "$food.name"},
// // 			{"food_image", "$food.food_image"},
// // 			{"table_number", "$table.table_number"},
// // 			{"table_id", "$table.table_id"},
// // 			{"order_id", "$order.order_id"},
// // 			{"price", "$food.price"},
// // 			{"quantity", 1},
// // 		}}}

// // 	groupStage := bson.D{{"$group", bson.D{{"_id", bson.D{{"order_id", "$order_id"}, {"table_id", "$table_id"}, {"table_number", "$table_number"}}}, {"payment_due", bson.D{{"$sum", "$amount"}}}, {"total_count", bson.D{{"$sum", 1}}}, {"order_items", bson.D{{"$push", "$$ROOT"}}}}}}

// // 	projectStage2 := bson.D{
// // 		{"$project", bson.D{

// // 			{"id", 0},
// // 			{"payment_due", 1},
// // 			{"total_count", 1},
// // 			{"table_number", "$_id.table_number"},
// // 			{"order_items", 1},
// // 		}}}

// // 	result, err := orderItemCollection.Aggregate(ctx, mongo.Pipeline{
// // 		matchStage,
// // 		lookupStage,
// // 		unwindStage,
// // 		lookupOrderStage,
// // 		unwindOrderStage,
// // 		lookupTableStage,
// // 		unwindTableStage,
// // 		projectStage,
// // 		groupStage,
// // 		projectStage2})

// // 	if err != nil {
// // 		panic(err)
// // 	}

// // 	if err = result.All(ctx, &OrderItems); err != nil {
// // 		panic(err)
// // 	}

// // 	defer cancel()

// // 	return OrderItems, err

// // }

// // /////////////////////////////////////// script ganti
// // kode orderItemsController.go
// func ItemsByOrder(id string) ([]primitive.M, error) {
// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

// 	matchStage := bson.D{{"$match", bson.D{{"order_id", id}}}}
// 	lookupStage := bson.D{{"$lookup", bson.D{{"from", "food"}, {"localField", "food_id"}, {"foreignField", "food_id"}, {"as", "food"}}}}
// 	unwindStage := bson.D{{"$unwind", bson.D{{"path", "$food"}, {"preserveNullAndEmptyArrays", true}}}}

// 	lookupOrderStage := bson.D{{"$lookup", bson.D{{"from", "order"}, {"localField", "order_id"}, {"foreignField", "order_id"}, {"as", "order"}}}}
// 	unwindOrderStage := bson.D{{"$unwind", bson.D{{"path", "$order"}, {"preserveNullAndEmptyArrays", true}}}}

// 	lookupTableStage := bson.D{{"$lookup", bson.D{{"from", "table"}, {"localField", "order.table_id"}, {"foreignField", "table_id"}, {"as", "table"}}}}
// 	unwindTableStage := bson.D{{"$unwind", bson.D{{"path", "$table"}, {"preserveNullAndEmptyArrays", true}}}}

// 	projectStage := bson.D{
// 		{"$project", bson.D{
// 			{"id", 0},
// 			{"amount", "$food.price"},
// 			{"total_count", 1},
// 			{"food_name", "$food.name"},
// 			{"food_image", "$food.food_image"},
// 			{"table_number", "$table.table_number"},
// 			{"table_id", "$table.table_id"},
// 			{"order_id", "$order.order_id"},
// 			{"price", "$food.price"},
// 			{"quantity", 1},
// 		}},
// 	}

// 	groupStage := bson.D{{"$group", bson.D{{"_id", bson.D{{"order_id", "$order_id"}, {"table_id", "$table_id"}, {"table_number", "$table_number"}}}, {"payment_due", bson.D{{"$sum", "$amount"}}}, {"total_count", bson.D{{"$sum", 1}}}, {"order_items", bson.D{{"$push", "$$ROOT"}}}}}}

// 	projectStage2 := bson.D{
// 		{"$project", bson.D{
// 			{"id", 0},
// 			{"payment_due", 1},
// 			{"total_count", 1},
// 			{"table_number", "$_id.table_number"},
// 			{"order_items", 1},
// 		}},
// 	}

// 	var orderItems []primitive.M
// 	result, err := orderItemCollection.Aggregate(ctx, mongo.Pipeline{
// 		matchStage,
// 		lookupStage,
// 		unwindStage,
// 		lookupOrderStage,
// 		unwindOrderStage,
// 		lookupTableStage,
// 		unwindTableStage,
// 		projectStage,
// 		groupStage,
// 		projectStage2,
// 	})

// 	if err != nil {
// 		return nil, err
// 	}

// 	if err = result.All(ctx, &orderItems); err != nil {
// 		return nil, err
// 	}

// 	defer cancel()

// 	return orderItems, nil
// }

// func GetOrderItem() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

// 		orderItemId := c.Param("order_item_id")
// 		var orderItem models.OrderItem

// 		err := orderItemCollection.FindOne(ctx, bson.M{"orderItem_id": orderItemId}).Decode(&orderItem)
// 		defer cancel()
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listing ordered item"})
// 			return
// 		}
// 		c.JSON(http.StatusOK, orderItem)
// 	}
// }

// func UpdateOrderItem() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

// 		var orderItem models.OrderItem

// 		orderItemId := c.Param("order_item_id")

// 		filter := bson.M{"order_item_id": orderItemId}

// 		var updateObj primitive.D

// 		if orderItem.Unit_price != nil {
// 			updateObj = append(updateObj, bson.E{"unit_price", *&orderItem.Unit_price})
// 		}

// 		if orderItem.Quantity != nil {
// 			updateObj = append(updateObj, bson.E{"quantity", *orderItem.Quantity})
// 		}

// 		if orderItem.Food_id != nil {
// 			updateObj = append(updateObj, bson.E{"food_id", *orderItem.Food_id})
// 		}

// 		orderItem.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
// 		updateObj = append(updateObj, bson.E{"updated_at", orderItem.Updated_at})

// 		upsert := true
// 		opt := options.UpdateOptions{
// 			Upsert: &upsert,
// 		}

// 		result, err := orderItemCollection.UpdateOne(
// 			ctx,
// 			filter,
// 			bson.D{
// 				{"$set", updateObj},
// 			},
// 			&opt,
// 		)

// 		if err != nil {
// 			msg := "Order item update failed"
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
// 			return
// 		}

// 		defer cancel()

// 		c.JSON(http.StatusOK, result)
// 	}
// }

// // var (
// // 	orderItemCollection *mongo.Collection = database.OpenCollection(database.Client, "order_item")
// // )

// func CreateOrderItem() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var orderItem models.OrderItem
// 		var order models.Order

// 		if err := c.BindJSON(&orderItem); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		validationErr := validate.Struct(orderItem)

// 		if validationErr != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
// 			return
// 		}

// 		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

// 		orderCollection := database.OpenCollection(database.Client, "order")

// 		err := orderCollection.FindOne(ctx, bson.M{"order_id": orderItem.OrderID}).Decode(&order)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "order not found"})
// 			cancel()
// 			return
// 		}

// 		orderItem.CreatedAt = time.Now()
// 		orderItem.UpdatedAt = time.Now()
// 		orderItem.ID = primitive.NewObjectID()

// 		_, err = orderItemCollection.InsertOne(ctx, orderItem)
// 		defer cancel()
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while creating order item"})
// 			return
// 		}

// 		c.JSON(http.StatusOK, orderItem)
// 	}
// }

// /////////////////////////////////////// script gangrti 2
package controller

import (
	"context"
	"golang-restaurant-management/database"
	"golang-restaurant-management/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderItemPack struct {
	Table_id    *string
	Order_items []models.OrderItem
}

var (
	orderItemCollection *mongo.Collection = database.OpenCollection(database.Client, "orderItem")
	tableColl           *mongo.Collection = database.OpenCollection(database.Client, "table")
)

func GetOrderItems(c *gin.Context) {
	orderID := c.Query("order_id")

	orderItems, err := ItemsByOrder(orderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get order items",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":     "Success",
		"order_items": orderItems,
	})
}

func GetOrderItemByID(c *gin.Context) {
	orderItemID := c.Param("order_item_id")

	orderItem, err := FindOrderItemByID(orderItemID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get order item",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":    "Success",
		"order_item": orderItem,
	})
}

func FindOrderItemByID(orderItemID string) (*models.OrderItem, error) {
	var ctx = context.Background()

	objectID, err := primitive.ObjectIDFromHex(orderItemID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objectID}

	var orderItem models.OrderItem
	err = orderItemCollection.FindOne(ctx, filter).Decode(&orderItem)
	if err != nil {
		return nil, err
	}

	return &orderItem, nil
}

func ItemsByOrder(id string) ([]primitive.M, error) {
	var ctx = context.Background()

	matchStage := bson.D{{"$match", bson.D{{"order_id", id}}}}
	lookupStage := bson.D{{"$lookup", bson.D{{"from", "food"}, {"localField", "food_id"}, {"foreignField", "food_id"}, {"as", "food"}}}}
	unwindStage := bson.D{{"$unwind", bson.D{{"path", "$food"}, {"preserveNullAndEmptyArrays", true}}}}

	lookupOrderStage := bson.D{{"$lookup", bson.D{{"from", "order"}, {"localField", "order_id"}, {"foreignField", "order_id"}, {"as", "order"}}}}
	unwindOrderStage := bson.D{{"$unwind", bson.D{{"path", "$order"}, {"preserveNullAndEmptyArrays", true}}}}

	lookupTableStage := bson.D{{"$lookup", bson.D{{"from", "table"}, {"localField", "order.table_id"}, {"foreignField", "table_id"}, {"as", "table"}}}}
	unwindTableStage := bson.D{{"$unwind", bson.D{{"path", "$table"}, {"preserveNullAndEmptyArrays", true}}}}

	projectStage := bson.D{
		{"$project", bson.D{
			{"id", 0},
			{"amount", "$food.price"},
			{"total_count", 1},
			{"food_name", "$food.name"},
			{"food_image", "$food.food_image"},
			{"table_number", "$table.table_number"},
			{"table_id", "$table.table_id"},
			{"order_id", "$order.order_id"},
			{"price", "$food.price"},
			{"quantity", 1},
		}},
	}

	groupStage := bson.D{{"$group", bson.D{{"_id", bson.D{{"order_id", "$order_id"}, {"table_id", "$table_id"}, {"table_number", "$table_number"}}}, {"payment_due", bson.D{{"$sum", "$amount"}}}, {"total_count", bson.D{{"$sum", 1}}}}}}

	pipeline := mongo.Pipeline{matchStage, lookupStage, unwindStage, lookupOrderStage, unwindOrderStage, lookupTableStage, unwindTableStage, projectStage, groupStage}

	cur, err := orderItemCollection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}

	var results []primitive.M
	if err := cur.All(ctx, &results); err != nil {
		return nil, err
	}

	return results, nil
}

func CreateOrderItem(c *gin.Context) {
	var orderItem models.OrderItem
	if err := c.ShouldBindJSON(&orderItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
		return
	}

	orderItem.OrderItemID = primitive.NewObjectID().Hex()
	orderItem.CreatedAt = time.Now()

	// Dapatkan order terkait berdasarkan table_id
	var table models.Table
	err := tableCollection.FindOne(context.TODO(), bson.M{"table_id": orderItem.TableID}).Decode(&table)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get table",
			"error":   err.Error(),
		})
		return
	}

	orderItem.OrderID = table.ID.Hex() // Set order_id pada order item

	_, err = orderItemCollection.InsertOne(context.TODO(), orderItem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create order item",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":    "Order item created successfully",
		"order_item": orderItem,
	})
}

func UpdateOrderItem(c *gin.Context) {
	orderItemID := c.Param("order_item_id")

	var orderItem models.OrderItem
	if err := c.ShouldBindJSON(&orderItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
		return
	}

	filter := bson.M{"order_item_id": orderItemID}
	update := bson.M{"$set": bson.M{
		"quantity":   orderItem.Quantity,
		"updated_at": time.Now(),
	}}

	_, err := orderItemCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to update order item",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Order item updated successfully",
	})
}

func DeleteOrderItem(c *gin.Context) {
	orderItemID := c.Param("order_item_id")

	_, err := orderItemCollection.DeleteOne(context.TODO(), bson.M{"order_item_id": orderItemID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to delete order item",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Order item deleted successfully",
	})
}
