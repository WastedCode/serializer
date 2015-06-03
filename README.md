# Crypt
A simple GoLang package to serialize data

[![Build Status](https://travis-ci.org/WastedCode/serializer.svg)](https://travis-ci.org/WastedCode/serializer)

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
