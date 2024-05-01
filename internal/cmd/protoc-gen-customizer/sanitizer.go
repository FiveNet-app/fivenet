package main

import (
	"strings"
	"text/template"

	pgs "github.com/lyft/protoc-gen-star/v2"
	pgsgo "github.com/lyft/protoc-gen-star/v2/lang/go"
)

// SanitizerPlugin
type SanitizerModule struct {
	*pgs.ModuleBase
	ctx pgsgo.Context
	tpl *template.Template
}

// Sanitizer returns an initialized SanitizerPlugin
func Sanitizer() *SanitizerModule { return &SanitizerModule{ModuleBase: &pgs.ModuleBase{}} }

func (p *SanitizerModule) InitContext(c pgs.BuildContext) {
	p.ModuleBase.InitContext(c)
	p.ctx = pgsgo.InitContext(c.Parameters())

	tpl := template.New("sanitizer").Funcs(map[string]interface{}{
		"package": p.ctx.PackageName,
		"name":    p.ctx.Name,
	})

	p.tpl = template.Must(tpl.Parse(sanitizerTpl))
}

// Name satisfies the generator.Plugin interface.
func (p *SanitizerModule) Name() string { return "sanitizer" }

func (p *SanitizerModule) Execute(targets map[string]pgs.File, pkgs map[string]pgs.Package) []pgs.Artifact {
	for _, t := range targets {
		p.generate(t)
	}

	return p.Artifacts()
}

func (p *SanitizerModule) generate(f pgs.File) {
	if len(f.Messages()) == 0 {
		return
	}

	data := struct {
		F    pgs.File
		FMap map[string]map[string]*Sanitize
	}{
		F:    f,
		FMap: map[string]map[string]*Sanitize{},
	}

	for _, m := range f.Messages() {
		data.FMap[string(m.Name())] = map[string]*Sanitize{}

		for _, f := range m.Fields() {
			if f.Type().ProtoType() != pgs.StringT {
				continue
			}

			comment := f.SourceCodeInfo().LeadingComments()
			comment = strings.TrimLeft(comment, " ")
			comment = strings.TrimRight(comment, "\n")

			if f.Type().ProtoType().IsNumeric() ||
				f.Type().IsMap() || f.Type().IsRepeated() {
				continue
			}

			if f.Type().ProtoType() == pgs.StringT {
				if !strings.HasPrefix(comment, "@sanitize") {
					continue
				}
			}

			s, err := p.parseComment(f, comment)
			if err != nil {
				p.Fail(err)
			}

			data.FMap[string(m.Name())][string(f.Name().UpperCamelCase())] = s
		}

		if len(data.FMap[string(m.Name())]) == 0 {
			delete(data.FMap, string(m.Name()))
		}
	}

	if len(data.FMap) > 0 {
		name := p.ctx.OutputPath(f).SetExt(".sanitizer.go")
		p.AddGeneratorTemplateFile(name.String(), p.tpl, data)
	}
}

func (p *SanitizerModule) parseComment(field pgs.Field, comment string) (*Sanitize, error) {
	comment = strings.TrimPrefix(comment, "@sanitize: ")
	comment = strings.TrimPrefix(comment, "@sanitize")

	s := &Sanitize{
		Name:     field.Name().UpperCamelCase().String(),
		Method:   "Sanitize",
		Optional: field.HasOptionalKeyword(),
		Deeper:   true,
	}

	if field.Type().ProtoType() != pgs.StringT {
		return s, nil
	}

	s.Deeper = false
	split := strings.Split(comment, ";")

	for i := 0; i < len(split); i++ {
		k, v, _ := strings.Cut(split[i], "=")
		if v == "" {
			continue
		}

		switch strings.ToLower(k) {
		case "method":
			s.Method = v
			continue
		}
	}

	return s, nil
}

const sanitizerTpl = `// Code generated by protoc-gen-customizer. DO NOT EDIT.
// source: {{ .F.InputPath }}

package {{ package .F }}

import (
    "github.com/fivenet-app/fivenet/pkg/htmlsanitizer"
)

{{ range $key, $fields := .FMap }}
func (m *{{ $key }}) Sanitize() error {
    {{ range $f := $fields }}
        {{ if $f.Deeper }}
    if v, ok := interface{}(m.Get{{ $f.Name }}()).(interface{ Sanitize() error }); ok {
        if err := v.Sanitize(); err != nil {
            return err
        }
    }
        {{ else }}

            {{ if $f.Optional }}
    if m.{{ $f.Name }} != nil {
            {{ end -}}
        {{ if $f.Optional }}*{{ end }}m.{{ $f.Name }} = htmlsanitizer.{{ $f.Method }}({{ if $f.Optional }}*{{ end }}m.{{ $f.Name }})
            {{- if $f.Optional  }}
    }
            {{- end }}

        {{ end }}

    {{ end }}

	return nil
}
{{ end }}
`

type Sanitize struct {
	Name     string
	Method   string
	Optional bool

	Deeper bool
}
