package account

import(
	 "context"
	 "github.com/segmentio/ksuid"
	)

type Service interface{
	PostAccount (ctx context.Context,name string)(*Account,error)
	GetAccount  (ctx context.Context,id string)(*Account,error)
	GetAccounts (ctx context.Context,skip uint64,take uint64)([]Account,error)
}

type Account struct{
	ID    string   `json:"id"`
	Name  string   `json:"name"`
}

// It holds a repository Repository field, which is an instance of the Repository interface (the data persistence layer).
//  This is an example of dependency injection, where the accountService depends on an abstraction (Repository) rather than a concrete implementation, making it more testable and flexible.
type accountService struct {
	repository Repository
}


// This is a constructor function for accountService
// This is where the Repository dependency is injected.
func newService(r Repository)Service{
	return &accountService{r}
}

func(s *accountService)	PostAccount (ctx context.Context,name string)(*Account,error){
	a := &Account{
		Name: name,
		ID: ksuid.New().String(),
	}
   if err := s.repository.PutAccount(ctx,*a);err!=nil{
		return nil,err
	 }

	 return a,nil
}

func(s *accountService)	GetAccount  (ctx context.Context,id string)(*Account,error){
  return s.repository.GetAccountById(ctx,id)
}

func(s *accountService)	GetAccounts (ctx context.Context,skip uint64,take uint64)([]Account,error){
	if take > 100 || (skip==0 && take ==0){
		take = 100
	}
  return  s.repository.ListAccounts(ctx,skip,take)
}