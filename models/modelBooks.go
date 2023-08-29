package models

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Books struct {
	Id       primitive.ObjectID `json:"_id" bson:"_id"`
	ISBN     int64              `json:"isbn" bson:"isbn"`
	Title    string             `json:"title"  bson:"title"`
	Author   string             `json:"author" bson:"author"`
	Summary  string             `json:"summary" bson:"summary"`
	ImageSrc string             `json:"image_src" bson:"image_src"`
	Price    struct {
		Currency string  `json:"currency" bson:"currency"`
		Value    float64 `json:"value" bson:"value"`
	} `json:"price" bson:"price"`
}

type FormBook struct {
	ISBN     int64  `json:"isbn" bson:"isbn" binding:"required"`
	Title    string `json:"title"  bson:"title" binding:"required"`
	Author   string `json:"author" bson:"author" binding:"required"`
	Summary  string `json:"summary" bson:"summary" binding:"required"`
	ImageSrc string `json:"image_src" bson:"image_src" binding:"required"`
	Price    struct {
		Currency string  `json:"currency" bson:"currency" binding:"required"`
		Value    float64 `json:"value" bson:"value" binding:"required"`
	} `json:"price" bson:"price"`
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

func (model Model) UpdateBook(bookId primitive.ObjectID, formBook Books) (Books, error) {
	collection := model.db.Database("books_db").Collection("books")

	ctx, cancel := defaultContext()
	defer cancel()

	opts := options.FindOneAndUpdate().SetUpsert(true)
	filter := bson.D{{Key: "_id", Value: bookId}}
	update := bson.D{{Key: "$set", Value: formBook}}

	var updatedDocument bson.M
	err := collection.FindOneAndUpdate(
		ctx,
		filter,
		update,
		opts,
	).Decode(&updatedDocument)
	if err != nil {
		return formBook, mongo.ErrNoDocuments
	}

	return formBook, nil
}

func (model Model) DeleteBook(bookId primitive.ObjectID) (int64, error) {
	collection := model.db.Database("books_db").Collection("books")

	ctx, cancel := defaultContext()
	defer cancel()

	query := bson.M{"_id": bookId}

	result, err := collection.DeleteOne(ctx, query)
	if err != nil {
		return 0, err
	}

	return result.DeletedCount, err
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

func (model Model) SelectBookById(id primitive.ObjectID) Books {
	var book Books

	collection := model.db.Database("books_db").Collection("books")

	ctx, cancel := defaultContext()
	defer cancel()

	query := bson.M{"_id": id}

	err := collection.FindOne(ctx, query).Decode(&book)
	if err != nil {
		return book
	}

	return book
}
