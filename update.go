package goskyline

type Update struct {
	ID             uint   `jsonapi:"primary,updates" json:"ID"`
	PackageName    string `jsonapi:"attr,packageName" json:"packageName"`
	CurrentVersion string `jsonapi:"attr,currentVersion" json:"currentVersion"`
	NewVersion     string `jsonapi:"attr,newVersion" json:"newVersion"`
	Repository     string `jsonapi:"attr,repository" json:"repository"`
	Security       bool   `jsonapi:"attr,security" json:"security"`
}
