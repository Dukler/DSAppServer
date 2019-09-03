package users

import (
	"DSAppServer/api"
	"DSAppServer/dbh"
	"github.com/jmoiron/sqlx"
)

type UserRepo interface {
	Create(episode *api.User) (*api.User, error)
	GetByID(id string) (*api.User, error)
	GetByEmailOrName(s string) (*api.User, error)
	GetTrending(limit int) ([]api.User, error)
}

const (
	userInsert = `
  INSERT INTO users (email, username, password)
  VALUES (:email, :username, :password)
  RETURNING email, username, password
  `
	userSelectBase = `
  SELECT *
  FROM users
  `
	userSelectByID	= userSelectBase + `WHERE id = $1`
	userSelectByEmail = userSelectBase + `WHERE email = $1`
)

type psqlUserRepo struct {
	create           *sqlx.NamedStmt
	selectByID       *sqlx.Stmt
	selectByEmail 	*sqlx.Stmt
	selectByToken   *sqlx.Stmt
}

func NewUserRepo () (*psqlUserRepo, error){
	repo := new(psqlUserRepo)
	db := dbh.GetDB()
	var err error
	repo.create, err  = db.PrepareNamed(userInsert)
	repo.selectByEmail, err  = db.Preparex(userSelectByEmail)
	repo.selectByID, err  = db.Preparex(userSelectByID)

	return repo, err
}

func (p *psqlUserRepo) Create(input *api.User) (*api.User, error) {
	var err error
	if _,err =  p.create.Exec(input); err != nil {
		return nil, err
	}
	return input, err
}

func (p *psqlUserRepo) GetByID(id string) (*api.User, error) {
	// implementation goes here
	return p.getOne(p.selectByID,id)
}

func (p *psqlUserRepo) GetByEmail(s string) (*api.User, error) {
	// implementation goes here
	return p.getOne(p.selectByEmail,s)
}

func (p *psqlUserRepo) getOne(stmt *sqlx.Stmt, args ...interface{}) (*api.User, error) {
	var user api.User

	if err :=  stmt.Get(&user, args...); err != nil {
		return nil, err
	}

	// do whatever computation you've gotta do

	return &user, nil
}