package gogen

const (
	// TYPE :
	TYPE = "type"
	// STRUCT ;
	STRUCT = "struct"
	// SPACE :
	SPACE = " "
	// FOURSPACE :
	FOURSPACE = "    "

	//LEFTBRACE :
	LEFTBRACE = "{"
	// RIGHTBRACE :
	RIGHTBRACE = "}"

	// BR :
	BR = "\n"
)

// Struct :
type Struct struct {
	Name   string
	Fields []Field
}

// Field :
type Field struct {
	Name string
	Type string
	Tag  string
}
