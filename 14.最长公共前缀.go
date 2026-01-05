import "sort"

/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {

	type KV struct {
		k string
		v int
	}

	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	kvs := make([]KV, 0, len(strs))

	for _, v := range strs {
		kvs = append(kvs, KV{k: v, v: len(v)})
	}

	// 按照值排序
	// sort.Slice(slice, lessFunc)
	sort.Slice(kvs, func(i, j int) bool {
		if kvs[i].v != kvs[j].v {
			return kvs[i].v < kvs[j].v
		}
		return kvs[i].k < kvs[j].k
	})

	var sign bool
	count := 0
	// fmt.Println(kvs)
	base := []rune(kvs[0].k)

	for i, rv := range base {
		for _, item := range kvs[1:] {
			r := []rune(item.k)
			if r[i] != rv {
				sign = true
				count = i
				break
			}
		}
		if sign {
			break
		}
		count = i + 1
	}

	return kvs[0].k[:count]
}

// @lc code=end
