package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	heroku "github.com/jonahgeorge/force-ssl-heroku"
	_ "github.com/lib/pq"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
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
	router.LoadHTMLGlob("static/templates/*.html")
	router.Static("/static", "static")
	router.StaticFile("sw.js", "./sw.js")
	router.StaticFile("manifest.webmanifest", "./manifest.webmanifest")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"navtitle": "Trail."})
	})

	router.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"navtitle": "Trail."})
	})

	router.GET("/tour", func(c *gin.Context) {
		if _, err := db.Exec("CREATE TABLE IF NOT EXISTS trees (id  SERIAL PRIMARY KEY, treename varchar(45) NOT NULL,latinname varchar(45), height numeric, age integer,description varchar(450) NOT NULL,location GEOMETRY(POINT,4326),origin VARCHAR(45), image TEXT)"); err != nil {
			c.String(http.StatusInternalServerError,
				fmt.Sprintf("Error creating database table: %q", err))
			return
		}

		//The following section of code handles the event in which a user scans a QR code of a specific tree (variable is passed in ? url param)
		treeNum := c.Query("id")
		fmt.Println("Tree ID is ?", treeNum)

		if treeNum != "" {
			rows, err := db.Query("SELECT treename, latinname, height, age, description, origin, img FROM trees WHERE id = $1", treeNum)
			if err != nil {
				c.String(http.StatusInternalServerError,
					fmt.Sprintf("Error reading trees: %q", err))
				return
			}
			defer rows.Close()
			var name string
			var latinname string
			var height int
			var age int
			var description string
			var origin string
			var img string
			for rows.Next() {
				if err := rows.Scan(&name, &latinname, &height, &age, &description, &origin, &img); err != nil {
					c.String(http.StatusInternalServerError,
						fmt.Sprintf("Error scanning trees: %q", err))
					return
				}
			}
			c.HTML(http.StatusOK, "tour.html", gin.H{"navtitle": "Tour.",
				"qr":          true,
				"id":          treeNum,
				"treename":    name,
				"latinname":   latinname,
				"height":      height,
				"age":         age,
				"description": description,
				"origin":      origin,
				"img":         img,
			})
		} else {
			c.HTML(http.StatusOK, "tour.html", gin.H{"navtitle": "Tour."})
		}
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

	router.GET("/location/:lat/:long", func(c *gin.Context) {
		lat := c.Param("lat")
		long := c.Param("long")
		if _, err := db.Exec("CREATE TABLE IF NOT EXISTS ticks (tick timestamp, lat real, long real)"); err != nil {
			c.String(http.StatusInternalServerError,
				fmt.Sprintf("Error creating database table: %q", err))
			return
		}

		if _, err := db.Exec("INSERT INTO ticks VALUES (now(),$1, $2)", lat, long); err != nil {
			c.String(http.StatusInternalServerError,
				fmt.Sprintf("Error incrementing tick: %q", err))
			return
		}
	})

	router.GET("/geofence/:lat/:long", func(c *gin.Context) {
		//lat := c.Param("lat")
		//long := c.Param("long")

		/*rows, err := db.Query("SELECT id, treename, latinname, height, age, description, origin FROM trees WHERE ST_DWithin ( geography (ST_Point(longitude,latitude)), geography (ST_Point($1, $2)), 60) limit 1", long, lat)
		if err != nil {
			c.String(http.StatusInternalServerError,
				fmt.Sprintf("Error reading trees: %q", err))
			return
		}
		defer rows.Close()
		var name string
		var latinname string
		var height int
		var age int
		var description string
		var origin string
		//var img string
		var id int
		for rows.Next() {
			if err := rows.Scan(&id, &name, &latinname, &height, &age, &description, &origin); err != nil {
				c.String(http.StatusInternalServerError,
					fmt.Sprintf("Error scanning trees: %q", err))
				return
			}
		}
		if name != "" { */
			/*c.HTML(http.StatusOK, "tour.html", gin.H{"navtitle": "Tour.",
				"geo":         true,
				"id":          id,
				"treename":    name,
				"latinname":   latinname,
				"height":      height,
				"age":         age,
				"description": description,
				"origin":      origin,
				"lat": lat,
				"long": long,
				//"img":         img,
			})*/
			fmt.Println("TEST")
		//}
		//c.HTML(http.StatusOK, "tour.html", gin.H{"navtitle": "Tour."}) 

	})

	router.Run(":" + port)
	heroku.ForceSsl(router)
}
