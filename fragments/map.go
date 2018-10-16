package fragments

import (
	"github.com/shousper/go-funcy/model"

	. "github.com/dave/jennifer/jen"
)

// MapOf produces a map type for the given types
//
//	type MapOfT map[K]T
//
func MapOf(f *File, m model.Model) {
	f.Add(Commentf("%s is a string map of %s", m.Map.Name, m.Type))
	f.Type().Id(m.Map.Name).
		Map(Id(m.Map.Key.Type)).Add(asType(m))
}

// MapOfKeys produces a slice of the keys of the map
//
//	func (m MapOfT) Keys() (keys []K) {
//		for key := range m {
//			keys = append(keys, key)
//		}
//		return
//	}
//
func MapOfKeys(f *File, m model.Model) {
	f.Add(Comment("Keys returns a slice of the map keys"))
	f.Func().
		Params(Id("m").Id(m.Map.Name)).
		Id("Keys").
		Params().
		Params(Id("keys").Index().Id(m.Map.Key.Type)).
		Block(
			For(Id("key").Op(":=").Range().Id("m")).Block(
				Id("keys").Op("=").Append(Id("keys"), Id("key")),
			),
			Return(),
		)
}

// MapOfValues produces a slice of the values of the map
//
//	func (m MapOfT) Values() (values []*T) {
//		for _, value := range m {
//			values = append(values, value)
//		}
//		return
//	}
//
func MapOfValues(f *File, m model.Model) {
	f.Add(Comment("Values returns a slice of the map values"))
	f.Func().
		Params(Id("m").Id(m.Map.Name)).
		Id("Values").
		Params().
		Params(Id("values").Index().Add(asType(m))).
		Block(
			For(List(Id("_"), Id("value")).Op(":=").Range().Id("m")).Block(
				Id("values").Op("=").Append(Id("values"), Id("value")),
			),
			Return(),
		)
}

// MapOfSelect produces a map with only the specified keys form the map
//
//	func (m MapOfT) Select(keys ...K) MapOfT {
//		result := make(MapOfT)
//		for _, key := range keys {
//			if value, exists := m[key]; exists {
//				result[key] = value
//			}
//		}
//		return result
//	}
//
func MapOfSelect(f *File, m model.Model) {
	f.Add(Comment("Select returns a map with only the given keys"))
	f.Func().
		Params(Id("m").Id(m.Map.Name)).
		Id("Select").
		Params(Id("keys").Op("...").Id(m.Map.Key.Type)).
		Id(m.Map.Name).
		Block(
			Id("result").Op(":=").Make(Id(m.Map.Name)),
			For(List(Id("_"), Id("key")).Op(":=").Range().Id("keys")).Block(
				If(
					List(Id("value"), Id("exists")).Op(":=").Id("m").Index(Id("key")),
					Id("exists"),
				).Block(
					Id("result").Index(Id("key")).Op("=").Id("value"),
				),
			),
			Return(Id("result")),
		)
}

// MapOfGroupBys produces maps regrouped by a field value
//
//	func (m MapOfT) GroupByF() map[J]T {
//		result := make(map[J]T)
//		for _, value := range m {
//			result[value.F] = append(result[value.F], value)
//		}
//		return result
//	}
//
func MapOfGroupBys(f *File, m model.Model) {
	for _, g := range m.GroupBys {
		fnName := "GroupBy" + g.Name
		outType := Map(toType(g)).Index().Add(asType(m))

		mappedField := Id("result").Index(Id("value").Dot(g.Accessor))

		f.Commentf("%s groups map values by %s", fnName, g.Name)
		f.Func().
			Params(Id("m").Id(m.Map.Name)).
			Id(fnName).
			Params().
			Add(outType).
			Block(
				Id("result").Op(":=").Make(outType),
				For(List(Id("_"), Id("value")).Op(":=").Range().Id("m")).Block(
					Add(mappedField).Op("=").Append(mappedField, Id("value")),
				),
				Return(Id("result")),
			)
	}
}
