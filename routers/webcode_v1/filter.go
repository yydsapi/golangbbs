// Copyright 2019 golangbbs Core Team.  All rights reserved.
// LICENSE: Use of this source code is governed by AGPL-3.0.
// license that can be found in the LICENSE file.
package webcodev1

/*
DFA：Determine finite automata.
Specific functions:
Keep sensitive words in map.
The sensitive words are filtered and the sensitive words are changed to "*".
Neglect meaningless symbols.
Sensitive Word Data Structure：
{  s：{
            isEnd: false
             e：{
                    isEnd：false
                    x：{
                              isEnd：true
                       }
                 }
       }
}
*/
//Generating a set of forbidden words
func AddSensitiveToMap(set map[string]interface{}) {
	for key := range set {
		str := []rune(key)
		nowMap := SensitiveWord
		for i := 0; i < len(str); i++ {
			if _, ok := nowMap[string(str[i])]; !ok { //If the key does not exist，
				thisMap := make(map[string]interface{})
				thisMap["isEnd"] = false
				nowMap[string(str[i])] = thisMap
				nowMap = thisMap
			} else {
				nowMap = nowMap[string(str[i])].(map[string]interface{})
			}

			if i == len(str)-1 {
				nowMap["isEnd"] = true
			}
		}

	}
}

// Conversion of sensitive words into*
func ChangeSensitiveWords(txt string, sensitive map[string]interface{}) (word string) {
	str := []rune(txt)
	nowMap := sensitive
	//logrus.Info(nowMap)
	start := -1
	tag := -1
	for i := 0; i < len(str); i++ {
		//logrus.Info(i,"------->"+string(str[i]))
		if _, ok := InvalidWord[(string(str[i]))]; ok {
			//logrus.Info("If the word is invalid, skip it directly")
			continue
		}
		if thisMap, ok := nowMap[string(str[i])].(map[string]interface{}); ok {
			//logrus.Info("The first character of a sensitive word")
			tag++
			if tag == 0 {
				start = i

			}
			//Determine whether it is the last word of a sensitive word
			if isEnd, _ := thisMap["isEnd"].(bool); isEnd {
				//Replace all sensitive words from the first to the last with *
				for y := start; y < i+1; y++ {
					str[y] = 42
					HasSensitiveWord = true
				}
				//Reset flag data
				nowMap = sensitive
				start = -1
				tag = -1

			} else { //If not the last one, assign the map it contains to nowMap
				nowMap = nowMap[string(str[i])].(map[string]interface{})
			}

		} else { //If the sensitive word is not fully matched, the sensitive word search is terminated. Continue judging from the second word at the beginning
			//logrus.Info("If the sensitive word is not fully matched, the sensitive word search is terminated. Continue judging from the second word at the beginning")
			if start != -1 {
				i = start
				//i = start + 1
			}
			//Reset flag data
			nowMap = sensitive
			start = -1
			tag = -1
		}
	}

	return string(str)
}
