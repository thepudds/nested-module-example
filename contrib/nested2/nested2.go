package nested2

import root "github.com/thepudds/nested-module-example/v2"

const Name = "nested2"

func init() { println("init", Name, "root is:", root.Name) }
