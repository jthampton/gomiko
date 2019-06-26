package cisco

import (
	"github.com/Ali-aqrabawi/gomiko/pkg/driver"
	"github.com/Ali-aqrabawi/gomiko/pkg/types"
)

type IOSDevice struct {
	Driver driver.IDriver
	Prompt string
	base   types.CiscoDevice
}

func (d *IOSDevice) OpenSession() error {
	return d.base.OpenSession()

}

func (d *IOSDevice) Disconnect() {

	d.base.Disconnect()

}

func (d *IOSDevice) SendCommand(cmd string) (string, error) {
	return d.base.SendCommand(cmd)

}

func (d *IOSDevice) SendConfigSet(cmds []string) (string, error) {
	return d.base.SendConfigSet(cmds)

}
func (d *IOSDevice) SetSecret(secret string) {
	d.base.SetSecret(secret)
}
