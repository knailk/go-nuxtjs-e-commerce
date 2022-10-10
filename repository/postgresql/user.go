package postgresql

import (
	"database/sql"

	"github.com/knailk/go-shopee/entity"
	"github.com/knailk/go-shopee/usecase/product"
)

type UserRepo struct {
	db *sql.DB
}

//create an implementation of product repository
func NewUserRepo(db *sql.DB) product.Repository{
	return &ProductRepo{
		db: db,
	}
}

//TODO implement method of User repository interface
//Querying, or Creating/ Inserting into any database will be stored here.

//get User by id
func (r *UserRepo) Get(id entity.ID) (*entity.User, error){
	stmt,err := r.db.Prepare(`select id,email,name,gender,phone where id = ?`)
	if err != nil {
		return nil, err
	}
	rows,err := stmt.Query(id)
	if err != nil{
		return nil, err
	}
	var u *entity.User
	rows.Scan(&u.UserId, &u.Email, &u.Name, &u.Gender, &u.Phone)
	return u,nil
}
//search list User by query
func (r *UserRepo) Search(query string) ([]*entity.User, error){
	stmt,err := r.db.Prepare(`select id,email,name,gender,phone where name like ? `)
	if err != nil {
		return nil, err
	}
	rows,err := stmt.Query("%" + query + "%")
	if err != nil{
		return nil, err
	}
	var result []*entity.User
	for rows.Next() {
		var u *entity.User
		rows.Scan(&u.UserId, &u.Email, &u.Name, &u.Gender, &u.Phone)
		result = append(result, u)
	}
	return result,nil
}
//TODO (continue here) get all user
func (r *UserRepo) List(entity.ID) ([]*entity.User, error){
	return nil,nil
}
//create a User
func (r *UserRepo) Create(e *entity.User) (entity.ID, error){
	return entity.NewID(),nil
}
//update a User
func (r *UserRepo) Update(e *entity.User) error {
	return nil
}
//delete a User
func (r *UserRepo) Delete(id entity.ID) error {
	return nil
}

