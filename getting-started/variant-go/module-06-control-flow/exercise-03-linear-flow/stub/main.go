package main

import "fmt"

// Linear flow
//
// This function has tangled control flow — the reader has to jump around
// to understand what happens. Rewrite it so it reads top to bottom:
//
//   1. Guard clauses at the top (reject bad input)
//   2. Data preparation in the middle
//   3. Result computation at the bottom
//
// Use named conditions where they help readability.
// Use a data table if you see a long mapping.
// Don't change the behavior.

type Ticket struct {
	EventName  string
	BasePrice  int // cents
	SeatTier   string // "floor", "lower", "upper", "balcony"
	IsVIP      bool
	Quantity   int
	PromoCode  string
}

var validPromoCodes = map[string]bool{
	"STUDENT": true,
	"EARLY":   true,
	"GROUP5":  true,
}

func calculateTicketPrice(t Ticket) string {
	result := ""
	if t.Quantity > 0 {
		if t.BasePrice > 0 {
			multiplier := 1.0
			if t.SeatTier == "floor" {
				multiplier = 2.5
			} else if t.SeatTier == "lower" {
				multiplier = 1.75
			} else if t.SeatTier == "upper" {
				multiplier = 1.25
			} else if t.SeatTier == "balcony" {
				multiplier = 1.0
			} else {
				return "Error: invalid seat tier"
			}

			seatPrice := int(float64(t.BasePrice) * multiplier)

			if t.IsVIP {
				seatPrice += 5000 // $50 VIP surcharge
			}

			total := seatPrice * t.Quantity

			if t.PromoCode != "" {
				if validPromoCodes[t.PromoCode] {
					if t.PromoCode == "STUDENT" {
						total = total - total*15/100
					} else if t.PromoCode == "EARLY" {
						total = total - total*10/100
					} else if t.PromoCode == "GROUP5" {
						if t.Quantity >= 5 {
							total = total - total*20/100
						} else {
							return "Error: GROUP5 requires 5+ tickets"
						}
					}
				} else {
					return "Error: invalid promo code"
				}
			}

			result = fmt.Sprintf("%s | %s%s x%d = $%d.%02d",
				t.EventName, t.SeatTier,
				func() string {
					if t.IsVIP {
						return " VIP"
					}
					return ""
				}(),
				t.Quantity, total/100, total%100)
		} else {
			result = "Error: price must be positive"
		}
	} else {
		result = "Error: quantity must be positive"
	}
	return result
}

func main() {
	tickets := []Ticket{
		{EventName: "Rock Show", BasePrice: 7500, SeatTier: "floor", IsVIP: true, Quantity: 2, PromoCode: ""},
		{EventName: "Jazz Night", BasePrice: 5000, SeatTier: "balcony", Quantity: 4, PromoCode: "STUDENT"},
		{EventName: "Comedy Hour", BasePrice: 3000, SeatTier: "upper", Quantity: 6, PromoCode: "GROUP5"},
		{EventName: "Opera", BasePrice: 10000, SeatTier: "lower", Quantity: 1, PromoCode: "EARLY"},
		{EventName: "Bad Event", BasePrice: 5000, SeatTier: "mezzanine", Quantity: 2},
		{EventName: "Free Event", BasePrice: 0, SeatTier: "floor", Quantity: 1},
	}

	for _, t := range tickets {
		fmt.Println(calculateTicketPrice(t))
	}
}
