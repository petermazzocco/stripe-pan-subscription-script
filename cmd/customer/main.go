package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/customer"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("failed to load env")
		return
	}

	// Change to PROD_TEST_KEY when ready for production customer ID's
	stripe.Key = os.Getenv("TEST_STRIPE_KEY")

	// Add your list of emails here
	emails := []string{"youremails@here.com"}

	params := &stripe.CustomerListParams{}

	for _, email := range emails {
		params.Filters.AddFilter("email", "", email)
		iter := customer.List(params)

		for iter.Next() {
			c := iter.Current()
			customer, ok := c.(*stripe.Customer)
			if !ok {
				fmt.Fprintf(os.Stderr, "Error: Could not convert to customer type\n")
				continue
			}
			fmt.Printf("Email: %s, Customer ID: %s\n", email, customer.ID)
			return
		}
		if err := iter.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "Error listing customer: %v\n", err)
			return
		}
		fmt.Println("No customer found with that email.")
	}
}
