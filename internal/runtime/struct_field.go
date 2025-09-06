package runtime

import (
	"reflect"
	"strings"
	"unicode"

	"github.com/jslyzt/go-json/internal/option/decode"
	"github.com/jslyzt/go-json/internal/option/encode"
)

func getTag(field reflect.StructField) string {
	return field.Tag.Get("json")
}

func IsIgnoredStructField(field reflect.StructField) bool {
	if field.PkgPath != "" {
		if field.Anonymous {
			t := field.Type
			if t.Kind() == reflect.Ptr {
				t = t.Elem()
			}
			if t.Kind() != reflect.Struct {
				return true
			}
		} else {
			// private field
			return true
		}
	}
	tag := getTag(field)
	return tag == "-"
}

type StructTag struct {
	Key         string
	IsTaggedKey bool
	IsOmitEmpty bool
	IsString    bool
	Field       reflect.StructField
}

type StructTags []*StructTag

func (t StructTags) ExistsKey(key string) bool {
	for _, tt := range t {
		if tt.Key == key {
			return true
		}
	}
	return false
}

func isValidTag(s string) bool {
	if s == "" {
		return false
	}
	for _, c := range s {
		switch {
		case strings.ContainsRune("!#$%&()*+-./:<=>?@[]^_{|}~ ", c):
			// Backslash and quote chars are reserved, but
			// otherwise any punctuation chars are allowed
			// in a tag name.
		case !unicode.IsLetter(c) && !unicode.IsDigit(c):
			return false
		}
	}
	return true
}

func StructTagFromField(field reflect.StructField, eflag *encode.OptionFlag, dflag *decode.OptionFlag) *StructTag {
	var (
		keyName = field.Name
		st      = &StructTag{Field: field}
		opts    = strings.Split(getTag(field), ",")
	)

	if len(opts) > 0 {
		if opts[0] != "" && isValidTag(opts[0]) {
			keyName = opts[0]
			st.IsTaggedKey = true
		}
	}
	st.Key = keyName

	if len(opts) > 1 {
		for _, opt := range opts[1:] {
			switch opt {
			case "omitempty":
				st.IsOmitEmpty = true
			case "string":
				st.IsString = true
			}
		}
	}

	if !st.IsString {
		kind := field.Type.Kind()
		if (kind == reflect.Int64 || kind == reflect.Uint64) &&
			((eflag != nil && (encode.Int64ToStringOption&(*eflag)) != 0) ||
				(dflag != nil && (decode.Int64ToStringOption&(*dflag)) != 0)) {
			st.IsString = true
		}
	}

	if st.IsOmitEmpty &&
		((eflag != nil && (encode.FieldNoOmitEmpty&(*eflag)) != 0) ||
			(dflag != nil && (decode.FieldNoOmitEmpty&(*dflag)) != 0)) {
		st.IsOmitEmpty = false
	}
	return st
}
