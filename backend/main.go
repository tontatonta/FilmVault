package main

import (
	"fmt"
	"net/http"
	"io"
)

func main() {

	url := "https://api.themoviedb.org/3/movie/movie_id?language=en-US"

	req, _ := http.NewRequest("GET", url, nil)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s",body)
	})
	fmt.Println("Server running on port 8000")
	http.ListenAndServe(":8000", nil)
}
