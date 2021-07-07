package controllers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/ahmedkrdzalic/StackAndFlow/database"
	"github.com/ahmedkrdzalic/StackAndFlow/models"
	"github.com/gofiber/fiber/v2"
)

func Newquestion(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var name string
	database.DB.Raw("SELECT name FROM users WHERE id = ?", data["questioner_id"]).Scan(&name)
	fmt.Printf(data["questioner_id"])
	question := models.Question{
		Questioner_id:   data["questioner_id"],
		Questioner_name: name,
		Title:           data["title"],
		Content:         data["content"],
		Upvotes:         0,
		Downvotes:       0,
		Date_time:       time.Now(),
	}

	database.DB.Create(&question)

	return c.JSON(question)

}

func Getquestion(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var question models.Question

	database.DB.Where("id = ?", data["id"]).First(&question)

	var answers []models.Answer

	database.DB.Raw("SELECT * FROM answers WHERE question_id = ?", data["id"]).Scan(&answers)

	if question.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "answers for that question not found",
		})
	}

	return c.JSON(fiber.Map{
		"question": question,
		"answers":  answers,
	})

}

func Getquestions(c *fiber.Ctx) error {
	var questions []models.Question

	database.DB.Raw("SELECT * FROM questions ORDER BY date_time DESC LIMIT 20").Scan(&questions)

	if questions == nil {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "questions not found",
		})
	}

	return c.JSON(&questions)
}

func Newanswer(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	Question_id, _ := strconv.ParseUint(data["question_id"], 0, 64)
	Answerer_id, _ := strconv.ParseUint(data["answerer_id"], 0, 64)

	var name string
	//I know that I have to use query builder but it is a bit complicated in GORM for now, and I do not have too much time.
	//I was defenitely using it in other projects before, but for now it is not required in this project requirements.
	database.DB.Raw("SELECT name FROM users WHERE id = ?", data["answerer_id"]).Scan(&name)

	answer := models.Answer{
		Question_id:   Question_id,
		Answerer_id:   Answerer_id,
		Answerer_name: name,
		Content:       data["content"],
		Date_time:     time.Now(),
	}

	database.DB.Create(&answer)

	return c.JSON(answer)

}
