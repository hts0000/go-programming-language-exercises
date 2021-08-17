package bank

// 数据竞争：无论什么时候，两个及以上goroutine并发访问同一个变量，且至少其中一个有写操作

var (
	deposits = make(chan int)
	balances = make(chan int)
)

func Deposits(amount int) { deposits <- amount }
func Balance() int        { return <-balances }

func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		}
	}
}

func init() {
	go teller()
}
