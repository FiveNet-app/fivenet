package main

import (
	"sort"
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
		Attributes            map[string]map[string]*Attr
	}{
		F:                     f,
		PermissionServiceKeys: []string{},
		Permissions:           map[string]map[string]*Perm{},
		PermissionRemap:       map[string]map[string]string{},
	}

	for _, s := range f.Services() {
		sName := string(s.Name())

		data.PermissionServiceKeys = append(data.PermissionServiceKeys, sName)
		p.Debugf("Service: %s (%s)", sName, data.PermissionServiceKeys)

		for _, m := range s.Methods() {
			mName := string(m.Name())
			mName = strings.TrimPrefix(mName, "services.")

			comment := m.SourceCodeInfo().LeadingComments()
			comment = strings.TrimLeft(comment, " ")
			if !strings.HasPrefix(comment, "@perm") {
				continue
			}
			comment = strings.TrimRight(comment, "\n")

			perm, err := p.parseComment(sName, mName, comment)
			if err != nil {
				p.Failf("failed to parse comment in %s method %s (comment: '%s'), error: %w", f.InputPath(), mName, comment, err)
				return
			}
			if perm == nil {
				p.Failf("failed to parse comment in %s method %s (comment: '%s')", f.InputPath(), mName, comment)
				return
			}

			if perm.Name != mName {
				remapServiceName := string(s.Name())
				if _, ok := data.PermissionRemap[remapServiceName]; !ok {
					data.PermissionRemap[remapServiceName] = map[string]string{}
				}
				if _, ok := data.PermissionRemap[remapServiceName][mName]; !ok {
					data.PermissionRemap[remapServiceName][mName] = perm.Name
					p.Debugf("Permission Remap added: %q -> %q\n", mName, perm.Name)
				} else {
					p.Debugf("Permission Remap already exists: %q -> %q\n", mName, perm.Name)
				}
			}

			if perm.Name == "SuperUser" || perm.Name == "Any" {
				continue
			}

			if _, ok := data.Permissions[sName]; !ok {
				data.Permissions[sName] = map[string]*Perm{}
			}
			if _, ok := data.Permissions[sName][perm.Name]; !ok {
				data.Permissions[sName][perm.Name] = perm
				p.Debugf("Permission added: %q - %+v\n", mName, perm)
			} else {
				p.Debugf("Permission already in list: %q - %+v\n", mName, perm)
			}
		}
	}

	if len(data.Permissions) == 0 && len(data.PermissionRemap) == 0 {
		return
	}

	sort.Strings(data.PermissionServiceKeys)

	name := p.ctx.OutputPath(f).SetExt(".perms.go")
	p.AddGeneratorTemplateFile(name.String(), p.tpl, data)
}

func (p *PermifyModule) parseComment(service string, method string, comment string) (*Perm, error) {
	comment = strings.TrimPrefix(comment, "@perm: ")
	comment = strings.TrimPrefix(comment, "@perm")

	perm := &Perm{
		Name: method,
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
		case "name":
			perm.Name = v
			continue
		case "attrs":
			for _, v := range strings.Split(v, "|") {
				attrSplit := strings.Split(v, "/")
				if len(attrSplit) <= 1 {
					p.Fail("Invalid attrs value found: ", v)
				}

				attrType := attrSplit[1]
				validValue := ""
				validList := strings.Split(attrSplit[1], ":")
				if len(validList) > 1 {
					attrType = validList[0]
					validValue = strings.Join(validList[1:], ":")
				}

				defaultValue := ""
				defaultSplit := strings.Split(validValue, "§")
				if len(defaultSplit) > 1 {
					validValue = defaultSplit[0]
					defaultValue = defaultSplit[1]
				}

				perm.Attrs = append(perm.Attrs, Attr{
					Key:     attrSplit[0],
					Type:    attrType,
					Valid:   validValue,
					Default: defaultValue,
				})
			}
			continue
		}
	}

	return perm, nil
}

const permifyTpl = `// Code generated by protoc-gen-customizer. DO NOT EDIT.
// source: {{ .F.File.InputPath }}

package {{ package .F }}

import (
    "github.com/galexrt/fivenet/pkg/config"
    "github.com/galexrt/fivenet/pkg/perms"
    "github.com/galexrt/fivenet/gen/go/proto/resources/permissions"
)

{{ with .PermissionServiceKeys }}
const (
{{ range $key, $sName := . -}}
    {{ $sName }}Perm perms.Category = "{{ $sName }}"
{{ end }}

{{ range $sName, $service := $.Permissions -}}
	    {{- range $perm := $service }}
	{{ $sName }}{{ $perm.Name }}Perm perms.Name = "{{ $perm.Name }}"
            {{- range $attr := $perm.Attrs }}
    {{ $sName }}{{ $perm.Name }}{{ $attr.Key }}PermField perms.Key = "{{ $attr.Key }}"
            {{- end }}
		{{- end }}
	{{- end }}
)
{{ end }}

{{ with .PermissionRemap }}
var PermsRemap = map[string]string{
    {{ range $service, $remap := . }}
	// Service: {{ $service }}
	{{ range $key, $target := $remap -}}
	"{{ $service }}/{{ $key }}": "{{- if and (ne $target "SuperUser") (ne $target "Any") }}{{ $service }}/{{ end }}{{ $target }}",
    {{ end }}
    {{ end }}
}

func (s *Server) GetPermsRemap() map[string]string {
    return PermsRemap
}

{{ end }}

func init() {
	perms.AddPermsToList([]*perms.Perm{
	{{ range $sName, $service := .Permissions -}}
		// Service: {{ $sName }}
		{{- range $perm := $service }}
		{
			Category: {{ $sName }}Perm,
			Name: {{ $sName }}{{ $perm.Name }}Perm,
            Attrs: []perms.Attr{
            {{- range $attr := $perm.Attrs }}
                {
                    Key: {{ $sName }}{{ $perm.Name }}{{ $attr.Key }}PermField,
                    Type: permissions.{{ $attr.Type }}AttributeType,
                    {{ with $attr.Valid -}}ValidValues: {{ $attr.Valid }},{{ end }}
                    {{ with $attr.Default -}}DefaultValues: {{ $attr.Default }},{{ end }}
                },
            {{- end }}
            },
		},
		{{- end }}
	{{- end }}
	})
}
`

type Perm struct {
	Name  string
	Attrs []Attr
}

type Attr struct {
	Key     string
	Type    string
	Valid   string
	Default string
}
