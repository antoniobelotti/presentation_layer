package models

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"log"
	"math"
	"os"
)

var db *sql.DB

var server = "afbd.database.windows.net"
var port = 1433
var user = "presentation_web_app_login"
var password = os.Getenv("DB_PWD")
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
	AlbumName        string
	ArtistName       string
	ImageURL         string
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

	sqlQuery := fmt.Sprintf(`
		SELECT album_name, artist_name, image_url, name, times_listened, track_duration
		FROM dbo.%s 
		ORDER BY times_listened DESC;`, dbTable)
	rows, err := db.Query(sqlQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var chart []GreatestHitRow

	for rows.Next() {
		var song GreatestHitRow

		err := rows.Scan(
			&song.AlbumName,
			&song.ArtistName,
			&song.ImageURL,
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

func GetAllUsernames() ([]string, error) {
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

	return usernames, nil
}

type PlaylistBasicInfo struct {
	PlaylistId       int
	NumSongs         int
	PlaylistDuration int
}

func GetPlaylistsBasicInfoByUsername(username string) ([]PlaylistBasicInfo, error) {
	sqlQuery := fmt.Sprintf(`
		SELECT playlist_id, num_tracks, playlist_duration
		FROM dbo.playlists
		WHERE username_checksum = CHECKSUM(N'%s')
		ORDER BY num_tracks DESC;`, username)
	rows, err := db.Query(sqlQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var playlists []PlaylistBasicInfo

	for rows.Next() {
		var pl PlaylistBasicInfo

		err := rows.Scan(
			&pl.PlaylistId,
			&pl.NumSongs,
			&pl.PlaylistDuration,
		)

		if err != nil {
			return nil, err
		}

		playlists = append(playlists, pl)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return playlists, nil
}

type PlaylistSong struct {
	SongProgressive int
	SongName        string
	ImageUrl        string
	SongDuration    int
	ArtistName      string
	AlbumName       string
}

func GetPlaylistsSongs(username string, playlistId string) ([]PlaylistSong, error) {
	sqlQuery := fmt.Sprintf(`
		SELECT position_inside_playlist, song_name, image_url, track_duration, artist_name, album_name
		FROM dbo.playlist_songs
		WHERE username_checksum = CHECKSUM(N'%s')
		AND playlist_id = %s
		ORDER BY position_inside_playlist ASC;`, username, playlistId)

	rows, err := db.Query(sqlQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var playlistsData []PlaylistSong

	for rows.Next() {
		var ps PlaylistSong

		err := rows.Scan(
			&ps.SongProgressive,
			&ps.SongName,
			&ps.ImageUrl,
			&ps.SongDuration,
			&ps.ArtistName,
			&ps.AlbumName,
		)

		if err != nil {
			return nil, err
		}

		playlistsData = append(playlistsData, ps)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return playlistsData, nil
}

type GeneralStats struct {
	NumUsers             int
	NumPlaylists         int
	AvgTracksPerPlaylist float64
	AvgPlaylistsPerUser  float64
	AvgPlaylistLength    float64
}

func GetGeneralStats() (*GeneralStats, error) {
	sqlQuery := `SELECT COUNT(*) FROM dbo.users;`
	rows, err := db.Query(sqlQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stats GeneralStats
	for rows.Next() {
		err := rows.Scan(&stats.NumUsers)
		if err != nil {
			return nil, err
		}
	}

	sqlQuery2 := `SELECT COUNT(*) FROM dbo.playlists;`
	rows2, err := db.Query(sqlQuery2)
	if err != nil {
		return nil, err
	}
	defer rows2.Close()
	for rows2.Next() {
		err := rows2.Scan(&stats.NumPlaylists)
		if err != nil {
			return nil, err
		}
	}

	sqlQuery3 := `SELECT avg_tracks_for_playlist, avg_playlists_per_user, avg_playlist_len FROM dbo.basic_stats;`
	rows3, err := db.Query(sqlQuery3)
	if err != nil {
		return nil, err
	}
	defer rows3.Close()
	for rows3.Next() {
		err := rows3.Scan(&stats.AvgTracksPerPlaylist, &stats.AvgPlaylistsPerUser, &stats.AvgPlaylistLength)
		if err != nil {
			return nil, err
		}
	}

	stats.AvgPlaylistLength = math.Round(stats.AvgPlaylistLength*100)/100
	stats.AvgPlaylistsPerUser = math.Round(stats.AvgPlaylistsPerUser*100)/100
	stats.AvgTracksPerPlaylist = math.Round(stats.AvgTracksPerPlaylist*100)/100
	return &stats, nil
}

type PlaylistsCountByLength struct {
	PlaylistLength int
	NumPlaylists   int
}

func GetPlaylistLengthDistribution() ([]PlaylistsCountByLength, error) {
	sqlQuery := `SELECT len_playlist,num_playlist FROM dbo.number_of_playlists_by_length_distribution;`
	rows, err := db.Query(sqlQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var distribution []PlaylistsCountByLength
	for rows.Next() {
		var entry PlaylistsCountByLength
		err := rows.Scan(&entry.PlaylistLength, &entry.NumPlaylists)
		if err != nil {
			return nil, err
		}

		distribution = append(distribution, entry)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return distribution, nil
}

type PlaylistsCountForUser struct {
	NumberOfPlaylists int
	NumberOfUsers     int
}

func GetPlaylistsByUserDistribution() ([]PlaylistsCountForUser, error) {
	sqlQuery := `SELECT number_of_playlists, number_of_users FROM dbo.number_of_playlists_per_user_distribution;`
	rows, err := db.Query(sqlQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var distribution []PlaylistsCountForUser
	for rows.Next() {
		var entry PlaylistsCountForUser
		err := rows.Scan(&entry.NumberOfPlaylists, &entry.NumberOfUsers)
		if err != nil {
			return nil, err
		}

		distribution = append(distribution, entry)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return distribution, nil
}

type NumberOfTracksPerPlaylist struct {
	NumberOfTracks    int
	NumberOfPlaylists int
}

func GetNumberOfTracksPerPlaylistDistribution() ([]NumberOfTracksPerPlaylist, error) {
	sqlQuery := `SELECT number_of_playlists, number_of_tracks FROM dbo.number_of_tracks_per_playlist_distribution;`
	rows, err := db.Query(sqlQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var distribution []NumberOfTracksPerPlaylist
	for rows.Next() {
		var entry NumberOfTracksPerPlaylist
		err := rows.Scan(&entry.NumberOfPlaylists, &entry.NumberOfTracks)
		if err != nil {
			return nil, err
		}

		distribution = append(distribution, entry)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return distribution, nil
}
