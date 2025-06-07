package nested3

import root "github.com/thepudds/nested-module-example/v2"

const Name = "nested3"

func init() { println("init", Name, "root is:", root.Name) }
