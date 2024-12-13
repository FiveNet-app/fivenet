package content

import (
	"bytes"
	"database/sql/driver"
	"fmt"
	"slices"
	"strings"

	"github.com/fivenet-app/fivenet/pkg/utils"
	"github.com/fivenet-app/fivenet/pkg/utils/protoutils"
	"github.com/yosssi/gohtml"
	"golang.org/x/net/html"
	"google.golang.org/protobuf/encoding/protojson"
)

type NodeType string

const (
	Version_v0 = "v0"

	RootNodeType NodeType = "root"

	ElementNodeType NodeType = "element"
	TextNodeType    NodeType = "text"
	CommentNodeType NodeType = "comment"
)

// Scan implements driver.Valuer for protobuf Content.
func (x *Content) Scan(value any) error {
	switch t := value.(type) {
	case string:
		if strings.HasPrefix(t, "{") {
			if err := protojson.Unmarshal([]byte(t), x); err != nil {
				return err
			}

			return x.Populate()
		}

		h, err := ParseHTML(t)
		if err != nil {
			return err
		}

		out, err := FromHTMLNode(h)
		if err != nil {
			return err
		}

		v := Version_v0
		x.Version = &v
		x.Content = out
		x.RawContent = &t

	case []byte:
		if strings.HasPrefix(string(t), "{") {
			if err := protojson.Unmarshal(t, x); err != nil {
				return err
			}

			return x.Populate()
		}

		hRaw := string(t)
		h, err := ParseHTML(hRaw)
		if err != nil {
			return err
		}

		out, err := FromHTMLNode(h)
		if err != nil {
			return err
		}

		v := Version_v0
		x.Version = &v
		x.Content = out
		x.RawContent = &hRaw

	default:
		return fmt.Errorf("invalid format for content")
	}

	return nil
}

// Value marshals the value into driver.Valuer.
func (x *Content) Value() (driver.Value, error) {
	if x == nil {
		return nil, nil
	}

	// If the raw content isn't nil, need to "encode" it to `JSONNode` for the `Content` field
	if x.RawContent != nil && *x.RawContent != "" {
		h, err := ParseHTML(*x.RawContent)
		if err != nil {
			return nil, err
		}

		x.Content, err = FromHTMLNode(h)
		if err != nil {
			return nil, err
		}
	}

	return protoutils.Marshal(&Content{
		Version: x.Version,
		Content: x.Content,
	})
}

func (x *Content) Populate() error {
	out, err := x.Content.ToHTMLP()
	if err != nil {
		return err
	}

	x.RawContent = &out

	return nil
}

func (n *JSONNode) populateFrom(htmlNode *html.Node) (*JSONNode, error) {
	if htmlNode.Parent == nil {
		n.Type = string(RootNodeType)
	} else {
		n.Type = string(ElementNodeType)
	}

	switch htmlNode.Type {
	case html.ElementNode:
		n.Tag = htmlNode.Data
		break

	case html.DocumentNode:
		break

	default:
		return nil, fmt.Errorf("given node needs to be an element or document")
	}

	if len(htmlNode.Attr) > 0 {
		n.Attributes = make(map[string]string)
		var a html.Attribute
		for _, a = range htmlNode.Attr {
			key := strings.ToLower(a.Key)
			switch key {
			case "id":
				// Skip empty id attribute
				if strings.TrimSpace(a.Val) == "" {
					continue
				}

				n.Id = a.Val

			case "class":
				// Skip empty class attribute
				if strings.TrimSpace(a.Val) == "" {
					continue
				}

				n.Class = a.Val

			case "style":
				// Skip empty class attribute
				if strings.TrimSpace(a.Val) == "" {
					continue
				}

				// Clean style options - maybe sort them in the future?
				val := strings.Replace(a.Val, " ", "", -1)
				if !strings.HasSuffix(val, ";") {
					val += ";"
				}

				n.Attributes[key] = val

			default:
				// Don't skip empty valued attributes as they might have a meaning on the frontend-side
				n.Attributes[key] = a.Val
			}
		}
	}

	var textBuffer bytes.Buffer
	e := htmlNode.FirstChild
	for e != nil {
		switch e.Type {
		case html.TextNode:
			n.Type = string(TextNodeType)

			trimmed := strings.TrimSpace(e.Data)
			if len(trimmed) > 0 {
				// mimic HTML text normalizing
				if textBuffer.Len() > 0 {
					textBuffer.WriteString(" ")
				}
				textBuffer.WriteString(trimmed)
			}

		case html.ElementNode:
			if n.Children == nil {
				n.Children = make([]*JSONNode, 0)
			}

			jsonElemNode := &JSONNode{}
			if _, err := jsonElemNode.populateFrom(e); err != nil {
				return nil, err
			}

			n.Children = append(n.Children, jsonElemNode)
		}

		e = e.NextSibling
	}

	if textBuffer.Len() > 0 {
		n.Text = textBuffer.String()
	}

	if strings.HasPrefix(n.Tag, "h") {
		if n.Id == "" && (len(n.Children) > 0 || n.Text != "") {
			if n.Text != "" {
				n.Id = utils.SlugNoDots(fmt.Sprintf("%s-%s", n.Tag, n.Text))
			} else {
				n.Id = utils.SlugNoDots(fmt.Sprintf("%s-%s", n.Tag, n.Children[0].Text))
			}
		}
	}

	return n, nil
}

