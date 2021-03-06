package trigger

import (
	"log"
	"sync"

	"github.com/nickysemenza/hyperion/color"
	"github.com/nickysemenza/hyperion/cue"
)

type trigger struct {
	source string
	id     int
}

type chanOfTriggers chan trigger

var (
	triggers chanOfTriggers
	once     sync.Once
)

func getTriggerChan() chanOfTriggers {
	once.Do(func() {
		triggers = make(chanOfTriggers, 100)
	})
	return triggers
}

//Action is called when an trigger needs to be fired
func Action(source string, id int) {
	c := getTriggerChan()
	c <- trigger{source, id}
}

//ProcessTriggers is a worker that processes triggers
func ProcessTriggers() {
	c := getTriggerChan()
	for t := range c {

		var newCues []cue.Cue
		log.Printf("new trigger! %v\n", t)
		if t.id == 1 {
			newCues = append(newCues, cue.NewSimple("hue1", color.FromString(color.Red)))
			newCues = append(newCues, cue.NewSimple("hue2", color.FromString(color.Blue)))
		}
		if t.id == 2 {
			newCues = append(newCues, cue.NewSimple("hue1", color.FromString(color.Green)))
		}
		if t.id == 3 {
			newCues = append(newCues, cue.NewSimple("hue1", color.FromString(color.Blue)))
		}
		if t.id == 4 {
			newCues = append(newCues, cue.NewSimple("hue1", color.FromString(color.Black)))
			newCues = append(newCues, cue.NewSimple("hue2", color.FromString(color.Black)))
		}

		for _, x := range newCues {
			stack := cue.GetCueMaster().GetDefaultCueStack()
			stack.EnQueueCue(x)
		}
	}
}
