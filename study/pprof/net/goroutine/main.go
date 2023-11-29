package main

import (
	"net/http"
	_ "net/http/pprof"
	"time"

	"github.com/gin-gonic/gin"
)

// ab -n 1000 -c 10 http://localhost:28080/live
func main() {
	runProfile()
	runGin()
}

func runProfile() {
	go func() {
		if err := http.ListenAndServe("localhost:6060", nil); err != nil {
			panic(err)
		}
	}()
}

func runGin() {
	e := gin.Default()
	e.GET("/live", liveHandler)
	if err := e.Run(":28080"); err != nil {
		panic(err)
	}
}

func liveHandler(c *gin.Context) {
	goroutineTaskService()
	c.JSON(http.StatusOK, map[string]int{"code": 0})
}

func goroutineTaskService() {
	ch := make(chan bool)
	go func() {
		time.Sleep(time.Millisecond)
		ch <- true
	}()
}
