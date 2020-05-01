# Bridge Pattern

合成/聚合复用原则（CARP）：尽量使用合成/聚合，尽量不要使用类继承。  
  
桥接模式：将抽象部分和它的实现部分分离，使他们都可以独立地变化(实现指的是抽象类和它的派生类用来实现自己的对象)。

# Implementation

示例实现了软件对不同品牌手机的松耦合的开发。

## Types

```
package bridge

type Soft interface {
	Run() error
}

type Brand interface {
	SetSoft(soft Soft)
	Run()
}
```

## Defferent Implementations

```
type huawei struct {
	soft bridge.Soft
}

func CreateHuaWei() bridge.Brand {
	return new(huawei)
}

func (hw *huawei) SetSoft(soft bridge.Soft) {
	// do something
	hw.soft = soft
}

func (hw *huawei) Run() {
	// do something
	hw.soft.Run()
}
```

```
type oppo struct {
	soft bridge.Soft
}

func CreateOppo() bridge.Brand {
	return new(oppo)
}

func (op *oppo) SetSoft(soft bridge.Soft) {
	// do something
	op.soft = soft
}

func (op *oppo) Run() {
	// do something
	op.soft.Run()
}
```

```
// 通讯录
type addressList struct {
}

func NewAddressList() bridge.Soft {
	return new(addressList)
}

func (al *addressList) Run() error {
	fmt.Println("Run the handset address list.")
	return nil
}
```

```
// 游戏
type game struct {
}

func NewGame() bridge.Soft {
	return new(game)
}

func (game *game) Run() error {
	fmt.Println("Run the handset game.")
	return nil
}
```

# Usage

```
hw := brand.CreateHuaWei()

hw.SetSoft(soft.NewGame())
hw.Run()

hw.SetSoft(soft.NewAddressList())
hw.Run()

oppo := brand.CreateOppo()

oppo.SetSoft(soft.NewGame())
oppo.Run()

oppo.SetSoft(soft.NewAddressList())
oppo.Run()
```

```
Run the handset game.
Run the handset address list.
Run the handset game.
Run the handset address list.
```

# Rules of Thumb

桥接模式理解：实现系统可能有多角度分类，每一种分类有可能变化，那么就把这种多角度分离出来让它们独立变化，减少他们的耦合。