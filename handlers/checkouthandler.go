package handlers

import (
	"github.com/gofiber/fiber/v2"
	jtoken "github.com/golang-jwt/jwt/v4"
	"github.com/raghavendrajha119/Ecommerce_website/config"
	"github.com/raghavendrajha119/Ecommerce_website/middlewares"
	"github.com/raghavendrajha119/Ecommerce_website/models"
)

func Checkout(c *fiber.Ctx) error {
	if tokenCookie := c.Cookies("jwt"); tokenCookie != "" {
		token, err := jtoken.Parse(tokenCookie, func(token *jtoken.Token) (interface{}, error) {
			return []byte(config.GetSecret()), nil
		})
		if err == nil && token.Valid {
			claims := token.Claims.(jtoken.MapClaims)
			userID, ok := claims["ID"].(float64)
			if ok {
				db, err := middlewares.OpenDB()
				if err != nil {
					return err
				}
				var cartItems []models.Cart
				if err := db.Where("user_id = ?", uint(userID)).Find(&cartItems).Error; err != nil {
					return err
				}
				for _, cartItem := range cartItems {
					var product models.Product
					if err := db.First(&product, cartItem.ProductID).Error; err != nil {
						return err
					}
					totalAmount := product.Price * float64(cartItem.Quantity)
					boughtProduct := models.BoughtProduct{
						UserID:      uint(userID),
						ProductID:   cartItem.ProductID,
						Quantity:    cartItem.Quantity,
						TotalAmount: totalAmount,
					}
					db.Create(&boughtProduct)
				}
				db.Delete(&cartItems)
				response := fiber.Map{
					"message":    "Purchase completed successfully.",
					"productIds": cartItems,
				}
				return c.JSON(response)
			}
		}
	}
	return c.Redirect("/login")
}
func GetBoughtProducts(c *fiber.Ctx) error {
	if tokenCookie := c.Cookies("jwt"); tokenCookie != "" {
		token, err := jtoken.Parse(tokenCookie, func(token *jtoken.Token) (interface{}, error) {
			return []byte(config.GetSecret()), nil
		})
		if err == nil && token.Valid {
			claims := token.Claims.(jtoken.MapClaims)
			userID, ok := claims["ID"].(float64)
			if ok {
				db, err := middlewares.OpenDB()
				if err != nil {
					return err
				}
				var boughtProducts []models.BoughtProduct
				if err := db.Where("user_id = ?", uint(userID)).Find(&boughtProducts).Error; err != nil {
					return err
				}
				var checkoutData []map[string]interface{}
				for _, boughtProduct := range boughtProducts {
					var product models.Product
					if err := db.First(&product, boughtProduct.ProductID).Error; err != nil {
						return err
					}
					itemData := map[string]interface{}{
						"ID":          product.ID,
						"Title":       product.Title,
						"Quantity":    boughtProduct.Quantity,
						"TotalAmount": boughtProduct.TotalAmount,
					}
					checkoutData = append(checkoutData, itemData)
				}
				return c.JSON(checkoutData)
			}
		}
	}
	return c.Redirect("/login")
}
