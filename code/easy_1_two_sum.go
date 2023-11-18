/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */
package code

// @lc code=start
func twoSum(nums []int, target int) []int {
	//辞書型配列の作成.要素はまだ入っていない.
	numsMap := make(map[int]int)

	//for rangeによってnumsの要素とそのindexを抽出しつつ要素の数だけfor文を回す.
	for numsIndex, numsValue := range nums {
		//numsMapに登録された値と一致するnumsMap[target-numsValue]が存在する場合numsMap[target-numsValue]は1を出力する.ない時は0を出力する.
		if numsMap[target - numsValue] != 0 {
			return []int{numsMap[target - numsValue] - 1, numsIndex}
		}
		//if文の条件式がfalseの際はnumsMapにその値を登録する.
		numsMap[numsValue] = numsIndex + 1
	}

	return 	[]int{}

}
// @lc code=end