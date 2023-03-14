package main

import (
	"fmt"
	"strconv"
	"strings"
	"text/template"

	pgsgo "github.com/lyft/protoc-gen-star/v2/lang/go"

	pgs "github.com/lyft/protoc-gen-star/v2"
)

// PermifyPlugin
type PermifyModule struct {
	*pgs.ModuleBase
	ctx pgsgo.Context
	tpl *template.Template
}

// Permify returns an initialized PermifyPlugin
func Permify() *PermifyModule { return &PermifyModule{ModuleBase: &pgs.ModuleBase{}} }

func (p *PermifyModule) InitContext(c pgs.BuildContext) {
	p.ModuleBase.InitContext(c)
	p.ctx = pgsgo.InitContext(c.Parameters())

	tpl := template.New("permify").Funcs(map[string]interface{}{
		"package": p.ctx.PackageName,
		"name":    p.ctx.Name,
	})

	p.tpl = template.Must(tpl.Parse(permifyTpl))
}

// Name satisfies the generator.Plugin interface.
func (p *PermifyModule) Name() string { return "permify" }

func (p *PermifyModule) Execute(targets map[string]pgs.File, pkgs map[string]pgs.Package) []pgs.Artifact {
	for _, t := range targets {
		p.generate(t)
	}

	return p.Artifacts()
}

func (p *PermifyModule) generate(f pgs.File) {
	if len(f.Services()) == 0 {
		return
	}

	data := struct {
		F                     pgs.File
		PermissionServiceKeys []string
		Permissions           map[string]map[string]*Perm
		PermissionRemap       map[string]map[string]string
	}{
		F:                     f,
		PermissionServiceKeys: []string{},
		Permissions:           map[string]map[string]*Perm{},
		PermissionRemap:       map[string]map[string]string{},
	}

	for _, s := range f.Services() {
		sName := string(s.Name())

		data.PermissionServiceKeys = append(data.PermissionServiceKeys, sName)

		for _, m := range s.Methods() {
			mName := string(m.Name())
			mName = strings.TrimPrefix(mName, "services.")

			comment := m.SourceCodeInfo().LeadingComments()
			comment = strings.TrimLeft(comment, " ")
			if !strings.HasPrefix(comment, "@permission") {
				continue
			}
			comment = strings.TrimRight(comment, "\n")

			perm, err := p.parseComment(sName, mName, comment)
			if err != nil {
				panic(fmt.Sprintf("failed to parse comment in %s method %s (comment: '%s'), error: %+v", f.InputPath(), mName, comment, err))
			}
			if perm == nil {
				panic(fmt.Sprintf("failed to parse comment in %s method %s (comment: '%s')", f.InputPath(), mName, comment))
			}

			if perm.Name != mName {
				remapServiceName := string(s.Name())
				if _, ok := data.PermissionRemap[remapServiceName]; !ok {
					data.PermissionRemap[remapServiceName] = map[string]string{}
				}
				if _, ok := data.PermissionRemap[remapServiceName][mName]; !ok {
					data.PermissionRemap[remapServiceName][mName] = perm.Name
				}
				p.Debugf("Permission Remap added: %q -> %q\n", mName, perm.Name)
			}

			if _, ok := data.Permissions[sName]; !ok {
				data.Permissions[sName] = map[string]*Perm{}
			}
			if _, ok := data.Permissions[sName][mName]; !ok {
				data.Permissions[sName][mName] = perm
			}
			p.Debugf("Permission added: %q - %+v\n", mName, perm)
		}
	}

	if len(data.Permissions) == 0 {
		return
	}

	name := p.ctx.OutputPath(f).SetExt(".perms.go")
	p.AddGeneratorTemplateFile(name.String(), p.tpl, data)
}

func (p *PermifyModule) parseComment(service string, method string, comment string) (*Perm, error) {
	comment = strings.TrimPrefix(comment, "@permission: ")
	comment = strings.TrimPrefix(comment, "@permission")

	perm := &Perm{
		Name:        method,
		PerJob:      false,
		Description: "",
		Fields:      []string{},
	}

	if comment == "" {
		return perm, nil
	}

	split := strings.Split(comment, ";")

	for i := 0; i < len(split); i++ {
		k, v, _ := strings.Cut(split[i], "=")
		if v == "" {
			continue
		}

		switch strings.ToLower(k) {
		case "key":
			perm.Key = v
			continue
		case "name":
			perm.Name = v
			continue
		case "perjob":
			bo, err := strconv.ParseBool(v)
			if err != nil {
				return nil, err
			}
			perm.PerJob = bo
			continue
		case "description":
			perm.Description = v
			continue
		case "fields":
			fields := strings.Split(v, ",")
			perm.Fields = fields
			continue
		}
	}

	return perm, nil
}

const permifyTpl = `// Code generated by protoc-gen-customizer. DO NOT EDIT.
// source: {{ .F.File.InputPath }}

package {{ package .F }}

import "github.com/galexrt/arpanet/pkg/perms"

{{ with .PermissionRemap }}
{{ range $service, $remap := . }}
var PermsRemap = map[string]string{
	// Service: {{ $service }}
	{{ range $key, $target := $remap -}}
	"{{ $service }}/{{ $key }}": "{{ $service }}/{{ $target }}",
{{ end }}
}
{{ end }}

func (s *Server) GetPermsRemap() map[string]string {
    return PermsRemap
}

{{ end }}

{{ with .PermissionServiceKeys }}
const (
{{ range $key, $sName := . -}}
    {{ $sName }}PermKey = "{{ $sName }}"
{{ end }}
)
{{ end }}

func init() {
	perms.AddPermsToList([]*perms.Perm{
	{{ range $sName, $service := .Permissions -}}
		// Service: {{ $sName }}
		{{- range $perm := $service }}
		{
			Key: {{ if $perm.Key }}"{{ $perm.Key }}"{{ else }}{{ $sName }}PermKey{{- end -}},
			Name: "{{ $perm.Name }}",
			{{- if $perm.PerJob }}
            PerJob: {{ $perm.PerJob }},{{ end }}
			{{- if $perm.Fields }}
            Fields: []string{ {{ range $i, $field := $perm.Fields }}"{{ $field }}",{{ end -}} },{{ end }}
		},
		{{- end }}
	{{- end }}
	})
}
`

type Perm struct {
	Key         string
	Name        string
	Description string
	Fields      []string
	PerJob      bool
}
