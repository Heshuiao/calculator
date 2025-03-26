package main

import (
	"fmt"
	"strings"
	"unicode"
	"strconv"
	"math"
)
//定义运算符优先级
var level = map[rune]int{
	'+':1,
	'-':1,
	'*':2,
	'/':2,
	'^':3,
}
//判断是否为可用运算符
func ifOperator(op rune) bool {
	if level[op] > 0 || op == 'l' || op == 'k' { //先这么测试下
		return true
	}
	return false
}
// 中缀转后缀
func infixToPostfix(str string) []string {
	var stack []string       // 符号栈
	var result []string      // 后缀表达式
	num := strings.Builder{} // 处理多位数

	i := 0
	for i < len(str) {
		ch := rune(str[i])
		if unicode.IsDigit(ch) || ch == '.' {
			num.WriteRune(ch)
		} else {
			if num.Len() > 0 { // 先把完整数字加入
				result = append(result, num.String())
				num.Reset()
			}
			if ch == '(' { // 处理 (
				stack = append(stack, string(ch))
			} else if ch == ')' { // 处理 )
				for len(stack) > 0 && stack[len(stack)-1] != "(" {
					result = append(result, stack[len(stack)-1])
					stack = stack[:len(stack)-1]
				}
				if len(stack) > 0 { // 弹出左括号
					stack = stack[:len(stack)-1]
				}
			} else if i+1 < len(str) && str[i:i+2] == "ln" { // 处理 ln
				stack = append(stack, "ln")
				i++ // 额外跳过 'n'
			} else if i+1 < len(str) && str[i:i+2] == "kf" { // 处理 kf
				stack = append(stack, "kf")
				i++ // 额外跳过 'f'
			} else if ifOperator(ch) {
				for len(stack) > 0 && stack[len(stack)-1] != "(" && level[ch] <= level[rune(stack[len(stack)-1][0])] {
					result = append(result, stack[len(stack)-1])
					stack = stack[:len(stack)-1]
				}
				stack = append(stack, string(ch))
			}
		}
		i++
	}
	if num.Len() > 0 { // 处理最后可能剩下的数字
		result = append(result, num.String())
	}
	for len(stack) > 0 {
		result = append(result, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}
	return result
}
//判断是否是数字并转换
func ifNum (str string) (bool,float64){
	num,err := strconv.ParseFloat(str,64)
	return err==nil,num
}
//识别操作符
func whichOp(op string, a float64, b float64) float64 {
	switch op {
	case "*":
		return b * a
	case "/":
		return b / a
	case "+":
		return b + a
	case "-":
		return b - a
	case "^":
		return math.Pow(b, a)
	case "ln":
		return math.Log(a) // 自然对数
	case "kf":
		return math.Sqrt(a) // 开平方
	}
	return 0
}
//计算函数
func calculate(str []string) float64 {
	var num_stack []float64
	for _, s := range str {
		isNum, num := ifNum(s)
		if isNum { // 数字直接入栈
			num_stack = append(num_stack, num)
		} else { // 处理运算符
			if s == "ln" || s == "kf" { // 处理单目运算符
				a := num_stack[len(num_stack)-1]
				num_stack = num_stack[:len(num_stack)-1]
				num_stack = append(num_stack, whichOp(s, a, 0)) // 传入a, 忽略b
			} else { // 处理双目运算符
				a := num_stack[len(num_stack)-1]
				num_stack = num_stack[:len(num_stack)-1]
				b := num_stack[len(num_stack)-1]
				num_stack = num_stack[:len(num_stack)-1]
				num_stack = append(num_stack, whichOp(s, a, b))
			}
		}
	}
	return num_stack[0]
}
func main() {
	var input string
	fmt.Println("请输入计算式：")
	fmt.Scanln(&input)
	Postfix := infixToPostfix(input)
	anwser := calculate(Postfix)
	fmt.Printf("结果是：%f",anwser)
}