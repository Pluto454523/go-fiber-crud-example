package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

// ** Variable for mock data Book struct
var books []Book

func main() {

	// ** New instant of fiber app
	app := fiber.New()

	// ** Apply CORS(Cross-origin) middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // Adjust this to be more restrictive if needed
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// ** Add mock data to books
	books = append(books, Book{ID: 1, Title: "Harry Potter", Author: "J.K Rolling"})
	books = append(books, Book{ID: 2, Title: "One piece", Author: "Eijiro Oda"})

	// ** Route
	app.Get("books", getBooks)
	app.Get("books/:id", getBookById)
	app.Post("books", createBook)
	app.Put("books/:id", updateBook)
	app.Delete("books/:id", deleteBook)

	// ** Route upload image
	app.Post("/upload", uploadImage)

	// ** Get Envoriment form .env file
	port := os.Getenv("PORT")

	// ** Listen app to port 3000
	app.Listen(":" + port)
}

// ** Read all
func getBooks(c *fiber.Ctx) error {

	return c.JSON(books)
}

// ** Read by id
func getBookById(c *fiber.Ctx) error {

	bookId, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	for i := range books {
		if books[i].ID == bookId {
			return c.JSON(books[i])
		}
	}
	return c.Status(fiber.StatusNotFound).SendString(fmt.Sprintf("book id %v not found", bookId))
}

// ** Create
func createBook(c *fiber.Ctx) error {

	// newBook := new(Book)
	newBook := &Book{}

	if err := c.BodyParser(newBook); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	newBook.ID = len(books) + 1

	// ** dereference newBook เนื่องจากถูกเก็บ pointer ไว้
	books = append(books, *newBook)

	return c.JSON(newBook)
}

// ** Update
func updateBook(c *fiber.Ctx) error {
	bookId, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	bookUpdate := &Book{}
	if err := c.BodyParser(bookUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	for i := range books {
		if books[i].ID == bookId {

			if bookUpdate.Title != "" {
				books[i].Title = bookUpdate.Title
			}

			if bookUpdate.Author != "" {
				books[i].Author = bookUpdate.Author
			}

			return c.JSON(books[i])
		}
	}

	return c.Status(fiber.StatusNotFound).SendString(fmt.Sprintf("book id %v not found", bookId))
}

// ** Delete
func deleteBook(c *fiber.Ctx) error {

	bookId, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	for i := range books {
		if books[i].ID == bookId {

			// ** books[i+1:]... เพื่อส่งค่าไปทีละตัวเพื่อ append
			books = append(books[:i], books[i+1:]...)
			return c.SendStatus(fiber.StatusOK)
		}
	}

	return c.Status(fiber.StatusNotFound).SendString(fmt.Sprintf("book id %v not found", bookId))
}

// ** Upload file
func uploadImage(c *fiber.Ctx) error {
	// Read file from request
	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	// Save the file to the server
	err = c.SaveFile(file, "./uploads/"+file.Filename)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.SendString("File uploaded successfully: " + file.Filename)
}
