package gogen

type T struct {
	A struct {
		A1 struct{}
	} `json:"A"`
	B struct {
		B1 struct{} `json:"B1"`
		B2 struct{}
	} `json:"B"`
	C int
	D string
}
