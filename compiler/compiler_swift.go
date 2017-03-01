package compiler

import (
	"strings"
	"sync"

	"github.com/samuel/go-thrift/parser"

	"github.com/Arror/CodeFly/context"
	"github.com/Arror/CodeFly/types"
)

const (
	enumTplName    = "SwiftEnum"
	structTplName  = "SwiftStruct"
	serviceTplName = "SwiftService"

	enumTplPath    = "templates/swift/enum.tpl"
	structTplPath  = "templates/swift/struct.tpl"
	serviceTplPath = "templates/swift/service.tpl"
)

var (
	_ctx *context.Context
)

// SwiftCompiler Swift Code Compiler
type SwiftCompiler struct{}

// SwiftCompilerAssistant Swift compiler assistant
type SwiftCompilerAssistant struct{}

// SwiftEnum Swift Enum
type SwiftEnum struct {
	*parser.Enum
	SCA SwiftCompilerAssistant
}

// SwiftStruct Swift Struct
type SwiftStruct struct {
	*parser.Struct
	SCA SwiftCompilerAssistant
}

// SwiftService Swift Service
type SwiftService struct {
	*parser.Service
	SCA SwiftCompilerAssistant
}

// Name Enum name
func (se *SwiftEnum) Name() string {
	return _ctx.Thrift.Namespaces[_ctx.Lang] + se.Enum.Name
}

// Name Struct name
func (ss *SwiftStruct) Name() string {
	return _ctx.Thrift.Namespaces[_ctx.Lang] + ss.Struct.Name
}

// Name Service name
func (ss *SwiftService) Name() string {
	return ss.Service.Name + "Service"
}

// MethodName Method Name
func (ss *SwiftService) MethodName(m *parser.Method) string {
	return strings.ToLower(m.Name[:1]) + m.Name[1:]
}

func (sc *SwiftCompiler) compile(ctx *context.Context) {

	_ctx = ctx

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, e := range ctx.Thrift.Enums {
			se := &SwiftEnum{
				Enum: e,
				SCA:  SwiftCompilerAssistant{},
			}
			fn := se.Name() + ".swift"
			err := ctx.ExportFiles(fn, enumTplName, enumTplPath, se)
			if err != nil {
				panic(err.Error())
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, s := range ctx.Thrift.Structs {
			ss := &SwiftStruct{
				Struct: s,
				SCA:    SwiftCompilerAssistant{},
			}
			fn := ss.Name() + ".swift"
			err := ctx.ExportFiles(fn, structTplName, structTplPath, ss)
			if err != nil {
				panic(err.Error())
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, s := range ctx.Thrift.Services {
			ss := &SwiftService{
				Service: s,
				SCA:     SwiftCompilerAssistant{},
			}
			fn := ss.Name() + ".swift"
			err := ctx.ExportFiles(fn, serviceTplName, serviceTplPath, ss)
			if err != nil {
				panic(err.Error())
			}
		}
	}()

	wg.Wait()
}

// FormatFiledName Format filed name
func (SCA SwiftCompilerAssistant) FormatFiledName(n string) string {

	name := n

	if !strings.Contains(name, "_") {
		return name
	}

	components := strings.Split(name, "_")

	name = ""

	for _, component := range components {
		if component != "" {
			name += (strings.ToUpper(component[:1]) + component[1:])
		}
	}

	if name == "" {
		panic("invaild filed name: " + n)
	}

	return strings.ToLower(name[:1]) + name[1:]
}

// TypeString Type string
func (SCA SwiftCompilerAssistant) TypeString(t *parser.Type) string {

	if t == nil {
		return swiftVoid
	}

	switch t.Name {
	case types.ThriftList:
		switch t.ValueType.Name {
		case types.ThriftList, types.ThriftSet, types.ThriftMap:
			panic("unsupported [[Type]]], [Key : Value] or Set<Type>")
		}
		return "[" + SCA.TypeString(t.ValueType) + "]"
	case types.ThriftMap, types.ThriftSet:
		panic("unsupported [Key : Value] or Set<Type>")
	}

	if base := mapping[t.Name]; base != "" {
		return base
	}

	components := strings.Split(t.Name, ".")

	count := len(components)

	var _thrift *parser.Thrift
	var _type string

	switch count {
	case 1:
		_thrift = _ctx.Thrift
		_type = components[0]
	case 2:
		key := _ctx.Thrift.Includes[components[0]]
		_thrift = _ctx.Thrifts[key]
		_type = components[1]
	}

	if _thrift == nil || _type == "" {
		panic("unsupported type " + t.Name)
	}

	return _thrift.Namespaces[_ctx.Lang] + _type
}

const (
	swiftInt    = "Int"
	swiftInt64  = "Int64"
	swiftDouble = "Double"
	swiftBool   = "Bool"
	swiftString = "String"
	swiftVoid   = "Void"
)

var mapping = map[string]string{
	types.ThriftI16:    swiftInt,
	types.ThriftI32:    swiftInt,
	types.ThriftI64:    swiftInt64,
	types.ThriftBool:   swiftBool,
	types.ThriftDouble: swiftDouble,
	types.ThriftString: swiftString,
	types.ThriftByte:   types.Unsupported,
	types.ThriftBinary: types.Unsupported,
}
