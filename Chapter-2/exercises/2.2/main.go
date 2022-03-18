package main

import (
	"flag"
	"fmt"

	"exercises-2.2/lenconv"
	"exercises-2.2/wgtconv"
)

func main() {
	mt := flag.Float64("m", 0.0, "Set length in meter.")
	ft := flag.Float64("f", 0.0, "Set length in feet.")
	lb := flag.Float64("p", 0.0, "Set weight in pound.")
	kg := flag.Float64("k", 0.0, "Set weight in kilos.")
	flag.Parse()

	// visit函数，遍历所有flag.Parse()函数解析成功的参数
	// 如：-f 10 -m 20 -k 100，会遍历10 20 100
	flag.Visit(func(f *flag.Flag) {
		switch f.Name {
		case "f":
			fmt.Printf("%s = %s\n", lenconv.Feet(*ft), lenconv.FtToM(lenconv.Feet(*ft)))
		case "k":
			fmt.Printf("%s = %s\n", wgtconv.Kilos(*kg), wgtconv.KgToPd(wgtconv.Kilos(*kg)))
		case "m":
			fmt.Printf("%s = %s\n", lenconv.Meter(*mt), lenconv.MToFt(lenconv.Meter(*mt)))
		case "p":
			fmt.Printf("%s = %s\n", wgtconv.Pound(*lb), wgtconv.PdToKg(wgtconv.Pound(*lb)))
		}
	})
}
