package middleware

import (
	"encoding/json"
	"errors"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"log"
	"net/http"
	"os"
	"skeleton-fiber-clean-architecture/infrastructure/logger"
	"time"
)

// OAuth2Config is the configuration for OAuth2
var OAuth2Config *oauth2.Config

var oauthStateString = "random-string"

// InitializeOAuth2Config initializes the OAuth2 configuration
func InitializeOAuth2Config() {
	OAuth2Config = &oauth2.Config{
		RedirectURL:  "http://localhost:3000/callback",
		ClientID:     os.Getenv("OAUTH2_CLIENT_ID"),
		ClientSecret: os.Getenv("OAUTH2_CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}
}

// HandleLogin redirects the user to the Google login page
func HandleLogin(c *fiber.Ctx) error {
	url := OAuth2Config.AuthCodeURL(oauthStateString)
	return c.Redirect(url, http.StatusTemporaryRedirect)
}

// HandleGoogleCallback handles the callback from Google after login
func HandleGoogleCallback(c *fiber.Ctx) error {
	state := c.Query("state")
	if state != oauthStateString {
		log.Printf("invalid oauth state, expected '%s', got '%s'\n", oauthStateString, state)
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	code := c.Query("code")
	token, err := OAuth2Config.Exchange(c.Context(), code)
	if err != nil {
		log.Printf("code exchange failed: %s\n", err.Error())
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		log.Printf("failed getting user info: %s\n", err.Error())
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}
	defer resp.Body.Close()

	userInfo := struct {
		Id      string `json:"id"`
		Name    string `json:"name"`
		Picture string `json:"picture"`
	}{}
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		log.Printf("failed decoding user info: %s\n", err.Error())
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	// Save user information in a cookie
	c.Cookie(&fiber.Cookie{
		Name:     "user_id",
		Value:    userInfo.Id,
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
	})

	return c.JSON(fiber.Map{
		"user": userInfo,
	})
}

// AuthMiddleware checks if the user is logged in
func AuthMiddleware(c *fiber.Ctx) error {
	userID := c.Cookies("user_id")
	if userID == "" {
		logger.LogError(errors.New("No user session found"))
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	// Additional logic to verify the user can be added here
	return c.Next()
}
