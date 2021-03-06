package homekit

import (
	"fmt"
	"log"

	"github.com/nickysemenza/hyperion/trigger"

	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
)

const numSwitches = 5

//Start starts the HomeKit services
func Start() {
	config := hc.Config{Pin: "10000000", StoragePath: "./_homekit_data"}

	//for now let's have N switches
	var switches []accessory.Switch
	for x := 1; x <= numSwitches; x++ {
		switchName := fmt.Sprintf("Switch %d", x)
		s := accessory.NewSwitch(accessory.Info{Name: switchName, Manufacturer: "hyperion"})
		switches = append(switches, *s)
	}

	//an array of their accessory attributes
	accessories := make([]*accessory.Accessory, len(switches))
	for i, s := range switches {
		accessories[i] = s.Accessory
	}

	bridge := accessory.NewBridge(accessory.Info{Name: "bridge1", Manufacturer: "Hyperion"})

	//start the server
	t, err := hc.NewIPTransport(config, bridge.Accessory, accessories...)
	if err != nil {
		log.Panic(err)
	}

	//add some handlers for the switches...
	for i := range switches {
		eachSwitch := switches[i]
		ix := i
		eachSwitch.Switch.On.OnValueRemoteUpdate(func(on bool) {
			//TODO: call some code...ID
			if on == true {
				trigger.Action("homekit-switch", ix+1)
				eachSwitch.Switch.On.SetValue(false)
			}
			// log.Printf("[homekit] changed: [%s] to %t", eachSwitch.Accessory.Info.Name.String.GetValue(), on)
		})
	}

	hc.OnTermination(func() {
		<-t.Stop()
		log.Println("Homekit services stopped.")
	})

	t.Start()
}
