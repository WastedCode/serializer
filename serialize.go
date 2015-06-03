package serializer

import (
    "bytes"
    "encoding/base64"
    "encoding/gob"
)

func SerializeToGob(value interface{}) ([]byte, error) {
    buf := new(bytes.Buffer)
    enc := gob.NewEncoder(buf)
    if err := enc.Encode(value); err != nil {
        return nil, err
    }
    return buf.Bytes(), nil
}

func DeserializeFromGob(serialized []byte, value interface{}) error {
    dec := gob.NewDecoder(bytes.NewBuffer(serialized))
    if err := dec.Decode(value); err != nil {
        return err
    }
    return nil
}

func SerializeInterfaceToString(value interface{}) (string, error) {
    bytes, err := SerializeToGob(value)
    if (err != nil) {
        return "", err
    }
    return ByteToBase64String(bytes), nil
}

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

func ByteToBase64String(bytes []byte) string {
    return base64.URLEncoding.EncodeToString(bytes)
}

func Base64StringToByte(str string) ([]byte, error) {
    return base64.URLEncoding.DecodeString(str)
}
