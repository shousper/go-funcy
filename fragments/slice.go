package fragments

import (
	"github.com/shousper/go-funcy/model"

	. "github.com/dave/jennifer/jen"
)

// SliceOf produces a slice type for the given types
//
//	type SliceOfT []T
//
func SliceOf(f *File, m model.Model) {
	f.Add(Commentf("%s is a string slice of %s", m.Slice.Name, m.Type))
	f.Type().Id(m.Slice.Name).
		Index().Add(asType(m))
}

// SliceOfAsMap produces a map from a slice grouping by its defined unique ID
//
//	func (s SliceOfT) AsMap() MapOfT {
//		result := make(MapOfT)
//		for _, value := range s {
//			result[value.K] = append(result[value.K], value)
//		}
//		return result
//	}
//
func SliceOfAsMap(f *File, m model.Model) {
	outType := Map(toType(*m.Map.Key)).Add(asType(m))

	f.Commentf("AsMap maps slice values by %s", m.Map.Key.Accessor)
	f.Func().
		Params(Id("s").Id(m.Slice.Name)).
		Id("AsMap").
		Params().
		Add(outType).
		Block(
			Id("result").Op(":=").Make(outType),
			For(List(Id("_"), Id("value")).Op(":=").Range().Id("s")).Block(
				Id("result").Index(Id("value").Dot(m.Map.Key.Accessor)).Op("=").Id("value"),
			),
			Return(Id("result")),
		)
}

// SliceOfGroupBys produce a maps of slice values grouped by a field value
//
//	func (s SliceOfT) GroupByF() map[J]T {
//		result := make(map[J]T)
//		for _, value := range s {
//			result[value.F] = append(result[value.F], value)
//		}
//		return result
//	}
//
func SliceOfGroupBys(f *File, m model.Model) {
	for _, g := range m.GroupBys {
		fnName := "GroupBy" + g.Name
		outType := Map(toType(g)).Index().Add(asType(m))

		mappedField := Id("result").Index(Id("value").Dot(g.Accessor))

		f.Commentf("%s groups slice values by %s", fnName, g.Accessor)
		f.Func().
			Params(Id("s").Id(m.Slice.Name)).
			Id(fnName).
			Params().
			Add(outType).
			Block(
				Id("result").Op(":=").Make(outType),
				For(List(Id("_"), Id("value")).Op(":=").Range().Id("s")).Block(
					Add(mappedField).Op("=").Append(mappedField, Id("value")),
				),
				Return(Id("result")),
			)
	}
}
