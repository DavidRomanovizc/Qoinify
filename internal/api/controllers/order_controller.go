package controllers

import (
	"github.com/DavidRomanovizc/Qoinify/pkg/orderbook"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"net/http"
	"strconv"
)

// GetPlaceOrder
// @Summary ping
// @Success 200 {string} pong "pong"
// @Router /api/ping/ [get]
func GetPlaceOrder(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

// HandlePlaceOrder processes a request to place a new order
// @Summary order
// @Description Places an order on the market or a limit order
// @Tags Orders
// @Accept json
// @Produce json
// @Param market body string true "Market"
// @Param type body string true "Type of the order (MARKET or LIMIT)"
// @Param bid body bool true "Direction of the order (BUY или SELL)"
// @Param price body string false "Price (only for limit order)"
// @Param size body string true "Volume"
// @Router /api/order/ [post]
func HandlePlaceOrder(ex *orderbook.Exchange) gin.HandlerFunc {
	return func(c *gin.Context) {
		var placeOrderData orderbook.PlaceOrderRequest
		if err := c.BindJSON(&placeOrderData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		market := orderbook.Market(placeOrderData.Market)
		ob, ok := ex.Orderbooks[market]
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "market not found",
			})
			return
		}
		order := orderbook.NewOrder(placeOrderData.Bid, placeOrderData.Size)

		isBid := false
		if order.Bid {
			isBid = true
		}

		if placeOrderData.Type == orderbook.MarketOrder {
			matches := ob.PlaceMarketOrder(order)
			matchesOrders := make([]*orderbook.MatchedOrder, len(matches))

			for i := 0; i < len(matches); i++ {
				id := matches[i].Bid.ID
				if isBid {
					id = matches[i].Ask.ID
				}
				matchesOrders[i] = &orderbook.MatchedOrder{
					ID:    id,
					Size:  matches[i].SizeFilled,
					Price: matches[i].Price,
				}
			}

			c.JSON(http.StatusOK, map[string]any{"matches": matchesOrders})

			return
		}

		price, err := decimal.NewFromString(placeOrderData.Price.String())
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "invalid price",
			})
			return
		}

		if price.LessThanOrEqual(decimal.Zero) {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "price must be greater than zero",
			})
			return
		}

		ob.PlaceLimitOrder(price, order)

		c.JSON(http.StatusOK, gin.H{
			"message": "order placed",
		})
	}
}

// HandleGetBook handles GET requests to retrieve the orderbook data for a particular market.
// @Summary Get orderbook data for a market
// @Description Get the aggregated bid and ask orders for a market
// @Tags orderbook
// @Param market path string true "Market name"
// @Produce json
// @Router /orderbook/{market} [get]
func HandleGetBook(ex *orderbook.Exchange) gin.HandlerFunc {
	return func(c *gin.Context) {
		market := orderbook.Market(c.Param("market"))
		ob, ok := ex.Orderbooks[market]
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "market not found",
			})
			return
		}

		orderbookData := orderbook.OrderbookData{
			TotalBidVolume: ob.BidTotalVolume(),
			TotalAskVolume: ob.AskTotalVolume(),
			Asks:           []*orderbook.Order{},
			Bids:           []*orderbook.Order{},
		}

		for _, limit := range ob.Asks() {
			for _, order := range limit.Orders {
				o := orderbook.Order{
					ID:        order.ID,
					Price:     limit.Price,
					Size:      order.Size,
					Bid:       order.Bid,
					Timestamp: order.Timestamp,
				}
				orderbookData.Asks = append(orderbookData.Asks, &o)
			}
		}

		for _, limit := range ob.Bids() {
			for _, order := range limit.Orders {
				o := orderbook.Order{
					ID:        order.ID,
					Price:     limit.Price,
					Size:      order.Size,
					Bid:       order.Bid,
					Timestamp: order.Timestamp,
				}
				orderbookData.Bids = append(orderbookData.Bids, &o)
			}
		}

		c.JSON(http.StatusOK, orderbookData)
	}
}

// HandleCancelOrder delete the order with the specified id
// @Summary Delete order
// @Description Deletes an order with the specified id
// @Tags Orders
// @Param id path int true "ID order"
// @Success 200 {string} string "order deleted"
// @Failure 400 {string} string "Invalid order id"
// @Failure 404 {string} string "Order not found"
// @Failure 500 {string} string "Internal server error"
// @Router /orders/{id} [delete]
func HandleCancelOrder(ex *orderbook.Exchange) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, _ := strconv.Atoi(idStr)

		ob := ex.Orderbooks[orderbook.MarketETH]
		order := ob.Orders[int64(id)]
		ob.CancelOrder(order)

		c.JSON(http.StatusOK, gin.H{
			"message": "order deleted",
		})

	}
}
