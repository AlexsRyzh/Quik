package router

import (
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/quik/backend/internal/api/token"
	"github.com/quik/backend/internal/router/auth"
	"github.com/quik/backend/internal/router/chat"
	"github.com/quik/backend/internal/router/comment"
	"github.com/quik/backend/internal/router/friend"
	"github.com/quik/backend/internal/router/image"
	"github.com/quik/backend/internal/router/like"
	"github.com/quik/backend/internal/router/post"
	"github.com/quik/backend/internal/router/user"
	"os"
)

func Init(e *echo.Echo) {
	authentication := e.Group("/auth")
	authentication.POST("/login", auth.Login)
	authentication.POST("/register", auth.Register)

	e.GET("/upload/:filename", image.Upload)
}

func InitSecurity(e *echo.Echo) {

	secure := e.Group("")

	secure.Use(echojwt.WithConfig(echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(token.JwtCustomClaims)
		},
		SigningKey: []byte(os.Getenv("SECRET_KEY")),
		ContextKey: "token",
	}))

	secure.POST("/auth/refresh", auth.Refresh)

	secure.GET("/users", user.GetAllOrByIDs)
	secure.GET("/users/:id", user.GetOne)
	secure.GET("/users-me", user.GetMyUser)
	secure.POST("/users/:id", user.UpdateOne)

	secure.GET("/posts-ids", post.GetPostIDs)
	secure.GET("/post-text/:id", post.GetPostText)
	secure.GET("/post-img/:id", post.GetPostImg)
	secure.GET("/post-is-like/:postID", post.GetPostIsLike)
	secure.GET("/posts/:id", post.GetMyPost)
	secure.GET("/posts/:id/count-likes", post.GetPostCountLike)
	secure.GET("/posts/:id/comments", post.GetPostComments)
	secure.GET("/posts/:id/count-comment", post.GetPostCountComment)
	secure.GET("/posts-ids-my", post.GetPostMyIDs)
	secure.POST("/posts", post.CreateOne)

	secure.POST("/likes/:postID", like.CreateOne)
	secure.DELETE("/likes/:postID", like.DeleteOne)

	secure.POST("/comments/:postID", comment.CreateOne)

	secure.POST("/friends/:id", friend.AddFriend)
	secure.DELETE("/friends/:id", friend.DeleteFriend)

	secure.GET("/ws", chat.Hello)
}
