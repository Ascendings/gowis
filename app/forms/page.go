package forms

// CreatePageForm - form used for creating a page
type CreatePageForm struct {
	URLSlug       string `form:"url_slug" binding:"Required;MinSize(2)"`
	PageContent   string `form:"page_content"`
	CommitMessage string `form:"commit_message"`
}
