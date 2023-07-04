package twosum

func twoSum(nums []int, target int) []int {
	result := []int{}

	m := make(map[int]int)
	
	for key, value := range nums {
		
		//	可以取得map值並判斷是否存在 結果回傳在exist
		j, exist := m[target-value]

		if exist {
			return []int{j, key}
		}

		m[value] = key
	}

    return result
}