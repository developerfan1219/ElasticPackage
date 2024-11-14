// handlers/order_handler.go
package handlers

import (
	"fmt"
	"net/http"

	"example.com/elasticpackage/models"
	"github.com/labstack/echo/v4"
)

type OrderHandler struct {
	OrderModel *models.OrderModel
}

type OrderPageRequest struct {
	Page   int    `json: "page"`
	Count  int    `json: "count"`
	Start  string `json: "start"`
	End    string `json: "end"`
	Search string `json: "search"`
}

// CreateOrder creates a new order
func (h *OrderHandler) CreateOrder(c echo.Context) error {
	var order models.Order
	if err := c.Bind(&order); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid input")
	}
	if err := h.OrderModel.CreateOrder(order); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, order)
}

// GetOrder retrieves an order by ID
func (h *OrderHandler) GetOrder(c echo.Context) error {
	id := c.Param("id")
	order, err := h.OrderModel.GetOrder(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Order not found")
	}
	return c.JSON(http.StatusOK, order)
}

// GetOrderByPage retrieves orders by page number and item number per page
func (h *OrderHandler) GetOrdersByPage(c echo.Context) error {

	postBody := new(OrderPageRequest)

	if err := c.Bind(postBody); err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid page request")
	}

	orders, err := h.OrderModel.GetOrdersByPage(postBody.Page, postBody.Count)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	total, err := h.OrderModel.GetTotalOrderCount()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]any{"items": orders, "total": total, "page": postBody.Page})
}

// GetOrderDetailsByPage retrieves order details by page number and item number per page
func (h *OrderHandler) GetOrderDetailsByPage(c echo.Context) error {

	postBody := new(OrderPageRequest)

	if err := c.Bind(postBody); err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid page request")
	}

	orders, err := h.OrderModel.GetOrderDetailsByPage(postBody.Page, postBody.Count, postBody.Search,
		postBody.Start, postBody.End)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	total, err := h.OrderModel.GetTotalOrderDetailCount(postBody.Search,
		postBody.Start, postBody.End)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]any{"items": orders, "total": total, "page": postBody.Page})
}

// UpdateOrder updates an existing order
func (h *OrderHandler) UpdateOrder(c echo.Context) error {
	id := c.Param("id")
	var order models.Order
	if err := c.Bind(&order); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid input")
	}
	order.ID = id
	if err := h.OrderModel.UpdateOrder(order); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, order)
}

// DeleteOrder deletes an order by ID
func (h *OrderHandler) DeleteOrder(c echo.Context) error {
	id := c.Param("id")
	if err := h.OrderModel.DeleteOrder(id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
