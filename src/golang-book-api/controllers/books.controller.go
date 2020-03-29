package controllers

import (
	"context"
	"encoding/json"
	"net/http"

	utils "golang-book-api/utils"

	m "golang-book-api/model"

	db "golang-book-api/database"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetBooks function retrieves all books
func GetBooks(w http.ResponseWriter, r *http.Request) {
	utils.LoggerIn("GetBooks")
	booksCollection := db.GetCollection("books")
	cursor, err := booksCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		utils.ErrorResponse(w, utils.ErrDatabase, "")
		return
	}
	books := []m.Book{}
	if err = cursor.All(context.TODO(), &books); err != nil {
		utils.ErrorResponse(w, utils.ErrQuery, "")
		return
	}
	utils.Logger.Debugf("Books array len - %d", len(books))
	json.NewEncoder(w).Encode(books)
	utils.LoggerOut("GetBooks")
}

// GetBook function retrieves book with given ID
func GetBook(w http.ResponseWriter, r *http.Request) {
	utils.LoggerIn("GetBook")

	params := mux.Vars(r)
	utils.Logger.Info("Getting Book with ID ", params["id"])
	objID, _ := primitive.ObjectIDFromHex(params["id"])

	booksCollection := db.GetCollection("books")
	var foundBook m.Book
	err := booksCollection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&foundBook)
	if err != nil {
		utils.ErrorResponse(w, utils.ErrNotFound, " Book with ID "+params["id"])
		return
	}

	utils.Logger.Debugf("Found Book [%+v]", foundBook)
	json.NewEncoder(w).Encode(foundBook)
	utils.LoggerOut("GetBook")
}

//CreateBook add new book
func CreateBook(w http.ResponseWriter, r *http.Request) {
	utils.LoggerIn("CreateBook")
	newBook, err := utils.CreateBookParseRequest(w, r)
	utils.Logger.Debug("Request ", newBook)
	if err != nil {
		return
	}
	if newBook.Title == "" {
		utils.ErrorResponse(w, utils.ErrArgsMissing, "Title")
		return
	}
	booksCollection := db.GetCollection("books")
	result, err := booksCollection.InsertOne(context.TODO(), newBook)
	if err != nil {
		utils.ErrorResponse(w, utils.ErrInsert, "")
		return
	}
	utils.Logger.Debugf("Book added successfully [BOOK - %+v]", newBook)
	json.NewEncoder(w).Encode(result.InsertedID)
	utils.LoggerOut("CreateBook")
}

//UpdateBook updates book
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	utils.LoggerIn("UpdateBook")
	params := mux.Vars(r)
	utils.Logger.Info("Getting Book with ID ", params["id"])
	objID, _ := primitive.ObjectIDFromHex(params["id"])

	updateBookObj, err := utils.CreateBookParseRequest(w, r)
	utils.Logger.Debug("Request ", updateBookObj)
	if err != nil {
		return
	}
	booksCollection := db.GetCollection("books")
	var foundBook m.Book
	err = booksCollection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&foundBook)
	if err != nil {
		utils.ErrorResponse(w, utils.ErrNotFound, " Book with ID "+params["id"])
		return
	}
	if updateBookObj.Title == "" {
		updateBookObj.Title = foundBook.Title
	}
	if updateBookObj.Category == "" {
		updateBookObj.Category = foundBook.Category
	}
	if nil == updateBookObj.Author {
		updateBookObj.Author = foundBook.Author
	}
	if len(updateBookObj.Tags) == 0 {
		updateBookObj.Tags = foundBook.Tags
	}
	res := booksCollection.FindOneAndUpdate(context.TODO(), bson.M{"_id": objID}, bson.M{"$set": updateBookObj})

	utils.Logger.Debugf("Book updated successfully [BOOK - %+v]", updateBookObj)
	json.NewEncoder(w).Encode(res)
	utils.LoggerOut("UpdateBook")
}

// DeleteBook function deletes book with given ID
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	utils.LoggerIn("DeleteBook")

	params := mux.Vars(r)
	utils.Logger.Info("Deleting Book with ID ", params["id"])
	objID, _ := primitive.ObjectIDFromHex(params["id"])
	booksCollection := db.GetCollection("books")

	var foundBook m.Book
	err := booksCollection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&foundBook)
	if err != nil {
		utils.ErrorResponse(w, utils.ErrNotFound, " Book with ID "+params["id"])
		return
	}
	deleteResult, err := booksCollection.DeleteOne(context.TODO(), bson.M{"_id": objID})
	if err != nil {
		utils.ErrorResponse(w, utils.ErrNotFound, " Book with ID "+params["id"])
		return
	}

	json.NewEncoder(w).Encode(deleteResult.DeletedCount)
	utils.LoggerOut("DeleteBook")
}
