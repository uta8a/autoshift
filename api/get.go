package handler

import (
	"encoding/csv"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method: ", r.Method)
	if r.Method == "POST" {
		header := []string{}
		// allline := [][]string{}
		allkey := []string{}
		allinput := make(map[string][][]string)
		colors := []string{"#EA5532", "#FFF100", "#0068B7", "#00A0E9", "#00A051", "#9FD9F6", "#E4007F", "#D3DEF1", "#187FC4", "#86B81B", "#EA5504", "#00693E", "#F39800", "#ED6C00", "#009E96", "#008DCB", "D4ECEA", "#1D2088", "#920783"}

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
		/*
			for day in Mon..Fri
				for e in A1..A10
					for time in 9:00..18:45
						make table[][]
				str += table
		*/
		days := len(elem)

		allstr := ""
		fmt.Println(allkey)
		for d := 0; d < days; d++ {
			outputTable := [][]string{}
			for _, ele := range allkey {
				E := []string{ele}
				for _, e := range allinput[ele][d][1:] {
					if e == "1" {
						E = append(E, ele)
					} else {
						E = append(E, "")
					}
				}
				outputTable = append(outputTable, E)
			}
			fmt.Println(outputTable)
			// append table
			allstr += `<h1>` + template.HTMLEscapeString(elem[d][0]) + `</h1>`
			allstr += `<table id="table` + strconv.Itoa(d) + `" class="table is-bordered">`
			allstr += `<thead><tr>`
			for _, ele := range header {
				allstr += `<td>` + template.HTMLEscapeString(ele) + `</td>`
			}
			allstr += `</tr></thead>`
			allstr += `<tbody>`
			for _, st := range outputTable {
				allstr += `<tr>`
				for _, ele := range st {
					allstr += `<td>` + template.HTMLEscapeString(ele) + `</td>`
				}
				allstr += `</tr>`
			}
			allstr += `</tbody>`
			allstr += `</table>`
		}
		allstr += `</body><script>document.querySelectorAll('td').forEach((v,i) => {`
		for i, el := range allkey {
			color := ""
			if i < len(colors) {
				color = colors[i]
			}
			allstr += `if (v.innerText == "` + template.HTMLEscapeString(el) + `") {v.style.backgroundColor = "` + color + `";}`
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
				  padding: 20px;
			  }
		  </style>
		</head>

		<body>
		  <h1>SHIFT Visualization</h1>`+allstr+`</html>`)
	} else {
		fmt.Fprint(w, "not post request")
	}
}
