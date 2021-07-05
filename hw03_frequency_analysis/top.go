package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

const topLimit = 10

type wordsCount struct {
	count int
	word  string
}

var prepareReplaceRegExp = regexp.MustCompile(`([^\p{L}]+-[^\p{L}]+|[^\p{L}\-]+)`)

func Top10(str string) []string {
	res := make([]string, 0, topLimit)

	preparedStr := prepareString(str)
	wordsSlice := strings.Fields(preparedStr)

	if len(wordsSlice) > 0 {
		wordsCnt := countWords(wordsSlice)

		sort.Slice(wordsCnt, func(i, j int) bool {
			if wordsCnt[i].count != wordsCnt[j].count {
				return wordsCnt[i].count > wordsCnt[j].count
			}
			return wordsCnt[i].word < wordsCnt[j].word
		})

		for i, v := range wordsCnt {
			if i < topLimit {
				res = append(res, v.word)
			}
		}
	}

	return res
}

func prepareString(str string) string {
	preparedStr := prepareReplaceRegExp.ReplaceAllString(str, " ")
	preparedStr = strings.ToLower(preparedStr)
	return preparedStr
}

func countWords(words []string) []wordsCount {
	indexMap := make(map[string]int)
	wordsCnt := make([]wordsCount, 0, len(words))

	for _, v := range words {
		if index, ok := indexMap[v]; ok {
			wordsCnt[index].count++
		} else {
			wordsCnt = append(wordsCnt, wordsCount{
				count: 1,
				word:  v,
			})
			indexMap[v] = len(wordsCnt) - 1
		}
	}
	return wordsCnt
}
