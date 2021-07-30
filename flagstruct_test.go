package flag

import (
	flag "github.com/spf13/pflag"
	"reflect"
	"testing"
)

type Building struct {
	Location    string   `flag:""`  // set defaults (i.e name=location, value="", usage="")
	Temperature int      `flag:"-"` // ignore this field
	IsOnFire    bool     `flag:"name=on-fire, value=false"`
	RoomNames   []string `flag:"value=Human Resources||Engineering||Administration"`
	Floors      []int    `flag:"usage=List of each floor number in this building"`
}

func TestFill(t *testing.T) {
	flagSet := flag.NewFlagSet("test", flag.ContinueOnError)
	building := Building{}

	Fill(flagSet, reflect.ValueOf(building), "")

	// test if all flags have been set
	if _, err := flagSet.GetString("location"); err != nil {
		t.Error("flag 'location' is not defined")
	}
	if _, err := flagSet.GetBool("on-fire"); err != nil {
		t.Error("flag 'on-fire' is not defined")
	}
	if _, err := flagSet.GetStringSlice("room-names"); err != nil {
		t.Error("flag 'room-names' is not defined")
	}
	if _, err := flagSet.GetIntSlice("floors"); err != nil {
		t.Error("flag 'floors' is not defined")
	}

	// test default values
	if location, _ := flagSet.GetString("location"); location != "" {
		t.Error("expected 'location' to be empty")
	}
	if onFire, _ := flagSet.GetBool("on-fire"); onFire != false {
		t.Error("expected 'on-fire' to be false")
	}
	if roomNames, _ := flagSet.GetStringSlice("room-names"); len(roomNames) != 3 {
		t.Error("expected 'room-names' to have 3 default values")
	}
	if floors, _ := flagSet.GetIntSlice("floors"); len(floors) != 0 {
		t.Error("expected 'floors' to have 0 default values")
	}

	// test ignored fields
	if _, err := flagSet.GetInt("temperature"); err == nil {
		t.Error("expected 'temperature' to be ignored")
	}
}

func TestExtract(t *testing.T) {
	flagSet := flag.NewFlagSet("test", flag.ContinueOnError)
	b := Building{}

	Fill(flagSet, reflect.ValueOf(Building{}), "")

	Extract(flagSet, reflect.ValueOf(&b).Elem(), "")

	if b.Location != "" {
		t.Error("invalid location")
	}
	if b.Temperature != 0 {
		t.Error("temperature should be 0")
	}
	if b.IsOnFire != false {
		t.Error("building should not be on fire")
	}
	if !reflect.DeepEqual(b.RoomNames, []string{"Human Resources", "Engineering", "Administration"}) {
		t.Error("invalid room names")
	}
	if len(b.Floors) != 0 {
		t.Error("invalid floors")
	}
}
