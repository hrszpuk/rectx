# rectx/utilities
This go package is for storing functions, types, and other reusable statements that many other packages may want to use.
Due to Go's cycle import errors, we cannot allow for multiple packages to import each other.
Each file stores a single function that can be used by other packages.
