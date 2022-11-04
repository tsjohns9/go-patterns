package adapter

import "fmt"

type LightningPort interface {
	Insert()
}

type Mac struct {
}

func (m *Mac) Insert() {
	fmt.Println("Lightning connector is plugged into mac machine.")
}

type Windows struct{}

func (w *Windows) insertIntoUSBPort() {
	fmt.Println("USB connector is plugged into windows machine.")
}

type WindowsAdapter struct {
	windowMachine *Windows
}

func (w *WindowsAdapter) Insert() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	w.windowMachine.insertIntoUSBPort()
}

type Client struct {
}

func (c *Client) InsertLightningConnectorIntoComputer(com LightningPort) {
	fmt.Println("Client inserts Lightning connector into computer.")
	com.Insert()
}

func main() {
	client := &Client{}
	mac := &Mac{}

	client.InsertLightningConnectorIntoComputer(mac)

	windowsMachine := &Windows{}
	windowsMachineAdapter := &WindowsAdapter{
		windowMachine: windowsMachine,
	}

	client.InsertLightningConnectorIntoComputer(windowsMachineAdapter)
}
