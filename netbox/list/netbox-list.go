package main

import (
	"fmt"
	"os"

	"github.com/quakelee/go-examples/netbox/libs"
	"github.com/quakelee/go-netbox/netbox/client/dcim"
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

	// List all of available Manufacturers saved in Netbox
	maparams := dcim.NewDcimManufacturersListParams()
	// Set limit to 0 is unlimit, default with paging setting of Netbox like 25
	limit := int64(0)
	maparams.SetLimit(&limit)
	mrs, err := c.Dcim.DcimManufacturersList(maparams, nil)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Number of Manufactures: %v\n", *(mrs.Payload.Count))
	fmt.Println("ID   Name                 Slug")
	fmt.Println("------------------------------------")
	for _, m := range mrs.Payload.Results {
		fmt.Printf("%-4d %-20s %-15s\n", m.ID, *m.Name, *m.Slug)
	}

	// List all of devices saved in Netbox.
	// Use manufacturer slug name as keyword, don't use Name value
	manus := []string{"arista", "cisco"}
	dparams := dcim.NewDcimDevicesListParams()
	nolimit := int64(0)
	dparams.SetLimit(&nolimit)
	for _, m := range manus {
		dparams.SetManufacturer(&m)
	}
	drs, err := c.Dcim.DcimDevicesList(dparams, nil)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Number of devices: %v\n", *(drs.Payload.Count))
	fmt.Println("ID   Name                        Manufacturer")
	fmt.Println("---------------------------------------------")
	for _, d := range drs.Payload.Results {
		fmt.Printf("%-4d %-32s %-10s\n", d.ID, d.Name, *d.DeviceType.Manufacturer.Name)
	}
}
