package main

type MinStack struct {
	data   []int
	helper []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	ms := MinStack{}
	ms.data = make([]int, 0)
	ms.helper = make([]int, 0)

	return ms
}

func (this *MinStack) Push(x int) {
	this.data = append(this.data, x)
	if len(this.helper) == 0 || x < this.helper[len(this.helper)-1] {
		this.helper = append(this.helper, x)
	} else {
		this.helper = append(this.helper, this.helper[len(this.helper)-1])
	}
}

func (this *MinStack) Pop() {
	if len(this.data) > 0 {
		this.data = this.data[:len(this.data)-1]
		this.helper = this.helper[:len(this.helper)-1]
	}

}

func (this *MinStack) Top() int {
	if len(this.data) > 0{
		return this.data[len(this.data)-1]
	}
	return 0
}

func (this *MinStack) GetMin() int {
	if len(this.helper) > 0{
		return this.helper[len(this.helper)-1]
	}
	return 0
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
