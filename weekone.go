package main

import (
	"log"
	"text/template"
	"net/http"
	"bytes"

	// Shortening the import reference name seems to make it a bit easier
	owm "github.com/briandowns/openweathermap"
)

func temp(web http.ResponseWriter, r *http.Request) {	
	const weatherTemplate = "right now the temperature is: {{.Main.Temp}}"

	OwnAPIKey := "c1ba0cb1a4985d1b360b9d471bcb36ed"

	location, err := owm.NewCurrent("c", "en", OwnAPIKey)
	if err != nil {
		log.Fatalln(err)
	}

	location.CurrentByName("Copenhagen")
	tmpl, err := template.New("weather").Parse(weatherTemplate)
	if err != nil {
		log.Fatalln(err)
	}

	var output bytes.Buffer

	err = tmpl.Execute(&output, location)
	if err != nil {
		log.Fatalln(err)
	}

	message := output.String()

	web.Write([]byte(message))

}


func main() {
	http.HandleFunc("/", temp)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
