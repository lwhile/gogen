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
	Name    string
	Fields  []Field
	nesting bool
	depth   int
	isArray bool
}

func (st *Struct) firstStr() string {
	return TYPE + SPACE + st.Name + SPACE + STRUCT + LEFTBRACE + BR
}

func (st *Struct) lastStr() string {
	var sp string
	for i := 0; i < (st.depth-1)*4; i++ {
		sp += " "
	}
	return sp + RIGHTBRACE + BR
}

func (st *Struct) spaceStr() string {
	var sp string
	for i := 0; i < st.depth*4; i++ {
		sp += " "
	}
	return sp
}

// Field :
type Field struct {
	Key   string
	Type  interface{}
	array bool
}
