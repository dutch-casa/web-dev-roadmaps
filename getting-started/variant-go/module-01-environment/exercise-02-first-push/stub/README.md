# Exercise: First push

Create a GitHub repo, write a Go program, and push it. By the end of this exercise you'll have a live repo on GitHub with working code in it.

## Steps

1. **Create a new directory and initialize a Go module:**
   ```
   mkdir hello-go
   cd hello-go
   go mod init hello-go
   ```

2. **Create `main.go`** with a program that prints your name and the current date. Use the `time` package to get the date — don't hardcode it.

3. **Run it** to make sure it works:
   ```
   go run .
   ```

4. **Initialize a git repo and make your first commit:**
   ```
   git init
   git add go.mod main.go
   git commit -m "first commit: hello-go prints name and date"
   ```

5. **Create a GitHub repo and push:**
   ```
   gh repo create hello-go --public --source=. --push
   ```

6. **Verify** — open the URL that `gh` prints. Your code should be on GitHub.

## What to check

- [ ] `go run .` prints your name and today's date
- [ ] The GitHub repo exists and contains `go.mod` and `main.go`
- [ ] The commit message describes what the code does, not what you did ("first commit" is fine for the very first one)
