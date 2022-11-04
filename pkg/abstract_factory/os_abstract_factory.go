package abstract_factory

type GUIFactory interface {
	CreateButton() Button
	CreateCheckbox() Checkbox
}

type Button interface {
	Paint()
}

type WinButton struct{}

func (w WinButton) Paint() {
}

type MacButton struct{}

func (w MacButton) Paint() {
}

type Checkbox interface {
	Paint()
}

type MacCheckbox struct{}

func (m MacCheckbox) Paint() {
}

type WinCheckbox struct{}

func (m WinCheckbox) Paint() {
}

type WinFactory struct{}

func (w WinFactory) CreateButton() Button {
	return nil
}

func (w WinFactory) CreateCheckbox() Checkbox {
	return nil
}

type MacFactory struct{}

func (w MacFactory) CreateButton() Button {
	return nil
}

func (w MacFactory) CreateCheckbox() Checkbox {
	return nil
}

type Application struct {
	Factory GUIFactory
	Button  Button
}

func (a *Application) CreateButton() {
	a.Button = a.Factory.CreateButton()
}

func (a *Application) Paint() {
	a.Button.Paint()
}

func NewApplication(factory GUIFactory) *Application {
	app := &Application{Factory: factory}
	return app
}

func main() {
	os := "mac"

	var factory GUIFactory

	if os == "mac" {
		factory = &MacFactory{}
	} else {
		factory = &WinFactory{}
	}

	app := NewApplication(factory)
	app.CreateButton()
	app.Paint()
}
