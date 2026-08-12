# Sudoku Solver

An interactive Sudoku solver with automatic difficulty categorization (Easy, Medium, Hard) and interactive solving capabilities, built in Go using the Ebitengine (Ebiten) game engine.

---

## 1. Native Desktop Mode

To build and run the application natively on your system (macOS/Linux/Windows):

### Prerequisites
* Go 1.25 or newer installed on your system.

### Build and Run
1. Run the application directly from the source:
   ```bash
   go run src/main.go
   ```
2. Or build a native binary:
   ```bash
   go build -o sudoku src/main.go
   ./sudoku
   ```

---

## 2. WebAssembly (WASM) Mode

To compile, host, and run the application in a web browser using WebAssembly:

### Prerequisites
* Go 1.25 or newer installed on your system.
* A local HTTP server utility (e.g., Python 3 or Go's `wasmserve`).

### Setup and Build
1. **Compile to WebAssembly**:
   Build the Go main entrypoint with target environment variables set to WebAssembly:
   ```bash
   env GOOS=js GOARCH=wasm go build -o sudoku.wasm ./src/main.go
   ```

2. **Copy the JS WASM bridge**:
   Copy the `wasm_exec.js` runtime bridge file matching your exact installed Go version to the root directory:
   ```bash
   cp $(go env GOROOT)/lib/wasm/wasm_exec.js .
   ```

3. **Serve index.html**:
   An `index.html` is provided at the root to load and execute the WASM application. Since browsers restrict local file access for WASM files (`file://` protocol), you must run a local HTTP server:
   
   * **Using Python 3**:
     ```bash
     python3 -m http.server 8080
     ```
     Navigate to [http://localhost:8080/](http://localhost:8080/) in your web browser.

   * **Using wasmserve**:
     ```bash
     go run github.com/hajimehoshi/wasmserve@latest ./src
     ```
     Navigate to [http://localhost:8080/](http://localhost:8080/) in your web browser.
