// Package readInventory reads and inventory xml file
package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
)

// ReadInventory reads an inventory xml
func main() {
	const (
		// A generic XML header suitable for use with the output of Marshal.
		// This is not automatically added to any output of this package,
		// it is provided as a convenience.
		Header = `<?xml version="1.0" encoding="UTF-8"?>` + "\n"
	)

	type Invn struct {
		ItemSid int64  `xml:"item_sid,attr"`
		Upc     string `xml:"upc,attr"`
	}

	type InvnSBS struct {
		SbsNo   string `xml:"sbs_no,attr"`
		Alu     string `xml:"alu,attr"`
		DcsCode string `xml:"dcs_code,attr"`

		Description1 string `xml:"description1,attr"`
		Description2 string `xml:"description2,attr"`
	}

	type Inventory struct {
		Invn    Invn    `xml:"INVN"`
		InvnSBS InvnSBS `xml:"INVN_SBS"`
	}

	type Document struct {
		XMLName    xml.Name    `xml:"DOCUMENT"`
		Inventorys []Inventory `xml:"INVENTORYS>INVENTORY"`
	}

	b, err := ioutil.ReadFile("inventory.xml") // b has type []byte
	// fmt.Println(string(b))
	if err != nil {
		log.Fatal(err)
	}
	v := Document{}
	err = xml.Unmarshal(b, &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	fmt.Println(v)
	output, err := xml.MarshalIndent(v, "", "   ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Println()
	fmt.Println("------- Output------- ")
	fmt.Println()
	output = []byte(xml.Header + string(output))
	fmt.Println(string(output))
	if err = ioutil.WriteFile("new.xml", output, 0644); err != nil {
		log.Fatal(err)
	}

}
