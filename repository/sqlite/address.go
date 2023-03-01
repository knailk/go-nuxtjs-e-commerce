package sqlite

import (
	"database/sql"
	"github.com/knailk/go-nuxtjs-e-commerce/app/entity"
)

type AddressRepo struct {
	db *sql.DB
}

func (r *AddressRepo) Get(email string) (*entity.Address, error){
	stmt, err:= r.db.Prepare(`select* from address where email = ?`)
	if err != nil {
		return nil,err
	}
	defer stmt.Close()
	rows, err := stmt.Query(email)
	var a entity.Address
	for rows.Next() {
		err = rows.Scan(&a.Email, &a.Country, &a.City, &a.District, &a.Ward, &a.Address, &a.CreatedAt, &a.UpdatedAt)
	}
	if err != nil {
		return nil, err
	}
	if a.Email == "" {
		return nil, entity.ErrNotFound
	}
	return &a, nil
}

func (r *AddressRepo) Add(address *entity.Address) error{
	stmt, err := r.db.Prepare(`
	insert into address 
	(email, country, city, district, ward, address, createdAt, updatedAt)
	values (?,?,?,?,?,?,?,?)`)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(
		address.Email,
		address.Country,
		address.City,
		address.District,
		address.Ward,
		address.Address,
		address.CreatedAt,
		address.UpdatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *AddressRepo) Update(address *entity.Address) error{
	stmt, err := r.db.Prepare(`
	update address set
	country = ?, 
	city = ?, 
	district = ?, 
	ward = ?, 
	address = ?, 
	updatedAt = ?
	where email = ?`)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(
		address.Country,
		address.City,
		address.District,
		address.Ward,
		address.Address,
		address.UpdatedAt,
		address.Email,
	)
	if err != nil {
		return err
	}
	return nil
}