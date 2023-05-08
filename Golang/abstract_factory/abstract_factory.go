package abstract_factory

import "runtime"

// Here is an example of how to define the abstract factory interface:
type WidgetFactory interface {
	CreateWindow() Window
	CreateButton() Button
	CreateMenu() Menu
}

// We can then define concrete implementations of this interface, such as
type WindowsFactory struct{}

func (w *WindowsFactory) CreateWindow() Window {
	return &WinWindow{}
}
func (w *WindowsFactory) CreateButton() Button {
	return &WinButton{}
}
func (w *WindowsFactory) CreateMenu() Menu {
	return &WinMenu{}
}

type MacFactory struct{}

func (m *MacFactory) CreateWindow() Window {
	return &MacWindow{}
}
func (m *MacFactory) CreateButton() Button {
	return &MacButton{}
}
func (m *MacFactory) CreateMenu() Menu {
	return &MacMenu{}
}

// In this example, we define two concrete factories,
// WindowsFactory and MacFactory, which implement the WidgetFactory interface.
// Each concrete factory is responsible for creating objects of a specific family.
// For example, the WindowsFactory creates objects with a Windows look and feel,
// while the MacFactory creates objects with a Mac look and feel.

// Then we initialise our factory during runtime.
func main() {
	var factory WidgetFactory

	if runtime.GOOS == "windows" {
		factory = &WindowsFactory{}
	} else {
		factory = &MacFactory{}
	}

	window := factory.CreateWindow()
	button := factory.CreateButton()
	menu := factory.CreateMenu()

	// Use the created objects...
}
