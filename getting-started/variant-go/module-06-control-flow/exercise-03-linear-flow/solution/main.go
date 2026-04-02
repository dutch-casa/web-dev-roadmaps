package main

import "fmt"

type Ticket struct {
	EventName string
	BasePrice int
	SeatTier  string
	IsVIP     bool
	Quantity  int
	PromoCode string
}

// Data tables replace branching.
var tierMultipliers = map[string]float64{
	"floor":   2.5,
	"lower":   1.75,
	"upper":   1.25,
	"balcony": 1.0,
}

var promoDiscountPercents = map[string]int{
	"STUDENT": 15,
	"EARLY":   10,
	"GROUP5":  20,
}

func calculateTicketPrice(t Ticket) string {
	// Guard clauses — reject bad input early.
	if t.Quantity <= 0 {
		return "Error: quantity must be positive"
	}
	if t.BasePrice <= 0 {
		return "Error: price must be positive"
	}
	multiplier, validTier := tierMultipliers[t.SeatTier]
	if !validTier {
		return "Error: invalid seat tier"
	}

	// Data preparation — compute the per-seat price.
	seatPrice := int(float64(t.BasePrice) * multiplier)
	if t.IsVIP {
		seatPrice += 5000
	}
	total := seatPrice * t.Quantity

	// Apply promo code if present.
	if t.PromoCode != "" {
		discountPercent, validPromo := promoDiscountPercents[t.PromoCode]
		if !validPromo {
			return "Error: invalid promo code"
		}
		isGroupPromoWithTooFew := t.PromoCode == "GROUP5" && t.Quantity < 5
		if isGroupPromoWithTooFew {
			return "Error: GROUP5 requires 5+ tickets"
		}
		total -= total * discountPercent / 100
	}

	// Format the result.
	vipLabel := ""
	if t.IsVIP {
		vipLabel = " VIP"
	}
	return fmt.Sprintf("%s | %s%s x%d = $%d.%02d",
		t.EventName, t.SeatTier, vipLabel, t.Quantity, total/100, total%100)
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
