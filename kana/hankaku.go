package kana

import "errors"

var kanaMap = map[rune]rune{
	'ｱ': 'ア',
	'ｲ': 'イ',
	'ｳ': 'ウ',
	'ｴ': 'エ',
	'ｵ': 'オ',

	'ｶ': 'カ',
	'ｷ': 'キ',
	'ｸ': 'ク',
	'ｹ': 'ケ',
	'ｺ': 'コ',

	'ｻ': 'サ',
	'ｼ': 'シ',
	'ｽ': 'ス',
	'ｾ': 'セ',
	'ｿ': 'ソ',

	'ﾀ': 'タ',
	'ﾁ': 'チ',
	'ﾂ': 'ツ',
	'ﾃ': 'テ',
	'ﾄ': 'ト',

	'ﾅ': 'ナ',
	'ﾆ': 'ニ',
	'ﾇ': 'ヌ',
	'ﾈ': 'ネ',
	'ﾉ': 'ノ',

	'ﾊ': 'ハ',
	'ﾋ': 'ヒ',
	'ﾌ': 'フ',
	'ﾍ': 'ヘ',
	'ﾎ': 'ホ',

	'ﾏ': 'マ',
	'ﾐ': 'ミ',
	'ﾑ': 'ム',
	'ﾒ': 'メ',
	'ﾓ': 'モ',

	'ﾔ': 'ヤ',
	'ﾕ': 'ユ',
	'ﾖ': 'ヨ',

	'ﾗ': 'ラ',
	'ﾘ': 'リ',
	'ﾙ': 'ル',
	'ﾚ': 'レ',
	'ﾛ': 'ロ',

	'ﾜ': 'ワ',
	'ｦ': 'ヲ',
	'ﾝ': 'ン',

	'ｧ': 'ァ',
	'ｨ': 'ィ',
	'ｩ': 'ゥ',
	'ｪ': 'ェ',
	'ｫ': 'ォ',

	'ｯ': 'ッ',

	'ｬ': 'ャ',
	'ｭ': 'ュ',
	'ｮ': 'ョ',
}

var dakuonMap = map[rune]rune{
	'ｶ': 'ガ',
	'ｷ': 'ギ',
	'ｸ': 'グ',
	'ｹ': 'ゲ',
	'ｺ': 'ゴ',

	'ｻ': 'ザ',
	'ｼ': 'ジ',
	'ｽ': 'ズ',
	'ｾ': 'ゼ',
	'ｿ': 'ゾ',

	'ﾀ': 'ダ',
	'ﾁ': 'ヂ',
	'ﾂ': 'ヅ',
	'ﾃ': 'デ',
	'ﾄ': 'ド',

	'ﾊ': 'バ',
	'ﾋ': 'ビ',
	'ﾌ': 'ブ',
	'ﾍ': 'ベ',
	'ﾎ': 'ボ',

	'ｳ': 'ヴ',
	'ﾜ': 'ヷ',
	'ｦ': 'ヺ',
}

var handakuonMap = map[rune]rune{
	'ﾊ': 'パ',
	'ﾋ': 'ピ',
	'ﾌ': 'プ',
	'ﾍ': 'ペ',
	'ﾎ': 'ポ',
}

func ToZenkaku(src string) (string, error) {
	chars := []rune(src)
	length := len(chars)
	var results []rune

	for i := 0; i < length; i++ {
		c := chars[i]
		if i+1 < length {
			n := chars[i+1]

			if n == 'ﾞ' {
				// 濁点
				if d, ok := dakuonMap[c]; ok {
					results = append(results, d)
					i++
					continue
				} else {
					return "", errors.New("invalid dakuon")
				}
			}

			if n == 'ﾟ' {
				// 半濁点
				if d, ok := handakuonMap[c]; ok {
					results = append(results, d)
					i++
					continue
				} else {
					return "", errors.New("invalid handakuon")
				}
			}
		}

		if c, ok := kanaMap[c]; ok {
			results = append(results, c)
		} else {
			return "", errors.New("invalid character")
		}
	}

	return string(results), nil
}
