package easy


/**
写一个 RecentCounter 类来计算特定时间范围内最近的请求。

请你实现 RecentCounter 类：

RecentCounter() 初始化计数器，请求数为 0 。
int ping(int t) 在时间 t 添加一个新请求，其中 t 表示以毫秒为单位的某个时间，并返回过去 3000 毫秒内发生的所有请求数（包括新请求）。确切地说，返回在 [t-3000, t] 内发生的请求数。
保证每次对 ping 的调用都使用比之前更大的 t 值。
 */
type RecentCounter struct {
	queue []int
}


func Constructor() RecentCounter {
	return RecentCounter {
		queue: make([]int, 0),
	}
}


func (this *RecentCounter) Ping(t int) int {
	this.queue = append(this.queue, t)
	l, ans := len(this.queue), 1
	for l -= 2; l >= 0; l-- {
		if t - this.queue[l] <= 3000 {
			ans++
		} else {
			this.queue = this.queue[l + 1:]
			break
		}
	}

	return ans
}
