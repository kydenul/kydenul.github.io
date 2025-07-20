# Interview Guide


{{< admonition type=abstract title="导语" open=true >}}
**这是导语部分**
{{< /admonition >}}

<!--more-->

## 基础语法

### 1. `=` 和 `:=` 的区别？

{{< admonition type=success title="答案" open=false >}}
`:=` 用于声明并初始化变量

`=` 用于赋值

```go
var foo int
foo = 1

// 等价于

foo := 1
```

{{< /admonition >}}

### 2. 指针的作用？

{{< admonition type=success title="答案" open=false >}}
{{< /admonition >}}

### 3. Go 允许多个返回值吗？

### 4. Go 有异常类型吗？

### 5. 什么是协程？

### 6. Go 中如何高效地拼接字符串？

### 7 什么是 rune 类型

### 8 如何判断 map 中是否包含某个 key ？

### 9 Go 支持默认参数或可选参数吗？

### 10 defer 的执行顺序

### 11 如何交换 2 个变量的值？

### 12 Go 语言 tag 的用处？

### 13 如何判断 2 个字符串切片（slice) 是相等的？

### 14 字符串打印时，%v 和 %+v 的区别

### 15 Go 语言中如何表示枚举值(enums)？

### 16 空 struct{} 的用途

## 实现原理

### 1 init() 函数是什么时候执行的？

### 2 Go 语言的局部变量分配在栈上还是堆上？

### 3 2 个 interface 可以比较吗 ？

### 4 2 个 nil 可能不相等吗？

### 5 简述 Go 语言GC(垃圾回收)的工作原理

### 6 函数返回局部变量的指针是否安全？

### 7 非接口非接口的任意类型 T() 都能够调用 *T 的方法吗？反过来呢？

## 并发编程

### 1 无缓冲的 channel 和有缓冲的 channel 的区别？

### 2 什么是协程泄露(Goroutine Leak)？

### 3 Go 可以限制运行时操作系统线程的数量吗？

## 代码输出

### 变量与常量

### 作用域

### defer 延迟调用


---

> Author: [kyden](https://github.com/kydenul)  
> URL: http://localhost:1313/posts/be50b37/  

