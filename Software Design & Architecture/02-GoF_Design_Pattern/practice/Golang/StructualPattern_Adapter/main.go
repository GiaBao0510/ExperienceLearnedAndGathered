package main

func main() {
	client := &Client{}
	mac := &Mac{}

	client.InsertIntoLightningPort(mac)

	windowsMachine := &Windows{}
	windowsAdapter := &WindowsAdapter{windowMachne: windowsMachine}

	client.InsertIntoLightningPort(windowsAdapter)

}