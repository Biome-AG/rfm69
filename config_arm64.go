// +build !walrus

package rfm69

// Configuration for NanoPi Neo2 LTS

const (
	spiDevice    = "/dev/spidev0.0"
	spiSpeed     = 6000000 // Hz
	interruptPin = 203     // GPIO for receive interrupts
	resetPin     = 6      // GPIO for hardware reset
)
