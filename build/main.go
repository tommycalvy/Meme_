package main

import (
  "net/http"
  "io/ioutil"
)

var (
  imageSearchUrl string = "https://www.google.com/searchbyimage?hl=en-US&image_url="
)

func imageSearch(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Access-Control-Allow-Origin", "localhost/")
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

func main() {
  http.HandleFunc("/image-search", imageSearch)
  appengine.Main() // Starts the server to receive requests
}
