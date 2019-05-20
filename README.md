GO常见问题

语言基础
1. 关键字“defer” 的作用，和如何使用？同一个方法中的若干个defer的执行顺序先后是什么？（初）

 > defer的作用类似于java里的final，也就是保证方法在推出之前，一定要执行哪些代码，一般是关闭句柄（文件/socket），解锁。另外，defer的执行顺序是堆栈形式（FILO）的，

```go
func main() {
    var wg sync.WaitGroup
    wg.Add(2)
    go printPrime("A")
    go printPrime("B")
    wg.Wait()
}

func printPrime(prefix string) {
    // Schedule the call to Done to tell main we are done.
    defer wg.Done()
}
```

2. go中如何处理Panic？比如破获Panic，打印Panic的信息并继续执行？（初）

 > 相关的内容可以参考 [Defer, Panic, and Recover](https://blog.golang.org/defer-panic-and-recover) 。没有 try...catch...finally的组合，取而代之的是“Defer, Panic, Recover”。用Defer声明拦截方法，在Panic发生的时候，Defer方法被调用，在这个方法内，调用Recover方法，取得Panic的信息并处理。*注意：例子中的"e"是不会输出的，因为一旦发生panic，该方法就是结束了，只会调用defer方法。*

 ```go
 func main() {
	handlePanic()
}

func genPanic() {
	fmt.Println("a")
	panic(55)
	fmt.Println("b") //not display
	fmt.Println("f") //not display
}

func handlePanic() {
	defer func() {
		fmt.Println("c")
		if err := recover(); err != nil {
			fmt.Println(err)
		}
		fmt.Println("d")
	}()
	genPanic()
	fmt.Println("e") // ** not display **
}
 ```

3. 如何定义一个类，并使这个类拥有一个类成员方法和自定义类成员变量？（初）

 > Go的实现和Java/C++大不相同。首先是定义一个复合类型，可以把它理解为类成员变量的声明，称之为struct，然后再定义一些方法，参数就是该struct的值引用或者指针引用。

 ```go
type Employee struct {  
    FirstName   string
    LastName    string
    TotalLeaves int
    LeavesTaken int
}

func (e Employee) LeavesRemaining() {  
    fmt.Printf("%s %s has %d leaves remaining", e.FirstName, e.LastName, (e.TotalLeaves - e.LeavesTaken))
}
 ```

4. 在定义类成员方法时，什么情况下使用指针参数，什么情况使用值参数？（中）

 > 如果是要修改实例的值，就用指针。只是打印的话，值引用就好。

 ```go
 func (e Employee) LeavesRemaining() {
    fmt.Printf("%s %s has %d leaves remaining", e.FirstName, e.LastName, (e.TotalLeaves - e.LeavesTaken))
 }
 func (e *Employee) RebalanceLeaves() {
    e.LeavesTaken = 12
    e.LeavesRemaining()
 }
 ```

5. go中如何实现继承？如何调用“父类”的方法？（中）

 > 严格来说go是没有继承的，实际上是“组合”。[“继承”和“重写”的示例](https://studygolang.com/articles/4113)，如果是调用“父类”方法的话可以：sam.Human.SayHi()

6. interface的作用是什么？（中）

 > [理解 Go interface 的 5 个关键点](https://sanyuesha.com/2017/07/22/how-to-understand-go-interface/)。它和java里的interface很像，但是实际意义和使用方式都不同。本质上是go是组合，而java/c++是继承。简单的说，具有某些方法的实例，就可以用某个interface类型的变量来指代，很像父类的“引用”或者指针。


7. 什么是method set？（高）

 > 由于interface类型既可以指向值也可以指向指针，但是这两个interface变量的方法集（method set）是不同的。具体可以参考[理解 Go interface 的 5 个关键点](https://sanyuesha.com/2017/07/22/how-to-understand-go-interface/)中的第5点。


项目维护
1. go有很多基础模块在国内不能访问，如何解决？（初）

 > 挂VPN，直接翻墙；翻墙以后，把对应的包下载回来，放在必要的位置；使用go modules，将golang.org相关的包replace为github上的镜像地址。

2. 在引用第三方依赖模块的时候，如何指定版本号？（中）

 > *我只用过modules，不知道以前是怎么玩的。*

3. 如何在同一台机器，同一个账户上，同时开发/维护两个项目？（中）

 > *我只用过modules，不知道以前是怎么玩的。*

4. 是否使用过go modules？（高）

 > [拜拜了，GOPATH君！新版本Golang的包管理入门教程](https://segmentfault.com/a/1190000018690961)

5. 项目如何编译打包，和发布？如何交叉编译？（高）

 > *TODO*
