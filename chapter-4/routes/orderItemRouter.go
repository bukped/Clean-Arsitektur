// package routes

// import (
// 	controller "golang-restaurant-management/controllers"

// 	"github.com/gin-gonic/gin"
// )

// func OrderItemRoutes(incomingRoutes *gin.Engine) {
// 	incomingRoutes.GET("/orderItems", controller.GetOrderItems())
// 	incomingRoutes.GET("/orderItems/:orderItem_id", controller.GetOrderItem())
// 	incomingRoutes.GET("/orderItems-order/:order_id", controller.GetOrderItemsByOrder())
// 	incomingRoutes.POST("/orderItems", controller.CreateOrderItem())
// 	incomingRoutes.PATCH("/orderItems/:orderItem_id", controller.UpdateOrderItem())
// }

// ////////////////////////////////////////////ganti
package routes

import (
	controllers "golang-restaurant-management/controllers"

	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orderItems", controllers.GetOrderItems)
	incomingRoutes.GET("/orderItems/:orderItemID", controllers.GetOrderItemByID)
	incomingRoutes.POST("/orderItems", controllers.CreateOrderItem)
	incomingRoutes.PATCH("/orderItems/:orderItemID", controllers.UpdateOrderItem)
	incomingRoutes.DELETE("/orderItems/:orderItemID", controllers.DeleteOrderItem)
}
