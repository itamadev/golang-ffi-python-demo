# golang-ffi-python-demo

This repo aims to show a simple example of how to use Golang to write a shared library and call it from Python,
this simple example can be used as a template for more complex projects and can be applied to more than just Python.

When implementing a library or tool FFI is a great way to allow other languages to use it without having to rewrite it,
but it can be a bit tricky to get it working, this repo aims to show how to do it, simply with a working example.

You can apply this pattern for your project to keep it simple and keep your repo clean.

## Prerequisites

- Golang
- Python 3.7+
- gcc
- make

## Running the example

To build the shared library run:

```bash
make build
```

Or jusr use the run command to build and run the example:

```bash
make run
```

## How it works

The `ffi` directory contains the `ffi.go` file which is the entry point for the shared library, it contains the `export` directive which tells the compiler to export the functions to the shared library.
Now we have an `ffi.py` and it's used as the interface letting Python users access to the built shared library which is in the `ffi` directory.

That way we can easily build the shared library and use it from Python while having both the Golang and Python code seperated or like this example shows, in the same repo but in different directories,
this makes the process and maintenance easier and easier to understand.