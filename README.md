# Go Quests [![Star on GitHub](https://img.shields.io/github/stars/lite-quests/go-quests?style=social)](https://github.com/lite-quests/go-quests)


A beginner-friendly set of small Go exercises (“quests”). Each quest is intentionally small so you can practice one Go concept at a time.

You implement functions in `solution.go` and confirm correctness with unit tests in `solution_test.go`.

---

### Requirements

- **Golang** (Install from the official site: `https://go.dev/dl/`)
- **Visual Studio Code** (Install from the official site: `https://code.visualstudio.com/download`)
- **Git** (optional, but recommended: `https://git-scm.com/install/`)


Check your Go version: 

```sh
go version
```

---

### Get the repo (two ways)

**Option A: Clone with Git (recommended)**
1. Open Visual studio code.
2. File → Open Folder → Select the folder where you want the repository to reside.
2. Toolbar → Terminal → New Terminal

```sh
git clone https://github.com/lite-quests/go-quests.git
cd go-quests # This command helps you go into the quest directory from the terminal
```

**Option B: Download ZIP (no Git required)**
1. Open the repo on GitHub: `https://github.com/lite-quests/go-quests`
2. Click **Code** → **Download ZIP**
3. Unzip the file
4. Visual Studio Code → Open Folder → Select the extracted/unzipped folder
---

### Code structure

- `quests/` — exercises you edit (each quest is a Go package)

- `quests/<n.topic>/README.md` — instructions for that quest

- `quests/<n.topic>/solution.go` — your implementation (edit this)

- `quests/<n.topic>/solution_test.go` — tests (do not change)

- `solutions/` — reference implementations (use to compare after you try)

---

### How to solve a quest (intended workflow)
1. Open the repo folder in VS Code (Or any other IDE of your choice)
2. Pick a quest folder under `quests/` (start with smaller numbers first)
3. Read `quests/<n.topic>/README.md` fully. Go through the reference links/docs mentioned to get a thorough idea of the concept.
4. Open the solution file:
   - Example: `quests/10.structs/solution.go`
5. Implement the required functions/methods. 
6. Run the quest tests:

```sh
go test -v ./quests/10.structs
```

6. Repeat until all tests pass. If stuck, inspect the failing test output first.

7. Only after you’ve tried: compare with the reference in `solutions/`.

This repo is designed so that tests teach you the spec: make changes → run tests → refine → repeat.


---

### Tips that save beginners a lot of time

- **Don’t print unless the quest asks**  
  Tests usually check return values or exact stdout bytes.

- **Read test failures carefully**: they usually tell you the exact edge case you missed.

- **Avoid changing test files**
  If tests don’t pass, fix `solution.go` (the tests define the requirements).

- **Prefer the intended concept over “test hacks”**  
  You might find a workaround that passes the tests, but the goal is to practice the concept the quest is designed to teach (the quest README will usually hint what to use) so please stick to it.


---

### Where are the answers?

- `solutions/` contains reference implementations for each quest.
- Best practice: try first, then compare after you pass (or when you’re stuck).


---

### Troubleshooting

- **“package not found”**: make sure you’re running test command from the correct directory.
- **Tests fail on stdout**: check newline requirements (`fmt.Print` vs `fmt.Println`). 
- **Changes don’t seem to affect tests**
  Clear the test cache:
  ```sh
  go clean -testcache
---

### Support
If these quests helped you learn Go, please consider starring the repo.

---

### Contact

For any issues, contact either:
- `yadlamaniymn2005@gmail.com`
- `ananyapappula2004@gmail.com`