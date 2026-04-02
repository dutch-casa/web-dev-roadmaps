# Commit archaeology

Run `setup.sh` first to create the sample repo. Then `cd archaeology-repo` and answer these questions using only git commands.

Write your answers below each question.

## Questions

**1. How many commits are in this repo?**

Command to use: `git log --oneline`

Answer:

**2. What was the commit message for the third commit?**

Answer:

**3. What file was added in the third commit that didn't exist before?**

Command to use: `git show <commit-hash> --stat`

Answer:

**4. What changed between the second and fourth commits in `main.go`?**

Command to use: `git diff <hash1> <hash2> -- main.go`

Answer:

**5. Which commit added the `listItems` function?**

Command to use: `git log -p -- main.go` (search through the patches)

Answer:

**6. What item was added in the most recent commit?**

Command to use: `git show HEAD`

Answer:
