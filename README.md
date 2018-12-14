### Reading and writing DLIS in go
Spec: http://w3.energistics.org/rp66/v1/Toc/main.html


#### NEXT

Re-write in maplang

Use abuse maps:
    - maps make it clear if something is missing



Think about the run.go Run(*V) *V, how does it work? It takes a *V which is slice of *V and think of them as one feeding the other. Values are simply get imbeded into the next? 


Let code deal with io.Reader, *V deals only with []byte to values

Default value objects are closures, that generate default value. 

Overhaul to keep the io.Reader around. LRS is a temporary transient structure that needs to go away. Need to get to object data as soon as possible.

Start overhaul with SUL.
    
eflr_parse.go - `parseSet()` must build the actual template to follow by the object.

How object would use the Template? How does attrib know it parces Template or object?

Once object data gets parsed just throw it into a chan with large buffer, on the other side collect it for disk persistance asap. Try to stay as wait free as possible.

#### older notes
    
repcode.go - use the funcs from `RepCode` var to build up the template.

reader.go - start with `NewDLISReader()` reading SUL as example. Everything is constructed as simple sequence of `func(in []byte) (Val, int)`. `Val` is universal value type. Calling function must know the expected return type.


### How to read the code
    
#### reader.go

Everything starts with reader.go `NewDLISReader()` which takes an `io.Reader` and returns `dlis.Reader`. Use `ReadVR()` of `dlis.Reader` to get next Visible Record. Then `ReadLRS()` to get next Logical Record Segment. (See example in reader_test.go).


#### IDEA

Construct the reader for each part as sequence of the functions based on the either predefined format as per spec, or construct it at run time based on the data read from the dlis. Such that prior data defines next reader.
    
