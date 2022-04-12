**NO LONGER MAINTAINED deprecated in favor of [go-flags](https://github.com/jessevdk/go-flags)**

# flagstruct

Write command line flags within your struct by utilizing tags like a pro.

Works standalone with [pflag](https://github.com/spf13/pflag) or with [cobra](https://github.com/spf13/cobra).

## Usage
```go
package main

import (
	flagstruct "github.com/sleeyax/flagstruct"
	flag "github.com/spf13/pflag"
	"reflect"
	"fmt"
)

// 1. Define flags on your struct using the `flag:"[name=x, value=y, usage=z]"` tag
type Pet struct {
	Name string `flag:"value=frank, usage=Name of the pet"`
	RealName string `flag:"-"` // hide this field
	MuchFluff bool `flag:"name=is-fluffy, value=false, usage=Mark pet as fluffy"`
}

func main() {
	// 2. Set the flags by parsing the struct tags
	flags := flag.NewFlagSet("pet-example", flag.ContinueOnError)
	flagstruct.Fill(flags, reflect.ValueOf(Pet{}), "") // (last param is a prefix)
	
	// 3. At this point you should do something with the flags.
	//    Flags are available to the user in the following format:
	//    $ ./example -h
	//      Global Flags:
	//          --name      Name of the pet
	//          --is-fluffy Mark pet as fluffy (default false)
	
	// 4. Parse the flags back into a struct
	var pet Pet
	flagstruct.Extract(flags, reflect.ValueOf(&pet).Elem(), "")
	
	// 5. Finally, 'pet' has the values specified through flags.
	fmt.Println(pet.Name) // name specified by --name or default 'Frank'
}
```
See the included [example](./example) using [cobra](https://github.com/spf13/cobra) for a full and more detailed example.
