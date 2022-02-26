package routes

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/ekholme/learning-go-projects/restaurant-orders/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var validate = validator.New()
var orderCollection *mongo.Collection = OpenCollection(Client, "orders")

//add an order
func AddOrder(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	var order models.Order

	if err := c.BindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validationErr := validate.Struct(order)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		return
	}

	order.ID = primitive.NewObjectID()

	result, err := orderCollection.InsertOne(ctx, order)
	if err != nil {
		msg := fmt.Sprintf("order item was not created")
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		return
	}

	defer cancel()
	c.JSON(http.StatusOK, result)
}

//get all orders
func GetOrders(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(),
		100*time.Second)
	var orders []bson.M

	cursor, err := orderCollection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err = cursor.All(ctx, &orders); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer cancel()

	fmt.Println(orders)

	c.JSON(http.StatusOK, orders)
}

func GetOrdersByWaiter(c *gin.Context) {
	waiter := c.Params.ByName("waiter")
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var orders []bson.M

	cursor, err := orderCollection.Find(ctx, bson.M{"server": waiter})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err = cursor.All(ctx, &orders); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer cancel()

	fmt.Println(orders)

	c.JSON(http.StatusOK, orders)
}

//RESUME AT GetOrderByID
