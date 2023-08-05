package templates

const (
	README_TEMPLATE = "README.tmpl"
	README_FILE     = "README.md"
)

type ReadmeInfo struct {
	AuthorName  string
	Description string
}

func NewReadmeInfo(authorName, description string) *ReadmeInfo {
	return &ReadmeInfo{
		AuthorName:  authorName,
		Description: description,
	}
}
