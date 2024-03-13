package picomm

import (
	"fmt"
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

type (
	// Watcher ...
	Watcher struct {
		Controller *Controller
	}
)

// NewWatcher ...
func NewWatcher(controller *Controller) *Watcher {
	return &Watcher{
		Controller: controller,
	}
}

// Run ...
func (w *Watcher) Run() {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		err := rpio.Open()
		if err != nil {
			fmt.Println(err)
			continue
		}
		pin := rpio.Pin(21)
		res := pin.Read() // Read state from pin (High / Low)
		fmt.Printf("state 21 %d\n", res)

		pin = rpio.Pin(22)
		res = pin.Read() // Read state from pin (High / Low)
		fmt.Printf("state 22 %d\n", res)

		rpio.Close()
	}
}
