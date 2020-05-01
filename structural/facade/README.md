# Facade Pattern

外观模式：为子系统中的一组接口提供一个一致的界面，此模式定义了一个高层接口，这个接口使得这一子系统更加容易使用。

# Implementation

```
// 外观类
type facade struct {
	one   *SubSystemOne
	two   *SubSystemTwo
	three *SubSystemThree
	four  *SubSystemFour
}

func CreateFacade() *facade {
	f := new(facade)
	f.one = new(SubSystemOne)
	f.two = new(SubSystemTwo)
	f.three = new(SubSystemThree)
	f.four = new(SubSystemFour)
	return f
}

func (p *facade) MethodA() {
	fmt.Println("方法组 MethodA()")
	p.one.MethodOne()
	p.two.MethodTwo()
	p.four.MethodFour()
}

func (p *facade) MethodB() {
	fmt.Println("方法组 MethodB()")
	p.two.MethodTwo()
	p.three.MethodThree()
}
```

# Usage

```
f := CreateFacade()

f.MethodA()

f.MethodB()
```

```
方法组 MethodA()
子系统方法一
子系统方法二
子系统方法四
方法组 MethodB()
子系统方法二
子系统方法三
```

# Rules of Thumb

第一，在设计初期阶段，应该要有意识的将不同的两个层分离，层与层之间建立外观Facade；  
第二，在开发阶段，子系统往往因为不断的重构烟花而变得越来越复杂，此时增加外观Facade可以提供一个简单的接口，减少它们的依赖；  
第三，在维护一个遗留的大型系统时，可能这个系统已经非常难以维护和扩展了，此时可以为新系统开发一个外观Facade类，来提供设计粗糙或高度复杂的遗留代码的比较清晰的简单接口，让新系统与Facade对象交互，Facade与遗留代码交互所有复杂的工作。