package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	heroku "github.com/jonahgeorge/force-ssl-heroku"
	_ "github.com/lib/pq"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"navtitle": "Trail."})
	})

	router.GET("/tour", func(c *gin.Context) {
		if _, err := db.Exec("CREATE TABLE IF NOT EXISTS trees (id  SERIAL PRIMARY KEY, treename varchar(45) NOT NULL, description varchar(450) NOT NULL,location GEOMETRY(POINT,4326))"); err != nil {
			c.String(http.StatusInternalServerError,
				fmt.Sprintf("Error creating database table: %q", err))
			return
		}
		c.HTML(http.StatusOK, "tour.html", gin.H{"navtitle": "Tour."})
	})

	router.GET("/map", func(c *gin.Context) {
		c.HTML(http.StatusOK, "map.html", gin.H{"navtitle": "Map."})
	})

	router.GET("/settings", func(c *gin.Context) {
		c.HTML(http.StatusOK, "settings.html", gin.H{"navtitle": "Settings."})
	})
	router.GET("/scan", func(c *gin.Context) {
		c.HTML(http.StatusOK, "scan.html", gin.H{"navtitle": "Scan."})
	})

	router.GET("/location", func(c *gin.Context) {
		c.ParseForm()
		lat := c.FormValue("lat")
		long := c.FormValue("long")
		if _, err := db.Exec("CREATE TABLE IF NOT EXISTS ticks (tick timestamp, lat real, long real)"); err != nil {
			c.String(http.StatusInternalServerError,
				fmt.Sprintf("Error creating database table: %q", err))
			return
		}

		if _, err := db.Exec("INSERT INTO ticks VALUES (now(),lat, long)"); err != nil {
			c.String(http.StatusInternalServerError,
				fmt.Sprintf("Error incrementing tick: %q", err))
			return
		}
	})

	router.Run(":" + port)
	heroku.ForceSsl(router)
}
