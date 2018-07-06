package factorish

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"

	dom "github.com/gowasm/go-js-dom"
	"github.com/gowasm/vdom"
)

var reg *registry

// App is just a special component
type App struct {
	Component // the root of the dom
}

// Component is a thing inside of the vdom
type Component struct {
	name string
	tpl  *template.Template
	node *vdom.Node
}

func (c Component) html(data interface{}) ([]byte, error) {
	b := new(bytes.Buffer)
	// data?
	comment := fmt.Sprintf("<!-- Component: %s -->\n", c.name)
	b.WriteString(comment)
	if err := c.tpl.Execute(b, data); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

// Render renders the component into a dom tree
func (c Component) Render(data interface{}) (*vdom.Tree, error) {
	// not caching
	html, err := c.html(data)
	if err != nil {
		return nil, err
	}
	return vdom.Parse(html)
}

type registry struct {
	components map[string]*Component
}

func newRegistry() *registry {
	c := map[string]*Component{}
	return &registry{
		components: c,
	}
}
func (r *registry) componentByName(name string) (*Component, error) {
	c, ok := r.components[name]
	if !ok {
		return nil, errors.New("component not registered")
	}
	return c, nil
}
func (r *registry) registerComponent(name, tplStr string) error {
	tpl, err := template.New(name).Parse(tplStr)
	if err != nil {
		return err
	}
	c := &Component{name: name, tpl: tpl}
	r.components[name] = c
	return nil
}
func ComponentByName(name string) (*Component, error) {
	checkRegistry()
	return reg.componentByName(name)
}
func RegisterComponent(name, tplStr string) error {
	checkRegistry()
	return reg.registerComponent(name, tplStr)
}
func checkRegistry() {
	if reg == nil {
		reg = newRegistry()
	}
}

// I think this is basically a
type NodeManager interface {
	NodeForVNode(n *vdom.Node) (dom.Node, error)
	Add(vdom.Node, dom.Node) error
}

type DomManager struct {
	mgr NodeManager
}

func (d DomManager) Apply(comp *Component, data interface{}) error {
	domNode, err := d.mgr.NodeForVNode(comp.node)
	if err != nil {
		return err
	}

	tree, err := comp.Render(data)
	if err != nil {
		return err
	}
	domNode.SetTextContent(string(tree.HTML()))
	return nil
}
