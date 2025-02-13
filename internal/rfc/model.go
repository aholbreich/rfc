package rfc

// RFCMetadata represents metadata for an individual RFC
type RFCMetadata struct {
	Title     string `yaml:"title"`
	Author    string `yaml:"author"`
	Status    string `yaml:"status"`
	CreatedAt string `yaml:"created_at"`
}

// TemplateData holds the data passed to the template
type TemplateData struct {
	ID        string
	Title     string
	Author    string
	Status    string
	CreatedAt string
}
