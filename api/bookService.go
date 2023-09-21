package api

import (
	"books/model"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

func getBooks(c *fiber.Ctx) error {
	cursor, err := model.MngInst.Collection.Find(c.Context(), bson.D{{}})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	var books []model.Book
	if err := cursor.All(c.Context(), &books); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}
	return c.JSON(books)
}

func getBook(c *fiber.Ctx) error {
	filter := bson.M{"isbn": c.AllParams()["isbn"]}
	var result model.Book
	if err := model.MngInst.Collection.FindOne(c.Context(), filter).Decode(&result); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}
	return c.JSON(result)
}

func addBook(c *fiber.Ctx) error {
	book := new(model.Book)

	if err := c.BodyParser(book); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}
	_, err := model.MngInst.Collection.InsertOne(c.Context(), book)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}
	return c.JSON(http.StatusOK)
}
