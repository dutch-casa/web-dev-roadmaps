// Data tables replace branching.
const tierMultipliers = new Map([
  ["floor", 2.5],
  ["lower", 1.75],
  ["upper", 1.25],
  ["balcony", 1.0],
]);

const promoDiscountPercents = new Map([
  ["STUDENT", 15],
  ["EARLY", 10],
  ["GROUP5", 20],
]);

const calculateTicketPrice = (ticket) => {
  // Guard clauses — reject bad input early.
  if (ticket.quantity <= 0) {
    return "Error: quantity must be positive";
  }
  if (ticket.basePrice <= 0) {
    return "Error: price must be positive";
  }
  const multiplier = tierMultipliers.get(ticket.seatTier);
  if (multiplier === undefined) {
    return "Error: invalid seat tier";
  }

  // Data preparation — compute the per-seat price.
  let seatPrice = Math.floor(ticket.basePrice * multiplier);
  if (ticket.isVIP) {
    seatPrice += 5000;
  }
  let total = seatPrice * ticket.quantity;

  // Apply promo code if present.
  if (ticket.promoCode !== "") {
    const discountPercent = promoDiscountPercents.get(ticket.promoCode);
    if (discountPercent === undefined) {
      return "Error: invalid promo code";
    }
    const isGroupPromoWithTooFew = ticket.promoCode === "GROUP5" && ticket.quantity < 5;
    if (isGroupPromoWithTooFew) {
      return "Error: GROUP5 requires 5+ tickets";
    }
    total -= Math.floor(total * discountPercent / 100);
  }

  // Format the result.
  const vipLabel = ticket.isVIP ? " VIP" : "";
  return `${ticket.eventName} | ${ticket.seatTier}${vipLabel} x${ticket.quantity} = $${Math.floor(total / 100)}.${String(total % 100).padStart(2, "0")}`;
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
