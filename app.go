package ginApp

import "github.com/gin-gonic/gin"

func Run() *gin.Engine {
    app := gin.Default()
    app.Run()
    return app
}
