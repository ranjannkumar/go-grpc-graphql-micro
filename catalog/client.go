package catalog

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ranjannkumar/go-grpc-grpahql-microservice/catalog/pb"
	"google.golang.org/grpc"
)

type Client struct {
	conn *grpc.ClientConn
	service pb.CatalogServiceClient
}

func NewClient(url string) (*Client, error) {
	var conn *grpc.ClientConn
	var err error

	maxRetries := 5
	retryInterval := 2 * time.Second
	totalTimeout := 15 * time.Second

	ctxRetry, cancelRetry := context.WithTimeout(context.Background(), totalTimeout)
	defer cancelRetry()

	for i := 0; i < maxRetries; i++ {
		log.Printf("Attempting to connect to Catalog service at %s (Attempt %d/%d)...", url, i+1, maxRetries)
		ctxDial, cancelDial := context.WithTimeout(context.Background(), retryInterval)
		defer cancelDial()

		conn, err = grpc.DialContext(ctxDial, url, grpc.WithInsecure(), grpc.WithBlock())
		if err == nil {
			log.Printf("Successfully connected to Catalog service at %s", url)
			c := pb.NewCatalogServiceClient(conn)
			return &Client{conn, c}, nil
		}

		log.Printf("Failed to connect to Catalog service: %v. Retrying in %v...", err, retryInterval)
		select {
		case <-time.After(retryInterval):
		case <-ctxRetry.Done():
			return nil, fmt.Errorf("failed to connect to Catalog service after multiple retries (total timeout reached): %w", ctxRetry.Err())
		}
	}

	return nil, fmt.Errorf("failed to connect to Catalog service after %d retries: %w", maxRetries, err)
}

func(c *Client) Close(){
	c.conn.Close()
}

func(c *Client) PostProduct(ctx context.Context,name,description string,price float64)(*Product,error){
	r,err := c.service.PostProduct(
		ctx,
		&pb.PostProductRequest{
			Name: name,
			Description: description,
			Price: price,
		},
	)
	if err !=nil{
		return nil,err
	}

	return &Product{
		ID: r.Product.Id,
		Name: r.Product.Name,
		Description: r.Product.Description,
		Price: r.Product.Price,
	},nil

}

func(c *Client) GetProduct(ctx context.Context,id string)(*Product,error){
	r,err:= c.service.GetProduct(
		ctx,
		&pb.GetProductRequest{
			Id:id,
		},
	)
	if err!=nil{
		return nil,err
	}

	return &Product{
		ID: r.Product.Id,
		Name: r.Product.Name,
		Description: r.Product.Description,
		Price: r.Product.Price,
	},nil
}

func(c *Client) GetProducts(ctx context.Context,skip uint64,take uint64,ids []string,query string)([]Product,error){
	r,err := c.service.GetProducts(
		ctx,
		&pb.GetProductsRequest{
			Ids: ids,
			Skip: skip,
			Take: take,
			Query: query,
		},
	)
	if err!=nil{
		return nil,err
	}
	products := []Product{}
	for _,p:= range r.Products{
		products=append(products,Product{
			ID: p.Id,
			Name: p.Name,
			Description: p.Description,
			Price: p.Price,
		})
	}

	return products,nil

}