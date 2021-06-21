package config

import (
	"reflect"
	"testing"

	sdkArgs "github.com/newrelic/infra-integrations-sdk/args"
)

func TestGetItemsFromArgs(t *testing.T) {
	type argumentList struct {
		sdkArgs.DefaultArgumentList
		Hostname       string       `default:"localhost" help:"lorem ipsum." config:"input;20"`
		Port           int          `default:"6379" help:"lorem ipsum." config:"input;5"`
		UnixSocketPath string       `default:"" help:"lorem ipsum." config:"input;20"`
		Keys           sdkArgs.JSON `default:"" help:"lorem ipsum." config:"input;100"`
		Mode           string       `default:"modeA" help:"lorem ipsum." config:"dropdown;modeA,modeB"`
		Password       string       `help:"lorem ipsum." config:"password;20"`
		UseUnixSocket  bool         `default:"false" help:"lorem ipsum." config:"checkbox"`
	}
	var args argumentList

	expected := []Item{
		InputOrPasswordItem{
			Item:   SimpleItem{label: "hostname", defaults: "localhost", comment: "lorem ipsum.", itemType: InputField},
			length: 20,
		},
		InputOrPasswordItem{
			Item:   SimpleItem{label: "port", defaults: "6379", comment: "lorem ipsum.", itemType: InputField},
			length: 5,
		},
		InputOrPasswordItem{
			Item:   SimpleItem{label: "unix_socket_path", defaults: "", comment: "lorem ipsum.", itemType: InputField},
			length: 20,
		},
		InputOrPasswordItem{
			Item:   SimpleItem{label: "keys", defaults: "", comment: "lorem ipsum.", itemType: InputField},
			length: 100,
		},
		DropDownItem{
			Item:    SimpleItem{label: "mode", defaults: "modeA", comment: "lorem ipsum.", itemType: DropDown},
			options: []string{"modeA", "modeB"},
		},
		InputOrPasswordItem{
			Item:   SimpleItem{label: "password", defaults: "", comment: "lorem ipsum.", itemType: PasswordField},
			length: 20,
		},
		SimpleItem{label: "use_unix_socket", defaults: "false", comment: "lorem ipsum.", itemType: CheckBox},
	}

	got, err := GetItemsFromArgs(&args)
	if err != nil {
		t.Errorf("expected args list: %+v not equal to result: %+v", expected, args)
	}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("expected args list: %+v not equal to result: %+v", expected, got)
	}
}
