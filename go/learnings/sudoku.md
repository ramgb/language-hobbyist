# Sudoku Solver Learnings & Bitmask Optimization

This document captures the design decisions, mathematical principles, and performance learnings from optimizing the Go Sudoku solver.

---

## 1. The Core Performance Bottleneck

The original backtracking solver represented cell candidate sets using `map[int]bool` for each of the 81 cells:
*   **Heap Allocations**: In Go, maps are pointers to heap-allocated hash tables. Modifying or re-creating maps during recursive backtracking generated millions of short-lived objects on the heap, triggering continuous Garbage Collection (GC) pauses.
*   **Performance Overhead**: Map operations (lookups, deletions, and insertions) require key hashing and bucket searches. When repeated millions of times during depth-first searches, this overhead degrades execution times.

---

## 2. The Solution: Bitmask Optimization

We replaced the `map[int]bool` grid with a stack-allocated grid of 16-bit integers (`[9][9]uint16`). Because each cell has at most 9 possible candidates (digits 1 to 9), we can represent the presence or absence of a candidate using individual bits.

### The Bitwise Left-Shift Operator (`1 << val`)
To map a digit $N \in [1, 9]$ to a specific bit position, we left-shift the binary representation of `1` by $N$ bits:

*   `1 << 1` $\rightarrow$ `0b0000000000000010` (denotes that **1** is a candidate)
*   `1 << 2` $\rightarrow$ `0b0000000000000100` (denotes that **2** is a candidate)
*   `1 << 9` $\rightarrow$ `0b0000001000000000` (denotes that **9** is a candidate)

Using this mapping, the set of all possible candidates $\{1, 2, 3, 4, 5, 6, 7, 8, 9\}$ is represented by the integer `0x3FE` (binary `0b0000001111111110`), where bits 1 through 9 are all set to `1`.

---

## 3. Key Bitwise Math Operations in Code

The following standard bitwise operators were used to manipulate candidate sets:

### A. Initializing Candidates
All cells are initially assumed to have all digits 1-9 as potential candidates:
```go
var guesses [9][9]uint16
for i := 0; i < 9; i++ {
    for j := 0; j < 9; j++ {
        guesses[i][j] = 0x3FE // 0b0000001111111110
    }
}
```

### B. Checking If a Candidate is Present
To check if digit `val` is a viable candidate for cell `(x, y)`:
```go
if (guesses[x][y] & (1 << val)) != 0 {
    // val is a candidate
}
```
*The bitwise AND (`&`) returns a non-zero value if and only if the bit at position `val` is set to `1`.*

### C. Setting a Single Determined Value
When assigning a digit `val` to cell `(x, y)` (either during initialization or search step):
```go
guesses[x][y] = 1 << val
```
*This clears all other bits, indicating that only `val` is possible.*

### D. Eliminating/Clearing a Candidate
To eliminate digit `val` from a neighbor's candidates during constraint propagation:
```go
guesses[nx][ny] &= ^(1 << val)
```
*The bitwise NOT (`^`) inverts the single-bit mask (e.g. turning `0b0100` into `0b1011`). The AND-ASSIGN (`&=`) then sets the bit at position `val` to `0` while preserving all other bits.*

### E. Extracting the Solved Value
To find which digit is set in a solved cell's mask:
```go
func getOnlyValue(mask uint16) int {
	for v := 1; v <= 9; v++ {
		if (mask & (1 << v)) != 0 {
			return v
		}
	}
	return -1
}
```

---

## 4. Trivial Backtracking State (Value Copying)

Because a `[9][9]uint16` array consumes exactly **162 bytes** of memory, we can copy the entire board state in a single instruction by passing it by value:

```go
nextGuesses := guesses
if setAndPropagate(x, y, val, &nextGuesses) {
    solved, result := s.solveInternal(x, y+1, nextGuesses)
    if solved {
        return true, result
    }
}
```

*   **No Allocation**: Assigning array structures copies them on the stack, which generates zero heap allocations.
*   **Implicit Backtracking**: If a path fails, the local `nextGuesses` is simply discarded when the function returns. The parent stack frame remains completely untouched, eliminating the need to maintain an undo log or write manual restoration code.

---

## 5. Performance Improvements Realized

Running benchmark suites comparing the two implementations yielded the following results:

| Metric | Original Map-based Solver | Optimized Bitmask Solver | Improvement |
| :--- | :--- | :--- | :--- |
| **Solving Time** | ~293 microseconds | ~11.3 microseconds | **25x Speedup** |
| **Heap Allocations** | Thousands of allocs/op | **0 B/op (0 allocs/op)** | **100% Allocation-free** |
