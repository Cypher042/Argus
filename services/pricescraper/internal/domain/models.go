package domain

import "time"

type AmazonProduct struct {
	ProductName string
	Currency    string    
	Price     float64   
	Source    string    
	Timestamp time.Time 
	URL string
}
