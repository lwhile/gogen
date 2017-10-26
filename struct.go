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
	Key  string
	Type interface{}
}

// // Type :
// type Type struct {
// 	nesting bool
// 	Name    interface{}
// }

// // TypeStr return string form of f.Type
// func (f *Field) TypeStr() string {
// 	switch f.
// }
