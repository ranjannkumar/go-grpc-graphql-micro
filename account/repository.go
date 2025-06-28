package account

import (
	"database/sql"
)
// skip (for pagination offset), and take (for pagination limit)
type Repository interface {
	Close()
	PutAccount(ctx context.Context,a Account)error
	GetAccountById(ctx context.Context,id string)(*Account,err)
	ListAccounts(ctx context.Context,skip uint64,take uint64)([]Account,err)
}

type postgresRespository struct {
	db *sql.DB
}

func NewPostgresRepository(url string)(Repository,error){
	db,err := sql.Open("postgres",url)
	if err!=nil{
		return  nil,err
	}

	err = db.Ping()
	if err!=nil{
		return nil,err
	}
	return &postgresRespository{db},nil
}

