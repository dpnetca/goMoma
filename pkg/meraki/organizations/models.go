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
