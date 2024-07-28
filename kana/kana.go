package kana

// ToHiragana テキストに含まれる全てのカタカナを平仮名に変換します
func ToHiragana(source string) string {
	chars := []rune(source)
	length := len(chars)
	results := make([]rune, length)

	for i := 0; i < length; i++ {
		c := chars[i]

		if c >= 'ァ' && c <= 'ヶ' {
			results[i] = c - 'ァ' + 'ぁ'
		} else {
			results[i] = c
		}
	}

	return string(results)
}

// ToKatakana テキストに含まれる全てのひらがなをカタカナに変換します
func ToKatakana(source string) string {
	chars := []rune(source)
	length := len(chars)
	results := make([]rune, length)

	for i := 0; i < length; i++ {
		c := chars[i]

		if c >= 'ぁ' && c <= 'ゖ' {
			results[i] = c - 'ぁ' + 'ァ'
		} else {
			results[i] = c
		}
	}

	return string(results)
}

// IsAllHiragana テキストに含まれる全ての文字がひらがなかどうか
func IsAllHiragana(text string) bool {
	chars := []rune(text)
	length := len(chars)

	for i := 0; i < length; i++ {
		c := chars[i]

		if (c < 'ぁ' || c > 'ゖ') && 'ー' != c {
			return false
		}
	}

	return true
}

// IsAllKatakana テキストに含まれる全ての文字がカタカナかどうか
func IsAllKatakana(text string) bool {
	chars := []rune(text)
	length := len(chars)

	for i := 0; i < length; i++ {
		c := chars[i]

		if (c < 'ァ' || c > 'ヺ') && 'ー' != c {
			return false
		}
	}

	return true
}
