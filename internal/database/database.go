package database

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/skuril-bobishku/song-library/internal/track"
	env "github.com/skuril-bobishku/song-library/pkg/systems"
)

//var db *sqlx.DB

func ConnectDB() *sqlx.DB {
	cfg := env.LoadDBConfig()

	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DBhost, cfg.DBport, cfg.DBuser, cfg.DBpassword, cfg.DBname, cfg.DBsslmode))
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func GetFilterFields(db *sqlx.DB, group string, song string, limit int, offset int) ([]track.Track, error) {
	var songs []track.Track

	query := `SELECT * FROM songs WHERE 1=1`

	if group != "" {
		query += " AND group_name = :group_name"
	}
	if song != "" {
		query += " AND song_name = :song_name"
	}

	query += " LIMIT :limit OFFSET :offset"

	namedParams := map[string]interface{}{
		"group_name": group,
		"song_name":  song,
		"limit":      limit,
		"offset":     offset,
	}

	err := db.Select(&songs, query, namedParams)
	return songs, err
}

func GetTextSong(db *sqlx.DB, id int) (string, error) {
	var text string
	query := `SELECT text FROM songs WHERE s_id = $1`
	err := db.Get(&text, query, id)
	return text, err
}

func AddSong(db *sqlx.DB, song track.Track) error {
	//query := `INSERT INTO songs (group_name, song_name, release_date, text, link) VALUES ($1, $2, $3, $4, $5)`
	//_, err := db.Exec(query, song.GroupName, song.SongName, song.ReleaseDate, song.Text, song.Link)

	query := `INSERT INTO songs (group_name, song_name) VALUES ($1, $2)`
	_, err := db.Exec(query, song.GroupName, song.SongName)

	return err
}

func ChangeSong(db *sqlx.DB, song track.Track) error {
	query := `UPDATE songs SET group_name = :group_name, song_name = :song_name, release_date = :release_date,
		text = :text, link = :link WHERE s_id = :s_id`
	_, err := db.NamedExec(query, song)
	return err
}

func DeleteSong(db *sqlx.DB, id int) error {
	query := `DELETE FROM songs WHERE s_id = $1`
	_, err := db.Exec(query, id)
	return err
}
