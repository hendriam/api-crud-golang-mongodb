# API CRUD USING GOLANG & MONGODB

I have configured this resource with mongodb and logging.

Kindly do the installation here [Mongodb](https://www.mongodb.com/docs/manual/installation/). For database, kindly adjust on the DNS in the lib/config.go file

For logging I use [Zerolog](https://pkg.go.dev/github.com/rs/zerolog)

For http server I use [GIN](https://gin-gonic.com/)

## Usage

```http
   
CREATE
POST http://localhost:8080/book/create
Content-Type: application/json

REQUEST
{
  "ISBN": 9781841499789,
  "title": "Bloodfire Quest",
  "author": "Terry Brooks",
  "summary": "The adventure that started in Wards of Faerie takes a thrilling new turn.",
  "image_src": "http://s.s-bol.com/imgbase0/imagebase/large/FC/7/0/0/7/9200000009027007.jpg",
  "price": {
    "currency": "EUR",
    "value": 24.5
  }
}

RESPONSE
{
	"code": 200,
	"success": true,
	"message": "Created succesfully",
	"data": {
		"_id": "64ecde332b77172101c40dd3",
		"isbn": 9781841499789,
		"title": "Bloodfire Quest",
		"author": "Terry Brooks",
		"summary": "The adventure that started in Wards of Faerie takes a thrilling new turn.",
		"image_src": "http://s.s-bol.com/imgbase0/imagebase/large/FC/7/0/0/7/9200000009027007.jpg",
		"price": {
			"currency": "EUR",
			"value": 24.5
		}
	}
}

READ
GET http://localhost:8080/books
Content-Type: application/json

RESPONSE
{
	"code": 200,
	"success": true,
	"message": "Data load succesfully",
	"data": [
		{
			"_id": "64eccc86cab5779bb6a6ea47",
			"isbn": 9789022558027,
			"title": "Magic staff",
			"author": "Terry Brooks",
			"summary": "Vijf eeuwen geleden werd de wereld door een noolottige demonenoorlog in de as gelegd. De overlevenden hebben een toevluchtsoord gevonden in een door magie beschermde vallei, maar nu staat een genadeloos leger op het punt de vallei binnen te vallen.",
			"image_src": "http://s.s-bol.com/imgbase0/imagebase/large/FC/2/2/5/2/9200000002212522.jpg",
			"price": {
				"currency": "EUR",
				"value": 17.5
			}
		},
		{
			"_id": "64ecde0f2b77172101c40dd2",
			"isbn": 9780552159722,
			"title": "Deception point",
			"author": "Dan Brown",
			"summary": "When a new NASA satellite detects evidence of an astonishingly rare object buried deep in the Arctic ice, the floundering space agency proclaims a much-needed victory.",
			"image_src": "http://s.s-bol.com/imgbase0/imagebase/large/FC/8/8/9/8/1001004006878988.jpg",
			"price": {
				"currency": "EUR",
				"value": 9.99
			}
		},
		{
			"_id": "64ecde332b77172101c40dd3",
			"isbn": 9781841499789,
			"title": "Bloodfire Quest",
			"author": "Terry Brooks",
			"summary": "The adventure that started in Wards of Faerie takes a thrilling new turn.",
			"image_src": "http://s.s-bol.com/imgbase0/imagebase/large/FC/7/0/0/7/9200000009027007.jpg",
			"price": {
				"currency": "EUR",
				"value": 24.5
			}
		}
	]
}

READ ONE
GET http://localhost:8080/book/64ecde332b77172101c40dd3
Content-Type: application/json

RESPONSE
{
	"code": 200,
	"success": true,
	"message": "Data load succesfully",
	"data": {
		"_id": "64ecde332b77172101c40dd3",
		"isbn": 9781841499789,
		"title": "Bloodfire Quest",
		"author": "Terry Brooks",
		"summary": "The adventure that started in Wards of Faerie takes a thrilling new turn.",
		"image_src": "http://s.s-bol.com/imgbase0/imagebase/large/FC/7/0/0/7/9200000009027007.jpg",
		"price": {
			"currency": "EUR",
			"value": 24.5
		}
	}
}

UPDATE
PUT http://localhost:8080/book/update/64eccc86cab5779bb6a6ea47
Content-Type: application/json

REQUEST
{
	"_id": "64eccc86cab5779bb6a6ea47",
	"ISBN": 9781841499789,
    "title": "Bloodfire Quest",
    "author": "Terry Brooks",
    "summary": "The adventure.",
    "image_src": "http://s.s-bol.com/imgbase0/imagebase/large/FC/7/0/0/7/9200000009027007.jpg",
    "price": {
      "currency": "EUR",
      "value": 24.5
    }
}

RESPONSE
{
	"code": 200,
	"success": true,
	"message": "updated succesfull",
	"data": {
		"_id": "64eccc86cab5779bb6a6ea47",
		"isbn": 9781841499789,
		"title": "Bloodfire Quest",
		"author": "Terry Brooks",
		"summary": "The adventure.",
		"image_src": "http://s.s-bol.com/imgbase0/imagebase/large/FC/7/0/0/7/9200000009027007.jpg",
		"price": {
			"currency": "EUR",
			"value": 24.5
		}
	}
}

DELETE
DELETE http://localhost:8080/book/delete/64eccc55cab5779bb6a6ea46
Content-Type: application/json

RESPONSE
{
	"code": 200,
	"success": true,
	"message": "delete succesfull",
	"data": null
}

```