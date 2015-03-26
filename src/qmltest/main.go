package main

import (
	"encoding/xml"
	"fmt"
	"gopkg.in/qml.v1"
	"io"
	"net/http"
	"os"
)

const (
	QML  = "share/ui/main.qml"
	BASE = "http://www.bing.com"
	ADDR = "/HPImageArchive.aspx?idx=0&n=1"
)

func main() {
	err := qml.Run(run)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERR: %v\n", err)
		os.Exit(1)
	}

}

func run() error {
	httpClient := new(http.Client)
	resp, err := httpClient.Get(BASE + ADDR)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	engine := qml.NewEngine()
	component, err := engine.LoadFile(QML)
	if err != nil {
		return err
	}
	win := component.CreateWindow(nil)
	win.Show()
	url, desc, _ := xmlParse(resp.Body)
	win.Set("descText", desc)
	win.Set("imgURL", BASE+url)
	win.Set("x", 200)
	win.Set("y", 200)
	win.Wait()
	return nil
}

func xmlParse(n io.ReadCloser) (url string, desc string, err error) {
	decoder := xml.NewDecoder(n)
	for {
		t, _ := decoder.Token()
		if t == nil {
			break
		}

		switch se := t.(type) {
		case xml.StartElement:
			switch se.Name.Local {
			case "url":
				decoder.DecodeElement(&url, &se)
			case "copyright":
				decoder.DecodeElement(&desc, &se)
			}
		}
	}
	err = nil
	return
}
