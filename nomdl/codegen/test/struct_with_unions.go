// This file was generated by nomdl/codegen.

package test

import (
	"github.com/attic-labs/noms/ref"
	"github.com/attic-labs/noms/types"
)

// This function builds up a Noms value that describes the type
// package implemented by this file and registers it with the global
// type package definition cache.
func __testPackageInFile_struct_with_unions_Ref() types.Ref {
	p := types.PackageDef{
		Types: types.MapOfStringToTypeRefDef{

			"StructWithUnions": __typeRefOfStructWithUnions(),
		},
	}.New()
	return types.Ref{R: types.RegisterPackage(&p)}
}

// StructWithUnions

type StructWithUnions struct {
	m types.Map
}

func NewStructWithUnions() StructWithUnions {
	return StructWithUnions{types.NewMap(
		types.NewString("$name"), types.NewString("StructWithUnions"),
		types.NewString("$type"), types.MakeTypeRef(types.NewString("StructWithUnions"), __testPackageInFile_struct_with_unions_Ref()),
		types.NewString("a"), New__unionOfBOfFloat64AndCOfString().NomsValue(),
		types.NewString("d"), New__unionOfEOfFloat64AndFOfString().NomsValue(),
	)}
}

type StructWithUnionsDef struct {
	A __unionOfBOfFloat64AndCOfStringDef
	D __unionOfEOfFloat64AndFOfStringDef
}

func (def StructWithUnionsDef) New() StructWithUnions {
	return StructWithUnions{
		types.NewMap(
			types.NewString("$name"), types.NewString("StructWithUnions"),
			types.NewString("$type"), types.MakeTypeRef(types.NewString("StructWithUnions"), __testPackageInFile_struct_with_unions_Ref()),
			types.NewString("a"), def.A.New().NomsValue(),
			types.NewString("d"), def.D.New().NomsValue(),
		)}
}

func (self StructWithUnions) Def() StructWithUnionsDef {
	return StructWithUnionsDef{
		__unionOfBOfFloat64AndCOfStringFromVal(self.m.Get(types.NewString("a"))).Def(),
		__unionOfEOfFloat64AndFOfStringFromVal(self.m.Get(types.NewString("d"))).Def(),
	}
}

// Creates and returns a Noms Value that describes StructWithUnions.
func __typeRefOfStructWithUnions() types.TypeRef {
	return types.MakeStructTypeRef(types.NewString("StructWithUnions"),
		types.NewList(
			types.NewString("a"), types.MakeStructTypeRef(types.NewString(""), nil, types.NewList(types.NewString("b"), types.MakePrimitiveTypeRef(types.Float64Kind),
				types.NewString("c"), types.MakePrimitiveTypeRef(types.StringKind))),
			types.NewString("d"), types.MakeStructTypeRef(types.NewString(""), nil, types.NewList(types.NewString("e"), types.MakePrimitiveTypeRef(types.Float64Kind),
				types.NewString("f"), types.MakePrimitiveTypeRef(types.StringKind))),
		),
		nil)

}

func StructWithUnionsFromVal(val types.Value) StructWithUnions {
	// TODO: Validate here
	return StructWithUnions{val.(types.Map)}
}

func (self StructWithUnions) NomsValue() types.Value {
	return self.m
}

func (self StructWithUnions) Equals(other StructWithUnions) bool {
	return self.m.Equals(other.m)
}

func (self StructWithUnions) Ref() ref.Ref {
	return self.m.Ref()
}

func (self StructWithUnions) Type() types.TypeRef {
	return self.m.Get(types.NewString("$type")).(types.TypeRef)
}

func (self StructWithUnions) A() __unionOfBOfFloat64AndCOfString {
	return __unionOfBOfFloat64AndCOfStringFromVal(self.m.Get(types.NewString("a")))
}

func (self StructWithUnions) SetA(val __unionOfBOfFloat64AndCOfString) StructWithUnions {
	return StructWithUnions{self.m.Set(types.NewString("a"), val.NomsValue())}
}

