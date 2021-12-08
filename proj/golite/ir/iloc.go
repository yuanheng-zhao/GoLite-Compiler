package ir

type OperandTy int

const (
	REGISTER OperandTy = iota
	IMMEDIATE
	ONEOPERAND
	GLOBALVAR
	VOID
)

type ApsrFlag int

const (
	GT ApsrFlag = iota
	LT
	GE
	LE
	EQ
	NE
	AL
)

type Instruction interface {
	GetTargets() []int // Get the registers targeted by this instruction

	GetSources() []int // Get the source registers for this instruction

	GetImmediate() *int // Get the immediate value (i.e., constant) of this instruction

	GetSourceString() string // Get the string component (global variable name) of this instruction

	GetLabel() string // Get the label for this instruction

	SetLabel(newLabel string) //Set the label for this instruction

	String() string // Return a string representation of this instruction

	TranslateToAssembly(map[int]int) []string
}

type FuncFrag struct {
	Label string        // Function name
	Body  []Instruction // Function body of ILOC instructions
	//Frame   *codegen.Frame	 // Activation Records (i.e., stack frame) for this function
}
