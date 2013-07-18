package strings

import (
	"strconv"
	"strings"
)

// Parse tries to create a native object from the given
// string, or just returns the string if nothing takes.
//
// Values wrapped in "quotes" or 'single quotes' will always
// be treated as a string, but the quotes will be removed.
//
// This method knows about all number types, and will always
// look for the smallest type to fit the number.  It also handles
// the boolean literals 'true' and 'false'.
func Parse(s string) interface{} {

	/*
	   Is it forced to be a string with quotes?
	*/
	if strings.HasPrefix(s, `"`) && strings.HasSuffix(s, `"`) {
		return strings.Trim(s, `"`)
	}
	if strings.HasPrefix(s, `'`) && strings.HasSuffix(s, `'`) {
		return strings.Trim(s, `'`)
	}

	/*
	   Check literals
	*/
	switch strings.ToLower(s) {
	/*
	   Booleans
	*/
	case "true":
		return true
	case "false":
		return false
	}

	/*
	   Numbers
	*/

	// try int (most common type)
	if val, err := strconv.ParseInt(s, 10, 0); err == nil {
		return int(val)
	}

	/*
	   ints
	*/

	// try int8
	if val, err := strconv.ParseInt(s, 10, 8); err == nil {
		return val
	}

	// try int16
	if val, err := strconv.ParseInt(s, 10, 16); err == nil {
		return val
	}

	// try int32
	if val, err := strconv.ParseInt(s, 10, 32); err == nil {
		return val
	}

	// try int64
	if val, err := strconv.ParseInt(s, 10, 64); err == nil {
		return val
	}

	/*
	   uints
	*/
	// try uint8
	if val, err := strconv.ParseUint(s, 10, 8); err == nil {
		return val
	}

	// try uint16
	if val, err := strconv.ParseUint(s, 10, 16); err == nil {
		return val
	}

	// try uint32
	if val, err := strconv.ParseUint(s, 10, 32); err == nil {
		return val
	}

	// try uint64
	if val, err := strconv.ParseUint(s, 10, 64); err == nil {
		return val
	}

	/*
	   floats
	*/

	// try float32
	if val, err := strconv.ParseFloat(s, 32); err == nil {
		return float32(val)
	}

	// try float64
	if val, err := strconv.ParseFloat(s, 64); err == nil {
		return val
	}

	/*
	   Nothing - just return the string
	*/
	return s
}