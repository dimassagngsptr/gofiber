package middlewares

import (
	"fmt"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func ExtractToken(c *fiber.Ctx) string {
	bearerToken := c.Get("Authorization")
	if strings.HasPrefix(bearerToken, "Bearer "){
		return strings.TrimPrefix(bearerToken, "Bearer ")
	}
	return ""
}

func JwtMiddleware() fiber.Handler{
	secretKey := os.Getenv("JWT_KEY")
	return func(c *fiber.Ctx) error{
		tokenString := ExtractToken(c)
		if tokenString == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":"Unauthorized",
			})
		}
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
						if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok	{
								return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
						}
				return []byte(secretKey), nil
			})
			if err != nil{
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"error": "Unauthorized",
				})
			}
			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Locals("userClaims", claims)
		} else {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}
			return c.Next()
	}
}

func ValidateSellerRole() fiber.Handler{
	return func(c *fiber.Ctx) error{
		claims:=GetUserClaims(c)
		if claims == nil{
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message":"Unauthorized",
			})
		}
		fmt.Println("role", claims)
		role, ok :=claims["role"].(string)
		if !ok || role != "seller"{
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"Forbidden":"You don't have permission to access",
			})
		}
		return c.Next()
	}
}

func GetUserClaims(c *fiber.Ctx) jwt.MapClaims {
	if claims, ok := c.Locals("userClaims").(jwt.MapClaims); ok {
		fmt.Println("claims middleware: ", claims)
		return claims
	}
	return nil
}