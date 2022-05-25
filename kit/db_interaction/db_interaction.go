package db_interaction

import (
	"database/sql"
	"fmt"

	"github.com/ryanjoy0000/newsfeed-app/kit/customErr"
	"github.com/ryanjoy0000/newsfeed-app/platform/newsFeed/models"
)

func ProcessRecords(db *sql.DB, q string) bool {
	// 1 - Create a prepared for execution
	stmt, err := db.Prepare(q)
	customErr.HandleErr(err)

	// 2 - Executes the prepared statement
	result, err := stmt.Exec()
	customErr.HandleErr(err)

	// 3 - Get the number of rows affected
	len, err := result.RowsAffected()
	customErr.HandleErr(err)

	fmt.Println("Executed with rows affected: ", len)
	return true
}

func FetchRecords(db *sql.DB, q string) []models.FeedItem {
	list := []models.FeedItem{}

	// 1 - Execute query to get rows
	rows, err := db.Query(q)
	customErr.HandleErr(err)

	// 2 - Prepare the next row for reading with Scan
	for rows.Next() {
		var feedItem models.FeedItem

		// 3 - Scan copies the columns in the current row into the values pointed at by dest.
		err = rows.Scan(
			&feedItem.ID,
			&feedItem.Content,
		)
		customErr.HandleErr(err)
		list = append(list, feedItem)
	}

	fmt.Println("Fetched total records: ", len(list))

	// 4 - Close rows after operations
	defer rows.Close()
	return list
}
