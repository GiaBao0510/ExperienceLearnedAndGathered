package main

import "fmt"

// Thuộc thành phần Factory
func getGun(gunType string) (IGun, error) {
	if gunType == "musket" {
		return newMusket(), nil
	}
	if gunType == "ak47" {
		return newAk47(), nil
	}
	return nil, fmt.Errorf("Gun type %s not found", gunType)
}