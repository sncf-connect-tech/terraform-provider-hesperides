package hesperides

type hesperidesPlatform struct {
	ApplicationName    string   `json:"application_name"`
	PlatformName       string   `json:"platform_name"`
	ApplicationVersion string   `json:"application_version"`
	Production         bool     `json:"production"`
	VersionId          int      `json:"version_id"`
	Modules            []string `json:"modules"`
}

type hesperidesModule struct {
	Name        string             `json:"name"`
	Version     string             `json:"version"`
	WorkingCopy bool               `json:"working_copy"`
	Technos     []hesperidesTechno `json:"technos"`
	VersionId   int                `json:"version_id"`
}

type hesperidesTechno struct {
	Name        string             `json:"name"`
	Version     string             `json:"version"`
	WorkingCopy bool               `json:"working_copy"`
	Template    hesperidesTemplate `json:"template"`
}

type hesperidesTemplate struct {
	Name      string                   `json:"name"`
	Namespace string                   `json:"namespace"`
	Filename  string                   `json:"filename"`
	Location  string                   `json:"location"`
	Content   string                   `json:"content"`
	Rights    hesperidesTemplateRights `json:"rights"`
	VersionId int                      `json:"version_id"`
}

type hesperidesTemplateRights struct {
	User  hesperidesTemplateFileRights `json:"user"`
	Group hesperidesTemplateFileRights `json:"group"`
	Other hesperidesTemplateFileRights `json:"other"`
}

type hesperidesTemplateFileRights struct {
	Read    bool `json:"read"`
	Write   bool `json:"write"`
	Execute bool `json:"execute"`
}
