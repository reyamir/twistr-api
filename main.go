package main

import (
	"context"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	twitterscraper "github.com/n0madic/twitter-scraper"
)

func main() {
	r := gin.Default()
	scraper := twitterscraper.New()

	r.Use(cors.Default())

	r.GET("/user/:handle", func(c *gin.Context) {
		user := c.Params.ByName("handle")
		resp := []string{}
		for tweet := range scraper.GetTweets(context.Background(), user, 2000) {
			if tweet.Error != nil {
				panic(tweet.Error)
			}
			resp = append(resp, tweet.Text)
		}
		c.JSON(http.StatusOK, gin.H{"data": resp})
	})

	r.Run()
}
