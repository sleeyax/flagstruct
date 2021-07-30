package flag

import (
	flag "github.com/spf13/pflag"
	"reflect"
	"strconv"
	"strings"
)

// Fill sets command flags by parsing the field tags of provided struct.
// Any fields that don't have a `flag:"name=x, value=y, usage=z"` tag will be skipped.
func Fill(flagSet *flag.FlagSet, v reflect.Value, prefix string) {
	for i := 0; i < v.NumField(); i++ {
		// get the current field's type
		fieldType := v.Type().Field(i)

		// recursively go over embedded structs
		if field := v.Field(i); field.Kind() == reflect.Struct {
			Fill(flagSet, field, strings.ToLower(fieldType.Name)+"-")
			continue
		}

		// get the tag value
		tag := fieldType.Tag.Get(tagName)

		// skip if tag is explicitly ignored
		if tag == "-" {
			continue
		}

		// convert tag string into its structural representation
		f := ParseTag(tag)

		// dynamically set name based on struct fieldType when unspecified
		if f.Name == "" {
			f.Name = ToFlag(fieldType.Name)
		}

		f.Name = prefix + f.Name

		// skip if flag is already defined
		// NOTE: pflag doesn't seem to support flag overrides...
		if flagSet.Lookup(f.Name) != nil {
			continue
		}

		// TODO: perhaps there's a better, faster way to populate the FlagSet in one go
		switch fieldType.Type.Kind() {
		case reflect.String:
			flagSet.String(f.Name, f.Value, f.Usage)
		case reflect.Bool:
			flagSet.Bool(f.Name, f.Value == "true", f.Usage)
		case reflect.Slice:
			var items []string
			if strings.Contains(f.Value, "||") {
				items = strings.Split(f.Value, "||")
				for j := range items {
					items[j] = strings.TrimSpace(items[j])
				}
			}
			switch fieldType.Type.Elem().Kind() {
			case reflect.String:
				flagSet.StringSlice(f.Name, items, f.Usage)
			case reflect.Int:
				integers := make([]int, len(items))
				for j := range items {
					integer, _ := strconv.Atoi(items[j])
					integers[j] = integer
				}
				flagSet.IntSlice(f.Name, integers, f.Usage)
			}
		}
	}
}

// Extract extracts flag values and populates all fields of given struct that match by `flag:` tags.
// This function basically does the inverse of Fill.
func Extract(flagSet *flag.FlagSet, v reflect.Value, prefix string) {
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := v.Type().Field(i)

		if field.Kind() == reflect.Struct {
			Extract(flagSet, field, strings.ToLower(fieldType.Name)+"-")
			continue
		}

		f := prefix + ToFlag(fieldType.Name)

		if !field.IsValid() || !field.CanSet() {
			continue
		}

		// TODO: perhaps there's a better, faster way to extract the FlagSet into the struct in one go
		switch fieldType.Type.Kind() {
		case reflect.String:
			if str, err := flagSet.GetString(f); err == nil {
				field.SetString(str)
			}
		case reflect.Bool:
			if b, err := flagSet.GetBool(f); err == nil {
				field.SetBool(b)
			}
		case reflect.Slice:
			switch fieldType.Type.Elem().Kind() {
			case reflect.String:
				if s, err := flagSet.GetStringSlice(f); err == nil {
					field.Set(reflect.ValueOf(s))
				}
			case reflect.Int:
				if s, err := flagSet.GetIntSlice(f); err == nil {
					field.Set(reflect.ValueOf(s))
				}
			}
		}
	}
}
