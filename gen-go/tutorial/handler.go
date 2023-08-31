package tutorial

import (
	"context"
	"fmt"
	"gen-go/shared"
	"strconv"
)

/**
handler. 应该是接口的实现类
接口 IDL：
service Calculator extends shared.SharedService {
	void ping(),
	i32 add(1:i32 num1, 2:i32 num2),
	i32 calculate(1:i32 logid, 2:Work w) throws (1:InvalidOperation ouch),
	oneway void zip()
}

service SharedService {
  SharedStruct getStruct(1: i32 key)
}

*/

type CalculatorHandler struct { // 一般 spring 里的 impl 也需要注入一个实例！
	log map[int]*shared.SharedStruct
}

func NewCalculatorHandler() *CalculatorHandler { // 新建 impl 实例。返回的是一个引用对象
	return &CalculatorHandler{log: make(map[int]*shared.SharedStruct)}
}

func (p *CalculatorHandler) Ping(ctx context.Context) (err error) { // 处理器的对象方法，实现了接口
	fmt.Print("ping()\n")
	return nil
}

func (p *CalculatorHandler) Add(ctx context.Context, num1 int32, num2 int32) (retval17 int32, err error) {
	fmt.Print("add(", num1, ",", num2, ")\n")
	return num1 + num2, nil
}

func (p *CalculatorHandler) Calculate(ctx context.Context, logid int32, w *Work) (val int32, err error) {
	fmt.Print("calculate(", logid, ", {", w.Op, ",", w.Num1, ",", w.Num2, "})\n")
	switch w.Op {
	case Operation_ADD:
		val = w.Num1 + w.Num2
	case Operation_SUBTRACT:
		val = w.Num1 - w.Num2
	case Operation_MULTIPLY:
		val = w.Num1 * w.Num2
	case Operation_DIVIDE:
		if w.Num2 == 0 {
			ouch := NewInvalidOperation()
			ouch.WhatOp = int32(w.Op)
			ouch.Why = "Cannot divide by 0"
			err = ouch
			return
		}
		val = w.Num1 / w.Num2
	default:
		ouch := NewInvalidOperation()
		ouch.WhatOp = int32(w.Op)
		ouch.Why = "Unknown operation"
		err = ouch
		return
	}

	entry := shared.NewSharedStruct()
	entry.Key = logid
	entry.Value = strconv.Itoa(int(val))
	k := int(logid)

	p.log[k] = entry
	return val, err
}

func (p *CalculatorHandler) GetStruct(ctx context.Context, key int32) (*shared.SharedStruct, error) {
	fmt.Print("getStruct(", key, ")\n")
	v := p.log[int(key)]
	return v, nil
}

func (p *CalculatorHandler) Zip(ctx context.Context) (err error) {
	fmt.Print("zip()\n")
	return nil
}
