package middleware

import (
	"encoding/base64"
	"fmt"
	"strings"
	"time"

	"toko-bangunan/config"
	"toko-bangunan/internal/protocols/http/response"
	rsaheader "toko-bangunan/internal/utils/rsa"
	userservice "toko-bangunan/src/modules/user/services"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/rs/zerolog/log"
)

type JwtMiddleware interface {
	JwtVerifyToken(c *fiber.Ctx) error
	JwtVerifyRefreshToken(c *fiber.Ctx) error
}

type JwtMiddlewareImpl struct {
	UserService userservice.UserService
}

func NewAuthMiddlewareImpl(userService userservice.UserService) JwtMiddleware {
	return &JwtMiddlewareImpl{UserService: userService}
}

func (t *JwtMiddlewareImpl) JwtVerifyToken(c *fiber.Ctx) error {
	jwtToken := strings.Replace(c.Get("Authorization"), fmt.Sprintf("%s ", "Bearer"), "", 1)
	if jwtToken == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(response.NewResponse("UNAUTHORIZED", 401, "youre not loggined", nil, nil))
	}
	// _, err := t.UserService.FindTokenByUserId(c.Context(), c.Get("id"))
	// if err != nil {
	// 	return c.Status(fiber.StatusUnauthorized).JSON(response.NewResponse("UNAUTHORIZED", 401, "youre not loggined", nil))
	// }

	decodedPublicKey, err := base64.StdEncoding.DecodeString(config.Get().Auth.JwtToken.AccessToken.PublicKey)
	if err != nil {
		log.Err(err).Msg("could not decode")
	}

	token, err := jwt.Parse(c.Get("Authorization"), func(t *jwt.Token) (interface{}, error) {
		tokenType := t.Claims.(jwt.MapClaims)["token_type"]
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method : %v", t.Header["alg"])
		} else if tokenType != "access_token" {
			return nil, fmt.Errorf("unexpected token type : %v", tokenType)
		} else {
			publicKey, err := rsaheader.ReadPublicKeyFromEnv(decodedPublicKey)
			if err != nil {
				log.Err(err).Msg("cannot read public key from env")
				return nil, nil
			}
			return publicKey, nil
		}
	})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(response.NewResponse("UNAUTHORIZED", 401, err.Error(), nil, nil))
	}

	if !token.Valid {
		log.Err(err).Msg("token not valid or err")
		return c.Status(fiber.StatusUnauthorized).JSON(response.NewResponse("UNAUTHORIZED", 401, "token is not valid", nil, nil))
	}
	claims, _ := token.Claims.(jwt.MapClaims)
	id := claims["id"].(string)
	if id == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(response.NewResponse("UNAUTHORIZED", 401, "token is not valid", nil, nil))
	}

	exp := claims["exp"].(float64)
	rawExp := int64(exp)
	if rawExp < time.Now().Unix() {
		return c.Status(fiber.StatusUnauthorized).JSON(response.NewResponse("UNAUTHORIZED", 401, "token has expired", nil, nil))
	}

	return c.Next()
}

func (middleware *JwtMiddlewareImpl) JwtVerifyRefreshToken(c *fiber.Ctx) error {
	jwtToken := strings.Replace(c.Get("Authorization"), fmt.Sprintf("%s ", "Bearer"), "", 1)
	if jwtToken == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(response.NewResponse("UNAUTHORIZED", 401, "youre not loggined", nil, nil))
	}

	decodedPublicKey, err := base64.StdEncoding.DecodeString(config.Get().Auth.JwtToken.RefreshToken.PublicKey)
	if err != nil {
		log.Err(err).Msg("could not decode")
	}

	token, err := jwt.Parse(c.Get("Authorization"), func(t *jwt.Token) (interface{}, error) {
		tokenType := t.Claims.(jwt.MapClaims)["token_type"]
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method : %v", t.Header["alg"])
		} else if tokenType != "refresh_token" {
			return nil, fmt.Errorf("unexpected token type : %v", tokenType)
		} else {
			publicKey, err := rsaheader.ReadPublicKeyFromEnv(decodedPublicKey)
			if err != nil {
				log.Err(err).Msg("err read private key rsa from env")
				return nil, nil
			}
			return publicKey, nil
		}
	})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(response.NewResponse("UNAUTHORIZED", 401, err.Error(), nil, nil))
	}

	if err != nil || !token.Valid {
		log.Err(err).Msg("token not valid or err")
		return c.Status(fiber.StatusUnauthorized).JSON(response.NewResponse("UNAUTHORIZED", 401, "Refresh token is not valid", nil, nil))
	}

	claims, _ := token.Claims.(jwt.MapClaims)
	id := claims["id"].(string)
	if id == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(response.NewResponse("UNAUTHORIZED", 401, "Refresh token is not valid", nil, nil))
	}

	exp := claims["exp"].(float64)
	rawExp := int64(exp)
	if rawExp < time.Now().Unix() {
		return c.Status(fiber.StatusUnauthorized).JSON(response.NewResponse("UNAUTHORIZED", 401, "Refresh token has expired", nil, nil))
	}

	// userFind, err := middleware.UserService.FindById(c.Context(), claims["id"].(string))
	// if err != nil {
	// 	return c.Status(fiber.StatusUnauthorized).JSON(response.NewResponse("UNAUTHORIZED", 401, "Refresh token is not valid", nil))
	// }

	// tokenFind, err := middleware.UserService.FindTokenByUserId(c.Context(), claims["id"].(string))
	// if err != nil {
	// 	return c.Status(fiber.StatusUnauthorized).JSON(response.NewResponse("UNAUTHORIZED", 401, "Refresh token is not valid", nil))
	// }

	// c.Locals("refreshToken", tokenFind.RefreshToken)
	return c.Next()
}
