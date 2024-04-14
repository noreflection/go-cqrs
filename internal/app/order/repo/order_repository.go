package repo

import (
	"database/sql"
	"github.com/pkg/errors"
	"go-cqrs/internal/domain"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) Create(order domain.Order) (string, error) {
	var orderID string
	err := r.db.QueryRow("INSERT INTO orders (product, quantity) VALUES ($1, $2) RETURNING id", order.Product, order.Quantity).Scan(&orderID)
	if err != nil {
		return "", errors.Wrap(err, "Failed to create order")
	}
	return orderID, nil
}

//func (r *OrderRepository) Create(order domain.Order) (string, error) {
//	// Insert a new order into the database and return the ID of the newly created order.
//	stmt, err := r.db.Prepare("INSERT INTO orders (...) VALUES (...)")
//	if err != nil {
//		return "", errors.Wrap(err, "failed to prepare statement for creating order")
//	}
//	defer stmt.Close()
//
//	_, err = stmt.Exec( /* values for other columns */ )
//	if err != nil {
//		return "", errors.Wrap(err, "failed to create order")
//	}
//	// Retrieve the auto-generated ID of the newly created order
//	orderID, err := r.getLastInsertedOrderID()
//	if err != nil {
//		return "", errors.Wrap(err, "failed to retrieve order ID")
//	}
//	return orderID, nil
//}

func (r *OrderRepository) Get(orderID int) (domain.Order, error) {
	// Retrieve an order by ID from the database.
	var order domain.Order
	println("here")
	println(orderID)
	err := r.db.QueryRow("SELECT * FROM Orders WHERE id = $1", orderID).Scan(&order.ID /* other fields */)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Order{}, errors.Wrap(err, "order not found")
		}
		return domain.Order{}, errors.Wrap(err, "failed to get order")
	}
	return order, nil
}

func (r *OrderRepository) Update(order domain.Order) error {
	// Update an existing order in the database.
	_, err := r.db.Exec("UPDATE Orders SET ... WHERE id = ?", order.ID /* other fields */)
	if err != nil {
		return errors.Wrap(err, "failed to update order")
	}
	return nil
}

//func (r *OrderRepository) Update(order domain.Order) error {
//	// Update an existing order in the database.
//	stmt, err := r.db.Prepare("UPDATE orders SET ... WHERE id = ?")
//	if err != nil {
//		return errors.Wrap(err, "failed to prepare statement for updating order")
//	}
//	defer stmt.Close()
//
//	_, err = stmt.Exec(order.ID, ...)
//	if err != nil {
//		return errors.Wrap(err, "failed to update order")
//	}
//	return nil
//}

func (r *OrderRepository) Delete(orderID string) error {
	// Delete an order by ID from the database.
	stmt, err := r.db.Prepare("DELETE FROM Orders WHERE id = ?")
	if err != nil {
		return errors.Wrap(err, "failed to prepare statement for deleting order")
	}
	defer stmt.Close()

	_, err = stmt.Exec(orderID)
	if err != nil {
		return errors.Wrap(err, "failed to delete order")
	}
	return nil
}
