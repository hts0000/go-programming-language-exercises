package customsort

import "sorting/track"

type CustomSort struct {
	T        []*track.Track
	LessFunc func(x, y *track.Track) bool
}

func (x CustomSort) Len() int           { return len(x.T) }
func (x CustomSort) Less(i, j int) bool { return x.LessFunc(x.T[i], x.T[j]) }
func (x CustomSort) Swap(i, j int)      { x.T[i], x.T[j] = x.T[j], x.T[i] }
