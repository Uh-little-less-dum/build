package slot_map

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"sync"

	target_paths "github.com/Uh-little-less-dum/build/pkg/targetPaths"
	string_utils "github.com/Uh-little-less-dum/go-utils/pkg/strings"
	"github.com/charmbracelet/log"
	"github.com/tidwall/gjson"
)

type SlotOutputFile struct {
	data              []string
	parentSlot        string
	subSlot           string
	slotItemData      gjson.Result
	pluginItemData    gjson.Result
	componentItemData gjson.Result
}

var replaceMeRegex *regexp.Regexp

var once sync.Once

func init() {
	once.Do(func() {
		r, err := regexp.Compile(`import\s+REPLACEME`)
		if err != nil {
			log.Fatal(err)
		}
		replaceMeRegex = r
	})
}

func (s *SlotOutputFile) TargetFileSubPath() string {
	data := s.slotItemData.Get("path").Str
	if data == "" {
		log.Fatalf("Cannot find slot data for %s/%s", s.parentSlot, s.subSlot)
	}
	return data
}

// slotMapItemData: parentSlot.subSlot item in the slotMap file, not the entire slotMapData json file.
// pluginItemData: pluginConfig.plugins item data
func NewSlotOutputFile(slotMapItemData, pluginItemData, componentItemData gjson.Result, parentSlot, subSlot string) *SlotOutputFile {
	return &SlotOutputFile{parentSlot: parentSlot, subSlot: subSlot, slotItemData: slotMapItemData, pluginItemData: pluginItemData, componentItemData: componentItemData}
}

func (s *SlotOutputFile) Data(filePath string) []string {
	if len(s.data) > 0 {
		return s.data
	}
	b, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	lines := string_utils.SplitLines(string(b))
	s.data = lines
	return lines
}

func (s *SlotOutputFile) ComponentName() string {
	return s.pluginItemData.Get("componentName").Str
}

func (s *SlotOutputFile) ExportsProps() string {
	return s.pluginItemData.Get("exportedPropsName").Str
}

func (s *SlotOutputFile) getPropsName() string {
	return fmt.Sprintf("%s%sProps", strings.ToUpper(s.subSlot[:1]), s.subSlot[1:])
}

func (s *SlotOutputFile) ImportedAs() string {
	return fmt.Sprintf("%s%s", strings.ToUpper(s.subSlot[:1]), s.subSlot[1:])
}

func (s *SlotOutputFile) removeREPLACEMEImport(filePath string) {
	lines := s.Data(filePath)
	var l []string
	for _, line := range lines {
		if !replaceMeRegex.MatchString(line) {
			l = append(l, line)
		}
	}
	s.data = l
}

func (s *SlotOutputFile) importLines(filePath string) (lines []string, lastIndex int) {
	data := s.Data(filePath)
	importLines := []string{}
	lastIdx := 0
	for i, l := range data {
		if strings.HasPrefix(strings.TrimSpace(l), "import ") {
			importLines = append(importLines, l)
			lastIdx = i
		}
	}
	return importLines, lastIdx
}

func (s *SlotOutputFile) appendImport(filePath string, importString string) {
	data := s.Data(filePath)
	_, idx := s.importLines(filePath)
	newLines := make([]string, len(data)+1)
	for i, l := range data {
		newLines = append(newLines, l)
		if i == idx {
			newLines = append(newLines, importString)
		}
	}
	s.data = newLines
}

func (s *SlotOutputFile) asString(filePath string) string {
	return string_utils.JoinLines(s.Data(filePath))
}

func (s *SlotOutputFile) replaceReplaceMe(targetFile string, importedAs string) {
	d := s.Data(targetFile)
	newLines := make([]string, len(d))
	for _, l := range d {
		newLines = append(newLines, strings.ReplaceAll(l, "REPLACEME", importedAs))
	}
	s.data = newLines
}

// FIX: Not yet implemented. Implement when on wifi and able to look over docs on Go regexps.
func (s *SlotOutputFile) AppendExportedType(targetFile string) {
	propsName := s.getPropsName()
	log.Error("This AppendExportedType method is not yet implemented.", propsName)
	d := s.Data(targetFile)
	newLines := make([]string, len(d))
	for _, l := range newLines {
		newLines = append(newLines, strings.ReplaceAll(l, "TemporaryComponentProps", propsName))
	}
	s.data = newLines
	// appendExportedType(componentName: string, exportedType: string) {
	//     let re = /^interface\s+TemporaryComponentProps\s*\{/gm
	//     let propsName = this.getPropsName(componentName)
	//     let lines = this.getLines().map((l) => {
	//         let t = l.trim()
	//         if(re.test(t)){
	//             return `interface ${propsName} extends ${exportedType} {}`
	//         }
	//         return l.replaceAll("TemporaryComponentProps", propsName)
	//     })
	//     this.content = lines.join("\n")
	// }
}

func (s *SlotOutputFile) importString(exportsProps, importedAs, exportedFrom string) string {
	if exportsProps == "" {
		return fmt.Sprintf("import %s from \"%s\"", importedAs, exportedFrom)
	}
	return fmt.Sprintf("import %s, { %s } from \"%s\"", importedAs, exportsProps, exportedFrom)
}

func (s *SlotOutputFile) WriteOutput(paths target_paths.TargetPaths) {
	targetFile := paths.JoinTargetDirString(s.TargetFileSubPath())
	importedAs := s.ImportedAs()
	exportedProps := s.ExportsProps()
	exportedFrom := fmt.Sprintf("%s/%s", s.pluginItemData.Get("pluginName").Str, s.componentItemData.Get("export"))
	if exportedProps != "" {
		s.AppendExportedType(targetFile)
	}
	s.removeREPLACEMEImport(targetFile)
	s.appendImport(targetFile, s.importString(exportedProps, importedAs, exportedFrom))
	s.replaceReplaceMe(targetFile, importedAs)
	val := s.asString(targetFile)
	err := os.WriteFile(targetFile, []byte(val), 0777)
	if err != nil {
		log.Fatal(err)
	}
}
