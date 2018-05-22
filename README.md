### Reading and writing DLIS in go
Spec: http://w3.energistics.org/rp66/v1/Toc/main.html


### How to read the code
    
#### reader.go

Everything starts with reader.go `NewDLISReader()` which takes an `io.Reader` and returns `dlis.Reader`. Use `ReadVR()` of `dlis.Reader` to get next Visible Record. Then `ReadLRS()` to get next Logical Record Segment. (See example in reader_test.go).


#### IDEA

Construct the reader for each part as sequence of the functions based on the either predefined format as per spec, or construct it at run time based on the data read from the dlis. Such that prior data defines next reader.

#### NEXT
eflr_parse.go - `parseSet()` must build the actual template to follow by the object.

    