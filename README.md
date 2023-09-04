# Package Size Calculator

This is a _Package Size_ challenge solving coded on GoLang.
This project aims to calculate and return the fewest number of package 
where the package sizes are dynamically retrieving from the user as well as the item count.

## Original Challenge

Imagine for a moment that one of our product lines ships in various pack sizes: 
* 250 Items
* 500 Items
* 1000 Items
* 2000 Items
* 5000 Items

Our customers can order any number of these items through our website, but they will always only be given complete packs.

Only whole packs can be sent. Packs cannot be broken open.

Within the constraints of Rule 1 above, send out no more items than necessary to fulfil the order.

Within the constraints of Rules 1 & 2 above, send out as few packs as possible to fulfil each order.

### Please look at the following 5 examples:

Items Ordered: 1
Correct number of packs: 1x250

Items Ordered: 250
Correct number of packs: 1x250

Items Ordered: 251
Correct number of packs: 1x500

Items Ordered: 501
Correct number of packs: 1x500, 1x250

Items Ordered: 12001
Correct number of packs: 2x5000, 1x2000, 1x250

Write an application that can calculate the number of packs we need to ship to the customer.
The API must be written in Golang & be usable by an HTTP API (by whichever method you choose).

### Optional:
Keep your application flexible so that pack sizes can be changed and added and
removed without having to change the code.
We look forward to receiving your response!

execution time limit: 4 seconds (go)

memory limit: 1 GB

## How To Run

##### Requirements

* [Go](https://go.dev/)

#### Run the App
```bash
git clone https://github.com/boraerciyas/package-size-calculator-go.git
git build main.go
git run main.go
```

App runs on port '10000' by default.

## REST API
The REST API to the app is described below.

### Calculate Packages

#### Request

`POST /calculatePackages`

    curl --location 'http://localhost:10000/calculatePackages' --header 'Content-Type: application/json' --data '{"packageSizeList": [250,500,1000,2500],"orderedItems": 2750}'

### Response

    HTTP/1.1 200 OK
    Date: Mon, 04 Sep 2023 12:36:30 GMT
    Status: 200 OK
    Connection: close
    Content-Type: application/json
    Location: /calculatePackages
    Content-Length: 127

    {"itemsOrdered":2750,"correctNumberOfPacks":[{"packageSize":2500,"numberOfPackage":1},{"packageSize":250,"numberOfPackage":1}]}

