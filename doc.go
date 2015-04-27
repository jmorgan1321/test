/*
test is a simplistic testing library that makes it easier and uniform to test
a variety of different situations (like panics, equality, or float equality).
test provides two functions (Assert and Expect) that are used to "validate"
different assertion functions that can be passed into one of those two
functions.

Here are examples are some examples of using the library:
    // example
    c := test.Checker(t)
    c.Assert(test.FloatEQ, 0.0, Origin.Mag())
    c.Expect(test.EQ, Origin, Origin.Div(5))
    c.Expect(test.PanicEQ, "div by zero", func() { Origin.Div(0) })

    // another example
    for i, tt := range tests {
        c := test.Checker(t, test.Summary("with test %v: %v", i, tt.summary))
        SerializeInPlace(obj, data)
        typ := obj.(*testType)
        c.Expect(test.EQ, 5, typ.I)
        c.Expect(test.FloatEQ, 3.14, typ.F)
        c.Expect(test.EQ, "string", typ.S)
        c.Expect(test.EQ, true, typ.B)
    }

*/
package test
