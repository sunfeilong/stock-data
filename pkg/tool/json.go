package tool

import "encoding/json"

func JSONToString(data interface{}) string {
    if marshal, err := json.Marshal(data); err == nil {
        return string(marshal)
    }
    return ""
}
