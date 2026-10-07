# Purpose

This repo is a GO Cli tool used to access my leetcode solutions.

This is so I can practice setting up a Golang Cli tool while also practicing my Leetcode problems

## Getting started

Tooling is managed with [mise](https://mise.jdx.dev). It pins Go, golangci-lint, hk, pkl and make to the versions in [mise.toml](mise.toml).

```bash
mise install
```

This installs the tools and then sets up the git hooks (`hk install --mise` runs as a post-install hook).

## Using the CLI

Build the binary into `build/` and run it:

```bash
make build
./build/leetcode-go list      # list all registered solutions
./build/leetcode-go run 1     # run the solution with LeetCode problem ID 1
```

Or skip the build step with `go run`:

```bash
go run . list
go run . run 1
```

## Project layout

```
main.go                 entry point, calls cmd.Execute()
cmd/root.go             CLI commands (list, run)
solutions/registry.go   Solution type and the registry the CLI reads from
solutions/<topic>/<problem>/
    solution.go         the solution, registered in init()
    solution_test.go    test cases for the solution
```

Solutions are grouped by topic, e.g. `solutions/hashtable/twosum`.

## Adding a solution

1. Create a package under `solutions/<topic>/<problem>/`.
2. In `solution.go`, register it from an `init()` function:

   ```go
   func init() {
       solutions.RegisterSolution(solutions.Solution{
           ID:         1,            // the LeetCode problem number
           Title:      "Two Sum",
           Difficulty: "Easy",
           Run:        func() { /* call your solution here */ },
       })
   }
   ```

3. Add a blank import for the package in `main.go` so its `init()` runs and the CLI can see it:

   ```go
   import _ "github.com/DrBeehre/Leetcode-go/solutions/hashtable/twosum"
   ```

4. Add table-driven tests in `solution_test.go`.

## Development

The [Makefile](Makefile) wraps the common tasks:

| Command       | What it does                                          |
| ------------- | ----------------------------------------------------- |
| `make`        | Runs everything below in order (fmt → build)          |
| `make fmt`    | Formats code with gofumpt + goimports                 |
| `make lint`   | Runs golangci-lint (config in `.golangci.yml`)        |
| `make critic` | Runs only the gocritic checks                         |
| `make test`   | Runs tests with the race detector                     |
| `make build`  | Builds the binary to `build/leetcode-go`              |
| `make clean`  | Deletes `build/`                                      |

### Git hooks

[hk](https://hk.jdx.dev) runs checks on every commit (config in [hk.pkl](hk.pkl)):

1. format staged Go files (fixes are re-staged automatically)
2. lint
3. test
4. build

Run the same checks by hand with:

```bash
hk check --all
```
