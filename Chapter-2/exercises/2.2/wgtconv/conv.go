package wgtconv

func PdToKg(lb Pound) Kilos { return Kilos(lb * LbToKg) }
func KgToPd(kg Kilos) Pound { return Pound(kg * KgToLb) }
