package simplejson

// returns the current implementation version
func Version() string { _ = "STUB: not implemented"; return "" }

type Json struct {
	data interface{}
}

// NewJson returns a pointer to a new `Json` object
// after unmarshaling `body` bytes
func NewJson(body []byte) (*Json, error) { _ = "STUB: not implemented"; return nil, nil }

// New returns a pointer to a new, empty `Json` object
func New() *Json { _ = "STUB: not implemented"; return nil }

// Interface returns the underlying data
func (j *Json) Interface() interface{} {
	_ = "STUB: not implemented"

	// Encode returns its marshaled data as `[]byte`
	return nil
}

func (j *Json) Encode() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// EncodePretty returns its marshaled data as `[]byte` with indentation
		nil
}

func (j *Json) EncodePretty() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Implements the json.Marshaler interface.
func (j *Json) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// Set modifies `Json` map by `key` and `value`
		// Useful for changing single key/value in a `Json` object easily.
		nil
}

func (j *Json) Set(key string, val interface{}) { _ = "STUB: not implemented"; return }

// SetPath modifies `Json`, recursively checking/creating map keys for the supplied path,
// and then finally writing in the value
func (j *Json) SetPath(branch []string, val interface{}) { _ = "STUB: not implemented"; return }

// in order to insert our branch, we need map[string]interface{}

// have to replace with something suitable

// key exists?

// make sure the value is the right sort of thing

// have to replace with something suitable

// add remaining k/v

// Del modifies `Json` map by deleting `key` if it is present.
func (j *Json) Del(key string) { _ = "STUB: not implemented"; return }

// Get returns a pointer to a new `Json` object
// for `key` in its `map` representation
//
// useful for chaining operations (to traverse a nested JSON):
//
//	js.Get("top_level").Get("dict").Get("value").Int()
func (j *Json) Get(key string) *Json { _ = "STUB: not implemented"; return nil }

// GetPath searches for the item as specified by the branch
// without the need to deep dive using Get()'s.
//
//	js.GetPath("top_level", "dict")
func (j *Json) GetPath(branch ...string) *Json { _ = "STUB: not implemented"; return nil }

// GetIndex returns a pointer to a new `Json` object
// for `index` in its `array` representation
//
// this is the analog to Get when accessing elements of
// a json array instead of a json object:
//
//	js.Get("top_level").Get("array").GetIndex(1).Get("key").Int()
func (j *Json) GetIndex(index int) *Json { _ = "STUB: not implemented"; return nil }

// CheckGet returns a pointer to a new `Json` object and
// a `bool` identifying success or failure
//
// useful for chained operations when success is important:
//
//	if data, ok := js.Get("top_level").CheckGet("inner"); ok {
//	    log.Println(data)
//	}
func (j *Json) CheckGet(key string) (*Json, bool) { _ = "STUB: not implemented"; return nil, false }

// Map type asserts to `map`
func (j *Json) Map() (map[string]interface{}, error) { _ = "STUB: not implemented"; return nil, nil }

// Array type asserts to an `array`
func (j *Json) Array() ([]interface{}, error) { _ = "STUB: not implemented"; return nil, nil }

// Bool type asserts to `bool`
func (j *Json) Bool() (bool, error) { _ = "STUB: not implemented"; return false, nil }

// String type asserts to `string`
func (j *Json) String() (string, error) { _ = "STUB: not implemented"; return "", nil }

// Bytes type asserts to `[]byte`
func (j *Json) Bytes() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// StringArray type asserts to an `array` of `string`
func (j *Json) StringArray() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

// MustArray guarantees the return of a `[]interface{}` (with optional default)
//
// useful when you want to interate over array values in a succinct manner:
//
//	for i, v := range js.Get("results").MustArray() {
//		fmt.Println(i, v)
//	}
func (j *Json) MustArray(args ...[]interface{}) []interface{} {
	_ = "STUB: not implemented"
	return nil
}

// MustMap guarantees the return of a `map[string]interface{}` (with optional default)
//
// useful when you want to interate over map values in a succinct manner:
//
//	for k, v := range js.Get("dictionary").MustMap() {
//		fmt.Println(k, v)
//	}
func (j *Json) MustMap(args ...map[string]interface{}) map[string]interface{} {
	_ = "STUB: not implemented"
	return nil
}

// MustString guarantees the return of a `string` (with optional default)
//
// useful when you explicitly want a `string` in a single value return context:
//
//	myFunc(js.Get("param1").MustString(), js.Get("optional_param").MustString("my_default"))
func (j *Json) MustString(args ...string) string { _ = "STUB: not implemented"; return "" }

// MustStringArray guarantees the return of a `[]string` (with optional default)
//
// useful when you want to interate over array values in a succinct manner:
//
//	for i, s := range js.Get("results").MustStringArray() {
//		fmt.Println(i, s)
//	}
func (j *Json) MustStringArray(args ...[]string) []string { _ = "STUB: not implemented"; return nil }

// MustInt guarantees the return of an `int` (with optional default)
//
// useful when you explicitly want an `int` in a single value return context:
//
//	myFunc(js.Get("param1").MustInt(), js.Get("optional_param").MustInt(5150))
func (j *Json) MustInt(args ...int) int { _ = "STUB: not implemented"; return 0 }

// MustFloat64 guarantees the return of a `float64` (with optional default)
//
// useful when you explicitly want a `float64` in a single value return context:
//
//	myFunc(js.Get("param1").MustFloat64(), js.Get("optional_param").MustFloat64(5.150))
func (j *Json) MustFloat64(args ...float64) float64 { _ = "STUB: not implemented"; return 0 }

// MustBool guarantees the return of a `bool` (with optional default)
//
// useful when you explicitly want a `bool` in a single value return context:
//
//	myFunc(js.Get("param1").MustBool(), js.Get("optional_param").MustBool(true))
func (j *Json) MustBool(args ...bool) bool { _ = "STUB: not implemented"; return false }

// MustInt64 guarantees the return of an `int64` (with optional default)
//
// useful when you explicitly want an `int64` in a single value return context:
//
//	myFunc(js.Get("param1").MustInt64(), js.Get("optional_param").MustInt64(5150))
func (j *Json) MustInt64(args ...int64) int64 { _ = "STUB: not implemented"; return 0 }

// MustUInt64 guarantees the return of an `uint64` (with optional default)
//
// useful when you explicitly want an `uint64` in a single value return context:
//
//	myFunc(js.Get("param1").MustUint64(), js.Get("optional_param").MustUint64(5150))
func (j *Json) MustUint64(args ...uint64) uint64 { _ = "STUB: not implemented"; return 0 }
