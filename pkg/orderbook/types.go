package orderbook

import (
	"github.com/shopspring/decimal"
)

type Market string
type OrderType string
type Limits []*Limit
type ByBestAsk struct{ Limits }
type ByBestBid struct{ Limits }
type Orders []*Order

const (
	MarketETH Market = "ETH"
)

const (
	MarketOrder OrderType = "MARKET"
	LimitOrder  OrderType = "LIMIT"
)

type Match struct {
	Ask        *Order
	Bid        *Order
	SizeFilled decimal.Decimal
	Price      decimal.Decimal
}

type Order struct {
	ID        int64
	Price     decimal.Decimal
	Size      decimal.Decimal
	Bid       bool
	Limit     *Limit
	Timestamp int64
}

type Limit struct {
	Price       decimal.Decimal
	Orders      Orders
	TotalVolume decimal.Decimal
}

type OrderBook struct {
	asks []*Limit
	bids []*Limit

	AskLimits map[decimal.Decimal]*Limit
	BidLimits map[decimal.Decimal]*Limit
	Orders    map[int64]*Order
}

type PlaceOrderRequest struct {
	Type   OrderType // Limit or market
	Bid    bool
	Size   decimal.Decimal
	Price  decimal.Decimal
	Market Market
}

type Exchange struct {
	Orderbooks map[Market]*OrderBook
}

type OrderbookData struct {
	TotalBidVolume decimal.Decimal
	TotalAskVolume decimal.Decimal
	Asks           []*Order
	Bids           []*Order
}

type MatchedOrder struct {
	ID    int64
	Price decimal.Decimal
	Size  decimal.Decimal
}
