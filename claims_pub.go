package sjwt

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// Add adds a name/value to claims
func (c Claims) Add(name string, value interface{}) { c[name] = value }

// Del deletes a name/value from claims
func (c Claims) Del(name string) { delete(c, name) }

// Has will let you know whether or not a claim exists
func (c Claims) Has(name string) bool { _, ok := c[name]; return ok }

// Get gets claim value
func (c Claims) Get(name string) interface{} { return c[name] }

// GetBool will get the boolean value on the Claims
func (c Claims) GetBool(name string) bool {
	if _, ok := c[name]; ok {
		switch val := c[name].(type) {
		case string:
			v, _ := strconv.ParseBool(val)
			return v
		case bool:
			return val
		}
	}
	return false
}

// GetStr will get the string value on the Claims
func (c Claims) GetStr(name string) string {
	if _, ok := c[name]; ok {
		switch val := c[name].(type) {
		case string:
			return val
		case float32:
			return strconv.FormatFloat(float64(val), 'f', -1, 32)
		case float64:
			return strconv.FormatFloat(val, 'f', -1, 64)
		default:
			return fmt.Sprintf("%v", val)
		}
	}
	return ""
}

// GetInt will get the int value on the Claims
func (c Claims) GetInt(name string) int {
	if _, ok := c[name]; ok {
		switch val := c[name].(type) {
		case float32:
			return int(val)
		case float64:
			return int(val)
		case string:
			v, _ := strconv.ParseInt(val, 10, 64)
			return int(v)
		case uint:
			return int(val)
		case int:
			return int(val)
		}
	}

	return 0
}

// GetFloat will get the float value on the Claims
func (c Claims) GetFloat(name string) float64 {
	if _, ok := c[name]; ok {
		switch val := c[name].(type) {
		case float64:
			return float64(val)
		case string:
			v, _ := strconv.ParseFloat(val, 64)
			return v
		case json.Number:
			v, _ := val.Float64()
			return v
		}
	}

	return 0
}
