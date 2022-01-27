package main

import (
	"fmt"
	"gee"
	"net/http"
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	ans := make([][]int,0)
	sort.Ints(candidates)
	backtracking(&ans, candidates, target,0,make([]int,0))
	return ans
}

func backtracking(ans *[][]int, candidates []int, target,index int, cur []int) {
	if target == 0 {
		*ans = append(*ans, cur[:])
		fmt.Println("ans:",*ans)
		return
	}

	for i := index; i < len(candidates); i++ {
		if target - candidates[i] < 0 {
			break
		}

		fmt.Println("ans:",*ans)
		fmt.Println(cur)
		cur = append(cur,candidates[i])
		fmt.Println(cur)

		fmt.Println("ans:",*ans)

		if len(*ans) > 0 {
			fmt.Printf("    %p\n",)
		}
		fmt.Println(target - candidates[i])
		backtracking(ans, candidates, target - candidates[i],i, cur)
		cur =cur[:len(cur) - 1]
	}
}

func app(a []int, b int)  {
	a = append(a,b)
}
func main() {
	r := gee.New()
	r.GET("/index", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})
	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *gee.Context) {
			c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
		})

		v1.GET("/hello", func(c *gee.Context) {
			// expect /hello?name=geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}
	v2 := r.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *gee.Context) {
			// expect /hello/geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *gee.Context) {
			c.JSON(http.StatusOK, gee.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})

	}

	r.Run(":9999")
}
