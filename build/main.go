package main

import (
  "html/template"
  "path/filepath"
  "io/ioutil"
  "net/http"
  "log"
  "os"
)

var (
  imageSearchUrl string
  tpl *template.Template
)

func init() {
  tpl = template.Must(template.ParseGlob("templates/*"))
  imageSearchUrl = "https://www.google.com/searchbyimage?hl=en-US&image_url="
}

func main() {
  http.HandleFunc("/image-search", imageSearch)
  http.HandleFunc("/", indexHandler)

  // Serve static files out of the public directory.
	// By configuring a static handler in app.yaml, App Engine serves all the
	// static content itself. As a result, the following two lines are in
	// effect for development only.
	public := http.StripPrefix("/public", http.FileServer(http.Dir("public")))
	http.Handle("/public/", public)

  // [START setting_port]
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
	// [END setting_port]
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Redirect(w, r, "/", http.StatusFound)
    return
  }
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func imageSearch(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Access-Control-Allow-Origin", "https://meme-242116.appspot.com")
  url, ok := r.URL.Query()["url"]

  if !ok || len(url[0]) < 1 {
    http.Error(w, "Url Param 'key' is missing", 500)
    return
  }

  resp, err := http.Get(imageSearchUrl + url[0])
  if err != nil {
    http.Error(w, "Error with http.Get for the image source", 500)
    return
  }
  defer resp.Body.Close()

  //buf := new(bytes.Buffer)
  //buf.ReadFrom(resp.Body)
  //newStr := buf.String()
  //fmt.Fprintf(w, resp.Body)

  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    http.Error(w, "Error with reading response body", 500)
    return
  }

  w.Write(body)

}
