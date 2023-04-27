package orderbook

import (
	"github.com/shopspring/decimal"
)

type Match struct {
	Ask        *Order
	Bid        *Order
	SizeFilled decimal.Decimal
	Price      decimal.Decimal
}

type Order struct {
	ID        int64
	Size      decimal.Decimal
	Bid       bool
	Limit     *Limit
	Timestamp int64
}

type Orders []*Order

type Limit struct {
	Price       decimal.Decimal
	Orders      Orders
	TotalVolume decimal.Decimal
}

type Limits []*Limit

type ByBestAsk struct{ Limits }

type ByBestBid struct{ Limits }

type OrderBook struct {
	asks []*Limit
	bids []*Limit

	AskLimits map[decimal.Decimal]*Limit
	BidLimits map[decimal.Decimal]*Limit
}
