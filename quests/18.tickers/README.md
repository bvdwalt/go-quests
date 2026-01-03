## Go Quest: Coordinated Tickers (`hello` & `world`)

### Objective

Use **multiple `time.Ticker`s and channels** to print messages at different fixed intervals, while ensuring the program terminates cleanly after a fixed duration.

---

### Requirements

1. Create **two tickers**:

   - `helloTicker`: ticks every **500 milliseconds**
   - `worldTicker`: ticks every **1000 milliseconds**

2. Behavior:

   - On every tick from `helloTicker`, print:

     ```
     hello
     ```

   - On every tick from `worldTicker`, print:

     ```
     world
     ```

3. Duration:

   - The program must run for **exactly 1 second**
   - After 1 second:

     - Both tickers must be **stopped**
     - The program must **exit gracefully** (no goroutine leaks)

4. Constraints:

   - Use `select` to listen to ticker channels
   - Do **not** use `time.Sleep` inside the goroutine loop
   - Use a `done` channel or equivalent signal to stop execution

---

### Expected Output (order may vary slightly due to scheduling)

Within ~3 second, you should observe:

```
hello
hello
world
hello
hello
world
hello
hello
world
```

Explanation:

- `hello` → at ~0.5s and ~1.0s
- `world` → at ~1.0s

---

### What This Quest Tests

- Understanding that **tickers emit values on channels**
- Correct use of `select` with **multiple time sources**
- Coordinating **concurrent periodic events**
- Proper **lifecycle management** of tickers and goroutines

---

## References

These explain the same ideas with examples:

- [https://gobyexample.com/tickers](https://gobyexample.com/tickers)

---

### Running the Tests

Once implemented, run:

```bash
go test ./quests/18.tickers -v
```
