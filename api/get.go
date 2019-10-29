package handler

import (
	"encoding/csv"
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method: ", r.Method)
	if r.Method == "POST" {
		r.ParseMultipartForm(32 << 20)
		file, _, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		reader := csv.NewReader(file)
		for {
			line, err := reader.Read()
			if err != nil {
				break
			}
			for _, word := range line {
				fmt.Fprint(w, word)
				fmt.Fprint(w, " ")
			}
			fmt.Fprint(w, "\n")
		}
		// fmt.Fprint(w, "%v", handler.Header)
		// fmt.Fprint(w, file)
	} else {
		fmt.Fprint(w, "not post request")
	}
}
