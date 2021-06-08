# Definitions

## Keywords

| Keyword      |                                                              |
| ------------ | ------------------------------------------------------------ |
| `func`       | The `func` keyword is how you define a function with a name and a body |
| `import`     | Used to import packages                                      |
| `package`    | A package is a collection of source files in the same directory that are compiled together. Functions, types, variables, and constants defined in one source file are visible to all other source files within the same package. [source](https://golang.org/doc/code#:~:text=Go%20programs%20are%20organized%20into,contains%20one%20or%20more%20modules.) |
| `module`     | A module is a collection of related Go packages that are released together. |
| `repository` | A repository contains one or more modules. A Go repository typically contains only one module, located at the root of the repository. A file named `go.mod` there declares the module path: the import path prefix for all packages within the module. The module contains the packages in the directory containing its `go.mod` file as well as subdirectories of that directory, up to the next subdirectory containing another `go.mod` file (if any). |
|              |                                                              |
|              |                                                              |
|              |                                                              |
|              |                                                              |

