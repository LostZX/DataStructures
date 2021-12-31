package main

import "fmt"

// 思考：没有系统学习过计算机基础的人写出来就是这样的，各种不通过，一道简单题改了6次，而且代码质量相当差 一不够健壮，二可维护性极差
// 解决：多看其他人的解题思路，找到通用的地方并学习解决

func isValid(s string) bool {
	flag := false
	tokenMap := make(map[string]string)
	stack := make([]string, 0)
	tokenMap["("] = ")"
	tokenMap["{"] = "}"
	tokenMap["["] = "]"

	// token必定成双出现才算正确
	if len(s)%2 != 0 {
		return false
	}

	for i := 0; i < len(s); i++ {
		c := rune(s[i])
		if string(c) == "{" || string(c) == "[" || string(c) == "(" {
			stack = append(stack, string(c))
		} else if len(stack) == 0 { // 增加 "]" 的判断
			flag = false
			break
		} else {
			top := stack[len(stack)-1]
			if tokenMap[top] == string(c) {
				flag = true
				stack = stack[:len(stack)-1]
			} else {
				// 如果匹配到一处错误，直接退出
				flag = false
				break
			}
		}
	}

	// 循环完如果还存在，则是错误的
	if len(stack) > 0 {
		return false
	}

	return flag
}

func main() {
	flag := isValid("([]){")
	fmt.Println(flag)
}