func (self StructWithUnions) D() __unionOfEOfFloat64AndFOfString {
	return __unionOfEOfFloat64AndFOfStringFromVal(self.m.Get(types.NewString("d")))
}

func (self StructWithUnions) SetD(val __unionOfEOfFloat64AndFOfString) StructWithUnions {
	return StructWithUnions{self.m.Set(types.NewString("d"), val.NomsValue())}
}

// __unionOfBOfFloat64AndCOfString

type __unionOfBOfFloat64AndCOfString struct {
	m types.Map
}

func New__unionOfBOfFloat64AndCOfString() __unionOfBOfFloat64AndCOfString {
	return __unionOfBOfFloat64AndCOfString{types.NewMap(
		types.NewString("$name"), types.NewString("__unionOfBOfFloat64AndCOfString"),
		types.NewString("$type"), types.MakeTypeRef(types.NewString("__unionOfBOfFloat64AndCOfString"), __testPackageInFile_struct_with_unions_Ref()),
		types.NewString("$unionIndex"), types.UInt32(0),
		types.NewString("$unionValue"), types.Float64(0),
	)}
}

type __unionOfBOfFloat64AndCOfStringDef struct {
	__unionIndex uint32
	__unionValue interface{}
}

func (def __unionOfBOfFloat64AndCOfStringDef) New() __unionOfBOfFloat64AndCOfString {
	return __unionOfBOfFloat64AndCOfString{
		types.NewMap(
			types.NewString("$name"), types.NewString("__unionOfBOfFloat64AndCOfString"),
			types.NewString("$type"), types.MakeTypeRef(types.NewString("__unionOfBOfFloat64AndCOfString"), __testPackageInFile_struct_with_unions_Ref()),
			types.NewString("$unionIndex"), types.UInt32(def.__unionIndex),
			types.NewString("$unionValue"), def.__unionDefToValue(),
		)}
}

func (self __unionOfBOfFloat64AndCOfString) Def() __unionOfBOfFloat64AndCOfStringDef {
	return __unionOfBOfFloat64AndCOfStringDef{
		uint32(self.m.Get(types.NewString("$unionIndex")).(types.UInt32)),
		self.__unionValueToDef(),
	}
}

func (def __unionOfBOfFloat64AndCOfStringDef) __unionDefToValue() types.Value {
	switch def.__unionIndex {
	case 0:
		return types.Float64(def.__unionValue.(float64))
	case 1:
		return types.NewString(def.__unionValue.(string))
	}
	panic("unreachable")
}

func (self __unionOfBOfFloat64AndCOfString) __unionValueToDef() interface{} {
	switch uint32(self.m.Get(types.NewString("$unionIndex")).(types.UInt32)) {
	case 0:
		return float64(self.m.Get(types.NewString("$unionValue")).(types.Float64))
	case 1:
		return self.m.Get(types.NewString("$unionValue")).(types.String).String()
	}
	panic("unreachable")
}

// Creates and returns a Noms Value that describes __unionOfBOfFloat64AndCOfString.
func __typeRefOf__unionOfBOfFloat64AndCOfString() types.TypeRef {
	return types.MakeStructTypeRef(types.NewString("__unionOfBOfFloat64AndCOfString"),
		types.NewList(),
		types.NewList(
			types.NewString("b"), types.MakePrimitiveTypeRef(types.Float64Kind),
			types.NewString("c"), types.MakePrimitiveTypeRef(types.StringKind),
		))

}

func __unionOfBOfFloat64AndCOfStringFromVal(val types.Value) __unionOfBOfFloat64AndCOfString {
	// TODO: Validate here
	return __unionOfBOfFloat64AndCOfString{val.(types.Map)}
}

func (self __unionOfBOfFloat64AndCOfString) NomsValue() types.Value {
	return self.m
}

func (self __unionOfBOfFloat64AndCOfString) Equals(other __unionOfBOfFloat64AndCOfString) bool {
	return self.m.Equals(other.m)
}

func (self __unionOfBOfFloat64AndCOfString) Ref() ref.Ref {
	return self.m.Ref()
}

