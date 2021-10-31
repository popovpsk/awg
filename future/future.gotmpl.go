package future

var futureTemplateStr = `{{- /*gotype: awg/future.Future*/ -}}
// {{.Name}}Async generated async wrapper for {{.Name}} function
func {{.GetRecvStr}} {{.Name}}Async({{$first := true}}{{range .Params}}{{if $first}}{{$first = false}}{{else}}, {{end}}{{.Name}} {{.T}}{{end}}) func() {{.ResultStr}} {
	var (
		done = make(chan struct{})
		{{range .Results}}
		{{.Name}} {{.T}}{{end}}
	)
	go func() {
		defer close(done)
		{{if .Results}}{{.GetVars}} = {{.GetCall}}{{else}}{{.GetCall}}{{end}}
	}()
	return func()  {{.ResultStr}} {
		<-done
		return {{.GetVars}}
	}
}
`

var callTemplateStr = `{{if .Recv}}{{.Recv.Name}}.{{end}}{{.Name}}({{$first := true}}{{range .Params}}{{if $first}}{{$first = false}}{{else}}, {{end}}{{.Name}}{{end}})`
var varsTemplateStr = `{{$first := true}}{{if .Results}}{{range .Results}}{{if $first}}{{$first = false}}{{else}}, {{end}}{{.Name}}{{end}}{{end}}`