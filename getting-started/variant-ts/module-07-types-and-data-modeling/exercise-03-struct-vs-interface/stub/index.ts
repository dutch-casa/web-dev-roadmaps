// Struct vs. interface (Type vs. interface in TypeScript)
//
// You're building a notification system that sends messages through
// different channels: email, SMS, and push notification.
//
// Each channel has different data requirements:
//   - Email needs: to address, subject, body
//   - SMS needs: phone number, message (max 160 chars)
//   - Push needs: device token, title, body, badge count
//
// Part 1: Implement this using ONLY types — one type per channel,
// and a standalone function for each that formats the notification
// as a string. No interface.
//
// Part 2: Implement the same thing using an interface called Channel
// with a method send(): string. Each channel type implements the interface.
// Then write a function that takes a Channel[] and sends all of them.
//
// Part 3: In a comment at the bottom, answer:
//   - When is the type-only approach better?
//   - When is the interface approach better?
//   - What would happen if you needed to add a new channel (e.g., Slack)?
//     How much code changes in each approach?

// --- Part 1: Type-only approach ---
// TODO: Define EmailNotification, SMSNotification, PushNotification types
// TODO: Write formatEmail, formatSMS, formatPush functions

// --- Part 2: Interface approach ---
// TODO: Define the Channel interface with a send(): string method
// TODO: Make each channel type implement Channel (as a class)
// TODO: Write sendAll(notifications: Channel[]) that sends each one

console.log("TODO: implement both approaches and compare them");
