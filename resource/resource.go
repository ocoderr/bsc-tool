package resource

import (
	"fyne.io/fyne/v2"
	"io/ioutil"
	"path"
	"path/filepath"
)

// GetResourceW get resource
func GetResourceW(fn string, fallback fyne.Resource) fyne.Resource {
	out := GetResource(fn)
	if out == nil {
		return fallback
	}
	return out
}

// GetResource returns a resource by name
func GetResource(fn string) fyne.Resource {
	bytes := assetRead(fn)
	if len(bytes) == 0 {
		return nil
	}
	name := filepath.Base(fn)
	return fyne.NewStaticResource(name, bytes)
}

func assetRead(fn string) []byte {
	bytes, err := ioutil.ReadFile(path.Join("assets", fn))
	if err != nil {
		return nil
	}
	return bytes
}
