package intermediate

import "fmt"

func main() {

	// --- General Formatting Verbs
	// %v   Prints the value in the default format
	// %#v  Prints the value om Go-syntax format
	// %T   Prints the type of the value
	// %%   Prints the '%' sign

	simpleFloat := 11_365.5
	simpleString := "Hello world!"

	fmt.Printf("simpleFloat: %v\n", simpleFloat)   // simpleFloat: 11365.5
	fmt.Printf("simpleFloat: %#v\n", simpleFloat)  // simpleFloat: 11365.5
	fmt.Printf("simpleFloat: %T\n", simpleFloat)   // simpleFloat: float64
	fmt.Printf("simpleFloat: %v%%\n", simpleFloat) // simpleFloat: 11365.5%

	fmt.Printf("simpleString: %v\n", simpleString)  // simpleString: Hello world!
	fmt.Printf("simpleString: %#v\n", simpleString) // simpleString: "Hello world!"
	fmt.Printf("simpleString: %T\n", simpleString)  // simpleString: string

	// --- Integer Formatting Verbs
	// %b    Base 2
	// %d    Base 10
	// %+d   Base 10 and always show sign
	// %o    Base 8
	// %O    Base 8, with leading 0o
	// %x    Base 16, lowercase
	// %X    Base 16, uppercase
	// %#x   Base 16, with leading 0x
	// %4d   Pad with spaces (width 4, right justified)
	// %-4d  Pad with spaces (width 4, left  justified)
	// %04d  Pad with zeroes (width 4)

	simpleInteger := 255

	fmt.Printf("\nsimpleInteger: %b\n", simpleInteger) // simpleInteger: 11111111
	fmt.Printf("simpleInteger: %d\n", simpleInteger)   // simpleInteger: 255
	fmt.Printf("simpleInteger: %+d\n", simpleInteger)  // simpleInteger: +255
	fmt.Printf("simpleInteger: %o\n", simpleInteger)   // simpleInteger: 377
	fmt.Printf("simpleInteger: %O\n", simpleInteger)   // simpleInteger: 0o377
	fmt.Printf("simpleInteger: %x\n", simpleInteger)   // simpleInteger: ff
	fmt.Printf("simpleInteger: %X\n", simpleInteger)   // simpleInteger: FF
	fmt.Printf("simpleInteger: %#X\n", simpleInteger)  // simpleInteger: 0XFF
	fmt.Printf("simpleInteger: %4d\n", simpleInteger)  // simpleInteger:  255
	fmt.Printf("simpleInteger: %-4d\n", simpleInteger) // simpleInteger: 255
	fmt.Printf("simpleInteger: %04d\n", simpleInteger) // simpleInteger: 0255

	// --- Strings Formatting Verbs
	// %s    Prints the value as plain string
	// %q    Prints the value as a double-quoted strings
	// %8s   Prints the value as plain string (width 8, right justified)
	// %-8s  Prints the value as plain string (width 8, left  justified)
	// %x    Prints the value as hez dump of byte values
	// % x   Prints the value as hez dump with spaces

	anotherString := "A cool blue gopher"

	fmt.Printf("\nanotherString: %s\n", anotherString) // A cool blue gopher
	fmt.Printf("anotherString: %q\n", anotherString)   // "A cool blue gopher"
	fmt.Printf("anotherString: %8s\n", anotherString)  // A cool blue gopher
	fmt.Printf("anotherString: %-8s\n", anotherString) // A cool blue gopher
	fmt.Printf("anotherString: %x\n", anotherString)   // 4120636f6f6c20626c756520676f70686572
	fmt.Printf("anotherString: % x\n", anotherString)  // 41 20 63 6f 6f 6c 20 62 6c 75 65 20 67 6f 70 68 65 72

	// --- Boolean Formatting Verbs
	// %t   Value of the boolean operator in true or false format (same as using %v)

	simpleTrue := true
	simpleFalse := false

	fmt.Printf("\nsimpleTrue: %t\n", simpleTrue) // simpleTrue: true
	fmt.Printf("simpleFalse: %t\n", simpleFalse) // simpleFalse: false

	// --- Float Formatting Verbs
	// Verb Description
	// %e    Scientific notation with 'e' as exponent
	// %f    Decimal point, no exponent
	// %.2f  Default width, precission 2
	// %6.2f Width 6, precission 2
	// %g	 Exponent as needed, only necessary digits

	anotherFloat := 3.14159265359

	fmt.Printf("\nanotherFloat: %e\n", anotherFloat)  // anotherFloat: 3.141593e+00
	fmt.Printf("anotherFloat: %f\n", anotherFloat)    // anotherFloat: 3.141593
	fmt.Printf("anotherFloat: %.2f\n", anotherFloat)  // anotherFloat: 3.14
	fmt.Printf("anotherFloat: %6.2f\n", anotherFloat) // anotherFloat:  3.14
	fmt.Printf("anotherFloat: %g\n", anotherFloat)    // anotherFloat: 3.14159265359

}
