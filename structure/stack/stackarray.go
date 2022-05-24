package stack

var stackArray []any

// 요소 추가
func stackPush(n any) {
	stackArray = append([]any{n}, stackArray...)
}

// 스택 길이 반환
func stackLength() int {
	return len(stackArray)
}

// 맨 앞 요소 반환
func stackPeak() any {
	return stackArray[0]
}

// 스택의 길이가 0일 경우 true, 그렇지 않을 경우 false 반환
func stackEmpty() bool {
	return len(stackArray) == 0
}

// 맨 앞 요소 제거
func stackPop() any {
	pop := stackArray[0]
	stackArray = stackArray[1:]
	return pop
}
