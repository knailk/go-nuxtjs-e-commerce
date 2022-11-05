package sqlite

import (
	"database/sql"
	"time"

	"github.com/knailk/go-shopee/app/entity"
	"github.com/knailk/go-shopee/app/usecase/user"
)

type UserRepo struct {
	db *sql.DB
}

// NewUserRepo create an implementation of product repository.
func NewUserRepo(db *sql.DB) user.Repository {
	return &UserRepo{
		db: db,
	}
}

// Get User by id.
func (r *UserRepo) Get(id entity.ID) (*entity.User, error) {
	stmt, err := r.db.Prepare(`select id,email,name,gender,phone,createdAt, updatedAt from user where id = ? and isDeleted = 0`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	var u entity.User
	for rows.Next() {
		err = rows.Scan(&u.UserId, &u.Email, &u.Name, &u.Gender, &u.Phone, &u.CreatedAt, &u.UpdatedAt)
	}
	if err != nil {
		return nil, err
	}
	return &u, nil
}

// Search User by query.
func (r *UserRepo) Search(query string) ([]*entity.User, error) {
	stmt, err := r.db.Prepare(`select id,email,name,gender,phone,createdAt, updatedAt from user where name like ? and isDeleted = 0`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query("%" + query + "%")
	if err != nil {
		return nil, err
	}
	var result []*entity.User
	for rows.Next() {
		var u entity.User
		rows.Scan(&u.UserId, &u.Email, &u.Name, &u.Gender, &u.Phone, &u.CreatedAt, &u.UpdatedAt)
		result = append(result, &u)
	}
	return result, nil
}

// List get all user
func (r *UserRepo) List() ([]*entity.User, error) {
	stmt, err := r.db.Prepare(`select id,email,name,gender,phone,createdAt, updatedAt from user where isDeleted = 0`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	var result []*entity.User
	for rows.Next() {
		var u entity.User
		rows.Scan(&u.UserId, &u.Email, &u.Name, &u.Gender, &u.Phone, &u.CreatedAt, &u.UpdatedAt)
		result = append(result, &u)
	}
	return result, nil
}

//Auth get user by email
func (r *UserRepo) Auth(mail string)(*entity.User, error){
	stmt, err := r.db.Prepare(`select id,email,password,name,gender,phone,role,createdAt, updatedAt from user where email = ? and isDeleted = 0`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(mail)
	if err != nil {
		return nil, err
	}
	var u entity.User
	for rows.Next() {
		err = rows.Scan(&u.UserId, &u.Email,&u.Password, &u.Name, &u.Gender, &u.Phone,&u.Role, &u.CreatedAt, &u.UpdatedAt)
	}
	if err != nil {
		return nil, err
	}
	return &u, nil
}
// Create a User.
func (r *UserRepo) Create(e *entity.User) (entity.ID, error) {
	stmt, err := r.db.Prepare(`
		insert into user (id, email, password, name, gender, phone, role, createdAt, updatedAt, isDeleted) 
		values(?,?,?,?,?,?,?,?,?,?)`)
	if err != nil {
		return e.UserId, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(
		e.UserId,
		e.Email,
		e.Password,
		e.Name,
		e.Gender,
		e.Phone,
		e.Role,
		e.CreatedAt,
		e.UpdatedAt,
		false,
	)
	if err != nil {
		return e.UserId, err
	}
	return e.UserId, nil
}

// Update a User.
func (r *UserRepo) Update(e *entity.User) error {
	stmt, err := r.db.Prepare(`
		update user set 
		email = ?, 
		password = ?, 
		name = ?, 
		gender = ?, 
		phone = ?, 
		updatedAt = ?
		where id = ? and isDeleted = 0`)
	if err != nil {
		return err
	}
	passw, err := entity.GeneratePassword(e.Password)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(
		e.Email,
		passw,
		e.Name,
		e.Gender,
		e.Phone,
		time.Now().Format(time.RFC3339),
		e.UserId,
	)
	if err != nil {
		return err
	}
	err = stmt.Close()
	if err != nil {
		return err
	}
	return nil
}

// Delete a User.
func (r *UserRepo) Delete(id entity.ID) error {
	_, err := r.db.Exec("update user set isDeleted = true where id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
