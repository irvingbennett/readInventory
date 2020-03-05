// Package readInventory reads and inventory xml file
package main

import (
	"encoding/csv"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type CsvLine struct {
	dcs_code   string
	dep        string
	dep_desc   string
	class      string
	class_desc string
}

// ReadInventory reads an inventory xml
func main() {
	lines, err := ReadCsv("dcs.csv")
	if err != nil {
		panic(err)
	}

	dcsMap := make(map[string]CsvLine)
	for _, line := range lines {
		data := CsvLine{
			dcs_code:   line[0],
			dep:        line[1],
			dep_desc:   line[2],
			class:      line[3],
			class_desc: line[4],
		}
		dcsMap[data.dcs_code] = data
		fmt.Println(data.dcs_code + " " + data.dep_desc + " " + data.class_desc)
	}
	// fmt.Println(dcsMap)

	const (
		// A generic XML header suitable for use with the output of Marshal.
		// This is not automatically added to any output of this package,
		// it is provided as a convenience.
		Header = `<?xml version="1.0" encoding="UTF-8"?>` + "\n"
	)
	type DOCUMENT struct {
		XMLName xml.Name `xml:"DOCUMENT"`
		Text    string   `xml:",chardata"`
		DCSS    struct {
			Text string `xml:",chardata"`
			DCS  []struct {
				Text           string `xml:",chardata"`
				DcsCode        string `xml:"dcs_code,attr"`
				SbsNo          string `xml:"sbs_no,attr"`
				DName          string `xml:"d_name,attr"`
				CName          string `xml:"c_name,attr"`
				SName          string `xml:"s_name,attr"`
				DLongName      string `xml:"d_long_name,attr"`
				CLongName      string `xml:"c_long_name,attr"`
				SLongName      string `xml:"s_long_name,attr"`
				UseQtyDecimals string `xml:"use_qty_decimals,attr"`
				TaxCode        string `xml:"tax_code,attr"`
				MarginType     string `xml:"margin_type,attr"`
				MarginValue    string `xml:"margin_value,attr"`
				Active         string `xml:"active,attr"`
				Regional       string `xml:"regional,attr"`
				PtrnName       string `xml:"ptrn_name,attr"`
				DocDesign      string `xml:"doc_design,attr"`
			} `xml:"DCS"`
		} `xml:"DCSS"`
	}

	b, err := ioutil.ReadFile("dcs.xml") // b has type []byte
	// fmt.Println(string(b))
	if err != nil {
		log.Fatal(err)
	}
	v := DOCUMENT{}
	err = xml.Unmarshal(b, &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	v.Text = ""
	fmt.Println(v.DCSS.DCS[0].DcsCode)
	v.DCSS.Text = ""
	for i, dep := range v.DCSS.DCS {
		dname := strings.TrimSpace(dcsMap[dep.DcsCode].dep_desc)
		if len(dname) > 18 {
			dname = dname[:18]
		}
		v.DCSS.DCS[i].DName = dname
		v.DCSS.DCS[i].DLongName = strings.TrimSpace(dcsMap[dep.DcsCode].dep_desc)
		cname := strings.TrimSpace(dcsMap[dep.DcsCode].class_desc)
		if len(dname) > 18 {
			cname = cname[:18]
		}
		v.DCSS.DCS[i].CName = cname
		v.DCSS.DCS[i].CLongName = strings.TrimSpace(dcsMap[dep.DcsCode].class_desc)
		v.DCSS.DCS[i].Active = "1"

	}

	output, err := xml.MarshalIndent(v, "", "   ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Println()
	fmt.Println("------- Output------- ")
	fmt.Println()
	output = []byte(xml.Header + string(output))
	// fmt.Println(string(output))
	if err = ioutil.WriteFile("new.xml", output, 0644); err != nil {
		log.Fatal(err)
	}

}

// ReadCsv accepts a file and returns its content as a multi-dimentional type
// with lines and each column. Only parses to string type.
func ReadCsv(filename string) ([][]string, error) {

	// Open CSV file
	f, err := os.Open(filename)
	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()

	// Read File into a Variable
	dcs := csv.NewReader(f)
	dcs.Comma = ';'

	lines, err := dcs.ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return lines, nil
}
