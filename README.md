# Crypt
A simple GoLang package to serialize data

[![Build Status](https://travis-ci.org/WastedCode/serializer.svg)](https://travis-ci.org/WastedCode/serializer) [![GoDoc](https://godoc.org/github.com/WastedCode/serializer?status.svg)](https://godoc.org/github.com/WastedCode/serializer)

## Example Usage
```
type TestingStruct struct {
    Num int
    ValueStr string
}

value := TestingStruct(1, "Hello")
serialized, err := SerializeInterfaceToString(value)

deserialized := TestingStruct{}
err = DeserializeStringToInterface(serialized)
```
