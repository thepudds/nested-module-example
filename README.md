# nested-module-example
Simple example of nested Go modules, including incrementing to v2

### History

This is a small demonstration repo related to [golang/go#74016](https://github.com/golang/go/issues/74016).

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
   * The three nested modules will become:
      * `module github.com/thepudds/nested-module-example/contrib/nested1/v2`
      * `module github.com/thepudds/nested-module-example/contrib/nested2/v2`
      * `module github.com/thepudds/nested-module-example/contrib/nested3/v2`
      * These will be tagged `contrib/nested1/v2.0.0` and similar for nested2 and nested3
