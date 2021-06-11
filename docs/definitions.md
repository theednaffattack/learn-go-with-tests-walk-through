# Definitions

## Keywords

| Keyword      |                                                              |
| ------------ | ------------------------------------------------------------ |
| `func`       | The `func` keyword is how you define a function with a name and a body |
| `import`     | Used to import packages                                      |
| `package`    | A package is a collection of source files in the same directory that are compiled together. Functions, types, variables, and constants defined in one source file are visible to all other source files within the same package. [source](https://golang.org/doc/code#:~:text=Go%20programs%20are%20organized%20into,contains%20one%20or%20more%20modules.) |
| `module`     | A module is a collection of related Go packages that are released together. |
| `repository` | A repository contains one or more modules. A Go repository typically contains only one module, located at the root of the repository. A file named `go.mod` there declares the module path: the import path prefix for all packages within the module. The module contains the packages in the directory containing its `go.mod` file as well as subdirectories of that directory, up to the next subdirectory containing another `go.mod` file (if any). |
| `len`        | The len built-in function returns the length of v, according to its type:<br />- Array: the number of elements in v.<br />- Pointer to array: the number of elements in *v (even if v is nil).<br />- Slice, or map: the number of elements in v; if v is nil, len(v) is zero.<br />- String: the number of bytes in v.<br />- Channel: the number of elements queued (unread) in the channel buffer;<br />- if v is nil, len(v) is zero.<br/>For some arguments, such as a string literal or a simple array expression, the result can be a constant. See the Go language specification's "Length and capacity" section for details. |
| `make`       | The make built-in function allocates and initializes an object of type slice, map, or chan (only). Like new, the first argument is a type, not a value. Unlike new, make's return type is the same as the type of its argument, not a pointer to it. The specification of the result depends on the type: |
| `append`     | The append built-in function appends elements to the end of a slice. If it has sufficient capacity, the destination is resliced to accommodate the new elements. If it does not, a new underlying array will be allocated. Append returns the updated slice. It is therefore necessary to store the result of append, often in the variable holding the slice itself:<br />```slice = append(slice, elem1, elem2) slice = append(slice, anotherSlice...)```<br /><br />```slice = append([]byte("hello "), "world"...)``` |
|              |                                                              |
|              |                                                              |

