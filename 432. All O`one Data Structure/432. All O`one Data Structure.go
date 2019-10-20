package leetcode

import (
	"container/list"
)

// 使用 map + list.List
type AllOne struct {
	keyCountMap map[string]*list.Element
	countList   *list.List
}

// 作为 list.Element 的 Value 成员
type elementValue struct {
	value   int
	keysMap map[string]struct{}
}

func Constructor() AllOne {
	return AllOne{make(map[string]*list.Element), list.New()}
}

func (this *AllOne) Inc(key string) {
	element, ok := this.keyCountMap[key]
	if !ok {
		// key不存在，则将key加入值为1的element
		back := this.countList.Back()
		if back != nil && back.Value.(*elementValue).value == 1 {
			// 存在值为1的element，把key添加进去
			back.Value.(*elementValue).keysMap[key] = struct{}{}
		} else {
			// 不存在值为1的element，list末尾添加值为1的element
			this.countList.PushBack(&elementValue{1, map[string]struct{}{key: {}}})
		}
		this.keyCountMap[key] = this.countList.Back()
	} else {
		// key存在，则转移key
		value := element.Value.(*elementValue).value
		prev := element.Prev()
		if prev != nil && prev.Value.(*elementValue).value == value+1 {
			prev.Value.(*elementValue).keysMap[key] = struct{}{}
		} else {
			this.countList.InsertBefore(&elementValue{value + 1, map[string]struct{}{key: {}}}, element)
		}
		this.keyCountMap[key] = element.Prev()
		keyMap := element.Value.(*elementValue).keysMap
		delete(keyMap, key)
		if len(keyMap) == 0 {
			this.countList.Remove(element)
		}
	}
}

func (this *AllOne) Dec(key string) {
	element, ok := this.keyCountMap[key]
	if !ok {
		return
	}
	value := element.Value.(*elementValue).value
	next := element.Next()
	if value != 1 {
		if next != nil && next.Value.(*elementValue).value == value-1 {
			next.Value.(*elementValue).keysMap[key] = struct{}{}
		} else {
			this.countList.InsertAfter(&elementValue{value - 1, map[string]struct{}{key: {}}}, element)
		}
		this.keyCountMap[key] = element.Next()
	} else {
		delete(this.keyCountMap, key)
	}

	keyMap := element.Value.(*elementValue).keysMap
	delete(keyMap, key)
	if len(keyMap) == 0 {
		this.countList.Remove(element)
	}
}

func (this *AllOne) GetMaxKey() string {
	element := this.countList.Front()
	if element == nil {
		return ""
	}
	for key := range element.Value.(*elementValue).keysMap {
		return key
	}
	return ""
}

func (this *AllOne) GetMinKey() string {
	element := this.countList.Back()
	if element == nil {
		return ""
	}
	for key := range element.Value.(*elementValue).keysMap {
		return key
	}
	return ""
}
