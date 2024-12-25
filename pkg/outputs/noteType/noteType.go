package note_type

import (
	"path/filepath"
	"strings"
	"sync"

	outputs_note_type_page "github.com/Uh-little-less-dum/build/pkg/outputs/noteType/noteTypePage"
	outputs_note_type_search_page "github.com/Uh-little-less-dum/build/pkg/outputs/noteType/noteTypeSearchPage"
	target_paths "github.com/Uh-little-less-dum/build/pkg/targetPaths"
	schemas_app_config "github.com/Uh-little-less-dum/go-utils/pkg/schemastructs/ulldAppConfig"
	"github.com/charmbracelet/log"
)

func getNoteTypePaths(noteUrl string, pathData target_paths.TargetPaths) (noteTypeSearchPage string, noteTypePage string) {
	items := []string{pathData.AppDir()}
	items = append(items, strings.Split(noteUrl, string(filepath.Separator))...)
	dirPath := filepath.Join(items...)
	return filepath.Join(dirPath, "page.tsx"), filepath.Join(dirPath, "[...slug]", "page.tsx")
}

func WriteNoteTypeOutputs(data []schemas_app_config.NoteType, pathData target_paths.TargetPaths, wg *sync.WaitGroup) {
	searchWriter, err := outputs_note_type_search_page.NewPageWriter()
	if err != nil {
		log.Fatal(err)
	}
	pageWriter, err := outputs_note_type_page.NewPageWriter()
	if err != nil {
		log.Fatal(err)
	}
	for _, nt := range data {
		searchPath, noteTypePagePath := getNoteTypePaths(*nt.URL, pathData)
		wg.Add(1)
		go func() {
			defer wg.Done()
			searchPage := outputs_note_type_search_page.NewNoteTypeSearchPage(*nt.ID, *nt.URL)
			searchWriter.WriteToPath(searchPath, searchPage)
		}()
		wg.Add(1)
		go func() {
			defer wg.Done()
			individualNotePage := outputs_note_type_page.NewNoteTypePage(*nt.ID)
			pageWriter.WriteToPath(noteTypePagePath, individualNotePage)
		}()
	}
}
