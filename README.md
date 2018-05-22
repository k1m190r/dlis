### Reading and writing DLIS in go
Spec: http://w3.energistics.org/rp66/v1/Toc/main.html


### How to read the code
    
#### reader.go

Everything starts with reader.go `NewDLISReader()` which takes an `io.Reader` and returns `dlis.Reader`. 