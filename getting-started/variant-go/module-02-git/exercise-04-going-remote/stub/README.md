# Exercise: Going remote

Work with a real GitHub repo. You'll fork, branch, push, and open a pull request.

## Setup

You need a GitHub account and the `gh` CLI authenticated:
```
gh auth login
```

## Steps

1. **Create a new repo for this exercise:**
   ```
   mkdir git-remote-practice
   cd git-remote-practice
   go mod init git-remote-practice
   ```

2. **Write a small program** — anything. A temperature converter, a tip calculator, whatever. Keep it under 30 lines. Commit it:
   ```
   git init
   git add .
   git commit -m "initial commit: <describe what your program does>"
   ```

3. **Push to GitHub:**
   ```
   gh repo create git-remote-practice --public --source=. --push
   ```

4. **Create a feature branch and add something:**
   ```
   git switch -c add-feature
   ```
   Add a feature to your program (another function, a different output mode, whatever). Commit it.

5. **Push the branch:**
   ```
   git push -u origin add-feature
   ```

6. **Open a pull request:**
   ```
   gh pr create --title "Add <your feature>" --body "Describe what this adds and why."
   ```

7. **Review your own PR on GitHub.** Look at the diff. Read the description. Is it clear what changed?

8. **Merge the PR:**
   ```
   gh pr merge --merge
   ```

9. **Pull the merge into your local main:**
   ```
   git switch main
   git pull
   ```

## What to check

- [ ] Your repo is on GitHub with at least 2 commits
- [ ] A merged PR exists in the repo's PR list
- [ ] `git log --oneline` on main shows the merge
- [ ] You can explain: what is a pull request and why would a team use them?

## Pair exercise (if at a club meeting)

Pair up with another member. Fork their repo, create a branch, add something (a new function, a comment, a test), push, and open a PR to *their* repo. They review and merge it. Then swap.

This is how real collaboration works.
