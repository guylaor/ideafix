package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func (v *VideoItem) getTime() string {
	// mysql received dates as 'YYYY-MM-DD HH:MM:SS' - youtube format is 2016-12-30T07:40:49.000Z'
	dateStr := ""
	dateStr = v.PublishDate[:10] + " " + v.PublishDate[11:19]
	return dateStr
}

func (v *VideoItem) upsert() {

	db, err := sql.Open("mysql", fmt.Sprintf("%s", config["mysql-connection"]))
	checkError(err)

	query, err := db.Prepare("SELECT count(id) FROM videos WHERE youtube_id =?")
	checkError(err)

	var qid int
	err = query.QueryRow(v.YoutubueID).Scan(&qid)
	checkError(err)

	if qid != 0 {
		fmt.Println("video already exists")
		return
	}

	sql, err := db.Prepare("INSERT INTO videos set title=?, publish_date=?, youtube_id=?, description=?, thumb_small=?, thumb_medium=?, thumb_high=?, thumb_standard=?")
	checkError(err)

	res, err := sql.Exec(v.Title, v.getTime(), v.YoutubueID, v.Description, v.ThumbSmall, v.ThumbMedium, v.ThumbHigh, v.ThumbStandard)
	checkError(err)

	id, err := res.LastInsertId()
	checkError(err)

	fmt.Println(id)
}
