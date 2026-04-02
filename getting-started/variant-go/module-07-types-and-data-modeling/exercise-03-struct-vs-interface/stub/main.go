package main

import "fmt"

// Struct vs. interface
//
// You're building a notification system that sends messages through
// different channels: email, SMS, and push notification.
//
// Each channel has different data requirements:
//   - Email needs: to address, subject, body
//   - SMS needs: phone number, message (max 160 chars)
//   - Push needs: device token, title, body, badge count
//
// Part 1: Implement this using ONLY structs — one struct per channel,
// and a function for each that formats the notification as a string.
// No interfaces.
//
// Part 2: Implement the same thing using an interface called Notification
// with a method Send() string. Each channel type implements the interface.
// Then write a function that takes a []Notification and sends all of them.
//
// Part 3: In a comment at the bottom, answer:
//   - When is the struct-only approach better?
//   - When is the interface approach better?
//   - What would happen if you needed to add a new channel (e.g., Slack)?
//     How much code changes in each approach?

// --- Part 1: Struct-only approach ---
// TODO: Define EmailNotification, SMSNotification, PushNotification structs
// TODO: Write formatEmail, formatSMS, formatPush functions

// --- Part 2: Interface approach ---
// TODO: Define the Notification interface with a Send() string method
// TODO: Make each channel type implement Notification
// TODO: Write sendAll(notifications []Notification) that sends each one

func main() {
	fmt.Println("TODO: implement both approaches and compare them")
}