func (n *JSONNode) populateTo(htmlNode *html.Node) *JSONNode {
	if n.Tag != "" {
		htmlNode.Data = n.Tag
		htmlNode.Type = html.ElementNode
	} else {
		htmlNode.Type = html.DocumentNode
	}

	if n.Class != "" {
		htmlNode.Attr = append(htmlNode.Attr, html.Attribute{
			Key: "class",
			Val: n.Class,
		})
	}
	if n.Id != "" {
		htmlNode.Attr = append(htmlNode.Attr, html.Attribute{
			Key: "id",
			Val: n.Id,
		})
	}

	keys := make([]string, 0, len(n.Attributes))
	for k := range n.Attributes {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	for _, k := range keys {
		htmlNode.Attr = append(htmlNode.Attr, html.Attribute{
			Key: k,
			Val: n.Attributes[k],
		})
	}

	if n.Text != "" {
		htmlNode.AppendChild(&html.Node{
			Type: html.TextNode,
			Data: n.Text,
		})
	}

	for _, e := range n.Children {
		htmlElem := &html.Node{}
		e.populateTo(htmlElem)
		htmlNode.AppendChild(htmlElem)
	}

	return n
}

func ParseHTML(in string) (*html.Node, error) {
	d, err := html.Parse(strings.NewReader(in))
	if err != nil {
		return nil, err
	}

	var body *html.Node
	var crawler func(*html.Node)
	crawler = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "body" {
			body = node
			body.Data = "body"
			body.Parent = nil
			body.PrevSibling = nil
			body.NextSibling = nil
			return
		}

		for child := node.FirstChild; child != nil; child = child.NextSibling {
			crawler(child)
		}
	}

	crawler(d.LastChild)

	if body != nil {
		return body, nil
	}

	return d, nil
}

// FromHTMLNode
func FromHTMLNode(node *html.Node) (*JSONNode, error) {
	jNode := &JSONNode{}
	jNode.populateFrom(node)

	return jNode, nil
}

// ToHTMLNode
func (n *JSONNode) ToHTMLNode() (*html.Node, error) {
	node := &html.Node{}

	n.populateTo(node)

	return node, nil
}

// ToHTML HTML potentially not pretty
func (n *JSONNode) ToHTML() (string, error) {
	h, err := n.ToHTMLNode()
	if err != nil {
		return "", err
	}

	htmlBuff := &bytes.Buffer{}
	if err := html.Render(htmlBuff, h); err != nil {
		return "", err
	}

	out := htmlBuff.String()
	out = strings.ReplaceAll(out, "<body>", "")
	out = strings.ReplaceAll(out, "</body>", "")

	return out, nil
}

// ToHTMLP Pretty HTML
func (n *JSONNode) ToHTMLP() (string, error) {
	h, err := n.ToHTML()
	if err != nil {
		return "", err
	}

	h = strings.ReplaceAll(h, "<body>\n", "")
	h = strings.ReplaceAll(h, "\n</body>", "")

	return gohtml.Format(h), nil
}

func PrettyHTML(in string) (string, error) {
	in = strings.ReplaceAll(in, "<body>\n", "")
	in = strings.ReplaceAll(in, "\n</body>", "")

	return gohtml.Format(in), nil
}
