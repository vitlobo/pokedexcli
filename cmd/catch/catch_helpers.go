package catch

import (
	"strings"
)

//
// ====== Helper Functions ======
//

func getFlavor(entries []struct {
    FlavorText string `json:"flavor_text"`
    Language   struct{ Name string `json:"name"` } `json:"language"`
    Version    struct{ Name string `json:"name"` } `json:"version"`
}) string {
    for i := len(entries) - 1; i >= 0; i-- {
        if entries[i].Language.Name == "en" {
            s := strings.ReplaceAll(entries[i].FlavorText, "\n", " ")
            s = strings.ReplaceAll(s, "\f", " ")
            return strings.TrimSpace(s)
        }
    }
    return ""
}