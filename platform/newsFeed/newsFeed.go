package newsFeed

import (
	"database/sql"

	"github.com/ryanjoy0000/newsfeed-app/kit/db_interaction"
	"github.com/ryanjoy0000/newsfeed-app/platform/newsFeed/models"
)

type FeedSystem struct {
	DB *sql.DB
}

func (feedSys *FeedSystem) Get() []models.FeedItem {
	q := `
	SELECT * FROM newsfeed;
	`
	result := db_interaction.FetchRecords(feedSys.DB, q)
	return result
}

func (feedSys *FeedSystem) Add(item models.FeedItem) {
	q := `
	INSERT INTO newsfeed(content)
	VALUES("` + item.Content + `")
	`

	db_interaction.ProcessRecords(feedSys.DB, q)
}

func CreateNewFeedSys(db *sql.DB) *FeedSystem {

	q := `
	CREATE TABLE IF NOT EXISTS newsfeed(
		ID INT AUTO_INCREMENT NOT NULL,
		content VARCHAR(500),
		PRIMARY KEY(ID)
	);
	`

	db_interaction.ProcessRecords(db, q)

	return &FeedSystem{
		DB: db,
	}
}
