package main

import "fmt"

// 可视化组件接口
type VisualComponent interface {
	Draw()
	Resize()
}

type Decorator struct {
	Component VisualComponent
}

func NewDecorator(comp VisualComponent) VisualComponent {
	return &Decorator{
		Component: comp,
	}
}

func (dec *Decorator) Draw() {
	dec.Component.Draw()
}

func (dec *Decorator) Resize() {
	dec.Component.Resize()
}

// 边框装饰器
type BorderDecorator struct {
	Decorator
	width int
}

func NewBorderDecorator(comp VisualComponent, borderWidth int) VisualComponent{
	bd := new(BorderDecorator)
	bd.Component = comp
	bd.width = borderWidth
	return bd
}

func (dec *BorderDecorator) Draw() {
	dec.Decorator.Draw()
	dec.drawBorder()
}

func (dec *BorderDecorator) drawBorder() {
	fmt.Printf("绘制大小为 %d*%d 的边框\n", dec.width, dec.width)
}

// 滚动条装饰器
type ScrollDecorator struct {
	Decorator
}

func (dec *ScrollDecorator) Draw() {
	dec.Decorator.Draw()
	dec.drawScroll()
}

func (dec *ScrollDecorator) drawScroll() {
	fmt.Println("绘制滚动条")
}

// 阴影装饰器
type DropShadowDecorator struct {
	Decorator
}

func (dec *DropShadowDecorator) Draw() {
	dec.Decorator.Draw()
}

func (dec *DropShadowDecorator) drawDropShadow() {
	fmt.Println("绘制阴影效果")
}