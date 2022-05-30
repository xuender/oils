package i18n

import (
	"os"
	"strings"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func GetPrinter(tags ...language.Tag) *message.Printer {
	return message.NewPrinter(GetTag(strings.Split(os.Getenv("LANGUAGE"), ":"), tags...))
}

func GetTag(langs []string, tags ...language.Tag) language.Tag {
	if len(tags) > 0 {
		if len(langs) == 0 {
			return tags[0]
		}

		for _, tag := range tags {
			str := tag.String()
			local := strings.Split(tag.String(), "-")[0]

			for _, lang := range langs {
				if lang == str {
					return tag
				}

				if strings.Split(lang, "_")[0] == local {
					return tag
				}
			}
		}

		return tags[0]
	}

	for _, lang := range langs {
		if tag, err := language.Parse(lang); err == nil {
			return tag
		}
	}

	return language.English
}
