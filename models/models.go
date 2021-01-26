package models

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"log"
	"os"
)

var db *sql.DB

var server = "afbd.database.windows.net"
var port = 1433
var user = "antonio.belotti"
var password = os.Getenv("AZUREDB_PWD")
var database = "lastfm"

func InitDB() error {
	// Build connection string
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", server, user, password, port, database)
	var err error
	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	ctx := context.Background()
	return db.PingContext(ctx)
}

type GreatestHitRow struct {
	AlbumMBID        string
	AlbumName        string
	ArtistMBID       string
	ArtistName       string
	ImageURL         string
	SongMDIB         string
	SongName         string
	TimesListened    int
	TrackDurationSec int
}

func TodayGreatestHits() ([]GreatestHitRow, error) {

	rows, err := db.Query("SELECT * FROM dbo.today_gh ORDER BY times_listened DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var chart []GreatestHitRow

	for rows.Next() {
		var song GreatestHitRow

		err := rows.Scan(
			&song.AlbumMBID,
			&song.AlbumName,
			&song.ArtistMBID,
			&song.ArtistName,
			&song.ImageURL,
			&song.SongMDIB,
			&song.SongName,
			&song.TimesListened,
			&song.TrackDurationSec)

		if err != nil {
			return nil, err
		}

		chart = append(chart, song)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return chart, nil
}
