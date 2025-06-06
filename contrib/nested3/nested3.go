package nested1

import root "github.com/thepudds/nested-module-example"

const Name = "nested3"

func init() { println("init", Name, "root is:", root.Name) }
