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
// Use a data table (Map) if you see a long mapping.
// Don't change the behavior.
//
// Run with: bun run index.js

const validPromoCodes = new Set(["STUDENT", "EARLY", "GROUP5"]);

const calculateTicketPrice = (ticket) => {
  let result = "";
  if (ticket.quantity > 0) {
    if (ticket.basePrice > 0) {
      let multiplier = 1.0;
      if (ticket.seatTier === "floor") {
        multiplier = 2.5;
      } else if (ticket.seatTier === "lower") {
        multiplier = 1.75;
      } else if (ticket.seatTier === "upper") {
        multiplier = 1.25;
      } else if (ticket.seatTier === "balcony") {
        multiplier = 1.0;
      } else {
        return "Error: invalid seat tier";
      }

      let seatPrice = Math.floor(ticket.basePrice * multiplier);

      if (ticket.isVIP) {
        seatPrice += 5000; // $50 VIP surcharge
      }

      let total = seatPrice * ticket.quantity;

      if (ticket.promoCode !== "") {
        if (validPromoCodes.has(ticket.promoCode)) {
          if (ticket.promoCode === "STUDENT") {
            total = total - Math.floor(total * 15 / 100);
          } else if (ticket.promoCode === "EARLY") {
            total = total - Math.floor(total * 10 / 100);
          } else if (ticket.promoCode === "GROUP5") {
            if (ticket.quantity >= 5) {
              total = total - Math.floor(total * 20 / 100);
            } else {
              return "Error: GROUP5 requires 5+ tickets";
            }
          }
        } else {
          return "Error: invalid promo code";
        }
      }

      const vipLabel = ticket.isVIP ? " VIP" : "";
      result = `${ticket.eventName} | ${ticket.seatTier}${vipLabel} x${ticket.quantity} = $${Math.floor(total / 100)}.${String(total % 100).padStart(2, "0")}`;
    } else {
      result = "Error: price must be positive";
    }
  } else {
    result = "Error: quantity must be positive";
  }
  return result;
};

const tickets = [
  { eventName: "Rock Show", basePrice: 7500, seatTier: "floor", isVIP: true, quantity: 2, promoCode: "" },
  { eventName: "Jazz Night", basePrice: 5000, seatTier: "balcony", isVIP: false, quantity: 4, promoCode: "STUDENT" },
  { eventName: "Comedy Hour", basePrice: 3000, seatTier: "upper", isVIP: false, quantity: 6, promoCode: "GROUP5" },
  { eventName: "Opera", basePrice: 10000, seatTier: "lower", isVIP: false, quantity: 1, promoCode: "EARLY" },
  { eventName: "Bad Event", basePrice: 5000, seatTier: "mezzanine", isVIP: false, quantity: 2, promoCode: "" },
  { eventName: "Free Event", basePrice: 0, seatTier: "floor", isVIP: false, quantity: 1, promoCode: "" },
];

for (const ticket of tickets) {
  console.log(calculateTicketPrice(ticket));
}
