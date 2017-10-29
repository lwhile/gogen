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

type Test struct {
	key1 string
	key2 string
	key3 struct {
		key3A string
		key3B struct {
			key3BA string
		}
	}
}
