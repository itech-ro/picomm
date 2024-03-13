package picomm

import "fmt"

type (
	// Wpi2Bcm ...
	Wpi2Bcm struct {
		Pins map[int]int
	}
)

// NewWpi2Bcm ...
func NewWpi2Bcm() *Wpi2Bcm {
	wpi2bcm := make(map[int]int)

	wpi2bcm[1] = 18
	wpi2bcm[2] = 27
	wpi2bcm[3] = 22
	wpi2bcm[4] = 23
	wpi2bcm[5] = 24
	wpi2bcm[6] = 25
	wpi2bcm[7] = 4
	wpi2bcm[8] = 2
	wpi2bcm[9] = 3
	wpi2bcm[10] = 8
	wpi2bcm[11] = 7
	wpi2bcm[12] = 10
	wpi2bcm[13] = 9
	wpi2bcm[14] = 11
	wpi2bcm[15] = 14
	wpi2bcm[16] = 15
	wpi2bcm[21] = 5
	wpi2bcm[22] = 6
	wpi2bcm[23] = 13
	wpi2bcm[24] = 19
	wpi2bcm[25] = 26
	wpi2bcm[26] = 12
	wpi2bcm[27] = 16
	wpi2bcm[28] = 20
	wpi2bcm[29] = 21
	wpi2bcm[31] = 1

	return &Wpi2Bcm{
		Pins: wpi2bcm,
	}
}

// Convert ...
func (w2b *Wpi2Bcm) Convert(wpi int) (int, error) {
	bcm, found := w2b.Pins[wpi]

	if !found {
		return 0, fmt.Errorf("invalid wPI pin: %d", wpi)
	}

	return bcm, nil
}
