package utils

import (
	"crypto/rand"
	"errors"
	"path/filepath"
	"strings"
	"time"

	"github.com/BIC-Final-Project/backend/internal/auth/http/middlewares"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/lucsky/cuid"
)

func GetFileName(fileName string) string {
	file := filepath.Base(fileName)

	return file[:len(file)-len(filepath.Ext(file))]
}

func GetTimeNow() int64 {
	time := time.Now().Unix()

	return time
}

func GetCurrentAuthUser(c *fiber.Ctx) (claim *middlewares.JWTClaim, tokenString string, err error) {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return nil, "", errors.New("authorization header is not provided")
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 {
		return nil, "", errors.New("invalid authorization header format")
	}

	tokenString = headerParts[1]

	claim, err = middlewares.VerifyJWT(tokenString)
	if err != nil {
		return nil, "", err
	}

	return claim, tokenString, nil
}

func ErrorHandler(c *fiber.Ctx, err error) error {
	// Status code defaults to 500
	code := fiber.StatusInternalServerError

	// Retrieve the custom status code if it's a *fiber.Error
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	// Set Content-Type: text/plain; charset=utf-8
	c.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)

	// Return status code with error message
	return c.Status(code).JSON(fiber.Map{
		"status":  code,
		"message": err.Error(),
	})
}

func GenerateUUID() (string, error) {
	u, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	return u.String(), nil
}

func GenerateCUID() (string, error) {
	cuid, err := cuid.NewCrypto(rand.Reader)
	if err != nil {
		return "", err
	}

	return cuid, nil
}
