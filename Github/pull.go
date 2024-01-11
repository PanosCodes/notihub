package Github

type Pull struct {
	Title  string
	Author User `json:"user"`
}
