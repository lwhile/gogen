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
}

// Serialize will make the Field formated to []byte
// func (st *Struct) Serialize() (b []byte, err error) {
// 	if st == nil {
// 		err = fmt.Errorf("error to serialize a nil struct to []byte")
// 		log.Error(err)
// 		return
// 	}

// 	var bf bytes.Buffer
// 	if !st.nesting {
// 		if _, err = bf.Write(st.firstStr()); err != nil {
// 			return
// 		}
// 	}

// 	for _, f := range st.Fields {
// 		if _, err = bf.Write([]byte(FOURSPACE + f.Key + SPACE)); err != nil {
// 			return
// 		}
// 		if m, ok := isMap(f.Type); ok {

// 		}
// 	}
// 	return
// }

func (st *Struct) firstStr() []byte {
	return []byte(TYPE + SPACE + st.Name + SPACE + STRUCT + LEFTBRACE + BR)
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
