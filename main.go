package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Song struct {
	Name string `form:"name" uri:"name" binding:"required"`
}

type Score struct {
	Name  string `form:"name" uri:"name" binding:"required"`
	Score int    `form:"score"`
	Count int
}

var counter map[string]int
var totalScore map[string]*Score

func SongGrade(c *gin.Context) {
	var form Score
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}

	if totalScore[form.Name] == nil {
		totalScore[form.Name] = &Score{form.Name, 0, 0}
	}
	totalScore[form.Name].Score += form.Score
	totalScore[form.Name].Count++
	c.String(http.StatusOK, "done!")
}

func GetSongGrade(c *gin.Context) {
	var form Score
	if err := c.ShouldBindUri(&form); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}

	if totalScore[form.Name] == nil {
		c.String(http.StatusOK, "0")
		return
	}

	if totalScore[form.Name].Count == 0 {
		c.String(http.StatusOK, "0")
	} else {
		c.String(http.StatusOK, strconv.Itoa(totalScore[form.Name].Score/totalScore[form.Name].Count))
	}
}

func GetSongGrader(c *gin.Context) {
	var form Score
	if err := c.ShouldBindUri(&form); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}

	if totalScore[form.Name] == nil {
		c.String(http.StatusOK, "0")
		return
	}

	c.String(http.StatusOK, strconv.Itoa(totalScore[form.Name].Count))
}

func SongCounter(c *gin.Context) {
	var form Song
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	counter[form.Name]++
	c.String(http.StatusOK, "done!")
}

func GetSongCounter(c *gin.Context) {
	var form Song
	if err := c.ShouldBindUri(&form); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}

	returnVal := 0
	if counter[form.Name] != 0 {
		returnVal = counter[form.Name]
	}
	c.String(http.StatusOK, strconv.Itoa(returnVal))
}

func main() {
	r := gin.Default()
	counter = make(map[string]int)
	totalScore = make(map[string]*Score)
	r.POST("/grade", SongGrade)
	r.GET("/:name/grade", GetSongGrade)
	r.GET("/:name/grader", GetSongGrader)
	r.POST("/counter", SongCounter)
	r.GET("/:name/counter", GetSongCounter)
	r.Run(":39002")
}
