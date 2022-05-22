package i18n_test

import (
	"strings"
	"testing"

	"github.com/xuender/oils/assert"
	"github.com/xuender/oils/i18n"
	"golang.org/x/text/language"
)

func TestGetTag(t *testing.T) {
	t.Parallel()

	langs := strings.Split("zh_CN:en_US:en", ":")

	assert.Equal(t, "zh-CN", (i18n.GetTag(langs).String()))
	assert.Equal(t, "zh-CN", (i18n.GetTag([]string{"zh-CN"}).String()))
	assert.Equal(t, "zh-Hans", (i18n.GetTag([]string{"zh-Hans"}, language.SimplifiedChinese).String()))
	assert.Equal(t, "zh-Hans", (i18n.GetTag(langs, language.SimplifiedChinese).String()))
	assert.Equal(t, "zh-Hans", (i18n.GetTag([]string{"en"}, language.SimplifiedChinese).String()))
	assert.Equal(t, "zh-Hans", (i18n.GetTag(langs, language.SimplifiedChinese, language.English).String()))
	assert.Equal(t, "en", (i18n.GetTag([]string{}).String()))
	assert.Equal(t, "zh", (i18n.GetTag([]string{}, language.Chinese).String()))
}
