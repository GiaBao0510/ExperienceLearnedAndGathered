package main

import "fmt"

type WindowsAdapter struct {
	windowMachne *Windows
}

func (w *WindowsAdapter) InsertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	w.windowMachne.InsertIntoLightningPort()
}