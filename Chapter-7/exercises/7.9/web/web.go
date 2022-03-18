package web

import (
	"html/template"
	"net/http"
	"sort"
	"track-sort-web/tracks"
)

var tplt = template.Must(template.New("trackTable").Parse(`
<!DOCTYPE html>
<html>
  <head>
    <title>ex7.9</title>
      <style>
        table {
	      border-collapse: collapse;
        }
        td, th {
	      border: solid 1px;
	      padding: 0.5em;
          text-align: right;
        }
      </style>
  </head>
  <body>
    <table>
      <tr>
	  	<!-- 在这里设置请求字段 -->
	    <th><a href="./?by=title">Title</a></th>
	    <th><a href="./?by=artist">Artist</a></th>
	    <th><a href="./?by=album">Album</a></th>
	    <th><a href="./?by=year">Year</a></th>
	    <th><a href="./?by=length">Length</a></th>
	  </tr>
      {{range .}}
      <tr>
        <td>{{.Title}}</td>
        <td>{{.Artist}}</td>
        <td>{{.Album}}</td>
        <td>{{.Year}}</td>
        <td>{{.Length}}</td>
      </tr>
      {{end}}
    </table>
  </body>
</html>
`))

func Start() {
	track := tracks.NewTrack()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		col := r.URL.Query().Get("by") // 获取请求中的by字段，by字段是在html模板中设置的
		x := tracks.CustomSort{T: track}

		tplt.Execute(w, x.T) // 用track解析，并将解析结果写到http.ResponseWriter中

		switch col {
		case "title":
			x.LessFuncs = append(x.LessFuncs, tracks.ColTitle)
		case "artist":
			x.LessFuncs = append(x.LessFuncs, tracks.ColArtist)
		case "album":
			x.LessFuncs = append(x.LessFuncs, tracks.ColAlbum)
		case "year":
			x.LessFuncs = append(x.LessFuncs, tracks.ColYear)
		case "length":
			x.LessFuncs = append(x.LessFuncs, tracks.ColLength)
		case "title&length": // 模拟点击多列
			x.LessFuncs = append(x.LessFuncs, tracks.ColTitle)
			x.LessFuncs = append(x.LessFuncs, tracks.ColLength)
		default:
			x.LessFuncs = append(x.LessFuncs, tracks.ColTitle)
		}

		sort.Sort(x)
	})
	http.ListenAndServe("localhost:8000", nil)
}
