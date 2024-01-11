package Github

type Pull struct {
	ID     int
	Url    string
	Title  string
	Author User `json:"user"`
}
