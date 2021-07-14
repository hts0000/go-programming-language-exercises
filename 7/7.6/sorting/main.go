package main

import (
	"sort"
	"sorting/byartist"
	"sorting/byyear"
	"sorting/customsort"
	"sorting/track"
)

func main() {
	tracks := track.NewTracker()
	track.PrintTracks(tracks)

	// sort.Reverse只改变了ByArtist接口的Less方法，使其变成倒序，其他方法则继续使用原有的
	sort.Sort(sort.Reverse(byartist.ByArtist(tracks)))
	track.PrintTracks(tracks)

	sort.Sort(byartist.ByArtist(tracks))
	track.PrintTracks(tracks)

	sort.Sort(byyear.ByYear(tracks))
	track.PrintTracks(tracks)

	sort.Sort(customsort.CustomSort{T: tracks, LessFunc: func(x, y *track.Track) bool {
		if x.Title != y.Title {
			return x.Title < y.Title
		}
		if x.Year != y.Year {
			return x.Year < y.Year
		}
		if x.Length != y.Length {
			return x.Length < y.Length
		}
		return false
	}})
	track.PrintTracks(tracks)
}
