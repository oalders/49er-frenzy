package main

import (
	"math/rand"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob(filepath.Join("templates", "*tmpl"))
	r.GET("/", func(c *gin.Context) {
		x := uniqueList([]int{})
		y := uniqueList([]int{})

		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
			"x":     x,
			"y":     y,
		})
	})

	_ = r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func randomInt() int {
	b := 12
	a := 2

	return a + rand.Intn(b-a+1) // a ≤ n ≤ b
}

func uniqueList(list []int) []int {
	m := make(map[int]bool)

	for i := 0; i < 20; i++ {
		num := randomInt()
		if _, ok := m[num]; !ok {
			m[num] = true

			list = append(list, num)
		}
	}

	return list[:7]
}