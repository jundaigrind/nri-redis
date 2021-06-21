package config

import (
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/newrelic/infra-integrations-sdk/log"
)

var camel = regexp.MustCompile("(^[^A-Z]*|[A-Z]*)([A-Z][^A-Z]+|$)")

func underscore(s string) string {
	var a []string
	for _, sub := range camel.FindAllStringSubmatch(s, -1) {
		if sub[1] != "" {
			a = append(a, sub[1])
		}
		if sub[2] != "" {
			a = append(a, sub[2])
		}
	}
	return strings.ToLower(strings.Join(a, "_"))
}

func GetItemsFromArgs(args interface{}) ([]Item, error) {
	var items []Item
	val := reflect.ValueOf(args).Elem()

	for i := 0; i < val.NumField(); i++ {
		typeField := val.Type().Field(i)
		tag := typeField.Tag

		// The labelValue will take the field's name in underscore
		labelValue := underscore(typeField.Name)
		defaultValue := tag.Get("default")
		helpValue := tag.Get("help")
		configValue := tag.Get("config")

		if configValue != "" {
			c := strings.Split(configValue, ";")
			itemType, err := newItemType(c[0])
			if err != nil {
				log.Warn("argument %s with config tag but wrong item type", labelValue)
				continue
			}

			switch itemType {
			case DropDown:
				item, err := NewDropDownItem(labelValue, defaultValue, helpValue, itemType, strings.Split(c[1], ","))
				if err != nil {
					return nil, err
				}
				items = append(items, item)
			case InputField, PasswordField:
				length, err := strconv.Atoi(c[1])
				if err != nil {
					return nil, err
				}
				item, err := NewInputOrPasswordItem(labelValue, defaultValue, helpValue, itemType, length)
				if err != nil {
					return nil, err
				}
				items = append(items, item)
			default:
				items = append(items, NewSimpleItem(labelValue, defaultValue, helpValue, itemType))
			}
		}
	}
	return items, nil
}
