package kana

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToHiragana(t *testing.T) {
	data := map[string]string{
		"": "",

		"あいうえお": "アイウエオ",
		"かきくけこ": "カキクケコ",
		"さしすせそ": "サシスセソ",
		"たちつてと": "タチツテト",
		"はひふへほ": "ハヒフヘホ",
		"まみむめも": "マミムメモ",
		"やゆよ":   "ヤユヨ",
		"らりるれろ": "ラリルレロ",
		"わをん":   "ワヲン",

		"がぎぐげご": "ガギグゲゴ",
		"ざじずぜぞ": "ザジズゼゾ",
		"だぢづでど": "ダヂヅデド",
		"ばびぶべぼ": "バビブベボ",
		"ぱぴぷぺぽ": "パピプペポ",

		"ゃゅょぁぃぅぇぉっ": "ャュョァィゥェォッ",

		"ゔ": "ヴ",
	}

	for h, k := range data {
		assert.Equal(t, h, ToHiragana(k))
		assert.Equal(t, k, ToKatakana(h))
	}
}

func TestIsAllHiragana(t *testing.T) {
	assert.True(t, IsAllHiragana(""))

	assert.True(t, IsAllHiragana("あいうえお"))
	assert.True(t, IsAllHiragana("かきくけこ"))
	assert.True(t, IsAllHiragana("さしすせそ"))
	assert.True(t, IsAllHiragana("たちつてと"))
	assert.True(t, IsAllHiragana("なにぬねの"))
	assert.True(t, IsAllHiragana("はひふへほ"))
	assert.True(t, IsAllHiragana("まみむめも"))
	assert.True(t, IsAllHiragana("やゆよ"))
	assert.True(t, IsAllHiragana("らりるれろ"))
	assert.True(t, IsAllHiragana("わをん"))

	assert.True(t, IsAllHiragana("ろーま"))

	assert.False(t, IsAllHiragana("漢字あいうえお"))
	assert.False(t, IsAllHiragana("あいうえお漢字"))
	assert.False(t, IsAllHiragana("漢字あいうえお漢字"))

	assert.False(t, IsAllHiragana("あいうえおアイウエオ"))
	assert.False(t, IsAllHiragana("アイウエオあいうえお"))
	assert.False(t, IsAllHiragana("アイウエオあいうえおアイウエオ"))
}

func TestIsAllKatakana(t *testing.T) {
	assert.True(t, IsAllKatakana(""))

	assert.True(t, IsAllKatakana("アイウエオ"))
	assert.True(t, IsAllKatakana("カキクケコ"))
	assert.True(t, IsAllKatakana("サシスセソ"))
	assert.True(t, IsAllKatakana("タチツテト"))
	assert.True(t, IsAllKatakana("ナニヌネノ"))
	assert.True(t, IsAllKatakana("ハヒフヘホ"))
	assert.True(t, IsAllKatakana("マミムメモ"))
	assert.True(t, IsAllKatakana("ヤユヨ"))
	assert.True(t, IsAllKatakana("ラリルレロ"))
	assert.True(t, IsAllKatakana("ワヲン"))

	assert.True(t, IsAllKatakana("ローマ"))

	assert.False(t, IsAllKatakana("漢字アイウエオ"))
	assert.False(t, IsAllKatakana("アイウエオ漢字"))
	assert.False(t, IsAllKatakana("漢字アイウエオ漢字"))

	assert.False(t, IsAllKatakana("あいうえおアイウエオ"))
	assert.False(t, IsAllKatakana("アイウエオあいうえお"))
	assert.False(t, IsAllKatakana("あいうえおアイウエオあいうえお"))
}
