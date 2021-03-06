# Interoperability  

## Table of Contents 

- [Using Scripts](#using-scripts)
  - [Type Conversion Table](#type-conversion-table)
  - [User Types](#user-types)
  - [Importing Scripts](#importing-scripts)
- [Sandbox Environments](#sandbox-environments)
- [Compiler and VM](#compiler-and-vm)

## Using Scripts

Embedding and executing the Tengo code in Go is very easy. At a high level, this process is like:

- create a [Script](https://godoc.org/github.com/d5/tengo/script#Script) instance with your code,
- _optionally_ add some [Script Variables](https://godoc.org/github.com/d5/tengo/script#Variable) to Script,
- compile or directly run the script,
- retrieve _output_ values from the [Compiled](https://godoc.org/github.com/d5/tengo/script#Compiled) instance.

The following is an example where a Tengo script is compiled and run with no input/output variables.

```golang
import "github.com/d5/tengo/script"

var code = `
reduce := func(seq, fn) {
    s := 0
    for x in seq { fn(x, s) }
    return s
}

print(reduce([1, 2, 3], func(x, s) { s += x }))
`

func main() {
    s := script.New([]byte(code))
    if _, err := s.Run(); err != nil {
        panic(err)
    }
}
```

Here's another example where an input variable is added to the script, and, an output variable is accessed through [Variable.Int](https://godoc.org/github.com/d5/tengo/script#Variable.Int) function:

```golang
import (
	"fmt"

	"github.com/d5/tengo/script"
)

func main() {
	s := script.New([]byte(`a := b + 20`))

	// define variable 'b'
	_ = s.Add("b", 10)

	// compile the source
	c, err := s.Compile()
	if err != nil {
		panic(err)
	}

	// run the compiled bytecode
	// a compiled bytecode 'c' can be executed multiple times without re-compiling it
	if err := c.Run(); err != nil {
		panic(err)
	}

	// retrieve value of 'a'
	a := c.Get("a")
	fmt.Println(a.Int())           // prints "30"
	
	// re-run after replacing value of 'b'
	if err := c.Set("b", 20); err != nil {
		panic(err)
	}
	if err := c.Run(); err != nil {
		panic(err)
	}
	fmt.Println(c.Get("a").Int())  // prints "40"
}
```

A variable `b` is defined by the user before compilation using [Script.Add](https://godoc.org/github.com/d5/tengo/script#Script.Add) function. Then a compiled bytecode `c` is used to execute the bytecode and get the value of global variables. In this example, the value of global variable `a` is read using [Compiled.Get](https://godoc.org/github.com/d5/tengo/script#Compiled.Get) function. See [documentation](https://godoc.org/github.com/d5/tengo/script#Variable) for the full list of variable value functions.

Value of the global variables can be replaced using [Compiled.Set](https://godoc.org/github.com/d5/tengo/script#Compiled.Set) function. But it will return an error if you try to set the value of un-defined global variables _(e.g. trying to set the value of `x` in the example)_.  

### Type Conversion Table

When adding a Variable _([Script.Add](https://godoc.org/github.com/d5/tengo/script#Script.Add))_, Script converts Go values into Tengo values based on the following conversion table.

| Go Type | Tengo Type | Note |
| :--- | :--- | :--- |
|`nil`|`Undefined`||
|`string`|`String`||
|`int64`|`Int`||
|`int`|`Int`||
|`bool`|`Bool`||
|`rune`|`Char`||
|`byte`|`Char`||
|`float64`|`Float`||
|`[]byte`|`Bytes`||
|`time.Time`|`Time`||
|`error`|`Error{String}`|use `error.Error()` as String value|
|`map[string]Object`|`Map`||
|`map[string]interface{}`|`Map`|individual elements converted to Tengo objects|
|`[]Object`|`Array`||
|`[]interface{}`|`Array`|individual elements converted to Tengo objects|
|`Object`|`Object`|_(no type conversion performed)_|


### User Types

Users can add and use a custom user type in Tengo code by implementing [Object](https://godoc.org/github.com/d5/tengo/objects#Object) interface. Tengo runtime will treat the user types in the same way it does to the runtime types with no performance overhead. See [Object Types](https://github.com/d5/tengo/blob/master/docs/objects.md) for more details.

### Importing Scripts

A script can import and use another script in the same way it can load the standard library or the user module. `Script.AddModule` function adds another script as a named module.

```golang
mod1Script := script.New([]byte(`a := 5`))                  // mod1 script

mainScript := script.New([]byte(`print(import("mod1").a)`)) // main script
mainScript.AddModule("mod1", mod1Script)                    // add mod1 using name "mod1"
mainScript.Run()                                            // prints "5"
```

Note that the script modules added using `Script.AddModule` will be compiled and run right before the main script is compiled.   

## Sandbox Environments

To securely compile and execute _potentially_ unsafe script code, you can use the following Script functions.

#### Script.DisableBuiltinFunction(name string)

DisableBuiltinFunction disables and removes a builtin function from the compiler. Compiler will reports a compile-time error if the given name is referenced.

```golang
s := script.New([]byte(`print([1, 2, 3])`))

s.DisableBuiltinFunction("print") 

_, err := s.Run() // compile error 
```

#### Script.DisableStdModule(name string)

DisableStdModule disables a [standard library](https://github.com/d5/tengo/blob/master/docs/stdlib.md) module. Compile will report a compile-time error if the code tries to import the module with the given name.

```golang
s := script.New([]byte(`import("exec")`))

s.DisableStdModule("exec") 

_, err := s.Run() // compile error 
```

#### Script.SetUserModuleLoader(loader compiler.ModuleLoader)

SetUserModuleLoader replaces the default user-module loader of the compiler, which tries to read the source from a local file.  

```golang
s := script.New([]byte(`math := import("mod1"); a := math.foo()`))
 
s.SetUserModuleLoader(func(moduleName string) ([]byte, error) {
    if moduleName == "mod1" {
        return []byte(`foo := func() { return 5 }`), nil
    }

    return nil, errors.New("module not found")
})
```

## Compiler and VM

Although it's not recommended, you can directly create and run the Tengo [Parser](https://godoc.org/github.com/d5/tengo/compiler/parser#Parser), [Compiler](https://godoc.org/github.com/d5/tengo/compiler#Compiler), and [VM](https://godoc.org/github.com/d5/tengo/runtime#VM) for yourself instead of using Scripts and Script Variables. It's a bit more involved as you have to manage the symbol tables and global variables between them, but, basically that's what Script and Script Variable is doing internally.

_TODO: add more information here_
