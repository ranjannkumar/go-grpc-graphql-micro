package main

import "github.com/99designs/gqlgen/graphql"

type Server struct {
	// 	accountClient   *account.Client
	// 	catalogClient   *catalog.Client
	// 	orderClient     *order.Client
}

func NewGraphQLServer(accounturl, catalogurl, orderurl string) (*Server, error) {
// 	accountClient, err := account.NewClient(accounturl)
// 	if err != nil {
// 		return nil, err
// 	}
// 	catalogClient, err := catalog.NewClient(catalogurl)
// 	if err != nil {
// 		accountClient.Close()
// 		return nil, err
// 	}
// 	orderClient, err := order.NewClient(orderurl)
// 	if err != nil {
// 		accountClient.Close()
// 		catalogClient.Close()
// 		return nil, err
// 	}

	return &Server{
// 		accountClient,
// 		catalogClient,
// 		orderClient,
	}, nil
}

// func (s *Server) Mutation() MutationResolver {
// 	return &mutationResolver{
// 		server: s,
// 	}
// }

// func (s *Server) Query() QueryResolver {
// 	return &queryResolver{
// 		server: s,
// 	}
// }

// func (s *Server) Account() AccountResolver {
// 	return &accountResolver{
// 		server: s,
// 	}
// }

func (s *Server) ToExecutableSchema() graphql.ExecutableSchema{
	return NewExecutableSchema(Config{
		Resolvers: s,
	})
}