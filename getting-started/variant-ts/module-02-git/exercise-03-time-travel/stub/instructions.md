# Exercise: Time travel

Run `setup.sh` first. Then `cd bisect-repo` and run the program:

```
bun run index.js
```

Notice that `10 - 4` is printing the wrong answer. Something broke the `subtract` function, but you don't know which commit did it.

## Part 1: Find the bug with git bisect

`git bisect` does a binary search through your commit history to find exactly which commit introduced a problem.

1. Start bisect and tell Git the current commit is bad:
   ```
   git bisect start
   git bisect bad
   ```

2. Tell Git the first commit was good (use `git log --oneline` to find its hash):
   ```
   git bisect good <first-commit-hash>
   ```

3. Git will check out a middle commit. Test it:
   ```
   bun run index.js
   ```
   If `10 - 4` gives the right answer, type `git bisect good`. If it's wrong, type `git bisect bad`.

4. Repeat until Git tells you which commit introduced the bug.

5. When done:
   ```
   git bisect reset
   ```

**If you get confused:** Run `git bisect reset` to start over. It won't hurt anything. You can restart the bisect as many times as you need.

**Write down:** Which commit message is on the bad commit? Why is the message misleading?

## Part 2: Undo the bug with git revert

`git revert` creates a new commit that undoes the changes from a specific commit. History is preserved — you can see the bug was introduced AND that it was fixed.

```
git revert <bad-commit-hash>
```

Git will open an editor for the revert message. The default message is fine.

Verify the fix:
```
bun run index.js
```

`10 - 4` should now print `6`.

## What to check

- [ ] You can explain what `git bisect` does and when it's useful
- [ ] The bug is fixed and the revert commit exists in `git log`
- [ ] The original bad commit is still in history (we didn't erase it)
