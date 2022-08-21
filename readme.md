# Read and Write json files in golang without deprecated ioutils

I was looking for examples to read a json file into a struct, and then write the struct changes into the json file again.

I noticed many of the files used deprecated functions from ioutils

>Package ioutil implements some I/O utility functions.

>Deprecated: As of Go 1.16, the same functionality is now provided by package io or package os, and those implementations should be preferred in new code. See the specific function documentation for details.

So I wrote this small sample in main.go