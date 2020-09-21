package bank

var (
	sema    = make(chan struct{}, 1) // a binary semaphore guarding balance
	balance int
)

// Deposit 入账
func Deposit(amount int) {
	sema <- struct{}{} // acquire token
	balance = balance + amount
	<-sema // release token
}

// Balance 查询余额
func Balance() int {
	sema <- struct{}{} // acquire token
	b := balance
	<-sema // release token
	return b
}
