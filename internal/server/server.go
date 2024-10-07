package server

import (
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/skuril-bobishku/song-library/internal/database"
	"github.com/skuril-bobishku/song-library/internal/routing"
)

func StartHTTP(port string) {
	sb_songs := database.ConnectDB()

	http.HandleFunc("/", routing.StartPage)
	http.HandleFunc("/filter", func(w http.ResponseWriter, r *http.Request) {
		routing.GetFilteredData(sb_songs, w, r)
	})
	http.HandleFunc("/text", func(w http.ResponseWriter, r *http.Request) {
		routing.GetTextSong(sb_songs, w, r)
	})
	http.HandleFunc("/delete", func(w http.ResponseWriter, r *http.Request) {
		routing.DeleteSong(sb_songs, w, r)
	})
	http.HandleFunc("/change", func(w http.ResponseWriter, r *http.Request) {
		routing.ChangeSong(sb_songs, w, r)
	})
	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		routing.AddSong(sb_songs, w, r)
	})

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
