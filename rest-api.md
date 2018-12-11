# Public Rest API for Follow

## General information

* Endpoints:
    * Dev: https://dev.api.thefollow.org
    * Prod: https://api.thefollow.org
* All endpoints return either a JSON object or array
* Data is returned in **ascending** order. Oldest first, newest last
* All time and timestamp related fields are in milliseconds
* *HTTP 4XX* return codes are used for for malformed requests; the issue is on the sender's side
* *HTTP 429* return code is used when breaking a request rate limit
* *HTTP 418* return code is used when an IP has been auto-banned for continuing to send requests after receiving 429 codes
* *HTTP 5XX* return codes are used for internal errors; the issue is on Follow's side. It is important to **NOT** treat this as a failure operation; the execution status is **UNKNOWN** and could have been a success
* Any endpoint can return an ERROR; the error payload is as follows:
```
{
  "code": -567,
  "msg": "Existing login"
}
```
* For *GET* endpoints, parameters must be sent as a *query string*
* For *POST*, *PUT*, and *DELETE* endpoints, the parameters may be sent in the request body with content type *application/json*
* Parameters may be sent in any order
* Each route has a weight which determines for the number of requests each endpoint counts for. Heavier endpoints and endpoints that do operations on multiple symbols will have a heavier weight
* When a *429* is recieved, it's your obligation as an API to back off and not spam the API
* Repeatedly violating rate limits and/or failing to back off after receiving *429s* will result in an automated IP ban (HTTP status *418*)
* IP bans are tracked and **scale in duration** for repeat offenders, **from 2 minutes to 3 days**

## Security

* Each endpoint has a security type that determines the how you will interact with it
* *Private key* is *an Ethereum account private key*
* *TOKEN* is obtained in response to authentication for one session.
* *TOKEN* is passed into the Rest API via the *X-MBX-TOKEN* header
* *TOKENs* and *private keys* **are case sensitive**
* Security Types:
    * NONE - Endpoint can be accessed freely
    * SIGN - Endpoint requires sending a valid TOKEN and Signature
    * COMMON - Endpoint requires sending a valid TOKEN

### SIGNED Endpoint security

* *SIGNED* endpoints require an additional parameter, *signature*, to be sent in the *request body*.
* Endpoints use *HMAC SHA256* signatures. The *HMAC SHA256* signature is a keyed *HMAC SHA256* operation. Use your *Ethereum private key* as the key and *totalParams* as the value for the HMAC operation
* The *signature* is not **case sensitive**
* *totalParams* is defined as the *query string* concatenated with the *request body*

## Test

### Ping

* **Security:** None

* **GET** /api/v1/ping

* **Description:**
Test connectivity to the Rest API

* **Weight:** 1

* **Parameters:** NONE

* **Response body:** application/json
```
{}
```

## Entry part

### Registration

* **Security:** SIGN

* **POST** /api/v1/join

* **Description:**
Register user and add user credentials

* **Weight:** 1

* **Parameters:** NONE

* **Request body:** application/json

```
{
    "pubkey": "168bs115a2ee09042d83d7c5811b533620531f67",
    "username": "Follow1",
    "password": "Follow123",
    "signature": "090bs115a2ee09042d83d7c5811b421520531859"
}

```
* **Response body:** application/json
```
{}
```

### Authentication

* **Security:** NONE

* **POST** /api/v1/auth

* **Description:**
Authenticate user on server

* **Weight:** 1

* **Parameters:** NONE

* **Request body:** application/json

```
{
    "username": "Follow1",
    "password": "Follow123"
}

```
* **Response body:** application/json
```
{    
    "userId": 123456,
    "token": "168bs115a2ee09042d83d7c5811b533620531f67"
}
```

## User data

### Set user data

* **Security:** SIGN

* **POST** /api/v1/user/set

* **Description:**
Set user info

* **Weight:** 1

* **Parameters:** NONE

* **Request body:** application/json

```
{
    "name": {
        "first": "Vasiliy",
        "middle": "Vasilievich",
        "second": "Vasiliev"
    },
    "phone": "8-800-555-3535",
    "email": "follow123@gmail.com",
    "country": "Russia",
    "image_url": "https://image/png",
    "birthday": 1544422247000,
    "regdate": 1544422247000,
    "signature": "090bs115a2ee09042d83d7c5811b421520531859"
    
}

```
* **Response body:** application/json
```
{}
```

### Get user data (SECURE)

* **Security:** COMMON

* **GET** /api/v1/user/{id}

* **Description:**
Get user info

