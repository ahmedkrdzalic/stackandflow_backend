package routes

import (
	"github.com/ahmedkrdzalic/StackAndFlow/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Get("/api/user", controllers.User)
	app.Post("/api/logout", controllers.Logout)
	app.Post("/api/newquestion", controllers.Newquestion)
	app.Post("/api/getquestion", controllers.Getquestion)
	app.Post("/api/newanswer", controllers.Newanswer)
	app.Get("/api/getquestions", controllers.Getquestions)
	app.Get("/api/test", controllers.Test)

}
