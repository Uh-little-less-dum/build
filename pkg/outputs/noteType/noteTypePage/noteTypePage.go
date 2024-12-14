package outputs_note_type_page

type NoteTypePage struct {
	Id string
}

func (n NoteTypePage) getTemplateData() map[string]string {
	data := make(map[string]string)
	data["CATEGORY_ID"] = n.Id
	return data
}

func NewNoteTypePage(noteTypeId string) NoteTypePage {
	return NoteTypePage{
		Id: noteTypeId,
	}
}
