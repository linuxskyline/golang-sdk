package goskyline

type Update struct {
	PackageName    string `json:"packageName"`
	CurrentVersion string `json:"currentVersion"`
	NewVersion     string `json:"newVersion"`
	Security       bool   `json:"security"`
}
