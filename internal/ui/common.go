package ui

type Components interface {
	ButtonProperties | BoxProperties | TextProperties
}

type HoverState[T Components] struct {
	OnHover   func(T) T
	isHovered bool
}

// func (state *HoverState[T]) OnHover(callback func(T) T) {
// 	state.onHover = callback
// }

type FocusState[T Components] struct {
	onFocus   func(T) T
	isFocused bool
}

func (state *FocusState[T]) OnFocus(callback func(T) T) {
	state.onFocus = callback
}
