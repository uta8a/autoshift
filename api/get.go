package handler

import (
	"encoding/csv"
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method: ", r.Method)
	if r.Method == "POST" {
		allline := [][]string{}
		allinput := [][]string{}
		r.ParseMultipartForm(32 << 20)
		formdata := r.MultipartForm
		files := formdata.File["uploadfile"]
		for i := range files {
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
				allinput = append(allinput, line)
			}

		}
		// fmt.Printf("%#v", allline)
		// allline = append(allline, allinput[0][1:])
		// for _, _ := range allinput {

		// }
		allline = allinput
		allstr := ""
		allstr += `<table id="table1" class="table is-bordered">`
		allstr += `<thead><tr>`
		for _, ele := range allline[0] {
			allstr += `<td>` + ele + `</td>`
		}
		allstr += `</tr></thead>`
		allstr += `<tbody>`
		for i, st := range allline {
			if i == 0 {
				continue
			}
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
