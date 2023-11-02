package vocore

import (
	"github.com/google/gousb"
)

const (
	// Device config
	Width      = 800
	Height     = 480
	PixelSize  = 3
	VendorID   = 0xC872 // 4 inch vocore screen
	ProductID  = 0x1004 // 4 inch vocore screen
	EndpointID = 0x2    // endpoint address for data transfer
)

type VocoreScreen struct {
	dev  *gousb.Device
	intf *gousb.Interface
}

func InitializeScreen() (v *VocoreScreen, err error) {
	ctx := gousb.NewContext()
	defer ctx.Close()

	dev, err := ctx.OpenDeviceWithVIDPID(VendorID, ProductID)
	if err != nil {
		return
	}

	intf, _, err := dev.DefaultInterface()
	if err != nil {
		return
	}

	code, err := dev.Control(0x40, 0xB0, 0, 0, []byte{0x00, 0x29})
	if err != nil || code != 0 {
		return
	}

	v = &VocoreScreen{
		dev:  dev,
		intf: intf,
	}

	return
}

func (v *VocoreScreen) WriteToScreen(img []byte) (err error) {
	code, err := v.dev.Control(0x40, 0xB0, 0, 0, []byte{0x00, 0x2C, 0x00, 0x00, 0x1C, 0x11})
	if err != nil || code != 0 {
		return
	}

	err = sendPixelData(v.intf, img)
	if err != nil {
		return
	}
	return
}

func (v *VocoreScreen) Close() {
	v.dev.Close()
	v.intf.Close()
}
