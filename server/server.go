package server

//fiber
import (
	"Cvalidator/service"
	"fmt"

	"github.com/gofiber/fiber/v2"
)


func Setup(){
	app := fiber.New()
	app.Post("/", hadlerPost)
	app.Get("/fake", fakeCard  )
	app.Listen(":3000")

}

type CreditCardPayload struct {
	CreditCard string `json:"creditCard"`
}

func hadlerPost(c *fiber.Ctx) error {
	var payload CreditCardPayload

	// Parse o corpo da requisição no formato da estrutura CreditCardPayload
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request payload",
		})
	}

	// Verifica se o campo no JSON é "creditCard"
	if payload.CreditCard == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Field name should be 'creditCard'",
		})
	}

	// Obtém o número do cartão de crédito da estrutura
	creditCard := payload.CreditCard

	// Verifica a validade do número do cartão de crédito usando o algoritmo de Luhn
	isValid := service.LunhAlgorithm(creditCard)
	fmt.Println(creditCard)

	if isValid {
		return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
			"message": "Credit card is valid",
		})
	} else {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Credit card is not valid",
		})
	}
}


func fakeCard(c *fiber.Ctx) error {
	card, err :=   service.CreateFakeCreditCard()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error",
		})
	}
	validCard := service.LunhAlgorithm(card)

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"card": card,
		"valid": validCard,
	})

	
}