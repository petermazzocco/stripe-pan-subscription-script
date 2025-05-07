package internal

import (
	"fmt"
	"time"
)

func ParseDate(dateStr string) (time.Time, error) {
	formats := []string{
		"1/2/06 15:04",
		"1/2/2006 3:04:05 PM",
		"1/2/2006 15:04:05",
	}

	for _, format := range formats {
		if t, err := time.Parse(format, dateStr); err == nil {
			month := t.Month()
			day := t.Day()
			year := t.Year()

			validatedDate := time.Date(year, month, day, t.Hour(), t.Minute(), t.Second(), 0, t.Location())

			if validatedDate.Month() == month && validatedDate.Day() == day {
				return t, nil
			} else {
				return time.Time{}, fmt.Errorf("invalid date: %s (month %d has no day %d)", dateStr, month, day)
			}
		}
	}

	return time.Time{}, fmt.Errorf("could not parse date: %s", dateStr)
}

func CalculateNextBillingDate(purchaseDate time.Time) time.Time {
	currentTime := time.Now()

	purchaseDay := purchaseDate.Day()

	year := currentTime.Year()
	month := currentTime.Month()

	targetDate := time.Date(year, month, purchaseDay, 0, 0, 0, 0, time.Local)

	if targetDate.Before(currentTime) {
		if month == time.December {
			year++
			month = time.January
		} else {
			month++
		}
	}

	daysInMonth := time.Date(year, month+1, 0, 0, 0, 0, 0, time.Local).Day()

	if purchaseDay > daysInMonth {
		purchaseDay = daysInMonth
	}

	billingDate := time.Date(year, month, purchaseDay, 0, 0, 0, 0, time.Local)

	return billingDate
}
