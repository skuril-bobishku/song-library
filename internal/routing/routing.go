package routing

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/skuril-bobishku/song-library/internal/database"
	"github.com/skuril-bobishku/song-library/internal/track"
	env "github.com/skuril-bobishku/song-library/pkg/systems"
)

func StartPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Стартовая")
}

func GetFilteredData(db *sqlx.DB, w http.ResponseWriter, r *http.Request) {
	group := r.URL.Query().Get("group")
	song := r.URL.Query().Get("song")
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))

	songs, err := database.GetFilterFields(db, group, song, limit, offset)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(songs)
}

func GetTextSong(db *sqlx.DB, w http.ResponseWriter, r *http.Request) {
	songID, _ := strconv.Atoi(r.URL.Query().Get("id"))

	text, err := database.GetTextSong(db, songID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"text": text})
}

func AddSong(db *sqlx.DB, w http.ResponseWriter, r *http.Request) {
	var song track.Track
	if err := json.NewDecoder(r.Body).Decode(&song); err != nil {
		http.Error(w, "Ошибка декодирования JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	/*songDetails, err := fetchSongDetails(song.GroupName, song.SongName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	song.ReleaseDate = songDetails.ReleaseDate
	song.Text = songDetails.Text
	song.Link = songDetails.Link*/

	if err := database.AddSong(db, song); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(song)
}

func fetchSongDetails(group, song string) (*track.TrackDetail, error) {
	url := fmt.Sprintf("http://%s:%s/info?group=%s&song=%s", env.GetEnvString("HOST_IP"),
		env.GetEnvString("SERVER_PORT"), url.QueryEscape(group), url.QueryEscape(song))

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Ошибка получения деталей трека: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Ошибка получения деталей трека, статус: %d", resp.StatusCode)
	}

	var details track.TrackDetail
	if err := json.NewDecoder(resp.Body).Decode(&details); err != nil {
		return nil, fmt.Errorf("Ошибка парсинга ответа: %v", err)
	}

	return &details, nil
}

func ChangeSong(db *sqlx.DB, w http.ResponseWriter, r *http.Request) {
	var song track.Track
	if err := json.NewDecoder(r.Body).Decode(&song); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := database.ChangeSong(db, song); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(song)
}

func DeleteSong(db *sqlx.DB, w http.ResponseWriter, r *http.Request) {
	songID, _ := strconv.Atoi(r.URL.Query().Get("id"))

	if err := database.DeleteSong(db, songID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
