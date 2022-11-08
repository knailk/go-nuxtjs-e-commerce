package sqlite

import (
	"database/sql"

	"github.com/knailk/go-shopee/app/entity"
)

type AuthRepo struct {
	db *sql.DB
}



func (r *AuthRepo) SignIn(email string) (*entity.User, error) {
	stmt, err := r.db.Prepare(`select id,email,password,name,gender,phone,role,createdAt, updatedAt from user where email = ? and isDeleted = 0`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(email)
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

func (r *AuthRepo) SignUp(user *entity.User) (int64, error) {
	return 0,nil
}

func (r *AuthRepo) Logout(userID entity.ID) error {
	return nil
}
