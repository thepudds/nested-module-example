package nested2

import root "github.com/thepudds/nested-module-example"

const Name = "nested2"

func init() { println("init", Name, "root is:", root.Name) }
