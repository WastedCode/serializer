package serializer

import (
    "testing"
)

type testingInterface struct {
    Num int
    ValueStr string
}

func TestBase64Encoding(t *testing.T) {
    bytes := []byte("abcdef")
    encStr := ByteToBase64String(bytes)
    if (encStr != "YWJjZGVm") {
        t.Error("Unable to encode to Base 64")
    }

    calculatedBytes, _ := Base64StringToByte(encStr)
    if (string(calculatedBytes) != "abcdef") {
        t.Error("Unable to decode Base 64 string")
    }
}

func TestSerializeInterfaceUsingGob(t *testing.T) {
    testInterface := testingInterface{10, "abc"}

    serialized, err := SerializeToGob(testInterface)
    if (err != nil) {
        t.Error("Raised error while serializing interface")
        t.Error(err)
    }

    deserialized := testingInterface{}
    err = DeserializeFromGob(serialized, &deserialized)
    if (err != nil) {
        t.Error("Unable to deserialize bytes using gob")
    }

    if ((deserialized.ValueStr != testInterface.ValueStr) ||
        (deserialized.Num != testInterface.Num)) {
        t.Error("The deserialized values dont match the original data")
    }
}

func TestSerializeInterfaceUsingGobAndString(t *testing.T) {
    testInterface := testingInterface{100, "xyz"}
    serialized, err := SerializeInterfaceToString(testInterface)
    if (err != nil) {
        t.Error("Raised error while serializing interface")
        t.Error(err)
    }

    deserialized := testingInterface{}
    err = DeserializeStringToInterface(serialized, &deserialized)
    if (err != nil) {
        t.Error("Unable to deserialize bytes using gob")
    }

    if ((deserialized.ValueStr != testInterface.ValueStr) ||
        (deserialized.Num != testInterface.Num)) {
        t.Error("The deserialized values dont match the original data")
    }
}
