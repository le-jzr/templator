templator
=========

Template-based code generator generator.

The purpose of this program is to make it easier to write programs that generate programs.
To achieve this, templator generates code from a template.

The "syntax" of the templates is exceedingly simple.
Lines commented out using C++-style line comments are "logic".
They get copied verbatim to the generated code.
The logic is currently assumed to be written in Go.
Lines not commented out are the output of the generated generator.
Uncommented lines can contain snippets of code encased in C-style multi-line
comments. The snippets are assumed to be Go expressions, whose value gets inserted
into the generated generator's output.

Completely empty lines are ignored by templator.

Template example
----------------

Suppose you want to generate C code that writes out a given number of hello worlds.
You can do this as follows:

	// func generate(n int) {
	
	#include <stdio.h>
	
	int main(void) {
		//// Loop to emit multiple printf()s.
		// for i := 0; i < n; i++ {
		printf("This is hello world number /* i */.\n");
		// }
		return 0;
	}
	
	// }

Output of templator:

	func generate(n int) {
	fmt.Printf("#include <stdio.h>\n")
	fmt.Printf("int main(void) {\n")
	// Loop to emit multiple printf()s.
	for i := 0; i < n; i++ {
	fmt.Printf("	printf(\"This is hello world number %v.\\n\");\n", i)
	}
	fmt.Printf("return 0;\n")
	fmt.Printf("}\n")
	}

Output of calling generate(3):

	#include <stdio.h>
	int main(void) {
		printf("This is hello world number 0.\n");
		printf("This is hello world number 1.\n");
		printf("This is hello world number 2.\n");
		return 0;
	}

Number of features to notice:
 - Replacement is textual, snippets within strings or other lexical elements still count.
 - In ordinary code editors, you get a proper highlighting for the final output.
 - You can make actual comments that get inserted into the generator code.
