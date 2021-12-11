# GoLite Project - Go Huskies!

## MileStone 1 - Scanner

Testing for Scanner:
1. The tested files should be saved in the directory of `proj/golite/scanner`.
2. Usage of `-lex`: Run the command below from the directory `../proj-gohuskies/proj/golite/`,
   `go run golite.go -lex scanner/test1.golite`

Expected test result:
The expected output is displayed sequentially in the format of `{token_type}({line_number})`

Example Output:
```
Token.PACK(1)
Token.ID(1)       
Token.SEMICOLON(1)
Token.COMMENT(1)  
Token.IMPORT(3)   
Token.QTDMARK(3)  
Token.FMT(3)      
Token.QTDMARK(3)  
Token.SEMICOLON(3)
Token.ID(4)       
Token.ASSIGN(4)   
Token.NUM(4)      
Token.SEMICOLON(4)
Token.EOF(4)
```


## MileStone 2 - Parser and Semantic Analysis

Testing for Parser:
1. The tested files should be saved in the directory of `proj/golite/parser`.
2. Usage of `-ast`: Run the command below from the directory `../proj-gohuskies/proj/golite/`,
   `go run golite.go -ast parser/test1_parser.golite`

Expected test result:
The displayed output is expected to be the same as the content of the test file.

Example Output:
```
package main;
import "fmt";
func main () {
var a int;
a = 1+1;

}
```

## MileStone 3 - ILOC

Testing for ILOC:
1. The tested files should be saved in the directory of `proj/golite/iloc`.
2. Usage of `-iloc`: Run the command below from the directory `../proj-gohuskies/proj/golite/`,
   `go run golite.go -iloc iloc/test3_iloc.golite`


Expected test result:
The expected output is displayed sequentially of the ILOC instructions of the source code in test file.

Example Output:
```
main:
    mov r2,#3
    mov r1,r2
    mov r3,#6
    add r4,r1,r3
    mov r0,r4
    print r0
    b condLabel_L1
loopBody_L2:
    mov r5,#1
    sub r6,r1,r5
    mov r1,r6
    print r1
condLabel_L1:
    mov r7,#0
    mov r8,#0
    cmp r1,r7
    movgt r8,#1
    cmp r8,#1
    beq loopBody_L2
```

## MileStone 4 (Final Submission) - Assembly

Usage of `-S`:
1. go to directory `../proj-gohuskies/proj/golite/`
2. Example: `go run golite.go -S .\arm\test1_arm.golite`
3. check the directory of `../proj-gohuskies/proj/golite`, the output file should be in the same folder as `golite.go`

Example Output:

```
	.arch armv8-a
	.comm d,8,8
	.comm e,8,8
	.text
	.type main,%function
	.global main
	.p2align		2
main:
	sub sp,sp,16
	stp x29,x30,[sp]
	mov x29,sp
	sub sp,sp,#48
	mov x1,#7
	str x1,[x29,#-8]
	ldr x2,[x29,#-8]
	mov x1,x2
	str x1,[x29,#-16]
	mov x1,#3
	str x1,[x29,#-24]
	ldr x2,[x29,#-24]
	mov x1,x2
	str x1,[x29,#-32]
	ldr x1,[x29,#-16]
	ldr x2,[x29,#-32]
	add x3,x1,x2
	str x3,[x29,#-40]
	ldr x2,[x29,#-40]
	mov x1,x2
	str x1,[x29,#-48]
	add sp,sp,#48
	ldp x29,x30,[sp]
	add sp,sp,16
	ret
	.size main,(.-main)
```
