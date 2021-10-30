# proj

## Component - Scanner

Testing for Scanner:
1. The tested files should be saved in the directory of proj/golite/scanner. 
2. Run the command below from the directory golite, 
   go run golite.go -lex scanner/test1.golite

Expected test result: 
The expected output is displayed sequentially in the format of ```{token_type}({line_number})```

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
