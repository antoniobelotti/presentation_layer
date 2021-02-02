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

func GreatestHits(period string) ([]GreatestHitRow, error) {
	var dbTable string
	switch period {
	case "today":
		dbTable = "today_gh"
	case "last_week":
		dbTable = "last_week_gh"
	case "last_month":
		dbTable = "last_month_gh"
	}

	sqlQuery := fmt.Sprintf("SELECT * FROM dbo.%s ORDER BY times_listened DESC", dbTable)
	rows, err := db.Query(sqlQuery)
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


func GetAllUsernames() ([]string,error) {
	sqlQuery := fmt.Sprintf("SELECT username FROM [dbo].[users];")
	rows, err := db.Query(sqlQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var usernames []string

	for rows.Next() {
		var username string
		err := rows.Scan(&username)

		if err != nil {
			return nil, err
		}

		usernames = append(usernames, username)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return usernames,nil
}

type Playlists struct {
	Username string
}

func GetPlaylistsByUsername(username string) (string, error){
	return "ok",nil
}