* **Weight:** 1

* **Parameters:**
    * **id**:
        * *Type*: Integer
        * *In*: Path
        * *Description*: User id
        * *Example*: 123

* **Response body:** application/json
```
{
    "name": {
        "first": "Vasiliy",
        "middle": "Vasilievich",
        "second": "Vasiliev"
    },
    "username": "Follow123",
    "phone": "8-800-555-3535",
    "email": "follow123@gmail.com",
    "country": "Russia",
    "image_url": "https://image/png",
    "birthday": 1544422247000,
    "regdate": 1544422247000
}
```

### Get user data (NON-SECURE)

* **Security:** NONE

* **GET** /api/v1/user/{id}

* **Description:**
Get user info

* **Weight:** 1

* **Parameters:**
    * **id**:
        * *Type*: Integer
        * *In*: Path
        * *Description*: User id
        * *Example*: 123

* **Response body:** application/json
```
{
    "name": {
        "first": "Vasiliy",
        "middle": "Vasilievich",
        "second": "Vasiliev"
    },
    "username": "Follow123",
    "country": "Russia",
    "image_url": "https://image/png",
    "birthday": 1544422247000
}
```

## Trader data

### Get trader data

* **Security:** NONE

* **GET** /api/v1/user/{id}/trader

* **Description:**
Get trader info

* **Weight:** 1

* **Parameters:**
    * **id**:
        * *Type*: Integer
        * *In*: Path
        * *Description*: User id
        * *Example*: 123

* **Response body:** application/json
```
{
    "followersBalance": 1000000,
    "followersCount": 1,
    "followersChange": {
        "daily": -1.5
        "monthly": -5.5,
        "yearly": 5.1,
        "alltime": 5.1
    },
    "selfChange": {
        "daily": -1.5
        "monthly": -5.5,
        "yearly": 5.1,
        "alltime": 5.1
    }
}
```

### Get traders followers list

* **Security:** NONE

* **GET** /api/v1/user/{id}/trader/followers

* **Description:**
Get trader followers list

* **Weight:** 1

* **Parameters:**
    * **id**:
        * *Type*: Integer
        * *In*: Path
        * *Description*: User id
        * *Example*: 123

* **Response body:** application/json
```
[
    {
        "id": 3,
        "username": "HOWHOWHOW",
        "imageUrl": "https://image.png",
        "deposit": 100.0,
        "balance": 105.1,
        "startDate": 1544422247000,
        "change": {
            "daily": -1.5
            "monthly": -5.5,
            "yearly": 5.1,
            "alltime": 5.1
        }
    }
]
```

### Get trader balance data

* **Security:** NONE

* **GET** /api/v1/user/{id}/trader/balance/start={startTime}&end={endTime}&period={period}

* **Description:**
Get trader balance state

* **Weight:** 1

* **Parameters:**
    * **id**:
        * *Type*: Integer
        * *In*: Path
        * *Description*: User id
        * *Example*: 123
    * **startTime**:
        * *Type*: Integer
        * *In*: Path
        * *Description*: Start timestamp on graph
        * *Example*: 1544422247000
    * **endTime**:
        * *Type*: Integer
        * *In*: Path
        * *Description*: End timestamp on graph
        * *Example*: 1544422247001
    * **period**:
        * *Type*: Integer
        * *In*: Path
        * *Description*: Graph period
        * *Example*: 600

* **Response body:** application/json
```
{
    "currentBalance": 100,
    "states": [
        {
            "timestamp": 1544422247000,
            "balance": "99",
        },
        {
            "timestamp": 1544422247001,
            "balance": "100",
        }
    ]
}
```

### Get trader deals data

* **Security:** COMMON

* **GET** /api/v1/user/{id}/trader/deals/{page}

* **Description:**
Get trader deals list

* **Weight:** 1

* **Parameters:**
    * **id**:
        * *Type*: Integer
        * *In*: Path
        * *Description*: User id
        * *Example*: 123
    * **page**:
        * *Type*: Integer
        * *In*: Path
        * *Description*: Page of traders deals list
        * *Example*: 1

* **Response body:** application/json
```
[
    {
        "timestamp": 1544422247000,
        "symbol": "ETHW3S",
        "price": "1.00000000",
        "status": "filled",
        "side": "sell",
        "stopPrice": "0.0",
        "invested": "7.84" //precentage amount from capital
    }
]
```

## Follower

### Get follower data

* **Security:** NONE

* **GET** /api/v1/user/{id}/follower

* **Description:**
Get follower info

