package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Movie struct {
	Title string `json:"title"`
	Id    int    `json:"id"`
}

func MoviesRoutes(r fiber.Router) {
	movies := []*Movie{
		{Id: 1, Title: "Kimetsu no yaiba"},
		{Id: 2, Title: "Dragon ball z"},
	}

	r.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"movies": movies,
		})
	})

	r.Get("/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.JSON(fiber.Map{
				"err": err,
			})
		}

		movie := Movie{}

		for _, item := range movies {
			if item.Id == id {
				movie = *item
			}
		}

		return c.JSON(fiber.Map{
			"movie": movie,
		})
	})

	r.Post("/", func(c *fiber.Ctx) error {
		movie := Movie{}
		c.BodyParser(&movie)
		movie.Id = len(movies) + 1

		movies = append(movies, &movie)
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"movie": movie,
		})
	})

	r.Put("/:id", func(c *fiber.Ctx) error {
		movie := new(Movie)

		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "There's a error",
			})
		}

		c.BodyParser(&movie)
		movie.Id = id

		for _, item := range movies {
			if item.Id == id {
				item.Title = movie.Title
			}
		}

		fmt.Println(movies)

		return c.JSON(fiber.Map{
			"moviee": movie,
		})
	})

	r.Delete("/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "There's a error",
			})
		}

		for i, item := range movies {
			if item.Id == id {
				movies = append(movies[:i], movies[i+1:]...)
			}
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"id":  id,
			"msg": "Movie has deleted",
		})
	})
}
