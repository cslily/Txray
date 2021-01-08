package tools

import (
	"strings"
)

// 解析索引参数(关键字从1开始，返回的结果从0开始)
func IndexDeal(key string, length int) []int {
	if length == 0 {
		return []int{}
	}
	// 处理关键字 all
	if key == "all" {
		result := make([]int, length, length)
		for i := 0; i < length; i++ {
			result[i] = i
		}
		return result
	}

	indexList := make([]bool, length, length)
	for _, item := range strings.Split(key, ",") {
		item = strings.Trim(item, " ")
		if strings.Contains(item, "-") {
			k := strings.Split(item, "-")
			if len(k) == 2 {
				start, end := 1, length
				if k[0] == "" && IsUint(k[1]) {
					end = StrToInt(k[1])
				} else if IsUint(k[0]) && k[1] == "" {
					start = StrToInt(k[0])
				} else if IsUint(k[0]) && IsUint(k[1]) {
					start, end = StrToInt(k[0]), StrToInt(k[1])
					if start > end {
						start, end = end, start
					}
				}
				// 处理越界
				if start < 1 {
					start = 1
				}
				if end > length {
					end = length
				}
				for i := start - 1; i < end; i++ {
					indexList[i] = true
				}
			}
		} else if IsUint(item) {
			i := StrToInt(item)
			if i > 0 && i <= length {
				indexList[i-1] = true
			}
		}
	}

	// 生成索引数组（从0开始）
	result := make([]int, 0, length)
	for i, isSelect := range indexList {
		if isSelect {
			result = append(result, i)
		}
	}
	return result
}