func (self __unionOfBOfFloat64AndCOfString) Type() types.TypeRef {
	return self.m.Get(types.NewString("$type")).(types.TypeRef)
}

func (self __unionOfBOfFloat64AndCOfString) B() (val float64, ok bool) {
	if self.m.Get(types.NewString("$unionIndex")).(types.UInt32) != 0 {
		return
	}
	return float64(self.m.Get(types.NewString("$unionValue")).(types.Float64)), true
}

func (self __unionOfBOfFloat64AndCOfString) SetB(val float64) __unionOfBOfFloat64AndCOfString {
	return __unionOfBOfFloat64AndCOfString{self.m.Set(types.NewString("$unionIndex"), types.UInt32(0)).Set(types.NewString("$unionValue"), types.Float64(val))}
}

func (def __unionOfBOfFloat64AndCOfStringDef) B() (val float64, ok bool) {
	if def.__unionIndex != 0 {
		return
	}
	return def.__unionValue.(float64), true
}

func (def __unionOfBOfFloat64AndCOfStringDef) SetB(val float64) __unionOfBOfFloat64AndCOfStringDef {
	def.__unionIndex = 0
	def.__unionValue = val
	return def
}

func (self __unionOfBOfFloat64AndCOfString) C() (val string, ok bool) {
	if self.m.Get(types.NewString("$unionIndex")).(types.UInt32) != 1 {
		return
	}
	return self.m.Get(types.NewString("$unionValue")).(types.String).String(), true
}

func (self __unionOfBOfFloat64AndCOfString) SetC(val string) __unionOfBOfFloat64AndCOfString {
	return __unionOfBOfFloat64AndCOfString{self.m.Set(types.NewString("$unionIndex"), types.UInt32(1)).Set(types.NewString("$unionValue"), types.NewString(val))}
}

func (def __unionOfBOfFloat64AndCOfStringDef) C() (val string, ok bool) {
	if def.__unionIndex != 1 {
		return
	}
	return def.__unionValue.(string), true
}

func (def __unionOfBOfFloat64AndCOfStringDef) SetC(val string) __unionOfBOfFloat64AndCOfStringDef {
	def.__unionIndex = 1
	def.__unionValue = val
	return def
}

// __unionOfEOfFloat64AndFOfString

type __unionOfEOfFloat64AndFOfString struct {
	m types.Map
}

func New__unionOfEOfFloat64AndFOfString() __unionOfEOfFloat64AndFOfString {
	return __unionOfEOfFloat64AndFOfString{types.NewMap(
		types.NewString("$name"), types.NewString("__unionOfEOfFloat64AndFOfString"),
		types.NewString("$type"), types.MakeTypeRef(types.NewString("__unionOfEOfFloat64AndFOfString"), __testPackageInFile_struct_with_unions_Ref()),
		types.NewString("$unionIndex"), types.UInt32(0),
		types.NewString("$unionValue"), types.Float64(0),
	)}
}

type __unionOfEOfFloat64AndFOfStringDef struct {
	__unionIndex uint32
	__unionValue interface{}
}

func (def __unionOfEOfFloat64AndFOfStringDef) New() __unionOfEOfFloat64AndFOfString {
	return __unionOfEOfFloat64AndFOfString{
		types.NewMap(
			types.NewString("$name"), types.NewString("__unionOfEOfFloat64AndFOfString"),
			types.NewString("$type"), types.MakeTypeRef(types.NewString("__unionOfEOfFloat64AndFOfString"), __testPackageInFile_struct_with_unions_Ref()),
			types.NewString("$unionIndex"), types.UInt32(def.__unionIndex),
			types.NewString("$unionValue"), def.__unionDefToValue(),
		)}
}

func (self __unionOfEOfFloat64AndFOfString) Def() __unionOfEOfFloat64AndFOfStringDef {
	return __unionOfEOfFloat64AndFOfStringDef{
		uint32(self.m.Get(types.NewString("$unionIndex")).(types.UInt32)),
		self.__unionValueToDef(),
	}
}

