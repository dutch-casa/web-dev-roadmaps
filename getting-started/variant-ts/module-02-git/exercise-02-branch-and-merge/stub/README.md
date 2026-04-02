# Exercise: Branch and merge

You'll create a feature branch, make changes on it, and merge it back — including dealing with a conflict.

## Setup

Use the same `archaeology-repo` from Exercise 01 (or run `setup.sh` again if you deleted it).

## Part 1: Clean merge

1. Create and switch to a new branch:
   ```
   git switch -c add-discount
   ```

2. Edit `inventory.js` — add a `discount` field to each item object. Set `discount: 0` for all items except Notebook, which gets `discount: 0.10` (10% off).

3. Commit your changes:
   ```
   git add .
   git commit -m "add discount field to inventory items"
   ```

4. Switch back to main and merge:
   ```
   git switch main
   git merge add-discount
   ```

This should merge cleanly because nobody changed main while you were on your branch.

## Part 2: Merge conflict

1. Stay on main. Edit the greeting in `index.js` — change `"customer"` to `"shopper"`:
   ```js
   console.log(greet("shopper"));
   ```

2. Commit:
   ```
   git add .
   git commit -m "change greeting to shopper"
   ```

3. Create another branch:
   ```
   git switch -c new-greeting
   ```

4. Change the same line to use `"friend"` instead:
   ```js
   console.log(greet("friend"));
   ```

5. Commit:
   ```
   git add .
   git commit -m "change greeting to friend"
   ```

6. Switch to main and try to merge:
   ```
   git switch main
   git merge new-greeting
   ```

7. Git will report a conflict. Open `index.js` — you'll see the conflict markers. Pick the version you want (or combine them), remove the markers, then:
   ```
   git add index.js
   git commit -m "resolve greeting conflict: keep friend"
   ```

## What to check

- [ ] `git log --oneline --graph` shows the branch history
- [ ] The discount field exists in `inventory.js`
- [ ] The conflict is resolved and the repo is clean (`git status` shows nothing)
