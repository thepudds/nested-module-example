# nested-module-example
Simple example of nested Go modules, including incrementing to v2

### Overview

This is a small demonstration repo related to [golang/go#74016](https://github.com/golang/go/issues/74016).

The latest contrib package documentation as rendered by pkg.go.dev can be viewed at:
* https://pkg.go.dev/github.com/thepudds/nested-module-example/v2/contrib

### Repo layout

The current basic repo structure (ignoring some auxillary files like go.sum):

```
.
├── go.mod               // top-level module's go.mod file
├── root.go
└── contrib
    ├── contrib.go
    ├── nested1
    │   ├── go.mod       // a nested module's go.mod file
    │   └── nested1.go
    ├── nested2
    │   ├── go.mod       // a nested module's go.mod file
    │   └── nested2.go
    └── nested3
        ├── go.mod       // a nested module's go.mod file
        └── nested3.go
```

### History

1. There was a single v0.9.0 module in this repo
2. Then it was split into 4 total v1.0.0 modules in this repo:
   * The top-level module:
      * `module github.com/thepudds/nested-module-example`
      * This was tagged `v1.0.0`
   * Three nested modules:
      * `module github.com/thepudds/nested-module-example/contrib/nested1`
      * `module github.com/thepudds/nested-module-example/contrib/nested2`
      * `module github.com/thepudds/nested-module-example/contrib/nested3`
      * These were tagged `contrib/nested1/v1.0.0` and similar for nested2 and nested3
3. Then the 4 modules were moved to v2.0.0
   * The top-level module became:
      * `module github.com/thepudds/nested-module-example/v2`
      * This was tagged `v2.0.0`
   * The three nested modules became:
      * `module github.com/thepudds/nested-module-example/contrib/nested1/v2`
      * `module github.com/thepudds/nested-module-example/contrib/nested2/v2`
      * `module github.com/thepudds/nested-module-example/contrib/nested3/v2`
      * These were tagged `contrib/nested1/v2.0.0` and similar for nested2 and nested3
4. Then the top-level module was moved to v2.0.2 with some example package doc links added to the [contrib](https://pkg.go.dev/github.com/thepudds/nested-module-example/v2/contrib) package.
