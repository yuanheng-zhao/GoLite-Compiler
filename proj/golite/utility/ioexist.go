package utility

var printExist, printlnExist, scanExist bool

func IOInit() {
	printExist = false
	printlnExist = false
	scanExist = false
}

func SetPrint() {
	printExist = true
}

func GetPrint() bool {
	return printExist
}

func SetPrintln() {
	printlnExist = true
}

func GetPrintln() bool {
	return printlnExist
}

func SetScan() {
	scanExist = true
}

func GetScan() bool {
	return scanExist
}