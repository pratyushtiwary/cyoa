package templates

import "text/template"

var DefaultTemplate *template.Template

/* TEMPLATES */
const DEFAULT_TEMPLATE = `
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>Choose Your Own Adventure</title>
</head>
<body>
	<section class="page">
		<h1>{{.Title}}</h1>
		{{range .Paragraphs}}
		<p>{{.}}</p>
		{{end}}
		<ul>
		{{range .Options}}
			<li><a href="/{{.Chapter}}">{{.Text}}</a></li>
		{{end}}
		</ul>
	</section>
</body>
</html>
`

/* Package Init */
func init() {
	DefaultTemplate = template.Must(template.New("").Parse(DEFAULT_TEMPLATE))
}
