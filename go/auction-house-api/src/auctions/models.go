package auctions

import (
	"auction-house/src/config"
	"auction-house/src/logger"
)

type Auction struct {
	ID         string  `json:"id"`
	MinimumBid float64 `json:"minimumBid"`
	BuyOut     float64 `json:"buyOut"`
	Owner      string  `json:"owner"`
	Item       string  `json:"item"`
}

func SaveAuction(minimumBid float64, buyOut float64, owner string, item string) (err error) {
	_, err = config.DB.Exec("INSERT INTO auctions (minimum_bid, buy_out, user_id, item_id) VALUES ($1,$2,$3,$4)", minimumBid, buyOut, owner, item)

	if err != nil {
		logger.MongoLogger.Fatalf("server failed to handle: %v", err)
		return err
	}
	return nil
}

func AllAuctions() ([]Auction, error) {
	rows, err := config.DB.Query("SELECT * FROM auctions")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	auctions := make([]Auction, 0)
	for rows.Next() {
		auction := Auction{}
		err := rows.Scan(&auction.ID, &auction.MinimumBid, &auction.BuyOut, &auction.Owner, &auction.Item)
		if err != nil {
			logger.MongoLogger.Fatalf("server failed to handle: %v", err)
			return nil, err
		}
		auctions = append(auctions, auction)
	}
	if err = rows.Err(); err != nil {
		logger.MongoLogger.Fatalf("server failed to handle: %v", err)
		return nil, err
	}
	return auctions, nil
}
