//+build ignore

package main
import "fmt"
func main() {
	generate(3)
}
func generate(n int) {
fmt.Print("#include <stdio.h>\n")
fmt.Print("int main(void) {\n")
// Loop to emit multiple printf()s.
for i := 0; i < n; i++ {
fmt.Print("\tprintf(\"This is hello world number ")
fmt.Printf("%v",  i )
fmt.Print(".\\n\");\n")
}
fmt.Print("\treturn 0;\n")
fmt.Print("}\n")
}
