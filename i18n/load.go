package i18n

import (
	"embed"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func Load(locales embed.FS) error { return LoadDir(locales, ".") }

func LoadDir(locales embed.FS, dir string) error {
	files, err := locales.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			if err := LoadDir(locales, filepath.Join(dir, file.Name())); err != nil {
				return err
			}
		} else {
			if err := LoadFile(locales, dir, file.Name()); err != nil {
				return err
			}
		}
	}

	return nil
}

func LoadFile(locales embed.FS, paths ...string) error {
	data := map[string]string{}
	if _, err := toml.DecodeFS(locales, filepath.Join(paths...), &data); err != nil {
		return err
	}

	if tag, err := Tag(paths...); err == nil {
		if err := SetString(tag, data); err != nil {
			return err
		}
	}

	return nil
}

func Tag(paths ...string) (language.Tag, error) {
	for _, path := range paths {
		base := filepath.Base(path)

		if tag, err := language.Parse(base); err == nil {
			return tag, err
		}
	}

	file := strings.Split(paths[len(paths)-1], ".")[0]
	index := strings.Index(file, "_")

	return language.Parse(file[index+1:])
}

func SetString(tag language.Tag, data map[string]string) error {
	for key, value := range data {
		if err := message.SetString(tag, key, value); err != nil {
			return err
		}
	}

	return nil
}
