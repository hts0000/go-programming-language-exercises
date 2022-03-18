package tracks

type LessFunc func(i, j *Track) bool
type CustomSort struct {
	T         []*Track
	LessFuncs []LessFunc
}

func ColTitle(x, y *Track) bool  { return x.Title < y.Title }
func ColArtist(x, y *Track) bool { return x.Artist < y.Artist }
func ColAlbum(x, y *Track) bool  { return x.Album < y.Album }
func ColYear(x, y *Track) bool   { return x.Year < y.Year }
func ColLength(x, y *Track) bool { return x.Length < y.Length }

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
