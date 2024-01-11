package usersRepositories

import (
	"fmt"

	"github.com/Japanisnmm/GoBackend101/modules/users"
	"github.com/Japanisnmm/GoBackend101/modules/users/usersPatterns"
	"github.com/jmoiron/sqlx"
)

type IUsersRepository interface {
	InsertUser(req *users.UserRegisterReq, isAdmin bool) (*users.UserPassport, error)
	FindOneUserByEmail(email string) (*users.UserCredentialCheck,error)
}

type usersRepository struct {
	db *sqlx.DB
}

func UsersRepository  (db *sqlx.DB )  IUsersRepository {
	return &usersRepository{
		db : db,
	}
}

func (r *usersRepository) InsertUser (req *users.UserRegisterReq, isAdmin bool) (*users.UserPassport, error) {
	result := usersPatterns.InsertUser(r.db, req, isAdmin)

	var err error
	if isAdmin {
        result, err = result.Admin()
		if err != nil {
			return nil,err
		}
	} else {
        result, err = result.Customer()
		if err != nil {
			return nil,err
		}
	
	
	}
	//get result from insert
	user ,err := result.Result()
	if err != nil {
		return nil,err
	}
	return  user,nil

}

func (r *usersRepository) FindOneUserByEmail(email string) (*users.UserCredentialCheck, error) {
	query := `
	SELECT
		"id",
		"email",
		"password",
		"username",
		"role_id"
	FROM "users"
	WHERE "email" = $1;`

	user := new(users.UserCredentialCheck)
	if err := r.db.Get(user, query, email); err != nil {
		return nil, fmt.Errorf("user not found")
	}
	return user, nil
}