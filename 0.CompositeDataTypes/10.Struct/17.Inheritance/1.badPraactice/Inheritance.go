package main

import "fmt"

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
	Product                Product
	Ram                    int
	Cpu                    string
	ScreenSize             float32
	OperatingSystem        string
	OperatingSystemVersion string
}

type Mobile struct {
	ElectronicProducts ElectronicProducts
	SimcardCapacity    int
	SimcardType        string
	NetworkType        string
	CameraCount        int
}

type Laptop struct {
	ElectronicProducts ElectronicProducts
	UsbPortCount       int
	KeyboardType       string
	HasCdRom           bool
}

func main() {
	laptop := Laptop{}

	laptop.ElectronicProducts.Product.Name = "Asus Rog Zephyrus G16 - G-603-vv"

	fmt.Println(laptop)
}
