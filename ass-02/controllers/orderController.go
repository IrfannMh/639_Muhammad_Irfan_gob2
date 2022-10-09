package controllers

import (
	"ass-02/structs"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type InDB struct {
	DB *gorm.DB
}

// to get all data order
func (idb *InDB) GetOrder(c *gin.Context) {
	var (
		orders []structs.Order
		result gin.H
	)

	idb.DB.Preload("Items").Find(&orders)
	if len(orders) <= 0 {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": orders,
			"count":  len(orders),
		}
	}

	c.JSON(http.StatusOK, result)
}

// Create new data order
func (idb *InDB) CreateOrder(c *gin.Context) {
	var (
		order  structs.Order
		result gin.H
	)

	json.NewDecoder(c.Request.Body).Decode(&order)

	err := idb.DB.Model(structs.Order{}).Create(&order).Error

	if err != nil {
		result = gin.H{
			"status": "Failed",
			"result": err.Error(),
		}
	}
	result = gin.H{
		"data": order,
	}
	c.JSON(http.StatusOK, result)
}

// update data
func (idb *InDB) UpdateOrder(c *gin.Context) {
	var (
		order    structs.Order
		newOrder structs.Order
		result   gin.H
	)

	id := c.Param("orderId")

	json.NewDecoder(c.Request.Body).Decode(&newOrder)

	err := idb.DB.Where("id = ?", id).First(&order).Error
	if err != nil {
		result = gin.H{
			"result":  err.Error(),
			"message": "data not found",
		}
	}
	order.Items = newOrder.Items
	order.CustomerName = newOrder.CustomerName
	order.OrderedAt = newOrder.OrderedAt
	idb.DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&order)

	result = gin.H{
		"result": order,
	}

	c.JSON(http.StatusOK, result)

}

func (idb *InDB) DeleteOrder(c *gin.Context) {
	var order structs.Order

	id := c.Param("ordersId")

	err := idb.DB.First(&order, id).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"result": "Order Not Found",
		})
		return
	}

	err = idb.DB.Select("Items").Delete(&order).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"result": fmt.Sprintf("Error deleting order data by id: %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Order deleted",
	})
}
