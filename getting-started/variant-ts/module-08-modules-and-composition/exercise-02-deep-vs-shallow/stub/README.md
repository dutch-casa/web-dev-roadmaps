# Exercise: Deep vs. shallow modules

This exercise doesn't have code to write. It's a reading and analysis exercise.

## Setup

Look at the two implementations in the `option_a/` and `option_b/` directories. Both implement the same thing: a simple key-value cache with expiration.

## Your task

For each implementation, count:

1. **How many exported names** does the module have? (types, functions, classes, methods)
2. **How many concepts** must a caller understand to use the module correctly?
3. **How many lines of code** does the caller need to write to perform a basic get-set-expire workflow?

Then answer:

- Which implementation is **deeper** (simpler interface, more hidden complexity)?
- Which is **shallower** (more interface surface relative to functionality)?
- If you needed to change the internal storage from a Map to a database, which implementation requires changing caller code?
- Which one would you want to maintain?

Write your answers in `answers.md`.
