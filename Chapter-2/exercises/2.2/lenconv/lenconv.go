package lenconv

import "fmt"

type Meter float64
type Feet float64

const (
	MToF = 3.28
	FToM = 0.3048
)

func (m Meter) String() string { return fmt.Sprintf("%.4gm", m) }
func (ft Feet) String() string { return fmt.Sprintf("%.4gft", ft) }
