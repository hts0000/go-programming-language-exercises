package bank

// 数据竞争：无论什么时候，两个及以上goroutine并发访问同一个变量，且至少其中一个有写操作

// 练习 9.1
// 给gopl.io/ch9/bank1程序添加一个Withdraw(amount int)取款函数
// 其返回结果应该要表明事务是成功了还是因为没有足够资金失败了
// 这条消息会被发送给monitor的goroutine，且消息需要包含取款的额度和一个新的channel
// 这个新channel会被monitor goroutine来把boolean结果发回给Withdraw。

var (
	deposits = make(chan int)
	balances = make(chan int)
	withdraw = make(chan int)
	result   = make(chan bool)
)

func Deposits(amount int) { deposits <- amount }
func Balance() int        { return <-balances }
func WithDraw(amount int) bool {
	withdraw <- amount
	return <-result
}

func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case amount := <-withdraw:
			if amount <= balance {
				balance -= balance
				result <- true
			} else {
				result <- false
			}
		}
	}
}

func Withdraw(amount int) {

}

func init() {
	go teller()
}
