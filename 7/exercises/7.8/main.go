package main

// 练习 7.8
// 很多图形界面提供了一个有状态的多重排序表格插件
// 主要的排序键是最近一次点击过列头的列，第二个排序键是第二最近点击过列头的列，等等
// 定义一个sort.Interface的实现用在这样的表格中
// 比较这个实现方式和重复使用sort.Stable来排序的方式

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

type LessFunc func(i, j *Track) bool

type CustomSort struct {
	T         []*Track
	LessFuncs []LessFunc
}

// LessFunc
func colTitle(x, y *Track) bool  { return x.Title < y.Title }
func colArtist(x, y *Track) bool { return x.Artist < y.Artist }
func colAlbum(x, y *Track) bool  { return x.Album < y.Album }
func colYear(x, y *Track) bool   { return x.Year < y.Year }
func colLength(x, y *Track) bool { return x.Length < y.Length }

// sort.Interface
func (c CustomSort) Len() int      { return len(c.T) }
func (c CustomSort) Swap(i, j int) { c.T[i], c.T[j] = c.T[j], c.T[i] }
func (c CustomSort) Less(i, j int) bool {
	for _, fn := range c.LessFuncs {
		switch {
		case fn(c.T[i], c.T[j]):
			return true
		case fn(c.T[j], c.T[i]):
			return false
		}
	}
	// 所有列比较完后均相等，不需要交换
	return false
}

type byTitle []*Track
type byArtist []*Track
type byAlbum []*Track
type byYear []*Track
type byLengnth []*Track

func (c byTitle) Len() int           { return len(c) }
func (c byTitle) Less(i, j int) bool { return c[i].Title < c[j].Title }
func (c byTitle) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }

func (c byArtist) Len() int           { return len(c) }
func (c byArtist) Less(i, j int) bool { return c[i].Artist < c[j].Artist }
func (c byArtist) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }

func (c byAlbum) Len() int           { return len(c) }
func (c byAlbum) Less(i, j int) bool { return c[i].Album < c[j].Album }
func (c byAlbum) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }

func (c byYear) Len() int           { return len(c) }
func (c byYear) Less(i, j int) bool { return c[i].Year < c[j].Year }
func (c byYear) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }

func (c byLengnth) Len() int           { return len(c) }
func (c byLengnth) Less(i, j int) bool { return c[i].Length < c[j].Length }
func (c byLengnth) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }

func main() {
	tracks := NewTrack()
	// sort.Interface实现
	c := CustomSort{
		T: tracks,
		LessFuncs: []LessFunc{
			colTitle,
			colArtist,
			colAlbum,
			// colYear,
			// colLength,
		}}
	sort.Sort(c)
	PrintTracks(c.T)

	tracks = NewTrack()
	// 使用sort.Stable
	sort.Stable(byTitle(tracks))
	sort.Stable(byArtist(tracks))
	sort.Stable(byAlbum(tracks))
	// sort.Stable(byYear(tracks))
	// sort.Stable(byLengnth(tracks))
	PrintTracks(tracks)
}

func NewTrack() []*Track {
	return []*Track{
		{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
		{"Go", "Moby", "Moby", 1992, length("3m37s")},
		{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
		{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
		{"Go", "Hts0000", "Hts0001", 2021, length("6m00s")},
		{"Go", "Hts0000", "Hts0000", 2021, length("6m00s")},
	}
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func PrintTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "-----", "-----", "-----", "-----")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	fmt.Fprintf(tw, "\n")
	tw.Flush()
}
