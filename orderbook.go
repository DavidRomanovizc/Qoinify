package main

import (
	"fmt"
	"sort"
	"time"
)

type Match struct {
	Ask        *Order
	Bid        *Order
	SizeFilled float64
	Price      float64
}

type Order struct {
	Size      float64
	Limit     *Limit
	Timestamp int64
	Bit       bool
}

type Limit struct {
	Price       float64
	Orders      Orders
	TotalVolume float64
}

type Limits []*Limit

type ByBestAsk struct {
	Limits
}

type ByBestBit struct {
	Limits
}

type Orders []*Order

type Orderbook struct {
	Asks []*Limit
	Bids []*Limit

	AskLimits map[float64]*Limit
	BidLimits map[float64]*Limit
}

func NewLimit(price float64) *Limit {
	return &Limit{
		Price:  price,
		Orders: []*Order{},
	}
}

func NewOrder(bit bool, size float64) *Order {
	return &Order{
		Size:      size,
		Bit:       bit,
		Timestamp: time.Now().UnixNano(),
	}
}

func (l *Limit) AddOrder(o *Order) {
	o.Limit = l
	l.Orders = append(l.Orders, o)
	l.TotalVolume += o.Size
}

func (o *Order) String() string {
	return fmt.Sprintf("[size: %.2f]", o.Size)
}

func (l *Limit) DeleteOrder(o *Order) {
	for i := 0; i < len(l.Orders); i++ {
		if l.Orders[i] == o {
			l.Orders[i] = l.Orders[len(l.Orders)-1]
			l.Orders = l.Orders[:len(l.Orders)-1]
		}
	}
	o.Limit = nil
	l.TotalVolume -= o.Size

	sort.Sort(l.Orders)
}

func NewOrderBook() *Orderbook {
	return &Orderbook{
		Asks:      []*Limit{},
		Bids:      []*Limit{},
		AskLimits: make(map[float64]*Limit),
		BidLimits: make(map[float64]*Limit),
	}
}

func (ob *Orderbook) PlaceOrder(price float64, o *Order) []Match {
	// 1. Try to match the orders
	// Matching logic
	// 2. Add the rest of the order to the book

	if o.Size > 0.0 {
		ob.add(price, o)
	}

	return []Match{}
}

func (ob *Orderbook) add(price float64, o *Order) {
	var limit *Limit

	if o.Bit {
		limit = ob.BidLimits[price]
	} else {
		limit = ob.AskLimits[price]
	}

	if limit == nil {
		limit = NewLimit(price)
		limit.AddOrder(o)
		if o.Bit {
			ob.Bids = append(ob.Bids, limit)
			ob.BidLimits[price] = limit
		} else {
			ob.Asks = append(ob.Asks, limit)
			ob.AskLimits[price] = limit
		}
	}
}

func (l *Limit) String() string {
	return fmt.Sprintf("[price: %.2f | volume: %.2f]", l.Price, l.TotalVolume)
}

func (a ByBestAsk) Len() int {
	return len(a.Limits)
}

func (a ByBestAsk) Swap(i, j int) {
	a.Limits[i], a.Limits[j] = a.Limits[j], a.Limits[i]
}

func (a ByBestAsk) Less(i, j int) bool {
	return a.Limits[i].Price < a.Limits[j].Price
}

func (b ByBestBit) Len() int {
	return len(b.Limits)
}

func (b ByBestBit) Swap(i, j int) {
	b.Limits[i], b.Limits[j] = b.Limits[j], b.Limits[i]
}

func (b ByBestBit) Less(i, j int) bool {
	return b.Limits[i].Price > b.Limits[j].Price
}

func (o Orders) Len() int {
	return len(o)
}

func (o Orders) Swap(i, j int) {
	o[i], o[j] = o[j], o[i]
}

func (o Orders) Less(i, j int) bool {
	return o[i].Timestamp < o[j].Timestamp
}
