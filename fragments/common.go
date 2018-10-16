package fragments

import (
	. "github.com/dave/jennifer/jen"
	"github.com/shousper/go-funcy/model"
)

func asType(m model.Model) Code {
	if m.IsInterface {
		return Id(m.Type)
	}
	return Op("*").Id(m.Type)
}

func toType(f model.FieldModel) Code {
	s := &Statement{}
	if f.Pointer {
		s = Op("*")
	}
	if f.Qual != "" {
		return s.Qual(f.Qual, f.Type)
	}
	return s.Id(f.Type)
}
