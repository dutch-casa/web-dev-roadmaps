package main

import "fmt"

// ============================================================
// Part 1: Struct-only approach
// ============================================================

type EmailNotification struct {
	To      string
	Subject string
	Body    string
}

type SMSNotification struct {
	Phone   string
	Message string
}

type PushNotification struct {
	DeviceToken string
	Title       string
	Body        string
	Badge       int
}

func formatEmail(n EmailNotification) string {
	return fmt.Sprintf("[EMAIL] To: %s | Subject: %s | Body: %s", n.To, n.Subject, n.Body)
}

func formatSMS(n SMSNotification) string {
	msg := n.Message
	if len(msg) > 160 {
		msg = msg[:157] + "..."
	}
	return fmt.Sprintf("[SMS] To: %s | Message: %s", n.Phone, msg)
}

func formatPush(n PushNotification) string {
	return fmt.Sprintf("[PUSH] Device: %s | %s: %s (badge: %d)",
		n.DeviceToken, n.Title, n.Body, n.Badge)
}

// ============================================================
// Part 2: Interface approach
// ============================================================

// Notification is the interface. Any type with a Send() string method
// can be treated as a notification. This lets us write code that works
// with *any* notification channel without knowing the specific type.
type Notification interface {
	Send() string
}

// Each channel type implements the interface by defining Send().
func (n EmailNotification) Send() string {
	return formatEmail(n)
}

func (n SMSNotification) Send() string {
	return formatSMS(n)
}

func (n PushNotification) Send() string {
	return formatPush(n)
}

// sendAll works with any slice of Notifications.
// It doesn't know or care whether they're emails, SMS, or push.
// If you add a new channel tomorrow, this function doesn't change.
func sendAll(notifications []Notification) {
	for _, n := range notifications {
		fmt.Println(n.Send())
	}
}

func main() {
	fmt.Println("=== Part 1: Struct-only ===")
	fmt.Println(formatEmail(EmailNotification{
		To: "student@auburn.edu", Subject: "Welcome", Body: "You're in!",
	}))
	fmt.Println(formatSMS(SMSNotification{
		Phone: "555-0123", Message: "Your code is 847291",
	}))
	fmt.Println(formatPush(PushNotification{
		DeviceToken: "dev_abc123", Title: "New message", Body: "Hey there", Badge: 3,
	}))

	fmt.Println()
	fmt.Println("=== Part 2: Interface approach ===")
	notifications := []Notification{
		EmailNotification{To: "student@auburn.edu", Subject: "Reminder", Body: "Meeting at 3pm"},
		SMSNotification{Phone: "555-0456", Message: "Class canceled"},
		PushNotification{DeviceToken: "dev_xyz789", Title: "Grade posted", Body: "Check your grades", Badge: 1},
	}
	sendAll(notifications)
}

// Part 3: Discussion
//
// When is the struct-only approach better?
//   When you have a small, fixed number of types and you always know
//   exactly which type you're working with. No indirection, no interface
//   overhead, easy to follow. If you have 2-3 notification types and
//   they're always handled separately, structs are simpler.
//
// When is the interface approach better?
//   When you need to treat different types uniformly. sendAll() works
//   with any notification channel — it doesn't need a switch statement
//   or type assertion. If you're building a system where new channels
//   get added over time, the interface lets you extend without modifying
//   existing code.
//
// What happens when you add a new channel (e.g., Slack)?
//   Struct-only: you add a new struct, a new format function, and you
//   update every place that handles notifications (add a new case to
//   every switch/if-else).
//
//   Interface: you add a new struct that implements Send(). That's it.
//   sendAll() and any other code that works with []Notification doesn't
//   change at all.
//
// The tradeoff: interfaces add a level of indirection. When you read
// sendAll(), you can't see what Send() actually does without looking
// at the implementing types. That's fine when the abstraction is stable
// and the implementations are many. It's overhead when the abstraction
// is shaky and you only have two types.
