# Decorator Pattern

装饰模式：动态地给一个对象添加一些额外的职责，就增加功能来说，装饰模式比生成子类更为灵活。

# Implementation

```
type Object func(int) int

func CalculateDecorate(o Object) Object {
	return func(n int) int {
		fmt.Println("Starting the execution with the integer : ", n)

		result := o(n)

		fmt.Println("Execution is completed  with the result : ", result)

		return result
	}
}
```

# Usage

```
func Double(n int) int {
	return n * 2
}

func Square(n int) int {
	return n * n
}

func main() {
	o := CalculateDecorate(Double)
	o(10)

	o = CalculateDecorate(Square)
	o(10)
}
```

```
Starting the execution with the integer :  10
Execution is completed  with the result :  20
Starting the execution with the integer :  10
Execution is completed  with the result :  100
```

# Rules of Thumb

当系统需要新功能的时候，向旧的类中添加新的代码，这些新加的代码通常装饰了原有类的核心职责或主要行为。  
装饰模式可以将每个要装饰的功能放在独立的类中，并让这个类包装它所要装饰的对象，因此，当需要执行特殊行为时，客户代码就可以在运行时根据需要有选择地、按顺序地使用装饰功能包装对象。  
  
装饰模式优点：把类中的装饰功能从类中搬移去除，这样可以简化原有的类。可以有效地把类的核心职责和装饰功能分开，而且可以去除相关类中重复的装饰逻辑。