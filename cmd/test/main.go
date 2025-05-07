package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/petermazzocco/stripe-pan-sub-script/internal"
	"github.com/petermazzocco/stripe-pan-sub-script/models"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/subscription"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("failed to load env")
		return
	}

	// Change to PROD_STRIPE_KEY when ready for prod
	stripe.Key = os.Getenv("TEST_STRIPE_KEY")
	priceId := os.Getenv("PRICE_ID")

	// Change to your own .csv when ready for prod
	file, err := os.Open("test-customers.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)

	_, err = reader.Read()
	if err != nil {
		fmt.Println("Error reading header:", err)
		return
	}

	customerMap := make(map[string]models.Customer)

	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading records:", err)
		return
	}

	for _, record := range records {
		if len(record) < 5 {
			continue
		}

		purchaseDate, err := internal.ParseDate(record[0])
		if err != nil {
			fmt.Printf("Warning: Could not parse date '%s': %v\n", record[0], err)
			continue
		}

		customer := models.Customer{
			PurchaseDate: purchaseDate,
			Email:        record[1],
			CustomerID:   record[4],
		}

		customerMap[customer.Email] = customer
	}

	fmt.Printf("Loaded %d customers\n\n\n", len(customerMap))

	for email, customer := range customerMap {
		if email == "" {
			continue
		}
		fmt.Printf("Email: %s, PrevDate: %s, ID: %s\n",
			customer.Email,
			customer.PurchaseDate.Format("2006-01-02"),
			customer.CustomerID)

		nextBillingDate := internal.CalculateNextBillingDate(customer.PurchaseDate)
		fmt.Printf("Next billing date: %s\n", nextBillingDate.Format("2006-01-02"))

		params := &stripe.SubscriptionParams{
			Customer: stripe.String(customer.CustomerID),
			Items: []*stripe.SubscriptionItemsParams{
				{
					Price: stripe.String(priceId),
				},
			},
			BillingCycleAnchor: stripe.Int64(nextBillingDate.Unix()),
			ProrationBehavior:  stripe.String("none"),
		}

		r, err := subscription.New(params)
		if err != nil {
			log.Printf("  Error creating subscription: %v\n", err.Error())
			continue
		}
		fmt.Printf("  Subscription created: %s\n", r.ID)

		fmt.Printf(" Created subscription with Price ID: %s on %s\n\n\n",
			priceId,
			time.Unix(*params.BillingCycleAnchor, 0).Format("2006-01-02"))
	}

}
