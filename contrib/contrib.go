package contrib

import root "github.com/thepudds/nested-module-example"

const Name = "contrib"

func init() { println("init", Name, "root is:", root.Name) }
