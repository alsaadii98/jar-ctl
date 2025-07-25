package jar

type JarEntry struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Path    string `json:"path"`
	AddedAt string `json:"added_at"`
}
