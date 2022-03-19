package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	handle := func(w http.ResponseWriter, r *http.Request) {
		// FormValue会调用ParseForm()并取出value[0]
		v := r.FormValue("cycles")
		fmt.Fprintln(os.Stdout, v)
		if v == "" {
			lissajous(w, 5)
		} else {
			n, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}
			lissajous(w, n)
		}
	}

	test := func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			log.Fatal(err)
		}
		// r.Form 底层类型是 map[string][]string
		// map的key是 http://localhost:8765?cycles=20%20100%20100 中的cycles
		// value是20%20100%20100，%20是空格转义而成
		// 通常来说只有一个值，用切片做value是为了兼容多个value的情况
		if v, ok := r.Form["cycles"]; !ok {
			fmt.Fprintln(w, "Cannot find key cycles")
		} else {
			fmt.Fprintln(w, "values", v)
		}
	}
	http.HandleFunc("/", handle)
	http.HandleFunc("/test", test)
	log.Fatal(http.ListenAndServe("localhost:8765", nil))
}

func lissajous(out io.Writer, cycles int) {
	const (
		//cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	err := gif.EncodeAll(out, &anim)
	if err != nil {
		fmt.Fprintf(os.Stderr, "lissajous: %v\n", err)
		return
	}
}
