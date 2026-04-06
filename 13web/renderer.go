package main

import (
	"html/template"
	"net/http"
	"path"
	"path/filepath"
	"sync"
)

type TemplateRenderer struct {
	cache       map[string]*template.Template
	mutex       sync.RWMutex
	dev         bool // if we are in development we don't want to cache
	templateDir string
}

// seq returns [1, 2, ..., n] so templates can do {{range $i := seq 5}}
func seq(n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = i + 1
	}
	return s
}

func NewTemplateRenderer(templateDir string, isDev bool) *TemplateRenderer {
	return &TemplateRenderer{
		cache:       make(map[string]*template.Template),
		dev:         isDev,
		templateDir: templateDir,
	}
}

func (t *TemplateRenderer) parseTemplate(templateName string) (*template.Template, error) {
	templatePath := path.Join(t.templateDir, templateName)

	files := []string{templatePath}

	layoutPath := path.Join(t.templateDir, "layouts/*.html")
	layouts, err := filepath.Glob(layoutPath)

	if err == nil {
		files = append(files, layouts...)
	}

	partialPath := path.Join(t.templateDir, "partials/*.html")
	partials, err := filepath.Glob(partialPath)
	if err == nil {
		files = append(files, partials...)
	}

	tmpl, err := template.ParseFiles(files...)

	if err != nil {
		return nil, err
	}

	return tmpl, nil
}

func (t *TemplateRenderer) getTemplate(templateName string) (*template.Template, error) {
	if !t.dev {
		t.mutex.RLock()
		if tmpl, ok := t.cache[templateName]; ok {
			t.mutex.RUnlock()
			return tmpl, nil
		}
		t.mutex.RUnlock()
	}

	tmpl, err := t.parseTemplate(templateName)
	if err != nil {
		return nil, err
	}

	if !t.dev {
		t.mutex.Lock()
		t.cache[templateName] = tmpl
		t.mutex.Unlock()
	}

	return tmpl, nil
}

func (t *TemplateRenderer) Render(w http.ResponseWriter, name string, data interface{}) {
	tmpl, err := t.getTemplate(name)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "base.html", data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
