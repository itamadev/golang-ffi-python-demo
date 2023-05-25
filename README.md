# golang-ffi-python-demo

This repo aims to show a simple example of how to use Golang to write a shared library and call it from Python,
this simple example can be used as a template for more complex projects and can be applied to more than just Python.

When implementing a library or tool FFI is a great way to allow other languages to use it without having to rewrite it,
but it can be a bit tricky to get it working, this repo aims to show how to do it, simply with a working example.

You can apply this pattern for your project to keep it simple and keep your repo clean.

## Running the example

To run the example you need to have Golang installed and Python 3.7+ with the `cffi` package installed.

To build the shared library run:

```bash
make build
```

Or jusr use the run command to build and run the example:

```bash
make run
```
