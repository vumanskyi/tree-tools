A simple tool to display the tree structure of your folder
==========================================================

Display the tree structure with fiels:
```
go run main.go . -f
```

Result:
```
├───lorem
│       └───lorem.txt(9b)
└───project
│       └───test.txt(20b)
```

Display the tree structure only with folders (without files)

```
go run main.go .
```

Result:
```
├───lorem
└───project
```

Run test 
```
go test -v
```

Expected result:
```
=== RUN   TestTreeFull
--- PASS: TestTreeFull (0.00s)
=== RUN   TestTreeDir
--- PASS: TestTreeDir (0.00s)
PASS
ok      _/Users/vladumanskyi/go/src/conc        0.089s
```
