package models

import (
	"database/sql"
	"fmt"
	"strings"
)

type Order struct {
	ID         string `json:"id" form:"id" query:"id"`
	CustomerID string `json:"customer_id" form:"customer_id" query:"customer_id"`
	OrderName  string `json:"order_name" form:"order_name" query:"order_name"`
	CreatedAt  string `json:"created_at" form:"created_at" query:"created_at"`
}

type OrderDetail struct {
	ID              string `json:"id" form:"id" query:"id"`
	OrderName       string `json:"order_name" form:"order_name" query:"order_name"`
	CreatedAt       string `json:"created_at" form:"created_at" query:"created_at"`
	OrderProduct    string `json:"order_product" form:"order_product" query:"order_product"`
	TotalQuantity   string `json:"total_quantity" form:"total_quantity" query:"total_quantity"`
	DeliveredAmount string `json:"delivered_amount" form:"delivered_amount" query:"delivered_amount"`
	UserID          string `json:"user_id" form:"user_id" query:"user_id"`
	CompanyName     string `json:"company_name" form:"company_name" query:"company_name"`
}

type OrderModel struct {
	DB *sql.DB
}

// CreateOrder inserts a new order into the database
func (m *OrderModel) CreateOrder(order Order) error {
	query := "INSERT INTO orders (customer_id, order_name, created_at) VALUES ($1, $2, $3)"
	_, err := m.DB.Exec(query, order.CustomerID, order.OrderName, order.CreatedAt)
	if err != nil {
		return fmt.Errorf("could not insert order: %v", err)
	}
	return nil
}

// GetTotalOrderCount retrieves the total number of orders
func (m *OrderModel) GetTotalOrderCount() (int, error) {
	var count int
	query := "SELECT COUNT(*) FROM orders"
	err := m.DB.QueryRow(query).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("could not get order count: %v", err)
	}
	return count, nil
}

// GetTotalOrderDetailCount retrieves the total number of orders
func (m *OrderModel) GetTotalOrderDetailCount(search string, start string, end string) (int, error) {
	var count int
	query := `with customerList(user_id, company_name) as (
		select customers.user_id, customer_companies.company_name
		from customers, customer_companies
		where customers.company_id = customer_companies.company_id
	),

	orderItemList(order_id, order_product, total_quantity, delivered_amount) as (
		select order_items.order_id, order_items.product, order_items.quantity, 
		deliveries.delivered_quantity
		from order_items
		left join deliveries on order_items.id = deliveries.order_item_id
	)


	select count(*)
	from orders
	Left join orderItemList on orders.id = orderItemList.order_id
	left join customerList on orders.customer_id = customerList.user_id	
	`

	where := ""
	var condition []string
	if search != "" {
		condition = append(condition, fmt.Sprintf(" (orders.order_name like '%%%s%%' or customerList.user_id like '%%%s%%') ", search, search))
	}

	if start != "" {
		condition = append(condition, fmt.Sprintf(" orders.created_at > '%s' ", start))
	}

	if end != "" {
		condition = append(condition, fmt.Sprintf(" orders.created_at < '%s' ", end))
	}

	if len(condition) > 0 {
		where = strings.Join(condition[:], " and ")
		where = strings.Join([]string{" where ", where}, " ")
	}
	query = strings.Join([]string{query, where}, " ")

	err := m.DB.QueryRow(query).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("could not get order count: %v", err)
	}
	return count, nil
}

// GetOrder retrieves an order by ID
func (m *OrderModel) GetOrder(id string) (*Order, error) {
	var order Order
	query := "SELECT id, customer_id, order_name, created_at FROM orders WHERE id = $1"
	err := m.DB.QueryRow(query, id).Scan(&order.ID, &order.CustomerID, &order.OrderName, &order.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("could not find order: %v", err)
	}
	return &order, nil
}

// GetOrderByPage retrieves orders by page number
func (m *OrderModel) GetOrdersByPage(page int, count int) ([]Order, error) {
	var orders []Order

	// Calculate the offset based on the page and count
	offset := (page - 1) * count

	// Query to fetch orders with pagination
	query := "SELECT id, customer_id, order_name, created_at FROM orders LIMIT $1 OFFSET $2"
	rows, err := m.DB.Query(query, count, offset)
	if err != nil {
		return nil, fmt.Errorf("could not fetch orders: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var order Order
		if err := rows.Scan(&order.ID, &order.CustomerID, &order.OrderName, &order.CreatedAt); err != nil {
			return nil, fmt.Errorf("could not scan order: %v", err)
		}
		orders = append(orders, order)
	}
	return orders, nil
}

// GetOrderByPage retrieves orders by page number
func (m *OrderModel) GetOrderDetailsByPage(page int, count int, search string,
	start string, end string) ([]OrderDetail, error) {
	var orders []OrderDetail

	// Calculate the offset based on the page and count
	offset := (page - 1) * count

	// Query to fetch orders with pagination
	query := `with customerList(user_id, company_name) as (
		select customers.user_id, customer_companies.company_name
		from customers, customer_companies
		where customers.company_id = customer_companies.company_id
	),

	orderItemList(order_id, order_product, total_quantity, delivered_amount) as (
		select order_items.order_id, order_items.product, order_items.quantity, 
		deliveries.delivered_quantity
		from order_items
		left join deliveries on order_items.id = deliveries.order_item_id
	)


	select orders.id, orders.order_name, orders.created_at, orderItemList.order_product, 
	COALESCE(orderItemList.total_quantity, '') as total_quantity, COALESCE(orderItemList.delivered_amount, '') as delivered_amount,
	customerList.user_id, customerList.company_name 
	from orders
	Left join orderItemList on orders.id = orderItemList.order_id
	left join customerList on orders.customer_id = customerList.user_id
	`
	pagination := fmt.Sprintf(`Limit %d offset %d`, count, offset)

	where := ""
	var condition []string
	if search != "" {
		condition = append(condition, fmt.Sprintf(" (orders.order_name like '%%%s%%' or customerList.user_id like '%%%s%%') ", search, search))
	}

	if start != "" {
		condition = append(condition, fmt.Sprintf(" orders.created_at > '%s' ", start))
	}

	if end != "" {
		condition = append(condition, fmt.Sprintf(" orders.created_at < '%s' ", end))
	}

	if len(condition) > 0 {
		where = strings.Join(condition[:], " and ")
		where = strings.Join([]string{" where ", where}, " ")
	}
	query = strings.Join([]string{query, where, pagination}, " ")

	rows, err := m.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("could not fetch orders: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var order OrderDetail
		if err := rows.Scan(&order.ID, &order.OrderName, &order.CreatedAt,
			&order.OrderProduct, &order.TotalQuantity, &order.DeliveredAmount,
			&order.UserID, &order.CompanyName); err != nil {
			return nil, fmt.Errorf("could not scan order: %v", err)
		}

		orders = append(orders, order)
	}
	return orders, nil
}

// UpdateOrder updates an existing order
func (m *OrderModel) UpdateOrder(order Order) error {
	query := "UPDATE orders SET customer_id = $1, order_name = $2, created_at = $3 WHERE id = $4"
	_, err := m.DB.Exec(query, order.CustomerID, order.OrderName, order.CreatedAt, order.ID)
	if err != nil {
		return fmt.Errorf("could not update order: %v", err)
	}
	return nil
}

// DeleteOrder deletes an order by ID
func (m *OrderModel) DeleteOrder(id string) error {
	query := "DELETE FROM orders WHERE id = $1"
	_, err := m.DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("could not delete order: %v", err)
	}
	return nil
}
