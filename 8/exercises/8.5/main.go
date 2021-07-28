package main

import (
	"fmt"
	"image"
	"image/color"
	"math/cmplx"
	"runtime"
	"sync"
	"time"
)

// 练习 8.5
//使用一个已有的CPU绑定的顺序程序，比如在3.3节中我们写的Mandelbrot程序或者3.2节中的3-D surface计算程序，并将他们的主循环改为并发形式，使用channel来进行通信
//在多核计算机上这个程序得到了多少速度上的改进？使用多少个goroutine是最合适的呢？

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 1024, 1024
)

func main() {
	sqFn()
	coFn()
}

func sqFn() {
	t1 := time.Now()
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	// 写到本地有io操作，Linux下有缓存干扰
	//png.Encode(os.Stdout, img) // NOTE: ignoring errors
	fmt.Println("sqFn执行时间为：", time.Since(t1))
}

func coFn() {
	t1 := time.Now()
	workers := runtime.NumCPU()
	runtime.GOMAXPROCS(workers)
	rows := make(chan int, height)
	var wg sync.WaitGroup

	// 构造数据
	go func() {
		for row := 0; row < height; row++ {
			rows <- row
		}
		close(rows)
	}()

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// 根据当前机器的cpu核心数启动对应数量goroutinue
	// 每个goroutinue从rows管道中获取需要处理的py位置，开始绘制对应像素点
	// 从整体图片来看，图片上的像素点生成并不是顺序的，而更像是随机的
	for worker := 0; worker < workers; worker++ {
		wg.Add(1)
		go func() {
			for py := range rows {
				y := float64(py)/height*(ymax-ymin) + ymin
				for px := 0; px < width; px++ {
					x := float64(px)/width*(xmax-xmin) + xmin
					z := complex(x, y)
					// Image point (px, py) represents complex value z.
					img.Set(px, py, mandelbrot(z))
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()

	// 写到本地有io操作，Linux下有缓存干扰
	//png.Encode(os.Stdout, img) // NOTE: ignoring errors
	fmt.Println("coFn执行时间为：", time.Since(t1))
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
