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

		if placeOrderData.Type == orderbook.MarketOrder {
			matches := ob.PlaceMarketOrder(order)
			c.JSON(http.StatusOK, gin.H{
				"matches": len(matches),
			})
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

func HandleCancelOrder(ex *orderbook.Exchange) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, _ := strconv.Atoi(idStr)

		ob := ex.Orderbooks[orderbook.MarketETH]
		orderCanceled := false

		for _, limit := range ob.Asks() {
			for _, order := range limit.Orders {
				if order.ID == int64(id) {
					ob.CancelOrder(order)
					orderCanceled = true
				}

				if orderCanceled {
					c.JSON(200, gin.H{"msg": "order canceled"})
					return
				}
			}
		}

		for _, limit := range ob.Bids() {
			for _, order := range limit.Orders {
				if order.ID == int64(id) {
					ob.CancelOrder(order)
					orderCanceled = true
				}

				if orderCanceled {
					c.JSON(200, gin.H{"msg": "order canceled"})
					return
				}
			}
		}
	}
}
