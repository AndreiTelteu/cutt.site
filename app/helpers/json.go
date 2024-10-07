package helpers

import "encoding/json"

func PrettyJson(v interface{}) string {
	if v == nil {
		return ""
	}
	if b, err := json.MarshalIndent(v, "", "  "); err == nil {
		return string(b)
	}
	return ""
}

func Json(v interface{}) string {
	if v == nil {
		return ""
	}
	if b, err := json.Marshal(v); err == nil {
		return string(b)
	}
	return ""
}
