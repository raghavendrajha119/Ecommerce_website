package handlers

import (
	"github.com/gofiber/fiber/v2"
	jtoken "github.com/golang-jwt/jwt/v4"
	"github.com/raghavendrajha119/Ecommerce_website/config"
	"github.com/raghavendrajha119/Ecommerce_website/middlewares"
	"github.com/raghavendrajha119/Ecommerce_website/models"
)

func AddtoCart(c *fiber.Ctx) error {
	if tokenCookie := c.Cookies("jwt"); tokenCookie != "" {
		token, err := jtoken.Parse(tokenCookie, func(token *jtoken.Token) (interface{}, error) {
			return []byte(config.Secret), nil
		})
		if err == nil && token.Valid {
			claims := token.Claims.(jtoken.MapClaims)
			userID, okId := claims["ID"].(float64)
			if okId {

				var data map[string]uint
				if err := c.BodyParser(&data); err != nil {
					return err
				}
				productID := data["productId"]
				db, err := middlewares.OpenDB()
				if err != nil {
					return err
				}
				cart := models.Cart{
					UserID:    uint(userID),
					ProductID: productID,
					Quantity:  1,
				}
				db.Create(&cart)
				return c.JSON(fiber.Map{
					"message": "Item added to cart successfully.",
				})
			}
		}
	} else {
		return c.Redirect("/login")
	}
	return nil
}
func GetfromCart(c *fiber.Ctx) error {
	if tokenCookie := c.Cookies("jwt"); tokenCookie != "" {
		token, err := jtoken.Parse(tokenCookie, func(token *jtoken.Token) (interface{}, error) {
			return []byte(config.Secret), nil
		})
		if err == nil && token.Valid {
			claims := token.Claims.(jtoken.MapClaims)
			userID, okId := claims["ID"].(float64)
			if okId {
				db, err := middlewares.OpenDB()
				if err != nil {
					return err
				}
				var cartItems []models.Cart
				if err := db.Where("user_id = ?", uint(userID)).Find(&cartItems).Error; err != nil {
					return err
				}
				var cartData []map[string]interface{}
				for _, cartItem := range cartItems {
					var product models.Product
					if err := db.First(&product, cartItem.ProductID).Error; err != nil {
						return err
					}
					itemData := map[string]interface{}{
						"ID":       product.ID,
						"Title":    product.Title,
						"Category": product.Category,
						"Price":    product.Price,
						"Quantity": cartItem.Quantity,
					}
					cartData = append(cartData, itemData)

				}
				return c.JSON(cartData)
			}
		}
	}
	return c.Redirect("/login")
}
