package contrib

import root "github.com/thepudds/nested-module-example/v2"

const Name = "contrib"

func init() { println("init", Name, "root is:", root.Name) }
