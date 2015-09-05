package main

import (
	"log"
	"os"
	"text/template"
)

const tpl = `
<!DOCTYPE html>
<html>
<head>
  <meta charset="UTF-8">
</head>
<body>
  {{range .Items}}
    <div>
        <a href="{{.Link}}" target="_blank" style="text-decoration:none">
          <h3>{{.Title}}</h3>
       </a>
       <p style="text-align: center;"><em >{{.PubDate}}</em></p>
       <hr style="width: 70%;">
       {{.Description}}
       <hr><hr>
    </div>
    {{else}}
    <div><strong>no news</strong></div>{{end}}
</body>
</html>
`

func ParseToHtml(items []Item) {

	check := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}
	t, err := template.New("webpage").Parse(tpl)
	data := struct {
		Items []Item
	}{
		Items: items}
	file, err := os.Create("index.html")
	check(err)
	err = t.Execute(file, data)
	check(err)

}
