# Adapter Pattern

适配器模式：将一个类的接口转换成客户希望的另外一个接口。Adapter 模式使得原本由于接口不兼容而不能一起工作的那些类可以一起工作。

# Implementation

```
// 客户端所期待的接口，目标可以是具体的或抽象的类，也可以是接口（这里示范的是接口）
type Target interface {
	Request()
}

// 适配器：通过在内部封装一个 Adaptee 对象，把源接口装换成目标接口
type Adapter struct {
	Adp *Adaptee
}

func (o *Adapter) Request() {
	o.Adp.SpecificRequest()
}

// 需要适配的类，原接口 SpecificRequest 已无法满足现有接口 Request 的需求
type Adaptee struct {}

func (o *Adaptee) SpecificRequest() {
	fmt.Println("Hello, This is SpecificRequest.")
}
```

# Usage

```
var target Target
target = &Adapter{
    Adp: new(Adaptee),
}

target.Request()
```

```
Hello, This is SpecificRequest.
```

# Rules of Thumb

在系统的数据和行为都正确，但接口不符时，此时考虑使用适配器，目的是使控制范围之外的一个原有对象与某接口匹配，适配器模式主要应用于希望复用一些现存的类，但是接口又与复用环境要求不一致。