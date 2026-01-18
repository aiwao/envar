package envar

import (
	"encoding/json"
	"os"
	"strconv"
)

// GetInt parse env to int and return result
func GetInt(key string) (int, error) {
	conv, err := strconv.ParseInt(os.Getenv(key), 10, strconv.IntSize)
	if err != nil {
		return 0, err
	}
	return int(conv), nil
}

// GetInt8 parse env to int8 and return result
func GetInt8(key string) (int8, error) {
	conv, err := strconv.ParseInt(os.Getenv(key), 10, 8)
	if err != nil {
		return 0, err
	}
	return int8(conv), nil
}

// GetInt16 parse env to int16 and return result
func GetInt16(key string) (int16, error) {
	conv, err := strconv.ParseInt(os.Getenv(key), 10, 16)
	if err != nil {
		return 0, err
	}
	return int16(conv), nil
}

// GetInt32 parse env to int32 and return result
func GetInt32(key string) (int32, error) {
	conv, err := strconv.ParseInt(os.Getenv(key), 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(conv), nil
}

// GetInt64 parse env to int64 and return result
func GetInt64(key string) (int64, error) {
	return strconv.ParseInt(os.Getenv(key), 10, 64)
}

// GetUInt parse env to uint and return result
func GetUInt(key string) (uint, error) {
	conv, err := strconv.ParseUint(os.Getenv(key), 10, strconv.IntSize)
	if err != nil {
		return 0, err
	}
	return uint(conv), nil
}

// GetUInt8 parse env to uint8 and return result
func GetUInt8(key string) (uint8, error) {
	conv, err := strconv.ParseUint(os.Getenv(key), 10, 8)
	if err != nil {
		return 0, err
	}
	return uint8(conv), nil
}

// GetUInt16 parse env to uint16 and return result
func GetUInt16(key string) (uint16, error) {
	conv, err := strconv.ParseUint(os.Getenv(key), 10, 16)
	if err != nil {
		return 0, err
	}
	return uint16(conv), nil
}

// GetUInt32 parse env to uint32 and return result
func GetUInt32(key string) (uint32, error) {
	conv, err := strconv.ParseUint(os.Getenv(key), 10, 32)
	if err != nil {
		return 0, err
	}
	return uint32(conv), nil
}

// GetUInt64 parse env to uint64 and return result
func GetUInt64(key string) (uint64, error) {
	return strconv.ParseUint(os.Getenv(key), 10, 64)
}

// GetBool parse env to bool and return result
func GetBool(key string) (bool, error) {
	return strconv.ParseBool(os.Getenv(key))
}

// GetComplex64 parse env to complex64 and return result
func GetComplex64(key string) (complex64, error) {
	conv, err := strconv.ParseComplex(os.Getenv(key), 64)
	if err != nil {
		return 0i, err
	}
	return complex64(conv), err
}

// GetComplex128 parse env to complex128 and return result
func GetComplex128(key string) (complex128, error) {
	return strconv.ParseComplex(os.Getenv(key), 128)
}

// GetByte parse env to byte and return result
func GetByte(key string) (byte, error) {
	conv, err := strconv.ParseUint(os.Getenv(key), 10, 8)
	if err != nil {
		return 0, err
	}
	return byte(conv), nil
}

// GetIntv if converting value to int succeeds, set value
func GetIntv(key string, value *int) error {
	conv, err := strconv.ParseInt(os.Getenv(key), 10, strconv.IntSize)
	if err != nil {
		return err
	}
	*value = int(conv)
	return nil
}

// GetInt8v if converting value to int8 succeeds, set value
func GetInt8v(key string, value *int8) error {
	conv, err := strconv.ParseInt(os.Getenv(key), 10, 8)
	if err != nil {
		return err
	}
	*value = int8(conv)
	return nil
}

// GetInt16v if converting value to int16 succeeds, set value
func GetInt16v(key string, value *int16) error {
	conv, err := strconv.ParseInt(os.Getenv(key), 10, 16)
	if err != nil {
		return err
	}
	*value = int16(conv)
	return nil
}

// GetInt32v if converting value to int32 succeeds, set value
func GetInt32v(key string, value *int32) error {
	conv, err := strconv.ParseInt(os.Getenv(key), 10, 32)
	if err != nil {
		return err
	}
	*value = int32(conv)
	return nil
}

// GetInt64v if converting value to int64 succeeds, set value
func GetInt64v(key string, value *int64) error {
	conv, err := strconv.ParseInt(os.Getenv(key), 10, 64)
	if err != nil {
		return err
	}
	*value = conv
	return nil
}

// GetUIntv if converting value to uint succeeds, set value
func GetUIntv(key string, value *uint) error {
	conv, err := strconv.ParseUint(os.Getenv(key), 10, strconv.IntSize)
	if err != nil {
		return err
	}
	*value = uint(conv)
	return nil

}

// GetUInt8v if converting value to uint8 succeeds, set value
func GetUInt8v(key string, value *uint8) error {
	conv, err := strconv.ParseUint(os.Getenv(key), 10, 8)
	if err != nil {
		return err
	}
	*value = uint8(conv)
	return nil
}

// GetUInt16v if converting value to uint16 succeeds, set value
func GetUInt16v(key string, value *uint16) error {
	conv, err := strconv.ParseUint(os.Getenv(key), 10, 16)
	if err != nil {
		return err
	}
	*value = uint16(conv)
	return nil
}

// GetUInt32v if converting value to uint32 succeeds, set value
func GetUInt32v(key string, value *uint32) error {
	conv, err := strconv.ParseUint(os.Getenv(key), 10, 32)
	if err != nil {
		return err
	}
	*value = uint32(conv)
	return nil
}

// GetUInt64v if converting value to uint64 succeeds, set value
func GetUInt64v(key string, value *uint64) error {
	conv, err := strconv.ParseUint(os.Getenv(key), 10, 64)
	if err != nil {
		return err
	}
	*value = conv
	return nil
}

// GetBoolv if converting value to bool succeeds, set value
func GetBoolv(key string, value *bool) error {
	conv, err := strconv.ParseBool(os.Getenv(key))
	if err != nil {
		return err
	}
	*value = conv
	return nil
}

// GetComplex64v if converting value to complex64 succeeds, set value
func GetComplex64v(key string, value *complex64) error {
	conv, err := strconv.ParseComplex(os.Getenv(key), 64)
	if err != nil {
		return err
	}
	*value = complex64(conv)
	return nil
}

// GetComplex128v if converting value to complex128 succeeds, set value
func GetComplex128v(key string, value *complex128) error {
	conv, err := strconv.ParseComplex(os.Getenv(key), 128)
	if err != nil {
		return err
	}
	*value = conv
	return nil
}

// GetBytev if converting value to byte succeeds, set value
func GetBytev(key string, value *byte) error {
	conv, err := strconv.ParseUint(os.Getenv(key), 10, 8)
	if err != nil {
		return err
	}
	*value = byte(conv)
	return nil
}

// Unmarshal if unmarshal succeeds, set value
func Unmarshal(key string, value any) error {
	if err := json.Unmarshal([]byte(os.Getenv(key)), value); err != nil {
		return err
	}
	return nil
}
