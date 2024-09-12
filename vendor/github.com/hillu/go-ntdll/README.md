Go interface to NTDLL functions
===============================

[![GoDoc](https://godoc.org/github.com/hillu/go-ntdll?status.svg)](https://godoc.org/github.com/hillu/go-ntdll)
[![Go Report Card](https://goreportcard.com/badge/github.com/hillu/go-ntdll)](https://goreportcard.com/report/github.com/hillu/go-ntdll)

This package makes selected `NTDLL` functions directly available in Go
programs. At the moment, types and functions for accessing kernel
objects and the Registry are included. The goal is to, eventually,
cover all available functions.

Note
----

Please do not consider this code, particularly the autogenerated code,
as stable. Identifiers names may still be subject to change.

Basic functions, types
----------------------

NTDLL functions usually return an `NTSTATUS`, this is mapped into an
`NtStatus` type (underlying kind: `uint32`). Strings are usually
passed around as `UNICODE_STRING`, there is a type `UnicodeStringType`
for that.

Extending
---------

Function signatures and structure definitions are generated from
header file snippets as found in the MSDN. The transformation code can
be found in `mkcode.go`, it probably has bugs and does not cover all
cases.

To add more structures or function stubs or enumeration types, just
find the appropriate definition on MSDN and copy them into block
comments, prefixed by `type:`, `func:`, or `enum:`, respectively. `go
generate` can then be used to regenerate the `*_generated.go` files
containing structure definitions and the glue code around
`golang.org/x/sys/windows`.

License
-------

Copyright (C) 2016-2021  Hilko Bengen <bengen@hilluzination.de>

All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

* Redistributions of source code must retain the above copyright notice, this
  list of conditions and the following disclaimer.

* Redistributions in binary form must reproduce the above copyright notice,
  this list of conditions and the following disclaimer in the documentation
  and/or other materials provided with the distribution.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.