// ============================================================
// Part 1: Type-only approach
// ============================================================

type EmailNotification = {
  to: string;
  subject: string;
  body: string;
};

type SMSNotification = {
  phone: string;
  message: string;
};

type PushNotification = {
  deviceToken: string;
  title: string;
  body: string;
  badge: number;
};

const formatEmail = (n: EmailNotification): string =>
  `[EMAIL] To: ${n.to} | Subject: ${n.subject} | Body: ${n.body}`;

const formatSMS = (n: SMSNotification): string => {
  const msg = n.message.length > 160
    ? n.message.slice(0, 157) + "..."
    : n.message;
  return `[SMS] To: ${n.phone} | Message: ${msg}`;
};

const formatPush = (n: PushNotification): string =>
  `[PUSH] Device: ${n.deviceToken} | ${n.title}: ${n.body} (badge: ${n.badge})`;

// ============================================================
// Part 2: Interface approach
// ============================================================

// Channel is the interface. Any object with a send(): string
// method can be treated as a channel. This lets us write code
// that works with *any* channel type without knowing the specifics.
interface Channel {
  send(): string;
}

// Each channel type implements the interface by defining send().

class EmailChannel implements Channel {
  constructor(
    readonly to: string,
    readonly subject: string,
    readonly body: string,
  ) {}

  send(): string {
    return formatEmail(this);
  }
}

class SMSChannel implements Channel {
  constructor(
    readonly phone: string,
    readonly message: string,
  ) {}

  send(): string {
    return formatSMS(this);
  }
}

class PushChannel implements Channel {
  constructor(
    readonly deviceToken: string,
    readonly title: string,
    readonly body: string,
    readonly badge: number,
  ) {}

  send(): string {
    return formatPush(this);
  }
}

// sendAll works with any array of Notifications.
// It doesn't know or care whether they're emails, SMS, or push.
// If you add a new channel tomorrow, this function doesn't change.
const sendAll = (notifications: Channel[]): void => {
  for (const n of notifications) {
    console.log(n.send());
  }
};

// --- Main ---

console.log("=== Part 1: Type-only ===");
console.log(formatEmail({
  to: "student@auburn.edu", subject: "Welcome", body: "You're in!",
}));
console.log(formatSMS({
  phone: "555-0123", message: "Your code is 847291",
}));
console.log(formatPush({
  deviceToken: "dev_abc123", title: "New message", body: "Hey there", badge: 3,
}));

console.log();
console.log("=== Part 2: Interface approach ===");
const notifications: Channel[] = [
  new EmailChannel("student@auburn.edu", "Reminder", "Meeting at 3pm"),
  new SMSChannel("555-0456", "Class canceled"),
  new PushChannel("dev_xyz789", "Grade posted", "Check your grades", 1),
];
sendAll(notifications);

// Part 3: Discussion
//
// When is the type-only approach better?
//   When you have a small, fixed number of types and you always know
//   exactly which type you're working with. No indirection, no class
//   overhead, easy to follow. If you have 2-3 notification types and
//   they're always handled separately, standalone functions are simpler.
//
// When is the interface approach better?
//   When you need to treat different types uniformly. sendAll() works
//   with any notification channel — it doesn't need a switch statement
//   or type guard. If you're building a system where new channels
//   get added over time, the interface lets you extend without modifying
//   existing code.
//
// What happens when you add a new channel (e.g., Slack)?
//   Type-only: you add a new type, a new format function, and you
//   update every place that handles notifications (add a new case to
//   every switch/if-else chain).
//
//   Interface: you add a new class that implements send(). That's it.
//   sendAll() and any other code that works with Channel[] doesn't
//   change at all.
//
// The tradeoff: interfaces add a level of indirection. When you read
// sendAll(), you can't see what send() actually does without looking
// at the implementing classes. That's fine when the abstraction is
// stable and the implementations are many. It's overhead when the
// abstraction is shaky and you only have two types.
