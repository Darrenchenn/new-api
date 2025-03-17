package middleware

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/supertokens/supertokens-golang/recipe/session"
	"github.com/supertokens/supertokens-golang/recipe/session/sessmodels"
	"github.com/supertokens/supertokens-golang/supertokens"
)

// This is a function that wraps the supertokens verification function
// to work the gin
func verifySession(options *sessmodels.VerifySessionOptions) gin.HandlerFunc {
	return func(c *gin.Context) {
		session.VerifySession(options, func(rw http.ResponseWriter, r *http.Request) {
			c.Request = c.Request.WithContext(r.Context())
			c.Next()
		})(c.Writer, c.Request)
		// we call Abort so that the next handler in the chain is not called, unless we call Next explicitly
		c.Abort()
	}
}

func SetSuperTokensCors(server *gin.Engine) {
	server.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:4173"}, //todo: change this to the actual domain of the frontend
		AllowMethods: []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"},
		AllowHeaders: append([]string{"content-type"},
			supertokens.GetAllCORSHeaders()...),
		AllowCredentials: true,
	}))
	server.Use(func(c *gin.Context) {
		supertokens.Middleware(http.HandlerFunc(
			func(rw http.ResponseWriter, r *http.Request) {
				c.Next()
			})).ServeHTTP(c.Writer, c.Request)
		// we call Abort so that the next handler in the chain is not called, unless we call Next explicitly
		c.Abort()
	})
	// return headers in response for the frontend to access
	server.GET("/ping", func(c *gin.Context) {
		// Manually setting CORS headers in the response
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:4173") //todo: change this to the actual domain of the frontend
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		c.JSON(200, gin.H{"message": "pong"})
	})
	server.OPTIONS("/ping", func(c *gin.Context) {
		// Manually setting CORS headers in the response
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:4173") //todo: change this to the actual domain of the frontend
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		c.JSON(200, gin.H{"message": "pong"})
	})
}
