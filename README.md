# leetcode-go

Solutions to LeetCode problems written in Go 1.25, with tests

## Project structure

```text
leetcode-go/
├── problems/                 # problem solutions
│   ├── pXXXX_<problem_name>/
│   │   ├── solution.go
│   │   ├── solution_test.go
│   │   └── notes.md
│   └── ...
├── internal/
│   └── ds/                   # shared data structures (ListNode, TreeNode, etc.)
├── go.mod
└── README.md
```

## Running tests
```bash
go test ./...
```