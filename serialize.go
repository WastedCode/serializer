package serializer

import (
    "bytes"
    "encoding/base64"
    "encoding/gob"
)

// Serialized the given interface to an array of bytes
// using Gob for serialization
// Returns an array of serialized bytes or an error
// Error is nil if the serialization succeeds
func SerializeToGob(value interface{}) ([]byte, error) {
    buf := new(bytes.Buffer)
    enc := gob.NewEncoder(buf)
    if err := enc.Encode(value); err != nil {
        return nil, err
    }
    return buf.Bytes(), nil
}

// Deserialize bytes in gob encoding to an interface
// value contains the deserialized object
// Returns nil if succeeds
func DeserializeFromGob(serialized []byte, value interface{}) error {
    dec := gob.NewDecoder(bytes.NewBuffer(serialized))
    if err := dec.Decode(value); err != nil {
        return err
    }
    return nil
}

// Serializes the given interface into a string
// It first encodes in gob, and then returns a base64 encoded string for the bytes
func SerializeInterfaceToString(value interface{}) (string, error) {
    bytes, err := SerializeToGob(value)
    if (err != nil) {
        return "", err
    }
    return ByteToBase64String(bytes), nil
}

// Deserializes a base 64 string into the interface
// The input string must be the base64 encoding of gob output
func DeserializeStringToInterface(serialized string, value interface{}) error {
    bytes, err := Base64StringToByte(serialized)
    if (err != nil) {
        return err
    }
    err = DeserializeFromGob(bytes, value)
    if (err != nil) {
        return err
    }
    return nil
}

// Encode to Base64 using URLEncoding
func ByteToBase64String(bytes []byte) string {
    return base64.URLEncoding.EncodeToString(bytes)
}

// Decode the input base64 string into bytes
func Base64StringToByte(str string) ([]byte, error) {
    return base64.URLEncoding.DecodeString(str)
}
