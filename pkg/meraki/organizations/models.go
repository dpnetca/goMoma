package organizations

type api struct {
	Enabled bool
}

type licensing struct {
	Model string
}
type Organization struct {
	Id        string
	Name      string
	Url       string
	Api       api
	Licensing licensing
}

type Admin struct {
	Id        string `json:"id,omitempty"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	OrgAccess string `json:"orgAccess"`
}

type AdminResponse struct {
	Success      bool
	StatusCode   int
	ErrorMessage []string `json:"errors"`
	Admin        Admin
}
