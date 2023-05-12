package orderbook

func NewExchange() *Exchange {
	orderbooks := make(map[Market]*OrderBook)
	orderbooks[MarketETH] = NewOrderBook()
	return &Exchange{
		Orderbooks: orderbooks,
	}
}
