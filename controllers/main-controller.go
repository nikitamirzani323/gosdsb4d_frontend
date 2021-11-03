package controllers

import (
	"log"
	"net/http"
	"time"

	"bitbucket.org/isbtotogroup/sdsb4d-frontend/configs"
	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
)

const PATH string = configs.PATH_API

// type clientlistresult struct {
// 	Token string `json:"token"`
// }

type response struct {
	Status int         `json:"status"`
	Record interface{} `json:"record"`
}

func ListResultHome(c *fiber.Ctx) error {
	// client := new(clientlistresult)
	render_page := time.Now()
	// if err := c.BodyParser(client); err != nil {
	// 	c.Status(fiber.StatusBadRequest)
	// 	return c.JSON(fiber.Map{
	// 		"status":  fiber.StatusBadRequest,
	// 		"message": err.Error(),
	// 		"record":  nil,
	// 	})
	// }

	axios := resty.New()
	resp, err := axios.R().
		SetResult(response{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_company": "",
		}).
		Post(PATH + "api/resultsdsbday")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*response)
	return c.JSON(fiber.Map{
		"status": http.StatusOK,
		"record": result.Record,
		"time":   time.Since(render_page).String(),
	})
}
func ListResultHomeNight(c *fiber.Ctx) error {
	// client := new(clientlistresult)
	render_page := time.Now()
	// if err := c.BodyParser(client); err != nil {
	// 	c.Status(fiber.StatusBadRequest)
	// 	return c.JSON(fiber.Map{
	// 		"status":  fiber.StatusBadRequest,
	// 		"message": err.Error(),
	// 		"record":  nil,
	// 	})
	// }

	axios := resty.New()
	resp, err := axios.R().
		SetResult(response{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_company": "",
		}).
		Post(PATH + "api/resultsdsbnight")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*response)
	return c.JSON(fiber.Map{
		"status": http.StatusOK,
		"record": result.Record,
		"time":   time.Since(render_page).String(),
	})
}
