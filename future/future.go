package future

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/popovpsk/awg/types"
)

type Future struct {
	Params  []types.Variable
	Results []types.Variable
	Recv    *types.Variable
	Name    string

	ResultStr string
}

var futureTemplate = template.Must(template.New("future").Parse(futureTemplateStr))
var callTemplate = template.Must(template.New("call").Parse(callTemplateStr))
var varsTemplate = template.Must(template.New("vars").Parse(varsTemplateStr))

func (f *Future) generateFields() {
	if len(f.Results) > 0 {
		builder := strings.Builder{}
		builder.WriteString("(")
		for i, v := range f.Results {
			if i > 0 {
				builder.WriteString(", ")
			}
			builder.WriteString(fmt.Sprintf("%s %s ", v.Name, v.T))
			f.Results[i].Name = fmt.Sprintf("p%d", i)
		}
		builder.WriteString(")")
		str := builder.String()
		f.ResultStr = str
	}
}

func (f *Future) GetCall() string {
	return f.fillTemplate(callTemplate)
}

func (f *Future) GetVars() string {
	return f.fillTemplate(varsTemplate)
}

func (f *Future) GetRecvStr() string {
	if f.Recv != nil {
		return fmt.Sprintf("(%s %s)", f.Recv.Name, f.Recv.T)
	}
	return ""
}

func (f *Future) GenerateFunc() string {
	f.generateFields()
	b := bytes.NewBuffer([]byte{})
	_ = futureTemplate.Execute(b, f)
	return b.String()
}

func (f *Future) fillTemplate(tmpl *template.Template) string {
	b := bytes.NewBuffer(make([]byte, 0, 32))
	_ = tmpl.Execute(b, f)
	return b.String()
}
