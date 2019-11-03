package handler

import (
	"encoding/csv"
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method: ", r.Method)
	if r.Method == "POST" {
		header := []string{}
		allline := [][]string{}
		allinput := make(map[string][][]string)
		r.ParseMultipartForm(32 << 20)
		formdata := r.MultipartForm
		files := formdata.File["uploadfile"]
		var ind string = ""
		for i := range files {
			input := [][]string{}
			file, err := files[i].Open()
			if err != nil {
				break
			}
			defer file.Close()
			reader := csv.NewReader(file)
			for {
				line, err := reader.Read()
				if err != nil {
					break
				}
				if line[0] == "name" {
					ind = line[1]
					// fmt.Println(ind)
				} else if line[0] == "date" {
					header = line
					continue
				} else {
					input = append(input, line)
				}
			}
			if ind == "" {
				fmt.Println("not set key")
			}
			// fmt.Printf("%#v", &input)
			allinput[ind] = input
			// fmt.Println("ok?")
		}
		// for key, ele := range allinput {
		// 	fmt.Printf("%#v", ele)
		// 	row := []string{}
		// 	for i := 0; i < len(ele); i++ {
		// 		row = append(row, ele[i][0])
		// 		var init string
		// 		for j := 0; j < len(ele[0]); j++ {
		// 			if ele[i][j] == "1" {
		// 				init += key
		// 			} else if ele[i][j] == "0" {
		// 				continue
		// 			}
		// 		}
		// 		row = append(row, init)
		// 	}
		// 	allline = append(allline, row)
		// }
		elem := allinput[ind]
		// fmt.Printf("%#v", elem)
		for i := 0; i < len(elem); i++ {
			row := []string{}
			row = append(row, elem[i][0])
			for j := 1; j < len(elem[i]); j++ {
				init := ""
				for key, e := range allinput {
					if e[i][j] == "1" {
						init += key
					}
				}
				row = append(row, init)
			}
			allline = append(allline, row)
		}
		fmt.Printf("%#v", allline[0])
		allstr := ""
		allstr += `<table id="table1" class="table is-bordered">`
		allstr += `<thead><tr>`
		for _, ele := range header {
			allstr += `<td>` + ele + `</td>`
		}
		allstr += `</tr></thead>`
		allstr += `<tbody>`
		for _, st := range allline {
			allstr += `<tr>`
			for _, ele := range st {
				allstr += `<td>` + ele + `</td>`
			}
			allstr += `</tr>`
		}
		allstr += `</tbody>`
		allstr += `</table>`

		fmt.Fprint(w, `
		<!DOCTYPE html>
		<html>

		<head>
		  <title>AUTOSHIFT</title>
		  <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bulma@0.8.0/css/bulma.min.css">
		  <style>
		  	html {
				  overflow-x: scroll;
			  }
		  </style>
		</head>

		<body>
		  <h1>HELLO</h1>
		  `+allstr+`
		</body>
		<script>
			document.querySelectorAll('td').forEach((v,i) => {
				if (v.innerText == 1) {
					v.style.backgroundColor = "black";
				}
			});
		</script>

		</html>
		`)
	} else {
		fmt.Fprint(w, "not post request")
	}
}
