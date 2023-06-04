package orderbook

import (
	"fmt"
	"github.com/shopspring/decimal"
	"reflect"
	"testing"
)

func assert(t *testing.T, a, b interface{}) {
	if !reflect.DeepEqual(a, b) {
		t.Errorf("%+v != %+v", a, b)
	}
}

func TestLimit(t *testing.T) {
	l := NewLimit(decimal.NewFromInt(10_000))
	buyOrderA := NewOrder(true, decimal.NewFromInt(5))
	buyOrderB := NewOrder(true, decimal.NewFromInt(8))
	buyOrderC := NewOrder(true, decimal.NewFromInt(10))

	l.AddOrder(buyOrderA)
	l.AddOrder(buyOrderB)
	l.AddOrder(buyOrderC)

	l.DeleteOrder(buyOrderB)

	fmt.Println(l)
}

func TestPlaceLimitOrder(t *testing.T) {
	ob := NewOrderBook()

	sellOrderA := NewOrder(false, decimal.NewFromInt(10))
	sellOrderB := NewOrder(false, decimal.NewFromInt(5))
	ob.PlaceLimitOrder(decimal.NewFromInt(10_000), sellOrderA)
	ob.PlaceLimitOrder(decimal.NewFromInt(9_000), sellOrderB)

	assert(t, len(ob.asks), 2)
	assert(t, ob.Orders[sellOrderA.ID], sellOrderA)
	assert(t, ob.Orders[sellOrderB.ID], sellOrderB)
	assert(t, len(ob.Orders), 2)
}

func TestPlaceMarketOrder(t *testing.T) {
	ob := NewOrderBook()

	sellOrder := NewOrder(false, decimal.NewFromInt(20))
	ob.PlaceLimitOrder(decimal.NewFromInt(10_000), sellOrder)

	buyOrder := NewOrder(true, decimal.NewFromInt(10))
	matches := ob.PlaceMarketOrder(buyOrder)

	assert(t, len(matches), 1)
	assert(t, len(ob.asks), 1)
	assert(t, ob.AskTotalVolume().Equal(decimal.NewFromInt(10)), true)
	assert(t, matches[0].Ask, sellOrder)
	assert(t, matches[0].Bid, buyOrder)
	assert(t, matches[0].SizeFilled.Equal(decimal.NewFromInt(10)), true)
	assert(t, matches[0].Price.Equal(decimal.NewFromInt(10_000)), true)
	assert(t, buyOrder.IsFilled(), true)
}

func TestPlaceMarketOrderMultiFill(t *testing.T) {
	ob := NewOrderBook()

	buyOrderA := NewOrder(true, decimal.NewFromInt(5))
	buyOrderB := NewOrder(true, decimal.NewFromInt(8))
	buyOrderC := NewOrder(true, decimal.NewFromInt(10))
	buyOrderD := NewOrder(true, decimal.NewFromInt(1))

	ob.PlaceLimitOrder(decimal.NewFromInt(5_000), buyOrderC)
	ob.PlaceLimitOrder(decimal.NewFromInt(5_000), buyOrderD)
	ob.PlaceLimitOrder(decimal.NewFromInt(9_000), buyOrderB)
	ob.PlaceLimitOrder(decimal.NewFromInt(10_000), buyOrderA)

	assert(t, ob.BidTotalVolume(), decimal.NewFromFloat(24.00))

	sellOrder := NewOrder(false, decimal.NewFromInt(20))
	matches := ob.PlaceMarketOrder(sellOrder)

	assert(t, ob.BidTotalVolume(), decimal.NewFromFloat(4.0))
	assert(t, len(matches), 4)
	assert(t, len(ob.bids), 2)
}

func TestCancelOrder(t *testing.T) {
	ob := NewOrderBook()
	buyOrder := NewOrder(true, decimal.NewFromInt(4))

	ob.PlaceLimitOrder(decimal.NewFromInt(10_000), buyOrder)
	assert(t, ob.BidTotalVolume().Equal(decimal.NewFromInt(4)), true)

	ob.CancelOrder(buyOrder)
	assert(t, ob.BidTotalVolume().Equal(decimal.NewFromInt(0)), true)

	_, ok := ob.Orders[buyOrder.ID]
	assert(t, ok, false)
}
