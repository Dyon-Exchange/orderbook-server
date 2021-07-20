# orderbook-server

This repo is a HTTP wrapper around an open source orderbook engine with some modifications to allow creating new order books for different assets. 


## Usage

Can be run locally with `go run main.go`. Connect to `http://localhost:5341`.

With docker: 

`docker run -p 80:5341 orderbook-server`

There is a docker compose file included also. This allows you to have two containers running orderbooks so you can have separate ones for different environments, running on port `80` and `8080`. The Dockerfile is stored on my private registry but it should be moved to a Dyon owned one moving forward.

## API

Orders are represented in the orderbook with the following JSON schema. Any endpoints returning order details will be in the following format. 

``` JSON
{
    side: "ASK" | "BID",
    id: string,
    timestamp: Date,
    quantity: decimal,
    price: decimal
}
```

### Add Asset

`/addAsset` creates a new orderbook for an asset if it doesn't already exist.

Request:

```JSON
{
    "Asset": string
}
```

Response: `200` if success.

### Add Limit Order

`/addLimitOrder` adds a new limit order to the book

Request:

```JSON
{
    Asset: string,
    Side: "ASK" | "BID",
    OrderId: string, // unique id that identifies order in exchange database. Use whatever you want for this as long as its unique for each order. UUID is fine.
    Quantity: decimal,
    Price:  decimal
}
```

Response:

```JSON 
{
     Done: []Order, // any orders that were filled. This will include your order if it was filled.
     Partial: Order, // any order that was partially filled. This will be your order if it was partially filled,
     PartialQuantityProcessed: decimal // the amount in the partial order that was filled
}
```

### Add Market Order

`/addMarketOrder`adds a new market order to the orderbook

Request:

```JSON
{
    Asset: string,
    Side: "ASK"|"BID",
    OrderId: string,
    Quantity: decimal,
    Price: decimal
}
```

Response:

``` JSON
{
    Done: []Order,
    Partial: Order,
    PartialQuantityProcessed: decimal,
    QuantityLeft: decimal // the the remainder from your market order that wasn't filled 
}
```

### Cancel Order

`/cancelOrder` removes an order from the orderbook.

Request:

```JSON
{
    "OrderId": string,
    "Asset": string
}
```

Response:

``` JSON
{
    Order: Order // order that was canceled 
}
```

### Calculate Market Price

`/calculateMarketPrice` calculates the market price for a buy or sell of specified quantity of an asset.

Request:

```JSON
{
    "Asset": string,
    "Side": "ASK" | "BID",
    "Quantity": decimal
}
```

Response:

``` JSON
{
    Price: decimal
}
```

### Health Check

`/healthCheck` confirms the service is alive. In future this should ping the external datastore of the orderbook. 

Response:

```JSON
{
    "alive": true
}
```


### Get orders

`/getOrders`returns all the orderbooks on the server. Will just dump out the entire state of the orderbook so it will be complicated to make sense of it. Useful for debugging with a CTRL-F of a particular orderId.

### Reset

`/reset`resets the orderbook and deletes all orders. Useful for development and testing.

