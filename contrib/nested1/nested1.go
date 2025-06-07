package nested1

import root "github.com/thepudds/nested-module-example/v2"

const Name = "nested1"

func init() { println("init", Name, "root is:", root.Name) }
