### new

> `new()`는 메모리를 할당하되 초기화는 하지 않는다. 메모리를 할당하고 해당 객체에 제로(Zero Value)값을 설정하고 해당 객체에 대한 포인터를 반환하게 된다. Go에서는 Zeroed Storage라고 표현하는 듯하며 `new()`를 사용하여 반환된 값은 포인터이고, 해당 포인터가 가리키는 값은 각 타입에 대한 제로값이다.

```Go
m := new(MyType)
fmt.Printf("%T, %#v", m, m) // *main.MyType &main.MyType{}
```

### new vs make

> `new()`는 메모리를 할당하는데, 어떤 데이터 타입에 대해서는 할당으로는 부족하며 미리 사용할 준비가 되도록 초기화를 시켜주어야 하는 경우가 있는데 그럴때 `make()`함수를 쓴다. `make()`에 사용가능한 껏은 슬라이스, 맵, 채널이며 `new()`는 포인터를 반환하지만 `make()`는 포인터를 반환하지 않는다.

```Go
var intSlice []int
fmt.Printf("%T, %#v", intSlice, intSlice) // []int []int(nil)
```

```Go
intSlice = new([]int)
fmt.Printf("%T, %#v", intSlice, intSlice) // *[]int &[]int(nil)
```

```Go
intSlice := make([]int, 5)
fmt.Printf("%T, %#v", intSlice, intSlice) // []int []int{0, 0, 0, 0, 0}
```

### 배열과 슬라이스 차이

|           | 배열(array) | 슬라이스(slice) |
| :-------: | :---------: | :-------------: |
|   타입    |     값      |      참조       |
| 용량(cap) |  수정불가   |    수정가능     |
| 길이(len) |  수정불가   |    수정가능     |
| 비교연산  |    가능     |     불가능      |
|   호출    |  복사전달   |    참조전달     |

- `cap()`: 배열/슬라이스 용량
- `len()`: 배열/슬라이스 개수
