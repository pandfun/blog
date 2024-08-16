package controller

import (
	"blog/database"
	"blog/model"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func BlogList(c *fiber.Ctx) error {

	ctx := fiber.Map{
		"statusText": "Ok",
		"msg":        "Blog List",
	}

	db := database.DBConn

	var records []model.Blog

	db.Find(&records)

	ctx["blogRecords"] = records

	c.Status(http.StatusOK)
	return c.JSON(ctx)
}

func BlogCreate(c *fiber.Ctx) error {

	ctx := fiber.Map{
		"statusText": "Ok",
		"msg":        "Add a Blog",
	}

	record := new(model.Blog)

	if err := c.BodyParser(&record); err != nil {

		log.Println("Error in parsing the request ", err.Error())

		ctx["statusText"] = "Fail"
		ctx["msg"] = "Failed to parse the incoming request data"

		return c.JSON(ctx)
	}

	res := database.DBConn.Create(record)

	if res.Error != nil {

		log.Println("Error in saving data to database ", res.Error.Error())

		ctx["statusText"] = "Fail"
		ctx["msg"] = "Failed to save the data to the database"

		return c.JSON(ctx)
	}

	ctx["msg"] = "Record saved succesfully"
	ctx["data"] = record

	c.Status(http.StatusCreated)
	return c.JSON(ctx)
}

func BlogUpdate(c *fiber.Ctx) error {

	ctx := fiber.Map{
		"statusText": "Ok",
		"msg":        "Update Blog",
	}

	blogId := c.Params("id")

	var record model.Blog

	database.DBConn.First(&record, blogId)

	if record.ID == 0 {

		log.Println("Record not found")

		c.Status(http.StatusBadRequest)

		ctx["statusText"] = "Fail"
		ctx["msg"] = "No record found with matching ID"

		return c.JSON(ctx)
	}

	if err := c.BodyParser(&record); err != nil {

		log.Println("Error in parsing the request ", err.Error())

		ctx["statusText"] = "Fail"
		ctx["msg"] = "Failed to parse the incoming request data"

		return c.JSON(ctx)
	}

	res := database.DBConn.Save(record)

	if res.Error != nil {

		log.Println("Error in saving data to database ", res.Error.Error())

		ctx["statusText"] = "Fail"
		ctx["msg"] = "Failed to save the data to the database"

		return c.JSON(ctx)
	}

	ctx["msg"] = "Record has been updated!"
	ctx["data"] = record

	c.Status(http.StatusOK)
	return c.JSON(ctx)
}

func BlogDelete(c *fiber.Ctx) error {

	ctx := fiber.Map{
		"statusText": "Ok",
		"msg":        "Delete Blog",
	}

	c.Status(http.StatusOK)
	return c.JSON(ctx)
}
