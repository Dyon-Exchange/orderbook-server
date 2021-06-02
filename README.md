# orderbook-server

This repo is a HTTP wrapper around an open source orderbook engine from Github with some modifications to allow creating new order books for different assets. 

## API

### Add Asset

`/addAsset` creates a new orderbook for an asset if it doesn't already exist.

Request:

```JSON
{
    "Asset": string
}
```

Response: `200` if success.

### Add Order

`/addOrder` adds a new order to the book

Request:

```JSON
{
    "Type": "LIMIT" | "MARKET",
    Asset: string,
    Side: "ASK" | "BID",
    OrderId: string, // unique id that identifies order in exchange database
    Quantity: decimal,
    Price:  decimal
}
```

### Cancel OrderA

`/cancelOrder` removes an order from the orderbook.

Request:

```JSON
{
    "OrderId": string,
    "Asset": string
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

### Health Check

`/healthCheck` confirms the service is alive. In future this should ping the external datastore of the orderbook.

Response

```JSON
{
    "alive": true
}
```
