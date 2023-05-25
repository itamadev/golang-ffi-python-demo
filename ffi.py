import ctypes

# Load the shared library
lib = ctypes.CDLL("./engine/lib.so")

# Declare the argument and return types for the print function
lib.print.argtypes = [ctypes.c_char_p]
lib.print.restype = None

def libprint(name: str):
    lib.print(bytes(name, "utf-8"))
