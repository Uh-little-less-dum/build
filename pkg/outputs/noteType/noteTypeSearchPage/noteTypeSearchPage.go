package outputs_note_type_search_page

type NoteTypeSearchPage struct {
	CategoryId string
	BaseUrl    string
}

func (n NoteTypeSearchPage) getTemplateData() map[string]string {
	data := make(map[string]string)
	data["CATEGORY_ID"] = n.CategoryId
	data["BASE_URL"] = n.BaseUrl
	return data
}

func NewNoteTypeSearchPage(categoryId, baseUrl string) NoteTypeSearchPage {
	return NoteTypeSearchPage{
		CategoryId: categoryId,
		BaseUrl:    baseUrl,
	}
}
