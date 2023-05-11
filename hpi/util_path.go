package hpi

import "strings"

// GetNameFromPath is function to get only name from full path
func (c *Transform) GetNameFromPath(path string) (name string) {

	// remove full path
	name = path
	if strings.Contains(path, ".") {
		i := strings.LastIndex(path, ".") + 1
		if i < len(path) {
			name = path[i:]
		}
	}

	// remove array [i]
	if strings.Contains(name, "[") {
		i := strings.Index(name, "[")
		name = name[:i]
	}

	return
}

// GetParentFromPath is function to get only parent from full path
func (c *Transform) GetParentFromPath(path string) (parent string) {

	// Parent is root
	if !strings.Contains(path, ".") {
		return ""
	} else {

		// Finding parent
		i := strings.LastIndex(path, ".")
		if i < len(path) && i >= 0 {
			parent = path[0:i]
		}
		return parent
	}

}
