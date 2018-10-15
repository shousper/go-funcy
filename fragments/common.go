package fragments

import (
	"github.com/shousper/go-funcy/model"

	. "github.com/dave/jennifer/jen"
)

func asType(m model.Model) Code {
	if m.IsInterface {
		return Id(m.Type)
	} else {
		return Op("*").Id(m.Type)
	}
}
