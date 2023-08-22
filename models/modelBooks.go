package models

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Books struct {
	Id       primitive.ObjectID `json:"_id" bson:"_id"`
	ISBN     int64              `json:"isbn" bson:"isbn"`
	Title    string             `json:"title"  bson:"title"`
	Author   string             `json:"author" bson:"author"`
	Summary  string             `json:"summary" bson:"summary"`
	ImageSrc string             `json:"image_src" bson:"image_src"`
}

type FormBook struct {
	ISBN     int64  `json:"isbn" bson:"isbn" form:"isbn" binding:"required"`
	Title    string `json:"title" form:"title" bson:"title" binding:"required"`
	Author   string `json:"author" form:"author" bson:"author" binding:"required"`
	Summary  string `json:"summary" form:"summary" bson:"summary" binding:"required"`
	ImageSrc string `json:"image_src" form:"image_src" bson:"image_src"`
}

type Price struct {
	Currency string  `json:"currency" form:"currency" bson:"currency" binding:"required"`
	Value    float64 `json:"value" form:"value" bson:"value" binding:"required"`
}

type InsertOneResult struct {
	Id interface{} `json:"_id" bson:"_id"`
	FormBook
}

func (model Model) InsertBook(formBook FormBook) (InsertOneResult, error) {
	collection := model.db.Database("books_db").Collection("books")

	ctx, cancel := defaultContext()
	defer cancel()

	lastId, err := collection.InsertOne(ctx, formBook)
	if err != nil {
		return InsertOneResult{}, err
	}

	var insertedResult = InsertOneResult{
		Id:       lastId.InsertedID,
		FormBook: formBook,
	}

	return insertedResult, nil
}

func (model Model) SelectBooks() ([]Books, error) {
	var books []Books

	collection := model.db.Database("books_db").Collection("books")

	ctx, cancel := defaultContext()
	defer cancel()

	query := bson.M{}

	cursor, err := collection.Find(ctx, query)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var book Books
		if err := cursor.Decode(&book); err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, err
}
