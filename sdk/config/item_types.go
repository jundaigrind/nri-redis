package config

import "errors"

const (
	DropDown      = "dropdown"
	InputField    = "input"
	CheckBox      = "checkbox"
	PasswordField = "password"
)

var itemTypes = []string{DropDown, InputField, CheckBox, PasswordField}

var errWrongItemType = errors.New("the config ItemType is not correct")

type itemType string

func newItemType(fieldType string) (itemType, error) {
	for _, it := range itemTypes {
		if fieldType == it {
			return itemType(fieldType), nil
		}
	}

	return "", errWrongItemType
}

func (i itemType) Value() string {
	return string(i)
}
