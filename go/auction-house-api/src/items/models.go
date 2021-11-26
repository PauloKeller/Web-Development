package items

import (
	"auction-house/src/config"
	"auction-house/src/logger"
)

// Item represents a tradable item
type Item struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

//EraseItem its a simple delete item
func EraseItem(id string) (err error) {
	_, err = config.DB.Exec("DELETE FROM items WHERE item_id=$1", id)

	if err != nil {
		logger.MongoLogger.Fatalf("server failed to handle: %v", err)
		return err
	}
	return nil
}

//SaveItem a single Item and insert into database
func SaveItem(name string, description string) (err error) {

	_, err = config.DB.Exec("INSERT INTO items (name,description) VALUES ($1,$2)", name, description)

	if err != nil {
		logger.MongoLogger.Fatalf("server failed to handle: %v", err)
		return err
	}
	return nil
}

// AllItems retrieve all records os Item inside the data base
func AllItems() ([]Item, error) {
	rows, err := config.DB.Query("SELECT * FROM items")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := make([]Item, 0)
	for rows.Next() {
		item := Item{}
		err := rows.Scan(&item.ID, &item.Name, &item.Description)
		if err != nil {
			logger.MongoLogger.Fatalf("server failed to handle: %v", err)
			return nil, err
		}
		items = append(items, item)
	}
	if err = rows.Err(); err != nil {
		logger.MongoLogger.Fatalf("server failed to handle: %v", err)
		return nil, err
	}
	return items, nil
}
