package simple

import "strings"

// mostWordsFound 2114. 句子中的最多单词数
func mostWordsFound(sentences []string) int {
	ret := 1
	for i := range sentences {
		cnt := strings.Count(sentences[i], " ")
		ret = maxMostWordsFound(cnt, ret)

	}
	return ret+1
}

func maxMostWordsFound(a, b int) int {
	if a > b {
		return a
	}
	return b
}
