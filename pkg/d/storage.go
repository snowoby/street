package d

import "street/ent"

type FileForm struct {
	Filename string `json:"filename"`
	Mime     string `json:"mime" binding:"required"`
	Size     int    `json:"size" binding:"required"`
	Category string `json:"category" binding:"required"`
}

type File struct {
	*ent.File
	NoEdges
	ValueType
}

type Part struct {
	Part int `uri:"part" binding:"required"`
}

func FileFromEnt(f *ent.File) *File {
	return &File{
		File:      f,
		ValueType: ValueType{"file"},
	}
}
