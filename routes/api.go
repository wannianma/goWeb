package routes

import (
	"goWeb/handler/api/indoor"
	"goWeb/handler/api/ny"
	"goWeb/server"
)

// RegisterApiRoutes register api route to Gin
func RegisterApiRoutes(env *server.Env) {
	router := env.Gin
	// JSON-REST API Version 1
	v1 := router.Group("/v1")
	{
		v1.GET("login", account.Login)
		v1.POST("register", account.Register)
	}

	stream := router.Group("/stream")
	{
		stream.GET("/data", ny.StreamData)
		stream.GET("/start", ny.TeamStart)
		stream.GET("/questions", ny.GetQuestions)
		stream.GET("/team/:name", ny.PushData)
		stream.GET("/answer/:name", ny.PushAnswer)
		stream.GET("/reset", ny.ResetTeam)
	}
}
