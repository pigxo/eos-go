package database

import (
	"fmt"
	"reflect"
)

type Carnivore struct {
	Lion  int
	Tiger int
}

type House struct {
	Id       uint64 `multiIndex:"id,increment"`
	Area     uint64 `multiIndex:"orderedNonUnique"`
	Name     string
	Carnivore Carnivore `multiIndex:"orderedNonUnique"`
}

type IdType uint64
type Name uint64
type AccountName uint64
type PermissionName uint64
type ActionName uint64
type TableName uint64
type ScopeName uint64


type TableIdObject struct{
	ID 		IdType			`multiIndex:"id,increment,byScope"`
	Code 	AccountName		`multiIndex:"orderedNonUnique:less"`
	Scope 	ScopeName		`multiIndex:"orderedNonUnique,byTable:greater,byScope:less"`
	Table 	TableName		`multiIndex:"orderedNonUnique,byTable:less"`
	Payer 	AccountName		`multiIndex:"orderedNonUnique,byScope"`
	Count 	uint32
}

func logObj(data interface{}){
	space := "	"
	ref := reflect.ValueOf(data)
	if !ref.IsValid() ||  reflect.Indirect(ref).Kind() != reflect.Struct {
		fmt.Println("log obj valid")
		return
	}

	s := &ref
	if s.Kind() == reflect.Ptr {
		e := s.Elem()
		s = &e
	}
	if s.Kind() != reflect.Struct {
		fmt.Println("log obj valid")
		return
	}
	typ := s.Type()


	num := s.NumField()
	for i := 0; i  < num; i++{
		v := s.Field(i)
		t := typ.Field(i)
		fmt.Print(t.Name,space,v,space)
	}
	fmt.Println("")
}
