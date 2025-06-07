// Package contrib demonstrates some behavior of how pkg.go.dev renders nested modules.
//
// Here is sample text:
//
// # Contrib Modules
//
// You can find the documentation for the contrib modules via the following links:
//   - [nested1/v2] - a sample v2 nested module.
//   - [nested2/v2] - a sample v2 nested module.
//   - [nested3/v2] - a sample v2 nested module.
//
// (Or with whatever wording or [formatting] is desired. Note that pkg.go.dev does not
// automatically render links for these related modules below in a "Directories" section,
// but we have the control to add links here if we want.)
//
// [nested1/v2]: https://pkg.go.dev/github.com/thepudds/nested-module-example/contrib/nested1/v2
// [nested2/v2]: https://pkg.go.dev/github.com/thepudds/nested-module-example/contrib/nested2/v2
// [nested3/v2]: https://pkg.go.dev/github.com/thepudds/nested-module-example/contrib/nested3/v2
// [formatting]: https://go.dev/doc/comment#syntax
package contrib

import root "github.com/thepudds/nested-module-example/v2"

const Name = "contrib"

func init() { println("init", Name, "root is:", root.Name) }
