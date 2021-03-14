package main

import (
	"encoding/json"
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

type TreeJson struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Latinname   string `json:"latinname"`
	Height      int    `json:"height"`
	Age         int    `json:"age"`
	Description string `json:"description"`
	Origin      string `json:"origin"`
	Img         string `json:"img"`
}

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

		//if the variable value is not null then pull all tree info corresponding to that id from database
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
			//loop through the data recieved from the SELECT statement and store in the specific variables above
			for rows.Next() {
				if err := rows.Scan(&name, &latinname, &height, &age, &description, &origin, &img); err != nil {
					c.String(http.StatusInternalServerError,
						fmt.Sprintf("Error scanning trees: %q", err))
					return
				}
			}
			//serve the tour page, passing in the tree info variables
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
		} else { //if the QR code has no variable assigned to it just load the tour page
			c.HTML(http.StatusOK, "tour.html", gin.H{"navtitle": "Tour."})
		}
	})

	router.GET("/map", func(c *gin.Context) {
		c.HTML(http.StatusOK, "map.html", gin.H{"navtitle": "Map."})
	})

	router.GET("/settings", func(c *gin.Context) {
		c.HTML(http.StatusOK, "settings.html", gin.H{"navtitle": "Settings."})
	})

	//serves the raw json to the user
	router.GET("/TreesHWU.geojson", func(c *gin.Context) {
		c.File("static/TreesHWU.geojson")
	})

	router.GET("/ar", func(c *gin.Context) {
		c.HTML(http.StatusOK, "ar.html", gin.H{"navtitle": "Ar."})
	})
	router.GET("/scan", func(c *gin.Context) {
		c.HTML(http.StatusOK, "scan.html", gin.H{"navtitle": "Scan."})
	})

	//this function is used for testing geolocation updates
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

	//this function checks whether a user is within the geofence of a route - if not it tells them they are not near a route
	router.GET("/route/:lat/:long", func(c *gin.Context) {
		lat := c.Param("lat")
		long := c.Param("long")
		result := false
		var id int
		rows, err := db.Query("SELECT id FROM boundary WHERE ST_DWithin ( geography (ST_Point(longitude,latitude)), geography (ST_Point($1, $2)), 280)", long, lat)
		if err != nil {
			c.String(http.StatusInternalServerError,
				fmt.Sprintf("Error reading trees: %q", err))
			return
		}
		defer rows.Close()
		for rows.Next() {
			if err := rows.Scan(&id); err != nil {
				c.String(http.StatusInternalServerError,
					fmt.Sprintf("Error scanning trees: %q", err))
				return
			}
			result = true
		}

		c.JSON(200, result)

	})

	//this function is used for testing geolocation updates
	router.GET("/trigger/:lat/:long", func(c *gin.Context) {
		lat := c.Param("lat")
		long := c.Param("long")
		if _, err := db.Exec("CREATE TABLE IF NOT EXISTS trigger (tick timestamp, lat real, long real)"); err != nil {
			c.String(http.StatusInternalServerError,
				fmt.Sprintf("Error creating database table: %q", err))
			return
		}

		if _, err := db.Exec("INSERT INTO trigger VALUES (now(),$1, $2)", lat, long); err != nil {
			c.String(http.StatusInternalServerError,
				fmt.Sprintf("Error incrementing tick: %q", err))
			return
		}
	})

	router.POST("/geofence/:lat/:long/:visited", func(c *gin.Context) {
		lat := c.Param("lat")
		long := c.Param("long")
		visited := c.Param("visited")
		rows, err := db.Query("SELECT id, treename, latinname, height, age, description, origin, img FROM trees WHERE ST_DWithin ( geography (ST_Point(longitude,latitude)), geography (ST_Point($1, $2)), 20) AND id != $3 limit 1", long, lat, visited)
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
		var id string
		var success bool = false

		for rows.Next() {
			if err := rows.Scan(&id, &name, &latinname, &height, &age, &description, &origin, &img); err != nil {
				c.String(http.StatusInternalServerError,
					fmt.Sprintf("Error scanning trees: %q", err))
				return
			}
			success = true
		}

		treeJson := TreeJson{
			Id:          id,
			Name:        name,
			Latinname:   latinname,
			Height:      height,
			Age:         age,
			Description: description,
			Img:         img,
		}

		js, err := json.Marshal(treeJson)

		if success == true && id != visited {
			c.JSON(200, string(js))
		} else {
			c.JSON(200, "null")
		}

	})

	router.Run(":" + port)
	heroku.ForceSsl(router)
}
