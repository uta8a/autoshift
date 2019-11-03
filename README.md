package handler

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"strconv"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method: ", r.Method)
	if r.Method == "POST" {
		header := []string{}
		allline := [][]string{}
		allkey := []string{}
		allinput := make(map[string][][]string)
		colors := ["#EA5532", "#FFF100", "#0068B7", "#00A0E9", "#00A051", "#9FD9F6", "#E4007F", "#D3DEF1", "#187FC4", "#86B81B", "#EA5504", "#00693E", "#F39800", "#ED6C00", "#009E96", "#008DCB", "D4ECEA", "#1D2088", "#920783"]

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
			allinput[ind] = input
			allkey = append(allkey, ind)
		}
		elem := allinput[ind]
		for _, e := range allinput {
			for i := 0; i < len(elem); i++ {
				sum := 0
				for j := len(elem[i]) - 1; j >= 0; j-- {
					if e[i][j] == "1" {
						e[i][j] = strconv.Itoa(sum + 1)
						sum++
					} else {
						sum = 0
					}
				}
			}
		}
		for i := 0; i < len(elem); i++ {
			row := []string{}
			row = append(row, elem[i][0])
			for j := 1; j < len(elem[i]); j++ {
				init := ""
				max := 0
				for key, e := range allinput {
					iv, _ := strconv.Atoi(e[i][j])
					if max < iv {
						max = iv
						init = key
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
		allstr += `</body><script>document.querySelectorAll('td').forEach((v,i) => {`
		for i, el in allkey {
			color := ""
			if i<len(colors) {
				color = colors[i]
			}
			allstr += `if (v.innerText == "`+el+`") {v.style.backgroundColor = "`+color +`";}`
		}
		allstr += `});</script>`

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
		  <h1>HELLO</h1>`+allstr+`</html>`)
	} else {
		fmt.Fprint(w, "not post request")
	}
}