func (def __unionOfEOfFloat64AndFOfStringDef) __unionDefToValue() types.Value {
	switch def.__unionIndex {
	case 0:
		return types.Float64(def.__unionValue.(float64))
	case 1:
		return types.NewString(def.__unionValue.(string))
	}
	panic("unreachable")
}

func (self __unionOfEOfFloat64AndFOfString) __unionValueToDef() interface{} {
	switch uint32(self.m.Get(types.NewString("$unionIndex")).(types.UInt32)) {
	case 0:
		return float64(self.m.Get(types.NewString("$unionValue")).(types.Float64))
	case 1:
		return self.m.Get(types.NewString("$unionValue")).(types.String).String()
	}
	panic("unreachable")
}

// Creates and returns a Noms Value that describes __unionOfEOfFloat64AndFOfString.
func __typeRefOf__unionOfEOfFloat64AndFOfString() types.TypeRef {
	return types.MakeStructTypeRef(types.NewString("__unionOfEOfFloat64AndFOfString"),
		types.NewList(),
		types.NewList(
			types.NewString("e"), types.MakePrimitiveTypeRef(types.Float64Kind),
			types.NewString("f"), types.MakePrimitiveTypeRef(types.StringKind),
		))

}

func __unionOfEOfFloat64AndFOfStringFromVal(val types.Value) __unionOfEOfFloat64AndFOfString {
	// TODO: Validate here
	return __unionOfEOfFloat64AndFOfString{val.(types.Map)}
}

func (self __unionOfEOfFloat64AndFOfString) NomsValue() types.Value {
	return self.m
}

func (self __unionOfEOfFloat64AndFOfString) Equals(other __unionOfEOfFloat64AndFOfString) bool {
	return self.m.Equals(other.m)
}

func (self __unionOfEOfFloat64AndFOfString) Ref() ref.Ref {
	return self.m.Ref()
}

func (self __unionOfEOfFloat64AndFOfString) Type() types.TypeRef {
	return self.m.Get(types.NewString("$type")).(types.TypeRef)
}

func (self __unionOfEOfFloat64AndFOfString) E() (val float64, ok bool) {
	if self.m.Get(types.NewString("$unionIndex")).(types.UInt32) != 0 {
		return
	}
	return float64(self.m.Get(types.NewString("$unionValue")).(types.Float64)), true
}

func (self __unionOfEOfFloat64AndFOfString) SetE(val float64) __unionOfEOfFloat64AndFOfString {
	return __unionOfEOfFloat64AndFOfString{self.m.Set(types.NewString("$unionIndex"), types.UInt32(0)).Set(types.NewString("$unionValue"), types.Float64(val))}
}

func (def __unionOfEOfFloat64AndFOfStringDef) E() (val float64, ok bool) {
	if def.__unionIndex != 0 {
		return
	}
	return def.__unionValue.(float64), true
}

func (def __unionOfEOfFloat64AndFOfStringDef) SetE(val float64) __unionOfEOfFloat64AndFOfStringDef {
	def.__unionIndex = 0
	def.__unionValue = val
	return def
}

func (self __unionOfEOfFloat64AndFOfString) F() (val string, ok bool) {
	if self.m.Get(types.NewString("$unionIndex")).(types.UInt32) != 1 {
		return
	}
	return self.m.Get(types.NewString("$unionValue")).(types.String).String(), true
}

func (self __unionOfEOfFloat64AndFOfString) SetF(val string) __unionOfEOfFloat64AndFOfString {
	return __unionOfEOfFloat64AndFOfString{self.m.Set(types.NewString("$unionIndex"), types.UInt32(1)).Set(types.NewString("$unionValue"), types.NewString(val))}
}

func (def __unionOfEOfFloat64AndFOfStringDef) F() (val string, ok bool) {
	if def.__unionIndex != 1 {
		return
	}
	return def.__unionValue.(string), true
}

func (def __unionOfEOfFloat64AndFOfStringDef) SetF(val string) __unionOfEOfFloat64AndFOfStringDef {
	def.__unionIndex = 1
	def.__unionValue = val
	return def
}
