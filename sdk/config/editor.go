package config

import (
	"strconv"

	"github.com/rivo/tview"
)

type Editor struct {
	configItems []Item
}

func NewEditor(configItems []Item) Editor {
	return Editor{
		configItems: configItems,
	}
}

func (e Editor) Run() {
	app := tview.NewApplication()
	form := tview.NewForm()
	for _, c := range e.configItems {
		switch c.Type() {
		case DropDown:
			form.AddDropDown(c.Label(), c.(DropDownItem).Options(), c.(DropDownItem).Initial(), nil)
		case InputField:
			form.AddInputField(c.Label(), c.Default(), c.(InputOrPasswordItem).Length(), nil, nil)
		case PasswordField:
			form.AddPasswordField(c.Label(), c.Default(), c.(InputOrPasswordItem).Length(), '*', nil)
		case CheckBox:
			d, _ := strconv.ParseBool(c.Default())
			form.AddCheckbox(c.Label(), d, nil)
		default:
		}
	}

	form.AddButton("Save", nil).
		AddButton("Quit", func() {
			app.Stop()
		})

	form.SetBorder(true).SetTitle("Enter some data").SetTitleAlign(tview.AlignLeft)
	if err := app.SetRoot(form, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
