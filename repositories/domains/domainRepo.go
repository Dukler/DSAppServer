package domains

import (
	"DSAppServer/api"
	"DSAppServer/dbh"
	"github.com/jmoiron/sqlx"
)

type DomainRepo interface {
	Create(episode *api.User) (*api.User, error)
	GetByID(id string) (*api.User, error)
	GetByEmailOrName(s string) (*api.User, error)
	GetTrending(limit int) ([]api.User, error)
}

const (
	domainInsert = `
  INSERT INTO domains (name, app_id)
  VALUES (:name, :app_id)
  RETURNING name, app_id
  `
	domainSelectBase = `
	select 	domains.*,
			app.name app_name
	from domains
	left join app on domains.app_id = app.id 
  `
	domainSelectByID	= domainSelectBase + `WHERE domains.id = $1`
	domainSelectByName	= domainSelectBase + `WHERE domains.name = $1`
)

type psqlDomainRepo struct {
	create           *sqlx.NamedStmt
	selectByID       *sqlx.Stmt
	selectByName       *sqlx.Stmt
}

func NewDomainRepo () (*psqlDomainRepo, error){
	repo := new(psqlDomainRepo)
	db := dbh.GetDB()
	var err error
	repo.create, err  = db.PrepareNamed(domainInsert)
	repo.selectByID, err  = db.Preparex(domainSelectByID)
	repo.selectByName, err  = db.Preparex(domainSelectByName)

	return repo, err
}

func (p *psqlDomainRepo) Create(input *api.User) (*api.User, error) {
	// implementation goes here
	//_, err := p.create.Exec(input)
	//return input, err
	var err error
	if _,err =  p.create.Exec(input); err != nil {
		return nil, err
	}
	return input, err
}

func (p *psqlDomainRepo) GetByID(id string) (*api.Domain, error) {
	// implementation goes here
	return p.getOne(p.selectByID,id)
}

func (p *psqlDomainRepo) GetByName(name string) (*api.Domain, error) {
	// implementation goes here
	return p.getOne(p.selectByName,name)
}

func (p *psqlDomainRepo) getOne(stmt *sqlx.Stmt, args ...interface{}) (*api.Domain, error) {
	var domain api.Domain

	if err :=  stmt.Get(&domain, args...); err != nil {
		return nil, err
	}

	// do whatever computation you've gotta do

	return &domain, nil
}