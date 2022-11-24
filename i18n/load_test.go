package i18n_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/oils/i18n"
	"golang.org/x/text/language"
)

func TestTag(t *testing.T) {
	t.Parallel()

	tag, err := i18n.Tag("zh-Hans", "a.toml")

	assert.Nil(t, err)
	assert.Equal(t, language.SimplifiedChinese.String(), tag.String())

	tag, err = i18n.Tag("a_zh-Hans.toml")

	assert.Nil(t, err)
	assert.Equal(t, language.SimplifiedChinese.String(), tag.String())
}

func TestLoad(t *testing.T) {
	t.Parallel()

	assert.Nil(t, i18n.Load(locales))
}
