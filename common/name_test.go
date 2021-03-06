package common

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewName(t *testing.T) {
	name := "eosio.system"
	val := N(name)
	assert.Equal(t, val, uint64(6138663591228101920))
	//fmt.Printf("%d\n", val)
	name2 := S(6138663591228101920)
	//fmt.Println(name2)
	assert.Equal(t, name2, name)
}

func TestNameStr(t *testing.T) {
	name := "eosio.systemabdxs"
	val := N(name)
	fmt.Printf("%d\n", val)

}

func TestNameSuffix(t *testing.T) {
	name := N("eosio.token")
	check := N("token")
	test := NameSuffix(name)
	//fmt.Println(N(test))
	assert.Equal(t, test, check)
}
