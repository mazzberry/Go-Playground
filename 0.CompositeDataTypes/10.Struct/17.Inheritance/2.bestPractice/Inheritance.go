package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	Name   string
	Color  string
	Length int
	Width  int
	Weight int
	Price  int
	Brand  string
	MadeIn string
}

type ElectronicProducts struct {
	Product
	Ram                    int
	Cpu                    string
	ScreenSize             float32
	OperatingSystem        string
	OperatingSystemVersion string
}

type Mobile struct {
	ElectronicProducts
	SimcardCapacity int
	SimcardType     string
	NetworkType     string
	CameraCount     int
}

type Laptop struct {
	ElectronicProducts
	UsbPortCount int
	KeyboardType string
	HasCdRom     bool
}

func main() {
	mobile := Mobile{}
	mobile.Name = "Xiaomi Mi Note 11 Pro Plus"
	mobile.Brand = "Xiaomi"
	mobile.Color = "Grey"
	mobile.Ram = 8
	mobile.OperatingSystem = "Android"
	mobile.OperatingSystemVersion = "13"
	mobile.ScreenSize = 6.3
	//
	mobile.CameraCount = 3
	mobile.NetworkType = "5G"
	mobile.SimcardType = "Nano"
	mobile.SimcardCapacity = 2

	laptop := Laptop{}
	laptop.Name = "Asus Rog Zephyrus G16 - G-603-vv"
	laptop.Color = "grey"
	laptop.Length = 14
	laptop.Width = 9
	laptop.Weight = 2
	laptop.Price = 1600
	laptop.Brand = "Asus"
	laptop.MadeIn = "Taiwan"
	laptop.Ram = 32
	laptop.Cpu = "Intel core I7 13620H "
	laptop.OperatingSystem = "Windows"
	laptop.OperatingSystemVersion = "11"

	laptop.UsbPortCount = 2
	laptop.KeyboardType = "1.6"
	laptop.HasCdRom = false

	mobileJson, _ := json.Marshal(mobile)
	laptopJson, _ := json.Marshal(laptop)

	fmt.Println("Mobile\n",string(mobileJson), "\n")
	fmt.Println("Laptop\n", string(laptopJson), "\n")
}
