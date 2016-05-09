// by setting package as main, Go will compile this as an executable file.
// Any other package turns this into a library
package main

// These are your imports / libraries / frameworks
import (
    // this is Go's built-in sql library
	"database/sql"
	"log"
	"net/http"
	"os"

    // this allows us to run our web server
	"github.com/gin-gonic/gin"
    // this lets us connect to Postgres DB's
	_ "github.com/lib/pq"
)

var (
    // this is the pointer to the database we will be working with
    // this is a "global" variable (sorta kinda, but you can use it as such)
	db *sql.DB
)

func setupDb(){

}

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	var errd error
    // here we want to open a connection to the database using an environemnt variable.
    // This isn't the best technique, but it is the simplest one for heroku
	db, errd = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if errd != nil {
		log.Fatalf("Error opening database: %q", errd)
	}

    setupDb()

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("html/*.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.GET("/ping", func(c *gin.Context) {
		ping := db.Ping()
		if ping != nil {
            // our site can't handle http status codes, but I'll still put them in cause why not
			c.JSON(http.StatusInternalServerError, gin.H{"error": "true", "message": "db was not created. Check your DATABASE_URL"})
		} else {
			c.JSON(http.StatusOK, gin.H{"error": "false", "message": "db created"})
		}
	})

    router.GET("/query1", func(c *gin.Context) {
        // YOUR CODE HERE
        c.JSON(http.StatusOK, gin.H{

            })
    })

    router.GET("/query2", func(c *gin.Context) {
        // YOUR CODE HERE
        c.JSON(http.StatusOK, gin.H{

            })
    })

    router.GET("/query3", func(c *gin.Context) {
        // YOUR CODE HERE
        c.JSON(http.StatusOK, gin.H{

            })
    })

    router.GET("/query4", func(c *gin.Context) {
        // YOUR CODE HERE
        //rows, err := db.Query("SELECT * FROM users")
        // TODO: finish query testing
        c.JSON(http.StatusOK, gin.H{

            })
    })

    // NO code should go after this line. it won't ever reach that point
	router.Run(":" + port)
}

/*
Example of processing a GET request

// this will run whenever someone goes to last-first-lab7.herokuapp.com/EXAMPLE
router.GET("/EXAMPLE", func(c *gin.Context) {
    // process stuff
    // run queries
    // do math
    //decide what to return
    c.JSON(http.StatusOK, gin.H{
        "key": "value"
        }) // this returns a JSON file to the requestor
    // look at https://godoc.org/github.com/gin-gonic/gin to find other return types. JSON will be the most useful for this
})

*/
