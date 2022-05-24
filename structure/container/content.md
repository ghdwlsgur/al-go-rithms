### container/list 패키지에서 제공하는 연결리스트 함수

- func New() \*List `연결 리스트 생성`
- func (l *List) PushBack(v interface{}) *Element `연결 리스트 맨 뒤 데이터 추가`
- func (l *List) Front() *Element `연결 리스트의 맨 앞 데이터를 가져옴`
- func (l *List) Back() *Element `연결 리스트의 맨 뒤 데이터를 가져옴`

Go 언어의 연결 리스트는 이중 연결 리스트(Doubly linked list)로서 앞 뒤 방향으로 순회할 수 있으며 list.New 함수로 연결리스트를 생성한 뒤 PushBack 등의 함수로 데이터를 추가한다.

| 함수          | 설명                                 |
| :------------ | :----------------------------------- |
| PushBack      | 맨 뒤에 노드를 추가합니다.           |
| PushFront     | 맨 앞에 노드를 추가합니다.           |
| PushBackList  | 맨 뒤에 다른 리스트를 붙입니다.      |
| PushFrontList | 맨 앞에 다른 리스트를 붙입니다.      |
| InsertAfter   | 특정 노드 뒤에 노드를 추가합니다.    |
| InsertBefore  | 특정 노드 앞에 노드를 추가합니다.    |
| MoveAfter     | 노드를 특정 노드 뒤로 옮깁니다.      |
| MoveBefore    | 노드를 특정 노드 앞으로 옮깁니다.    |
| MoveToFront   | 노드를 맨 앞으로 옮깁니다.           |
| Len           | 리스트의 노드 개수(길이)를 구합니다. |
| Remove        | 특정 노드를 삭제합니다.              |
