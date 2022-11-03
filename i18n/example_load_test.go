package i18n_test

import (
	"embed"

	"github.com/xuender/oils/i18n"
)

//go:embed locales
var locales embed.FS

func ExampleLoad() {
	_ = i18n.Load(locales)
	tag, _ := i18n.Tag("zh_CN")
	printer := i18n.GetPrinter(tag)

	printer.Printf("hellohometown")

	// Output:
	// 你好家乡!
}
