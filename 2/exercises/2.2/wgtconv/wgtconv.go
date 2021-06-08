package wgtconv

import "fmt"

type Pound float64
type Kilos float64

const (
	LbToKg = 0.45359237
	KgToLb = 2.2046226
)

func (p Pound) String() string { return fmt.Sprintf("%.4glb", p) }
func (k Kilos) String() string { return fmt.Sprintf("%.4gkg", k) }