* **Weight:** 1

* **Parameters:**
    * **id**:
        * *Type*: Integer
        * *In*: Path
        * *Description*: User id
        * *Example*: 123

* **Response body:** application/json
```
{
    "deposit": 100,
    "balance": 105,
    "tradersCount": 1,
    "tradersChange": {
        "daily": -1.5
        "monthly": -5.5,
        "yearly": 5.1,
        "alltime": 5.1
    },
    "selfChange": {
        "daily": -1.5
        "monthly": -5.5,
        "yearly": 5.1,
        "alltime": 5.1
    }
}
```

### Get follower traders list

* **Security:** NONE

* **GET** /api/v1/user/{id}/follower/traders

* **Description:**
Get follower traders list

* **Weight:** 1

* **Parameters:**
    * **id**:
        * *Type*: Integer
        * *In*: Path
        * *Description*: User id
        * *Example*: 123

* **Response body:** application/json
```
[
    {
        "id": 3,
        "username": "HOWHOWHOW",
        "imageUrl": "https://image.png",
        "deposit": 100.0,
        "balance": 105.1,
        "startDate": 1544422247000,
        "change": {
            "daily": -1.5
            "monthly": -5.5,
            "yearly": 5.1,
            "alltime": 5.1
        }
    }
]
```

### Get follower balance data

* **Security:** NONE

* **GET** /api/v1/user/{id}/follower/balance/start={startTime}&end={endTime}&period={period}

* **Description:**
Get follower balance state

* **Weight:** 1

* **Parameters:**
    * **id**:
        * *Type*: Integer
        * *In*: Path
        * *Description*: User id
        * *Example*: 123
    * **startTime**:
        * *Type*: Integer
        * *In*: Path
        * *Description*: Start timestamp on graph
        * *Example*: 1544422247000
    * **endTime**:
        * *Type*: Integer
        * *In*: Path
        * *Description*: End timestamp on graph
        * *Example*: 1544422247001
    * **period**:
        * *Type*: Integer
        * *In*: Path
        * *Description*: Graph period
        * *Example*: 600

* **Response body:** application/json
```
{
    "currentBalance": 100,
    "states": [
        {
            "timestamp": 1544422247000,
            "balance": "99",
        },
        {
            "timestamp": 1544422247001,
            "balance": "100",
        }
    ]
}
```

## Social

### Traders list

* **Security:** NONE

* **GET** /api/v1/traders/{sort}/{limit}/{page}

* **Description:**
Get traders list

* **Weight:** 1

* **Parameters:**
    * **sort**:
        * *Type*: Integer
        * *In*: Path
        * *Description*: Sorting type id representation (ex: 1 - by profit per year, etc.)
        * *Example*: 1
    * **limit**:
        * *Type*: Integer
        * *In*: Path
        * *Description*: Amount per page
        * *Example*: 100
    * **page**:
        * *Type*: Integer
        * *In*: Path
        * *Description*: Page of traders list
        * *Example*: 10

* **Response body:** application/json
```
[
    "id": 3,
    "username": "HOWHOWHOW",
    "imageUrl": "https://image.png",
    "followersBalance": 1000000,
    "followersCount": 1,
    "followersChange": {
        "daily": -1.5
        "monthly": -5.5,
        "yearly": 5.1,
        "alltime": 5.1
    },
    "selfChange": {
        "daily": -1.5
        "monthly": -5.5,
        "yearly": 5.1,
        "alltime": 5.1
    }
]
```

## Transactions

### Deposit

* **Security:** SIGN

* **POST** /api/v1/deposit

* **Description:**
Deposit money to some follow-trade account

* **Weight:** 1

* **Parameters:** NONE

* **Request body:** application/json

```
{
    "amount": "14.04",
    "timestamp": 1544422247000,
    "followerId": 1,
    "traderId": 2,
    "txHash": "090bs115a2ee09042d83d7c5811b421520531859",
    "signature": "090bs115a2ee09042d83d7c5811b421520531812"
}
```

* **Response body:** application/json
```
{}
```

### Withdraw

* **Security:** SIGN

* **POST** /api/v1/withdraw

* **Description:**
Withdraw money from some follow-trade account

* **Weight:** 1

* **Parameters:** NONE

* **Request body:** application/json

```
{
    "amount": "14.04",
    "timestamp": 1544422247000,
    "followerId": 1,
    "traderId": 2,
    "signature": "090bs115a2ee09042d83d7c5811b421520531812"
}
```

* **Response body:** application/json
```
{}
```
