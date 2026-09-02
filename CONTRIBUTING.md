# Contributing Workflow: Branching & Merging

This document describes the standard workflow for pushing changes from a working branch into `main`.

## 1. Create a Branch

Always branch off the latest `main`:

```bash
git checkout main
git pull origin main
git checkout -b feature/short-description
```

**Naming convention:**
| Prefix | Use case |
|---|---|
| `feature/` | New functionality |
| `fix/` | Bug fixes |
| `chore/` | Maintenance, deps, config |
| `docs/` | Documentation only |

## 2. Commit Your Changes

- Keep commits small and focused on a single change.
- Write clear, descriptive commit messages (imperative mood recommended):

```bash
git add .
git commit -m "Add validation for login form inputs"
```

## 3. Push the Branch

```bash
git push origin feature/short-description
```

## 4. Open a Pull Request

- Go to the repository on GitHub and open a PR from your branch into `main`.
- **PR title:** concise summary of the change.
- **PR description should include:**
  - What changed and why
  - Any related issue number (e.g., `Closes #42`)
  - Screenshots/notes if relevant (UI changes, breaking changes, etc.)

## 5. Review & CI Checks

- Request at least one reviewer.
- Ensure all CI checks (tests, linting, build) pass before merge.
- Address review comments with additional commits (avoid force-pushing after review starts, unless agreed with reviewers).

## 6. Merge into Main

Once approved and checks pass, merge using the team's agreed strategy:

| Strategy | When to use |
|---|---|
| **Squash and merge** | Default — keeps `main` history clean, one commit per PR |
| **Merge commit** | When preserving full branch history matters |
| **Rebase and merge** | For linear history without a merge commit |

## 7. Clean Up

- Delete the feature branch after merging (GitHub can do this automatically via repo settings).
- Pull the latest `main` locally before starting your next branch.

```bash
git checkout main
git pull origin main
git branch -d feature/short-description
```

## Recommended Repo Settings

- **Branch protection on `main`:** disallow direct pushes; require PRs.
- **Required reviews:** at least 1 approval before merge.
- **Required status checks:** CI must pass before merge is allowed.
- **Auto-delete head branches:** enabled after merge.
