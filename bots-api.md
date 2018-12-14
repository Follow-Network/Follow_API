Bots API

# Bots API for Follow

[TOC]

## General information

## Security

### SIGNED endpoint security
+ SIGNED endpoints require an additional parameter, signature, to be sent in the request body.
+ Endpoints use HMAC SHA256 signatures. The HMAC SHA256 signature is a keyed HMAC SHA256 operation. Use your Ethereum private key as the key and totalParams as the value for the HMAC operation
+ The signature is not case sensitive
+ totalParams is defined as the query string concatenated with the request body
## Test

### **Ping**
+ **Security**: None
+ **GET** /api/v1/ping
+ **Description:** Test connectivity to the Bot API
+ **Weight:** 1
+ **Parameters:** NONE
+ **Response body:** application/json
```
{}
```

### **Check server time**
+ **Security**: None
+ **GET** /api/v1/time
+ **Description:** Test connectivity to the Rest API and get the current server time.
+ **Weight:** 1
+ **Parameters:** NONE
+ **Response body:** application/json
```
{
    "serverTime": 1499827319559
}
```

### **Get exchange info**
+ **Security**: None
+ **GET** /api/v1/exchange_info
+ **Description:** Get exchange info
+ **Weight:** 1
+ **Parameters:** NONE
+ **Response body:** application/json
```
{
    "timezone": "UTC",
      "serverTime": 1508631584636,
      "rateLimits": [
        // These are defined in the `ENUM definitions` section under `Rate limiters (rateLimitType)`.
        // All limits are optional.
      ],
      "exchangeFilters": [
        // There are defined in the `Filters` section.
        // All filters are optional.
      ],
      "symbols": [{
        "symbol": "ETHBTC",
        "status": "TRADING",
        "baseAsset": "ETH",
        "baseAssetPrecision": 8,
        "quoteAsset": "BTC",
        "quotePrecision": 8,
        "orderTypes": [
          // These are defined in the `ENUM definitions` section under `Order types (orderTypes)`.
          // All orderTypes are optional.
        ],
        "icebergAllowed": false,
        "filters": [
          // There are defined in the `Filters` section.
          // All filters are optional.
        ]
      }
    ]
}
```

## Public 

### **Get trades**
+ **Security**: None
+ **GET** /api/v1/get_trades/{pair}
+ **Description:** Get all trades on the pair
+ **Weight:** 1
+ **Parameters:**
  + Pair 
    + **Example:** BTC_USD
+ **Response body:** application/json
```
{
    "BTC_USD": [
        {
          "trade_id": 3,
          "type": "sell",
          "price": "100",
          "quantity": "1",
          "amount": "100",
          "date": 1435488248
        }
    ]
}
```
### **Get order book**
+ **Security**: None
+ **GET** /api/v1/get_orders/{pair}
+ **Description:** Get all trades on the pair
+ **Weight:** 1
+ **Parameters:**
  + Pair 
    + *Example:* BTC_USD
  + Limits (default: 100)

+ **Response body:** application/json
```
{
  "BTC_USD": {
    "ask_quantity": "3",
    "ask_amount": "500",
    "ask_top": "100",
    "bid_quantity": "1",
    "bid_amount": "99",
    "bid_top": "99",
    "ask": [[100,1,100],[200,2,400]],
    "bid": [[99,1,99]]
  }
}
```

## Authenticated

### **User info**
+ **Security**: Sign
+ **POST** /api/v1/user_info
+ **Description:** Get user information and balances
+ **Weight:** 1
+ **Parameters:** NONE
+ **Response body:** application/json
```
{
  "uid": 10542,
  "server_date": 1435518576,
  "balances": {
    "BTC": "970.994",
    "USD": "949.47"
  },
  "reserved": {
    "BTC": "3",
    "USD": "0.5"
  }
}
```

### **Get open orders**
+ **Security**: Sign
+ **POST** /api/v1/user_open_orders
+ **Description:** Get user open orders information
+ **Weight:** 1
+ **Parameters:** NONE
+ **Response body:** application/json
```
{
  "BTC_USD": [
    {
      "order_id": "14",
      "created": "1435517311",
      "type": "buy",
      "pair": "BTC_USD",
      "price": "100",
      "quantity": "1",
      "amount": "100"
    }
  ]
}
```
### **Get user trades**
+ **Security**: Sign
+ **POST** /api/v1/user_trades
+ **Description:** Get user trades information
+ **Weight:** 1
+ **Parameters:** 
   + Pair 
   + Limits (default: 100)
   + Offset - from the last trade (default: 0)
+ **Request body:** application/json
```
{
    "pair": "BTC_USD",
    "limit": "100",
    "offset": "0"
}
```
+ **Response body:** application/json
```
{
  "BTC_USD": [
    {
      "order_id": "14",
      "created": "1435517311",
      "type": "buy",
      "pair": "BTC_USD",
      "price": "100",
      "quantity": "1",
      "amount": "100"
    }
  ]
}
```

### **Order create**
+ **Security**: Sign
+ **POST** /api/v1/order_create
+ **Description:** Create order to buy/sell using set/market price
+ **Weight:** 1
+ **Parameters:** 
   + Pair 
   + Quantity (default: 100)
   + Price - from the last trade (default: 0)
   + Type
       + buy
       + sell
       + market_buy 
       + market_sell

+ **Request body:** application/json
```
{
    "pair": "BTC_USD",
    "quantity": "3",
    "price": "100",
    "type": "buy"
}
```
+ **Response body:** application/json
```
{
    "result": true,
    "error": "",
    "order_id": 123456
}
```

### **Order cancel**
+ **Security**: Sign
+ **POST** /api/v1/order_cancel
+ **Description:** Cancel order using its id
+ **Weight:** 1
+ **Parameters:** 
   + Order_id

+ **Request body:** application/json
```
{
    "order_id": "104235"
}
```
+ **Response body:** application/json
```
{
  "result": true,
  "error": ""
}
```