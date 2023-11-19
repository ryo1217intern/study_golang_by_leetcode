/*
 * @lc app=leetcode id=1768 lang=golang
 *
 * [1768] Merge Strings Alternately
 */
package code

// @lc code=start
func mergeAlternately(word1 string, word2 string) string {

	//仮にword1が最小の文字数として最小の文字数を代入
	minLength := len(word1)
	//残りのテキストを初期化
	remainText := ""
	//リターンする値の初期化
	resultText := ""

	//word2が最小の文字の際の処理を記述
	if len(word1) > len(word2) {
		minLength = len(word2)
		//ここで初めて残りのテキストに文字列を代入する
		remainText = word1[minLength:]
	} else {
		//word1が最小の場合remainTextはword2の残りとなる
		remainText = word2[minLength:]
	}

	//文字を順番に足していく
	for i := 0; i < minLength; i++ {
		resultText += string(word1[i]) + string(word2[i])
	}

	//残った文字をresultTextに追加する
	resultText += remainText

	return resultText
}
// @lc code=end

