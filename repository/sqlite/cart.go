package sqlite

import (
	"database/sql"

	"github.com/knailk/go-shopee/app/entity"
)

type CartRepo struct {
	db *sql.DB
}

//Get get all product in cart
func (r *CartRepo) GetAll(userId entity.ID) ([]*entity.Cart, error) {
	stmt, err := r.db.Prepare(`select userId,productId,quantity from cart where userId = ?`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(userId)
	if err != nil {
		return nil, err
	}
	var result []*entity.Cart
	for rows.Next(){
		var c entity.Cart
		rows.Scan(&c.UserId, &c.ProductId, &c.Quantity)
		result = append(result, &c)
	}
	return result, nil
}

//GetOne get one product in cart
func (r *CartRepo) GetOne(userId entity.ID, productId entity.ID) (*entity.Cart, error){
	stmt, err := r.db.Prepare(`select userId,productId,quantity from cart where userId = ? and productId = ?`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(userId, productId)
	if err != nil {
		return nil, err
	}
	var c entity.Cart
	for rows.Next(){
		rows.Scan(&c.UserId, &c.ProductId, &c.Quantity)
	}
	if err != nil {
		return nil, err
	}
	return &c, nil
}
//Add add a product to cart
func (r *CartRepo) Add(cart *entity.Cart) error {
	stmt, err := r.db.Prepare(`insert into cart (userId, productId, quantity) values(?,?,?)`)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_,err = stmt.Exec(cart.UserId,cart.ProductId,cart.Quantity)
	if err != nil {
		return  err
	}
	return nil
}

//Update will change quantity product of cart
func (r *CartRepo) Update(cart *entity.Cart) error {
	stmt, err := r.db.Prepare(`update cart set quantity = ?) where userId = ? and productId = ?`)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_,err = stmt.Exec(cart.Quantity,cart.UserId,cart.ProductId)
	if err != nil {
		return  err
	}
	return nil
}