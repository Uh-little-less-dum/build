package outputs_note_type_search_page

import (
	"embed"
	"io"
	"text/template"

	"github.com/Uh-little-less-dum/go-utils/pkg/fs/file"
	"github.com/charmbracelet/log"
)

var (
	//go:embed "templates/*"
	templateFiles embed.FS
)

type PageWriter struct {
	templ *template.Template
}

func NewPageWriter() (*PageWriter, error) {
	templ, err := template.ParseFS(templateFiles, "templates/*.gotsx")
	if err != nil {
		return nil, err
	}

	return &PageWriter{templ: templ}, nil
}

func (r *PageWriter) WriteOutput(w io.Writer, p NoteTypeSearchPage) error {

	if err := r.templ.ExecuteTemplate(w, "noteTypeSearchPage.gotsx", p.getTemplateData()); err != nil {
		return err
	}

	return nil
}

func (r PageWriter) WriteToPath(p string, data NoteTypeSearchPage) {
	f := file.NewFileItem(p)
	err := r.WriteOutput(f, data)
	if err != nil {
		log.Fatal(err)
	}
}
