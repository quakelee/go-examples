package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/quakelee/go-examples/netbox/libs"
	"github.com/quakelee/go-netbox/netbox/client/dcim"
	"github.com/quakelee/go-netbox/netbox/models"
)

// go-netbox is created by OpenAPI, so you can set environment variable SWAGGER_DEBUG=true
// to help debug. In Bash shell with `export SWAGGER_DEBUG=true`

func main() {
	apiToken := "<your token>"
	c, err := libs.NetboxConnect("netboxhost.example.com", 443, apiToken, "https", true)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	macregexp := regexp.MustCompile("([0-9a-fA-F][0-9a-fA-F]:){5}([0-9a-fA-F][0-9a-fA-F])")
	var wintf models.WritableInterface
	deviceID := int64(1)
	lagID := int64(1)
	macaddr := "00:50:00:00:00:99"
	modeID := int64(100)  //Available values listed in URL https://netboxhost.example.com/api/dcim/_choices/interface:mode/
	typeID := int64(1150) //Available values listed in URL https://netboxhost.example.com/api/dcim/_choices/interface:type/
	mtu := int64(1500)
	intName := "Ethernet1"
	untaggedVLAN := int64(1)

	// Fill up WritableInterface struct first
	wintf.Description = "description"
	wintf.Device = &deviceID
	wintf.Enabled = true
	if lagID > 0 {
		wintf.Lag = lagID
	}
	if macregexp.MatchString(macaddr) {
		wintf.MacAddress = macaddr
	}
	wintf.MgmtOnly = false
	wintf.Mode = modeID

	wintf.Mtu = &mtu
	wintf.Name = &intName
	wintf.Type = typeID
	wintf.TaggedVlans = []int64{}
	wintf.UntaggedVlan = untaggedVLAN

	// Create DcimInterfacesCreateParams object
	ic := dcim.NewDcimInterfacesCreateParams()
	// Fill data with WritableInterface object
	ic.SetData(&wintf)
	// Send REST call to Netbox get interface ID as return value
	ci, err := c.Dcim.DcimInterfacesCreate(ic, nil)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	fmt.Printf("New interface created in Netbox which ID is %d\n", ci.Payload.ID)
}
