package kana

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToZenkaku(t *testing.T) {
	data := map[string]string{
		"": "",

		"ｱｲｳｴｵ": "アイウエオ",
		"ｶｷｸｹｺ": "カキクケコ",
		"ｻｼｽｾｿ": "サシスセソ",
		"ﾀﾁﾂﾃﾄ": "タチツテト",
		"ﾅﾆﾇﾈﾉ": "ナニヌネノ",
		"ﾊﾋﾌﾍﾎ": "ハヒフヘホ",
		"ﾏﾐﾑﾒﾓ": "マミムメモ",
		"ﾔﾕﾖ":   "ヤユヨ",
		"ﾗﾘﾙﾚﾛ": "ラリルレロ",
		"ﾜｦﾝ":   "ワヲン",

		"ｶﾞｷﾞｸﾞｹﾞｺﾞ": "ガギグゲゴ",
		"ｻﾞｼﾞｽﾞｾﾞｿﾞ": "ザジズゼゾ",
		"ﾀﾞﾁﾞﾂﾞﾃﾞﾄﾞ": "ダヂヅデド",
		"ﾊﾞﾋﾞﾌﾞﾍﾞﾎﾞ": "バビブベボ",

		"ﾊﾟﾋﾟﾌﾟﾍﾟﾎﾟ": "パピプペポ",

		"ｬｭｮ":    "ャュョ",
		"ｧｨｩｪｫ":  "ァィゥェォ",
		"ｳﾞﾜﾞｦﾞ": "ヴヷヺ",
	}

	for k, v := range data {
		z, err := ToZenkaku(k)
		assert.Nil(t, err)
		assert.Equal(t, v, z)
	}
}
