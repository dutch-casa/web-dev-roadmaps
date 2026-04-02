# Exercise: First push

Create a GitHub repo, write a JavaScript program, and push it. By the end of this exercise you'll have a live repo on GitHub with working code in it.

## Steps

1. **Create a new directory and initialize a project:**
   ```
   mkdir hello-js
   cd hello-js
   bun init
   ```

2. **Edit `index.js`** (or `index.ts` — Bun creates one for you) with a program that prints your name and the current date. Use `new Date()` to get the date — don't hardcode it.

3. **Run it** to make sure it works:
   ```
   bun run index.js
   ```

4. **Initialize a git repo and make your first commit:**
   ```
   git init
   git add package.json index.js
   git commit -m "first commit: hello-js prints name and date"
   ```

5. **Create a GitHub repo and push:**
   ```
   gh repo create hello-js --public --source=. --push
   ```

6. **Verify** — open the URL that `gh` prints. Your code should be on GitHub.

## What to check

- [ ] `bun run index.js` prints your name and today's date
- [ ] The GitHub repo exists and contains `package.json` and `index.js`
- [ ] The commit message describes what the code does, not what you did ("first commit" is fine for the very first one)
