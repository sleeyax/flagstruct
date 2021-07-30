package main

import (
	"encoding/json"
	flagstruct "github.com/sleeyax/flagstruct"
	"github.com/spf13/cobra"
	"log"
	"reflect"
)

type Hobby struct {
	Name  string `flag:"value=Football, usage=Short name of the hobby"`
	IsFun bool   `flag:"value=true, usage=Mark this hobby as a fun one"`
}

type Person struct {
	FirstName           string   `flag:"value=david, usage=Your first name"`
	LastName            string   `flag:"value=beckham, usage=Your last name"`
	Height              int      `flag:""`
	FavoriteConsumables []string `flag:"name=favorite-foods, usage=List of favorite foods"`
	UnluckyNumbers      []int    `flag:"value=13||6, usage=List of your most unlucky numbers"`
	Hobby               Hobby
	IgnoreMe            bool `flag:"-"`
}

func newCmdCreatePerson() *cobra.Command {
	return &cobra.Command{
		Use:   "create",
		Short: "creates a person",
		Run: func(cmd *cobra.Command, args []string) {
			var p Person
			flagstruct.Extract(cmd.Flags(), reflect.ValueOf(&p).Elem(), "")

			// print object as JSON to stdout
			log.Println("Would create the following object:")
			data, err := json.MarshalIndent(p, "", "\t")
			if err != nil {
				log.Fatal(err)
			}
			log.Println(string(data))
		},
	}
}

func main() {
	cmdRoot := &cobra.Command{
		Use:   "example",
		Short: "flagstruct example - see https://github.com/sleeyax/flagstruct",
	}
	cmdRoot.PersistentFlags().SortFlags = false
	cmdRoot.Flags().SortFlags = false

	// Populate cmdRoot flags with values defined through struct tags.
	// All possible flags will be rendered once the user uses the '-h' or '--help' flag
	flagstruct.Fill(cmdRoot.PersistentFlags(), reflect.ValueOf(Person{}), "")

	// Register an example 'create person' command.
	// In a realistic scenario this command would store the object somewhere.
	// But for demonstration purposes we'll just print the resulting struct as JSON to stdout.
	cmdRoot.AddCommand(newCmdCreatePerson())

	cmdRoot.Execute()
}
