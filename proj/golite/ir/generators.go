package ir

import "fmt"

type registerGen struct{
	count int //The current counter for the label
}


func  NewRegister() int {
	retVal := rGen.count
	rGen.count += 1
	return retVal
}

type labelGen struct {
	count int //The current label number

}
func NewLabelWithPre(prefix string) string {
	retVal := fmt.Sprintf("%s_L%d",prefix,lGen.count)
	lGen.count += 1
	return retVal
}

func NewLabel() string {
	retVal := fmt.Sprintf("L%d",lGen.count)
	lGen.count += 1
	return retVal
}

var rGen  *registerGen
var lGen  *labelGen

// The init() function will only be called once per package. This is where you can setup singletons for types
func init() {
	rGen = &registerGen{0}
	lGen = &labelGen{0}
}