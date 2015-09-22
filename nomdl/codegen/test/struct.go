// This file was generated by nomdl/codegen.

package test

import (
	"github.com/attic-labs/noms/ref"
	"github.com/attic-labs/noms/types"
)

// This function builds up a Noms value that describes the type
// package implemented by this file and registers it with the global
// type package definition cache.
func __testPackageInFile_struct_Ref() types.Ref {
	p := types.PackageDef{
		Types: types.MapOfStringToTypeRefDef{

			"Struct": __typeRefOfStruct(),
		},
	}.New()
	return types.Ref{R: types.RegisterPackage(&p)}
}

// Struct

type Struct struct {
	m types.Map
}

func NewStruct() Struct {
	return Struct{types.NewMap(
		types.NewString("$name"), types.NewString("Struct"),
		types.NewString("$type"), types.MakeTypeRef(types.NewString("Struct"), __testPackageInFile_struct_Ref()),
		types.NewString("s"), types.NewString(""),
		types.NewString("b"), types.Bool(false),
	)}
}

type StructDef struct {
	S string
	B bool
}

func (def StructDef) New() Struct {
	return Struct{
		types.NewMap(
			types.NewString("$name"), types.NewString("Struct"),
			types.NewString("$type"), types.MakeTypeRef(types.NewString("Struct"), __testPackageInFile_struct_Ref()),
			types.NewString("s"), types.NewString(def.S),
			types.NewString("b"), types.Bool(def.B),
		)}
}

func (self Struct) Def() StructDef {
	return StructDef{
		self.m.Get(types.NewString("s")).(types.String).String(),
		bool(self.m.Get(types.NewString("b")).(types.Bool)),
	}
}

// Creates and returns a Noms Value that describes Struct.
func __typeRefOfStruct() types.TypeRef {
	return types.MakeStructTypeRef(types.NewString("Struct"),
		types.NewList(
			types.NewString("s"), types.MakePrimitiveTypeRef(types.StringKind),
			types.NewString("b"), types.MakePrimitiveTypeRef(types.BoolKind),
		),
		nil)

}

func StructFromVal(val types.Value) Struct {
	// TODO: Validate here
	return Struct{val.(types.Map)}
}

func (self Struct) NomsValue() types.Value {
	return self.m
}

func (self Struct) Equals(other Struct) bool {
	return self.m.Equals(other.m)
}

func (self Struct) Ref() ref.Ref {
	return self.m.Ref()
}

func (self Struct) Type() types.TypeRef {
	return self.m.Get(types.NewString("$type")).(types.TypeRef)
}

func (self Struct) S() string {
	return self.m.Get(types.NewString("s")).(types.String).String()
}

func (self Struct) SetS(val string) Struct {
	return Struct{self.m.Set(types.NewString("s"), types.NewString(val))}
}

func (self Struct) B() bool {
	return bool(self.m.Get(types.NewString("b")).(types.Bool))
}

func (self Struct) SetB(val bool) Struct {
	return Struct{self.m.Set(types.NewString("b"), types.Bool(val))}
}
