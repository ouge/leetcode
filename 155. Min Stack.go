package leetcode

/*
type MinStack struct {
	data []int
	min  []int
}

func Constructor() MinStack {
	return MinStack{nil, nil}
}

func (this *MinStack) Push(x int) {
	this.data = append(this.data, x)
	this.min = append(this.min, x)
	if len(this.min) >= 2 && this.min[len(this.min)-2] < x {
		this.min[len(this.min)-1] = this.min[len(this.min)-2]
	}
}

func (this *MinStack) Pop() {
	this.data = this.data[:len(this.data)-1]
	this.min = this.min[:len(this.min)-1]
}

func (this *MinStack) Top() int {
	return this.data[len(this.data)-1]
}

func (this *MinStack) GetMin() int {
	return this.min[len(this.min)-1]
}
*/

type MinStack struct {
	data []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{nil}
}

func (this *MinStack) Push(x int) {
	this.data = append(this.data, x)
	if len(this.data) >= 2 && this.data[len(this.data)-2] < x {
		this.data = append(this.data, this.data[len(this.data)-2])
	} else {
		this.data = append(this.data, x)
	}
}

func (this *MinStack) Pop() {
	this.data = this.data[:len(this.data)-2]
}

func (this *MinStack) Top() int {
	return this.data[len(this.data)-2]
}

func (this *MinStack) GetMin() int {
	return this.data[len(this.data)-1]
}
