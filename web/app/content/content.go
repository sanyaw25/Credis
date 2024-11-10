// handler.go
package content

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FileHandler(c *gin.Context) {
	content, err := ioutil.ReadFile("output.txt")
	if err != nil {
		log.Println("Error reading file:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not read the file"})
		return
	}

	tmpl, err := template.ParseFiles("output.html")
	if err != nil {
		log.Println("Error parsing template:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not load the HTML template"})
		return
	}

	c.Header("Content-Type", "text/html")
	err = tmpl.Execute(c.Writer, string(content))
	if err != nil {
		log.Println("Error rendering template:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not render the template"})
	}
}
