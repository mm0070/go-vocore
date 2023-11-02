package vocore

import (
	"github.com/google/gousb"
)

func setFrame(intf *gousb.Interface, data []byte) error {
	outEp, err := intf.OutEndpoint(EndpointID)
	if err != nil {
		return err
	}

	_, err = outEp.Write(data)
	return err
}

func sendPixelData(intf *gousb.Interface, pixelData []byte) error {
	err := setFrame(intf, pixelData)
	return err
}
