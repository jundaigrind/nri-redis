package config

import (
	"errors"
)

var (
	errWrongInputOrPasswordItemType = errors.New("the item type should be password or input")
	errWrongDropDownItemType        = errors.New("the item type should be dropdown")
)

type Item interface {
	Label() string
	Default() string
	Comment() string
	Type() string
}

type SimpleItem struct {
	label    string
	defaults string
	comment  string
	itemType itemType
}

type InputOrPasswordItem struct {
	Item
	length int
}

type DropDownItem struct {
	Item
	options []string
	initial int
}

func NewSimpleItem(label, defaults, comment string, itemType itemType) SimpleItem {
	return SimpleItem{
		label:    label,
		defaults: defaults,
		comment:  comment,
		itemType: itemType,
	}
}

func (i SimpleItem) Label() string {
	return i.label
}

func (i SimpleItem) Default() string {
	return i.defaults
}

func (i SimpleItem) Comment() string {
	return i.comment
}

func (i SimpleItem) Type() string {
	return i.itemType.Value()
}

func NewInputOrPasswordItem(label, defaults, comment string, itemType itemType, length int) (InputOrPasswordItem, error) {
	if itemType != PasswordField && itemType != InputField {
		return InputOrPasswordItem{}, errWrongInputOrPasswordItemType
	}
	item := NewSimpleItem(label, defaults, comment, itemType)

	return InputOrPasswordItem{
		Item:   item,
		length: length,
	}, nil
}

func (d InputOrPasswordItem) Length() int {
	return d.length
}

func NewDropDownItem(label, defaults, comment string, itemType itemType, options []string) (DropDownItem, error) {
	var initial int
	if itemType != DropDown {
		return DropDownItem{}, errWrongDropDownItemType
	}
	for i, item := range options {
		if item == defaults {
			initial = i
		}
	}
	item := NewSimpleItem(label, defaults, comment, itemType)

	return DropDownItem{
		Item:    item,
		options: options,
		initial: initial,
	}, nil
}

func (d DropDownItem) Options() []string {
	return d.options
}

func (d DropDownItem) Initial() int {
	return d.initial
}
