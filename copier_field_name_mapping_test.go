package copier_test

import (
	"reflect"
	"testing"

	"goyave.dev/copier"
)

func TestCustomFieldName(t *testing.T) {
	type User1 struct {
		ID      int64
		Name    string
		Address []string
	}

	type User2 struct {
		ID2      int64
		Name2    string
		Address2 []string
	}

	u1 := User1{ID: 1, Name: "1", Address: []string{"1"}}
	var u2 User2
	err := copier.CopyWithOption(&u2, u1, copier.Option{FieldNameMapping: []copier.FieldNameMapping{
		{SrcType: u1, DstType: u2,
			Mapping: map[string]string{
				"ID":      "ID2",
				"Name":    "Name2",
				"Address": "Address2"}},
	}})

	if err != nil {
		t.Fatal(err)
	}

	if u1.ID != u2.ID2 {
		t.Error("copy id failed.")
	}

	if u1.Name != u2.Name2 {
		t.Error("copy name failed.")
	}

	if !reflect.DeepEqual(u1.Address, u2.Address2) {
		t.Error("copy address failed.")
	}
}
