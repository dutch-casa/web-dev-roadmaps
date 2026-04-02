# Time travel — answers

**Which commit introduced the bug?**

The third commit: "refactor: clean up calc functions"

It changed `subtract` from `(a, b) => a - b` to `(a, b) => a + b`.

**Why is the message misleading?**

The message says "refactor: clean up calc functions" — which implies no behavior changed. But the commit actually broke `subtract`. This is why commit messages matter: if the message said "change subtract implementation," a reviewer might have caught it.

**The fix:**

```
git revert <hash-of-commit-3>
```

This creates a new commit that reverses the changes from commit 3, restoring `subtract` to `(a, b) => a - b`. The bad commit stays in history so you can see what happened. The revert commit explains the fix.
