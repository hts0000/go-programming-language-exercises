package lenconv

func MToFt(m Meter) Feet  { return Feet(m * MToF) }
func FtToM(ft Feet) Meter { return Meter(ft * FToM) }
