package models

import "time"
type Bookings struct {
	Bookings_ID 	string 		`gorm:"primaryKey"`
	CarsID        	string    	`json:"cars_id"`
	UserID        	string    	`json:"user_id"`
	CustomerID		string    	`json:"customer_id"`
	StartDate     	time.Time 	`json:"start_date"`
	EndDate       	time.Time 	`json:"end_date"`
	TotalPrice    	float64   	`json:"total_price"`
	Status        	string    	`json:"status"`

	Car  Cars  `json:"car" gorm:"foreignKey:CarsID;references:id"`
	User Users `json:"user" gorm:"foreignKey:UserID;references:id"`
	Customer Customers `json:"customer" gorm:"foreignKey:CustomerID;references:id"`
	
}