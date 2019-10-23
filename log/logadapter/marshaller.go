//
//    Copyright 2019 Insolar Technologies
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.
//

package logadapter

import (
	"fmt"
	"reflect"
	"sync"

	"github.com/insolar/insolar/insolar"
)

var marshallerFactory MarshallerFactory = &defaultLogObjectMarshallerFactory{}

type defaultLogObjectMarshallerFactory struct {
	mutex       sync.RWMutex
	marshallers map[reflect.Type]*typeMarshaller
	reporters   map[reflect.Type]FieldReporterFunc
	forceAddr   bool // enforce use of address/pointer-based access to fields
}

func (p *defaultLogObjectMarshallerFactory) RegisterFieldReporter(fieldType reflect.Type, fn FieldReporterFunc) {
	if fn == nil {
		panic("illegal value")
	}
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if p.reporters == nil {
		p.reporters = make(map[reflect.Type]FieldReporterFunc)
	}
	p.reporters[fieldType] = fn
}

func (p *defaultLogObjectMarshallerFactory) CreateLogObjectMarshaller(o reflect.Value) insolar.LogObjectMarshaller {
	if o.Kind() != reflect.Struct {
		panic("illegal value")
	}
	t := p.getTypeMarshaller(o.Type())
	return defaultLogObjectMarshaller{t, t.prepareValue(o)} // do prepare for a repeated use of marshaller
}

func (p *defaultLogObjectMarshallerFactory) getFieldReporter(t reflect.Type) FieldReporterFunc {
	p.mutex.RLock()
	fr := p.reporters[t]
	p.mutex.RUnlock()
	return fr
}

func (p *defaultLogObjectMarshallerFactory) getTypeMarshaller(t reflect.Type) *typeMarshaller {
	p.mutex.RLock()
	tm := p.marshallers[t]
	p.mutex.RUnlock()
	if tm != nil {
		return tm
	}

	tm = p.buildTypeMarshaller(t) // do before lock to reduce in-lock time

	p.mutex.Lock()
	defer p.mutex.Unlock()

	if p.marshallers == nil {
		p.marshallers = make(map[reflect.Type]*typeMarshaller)
	} else {
		tm2 := p.marshallers[t]
		if tm2 != nil {
			return tm2
		}
	}
	p.marshallers[t] = tm
	return tm
}

func (p *defaultLogObjectMarshallerFactory) buildTypeMarshaller(t reflect.Type) *typeMarshaller {
	n := t.NumField()
	if n <= 0 {
		return nil
	}

	tm := typeMarshaller{printNeedsAddr: p.forceAddr, reportNeedsAddr: p.forceAddr}

	if !tm.getFieldsOf(t, 0, p.getFieldReporter) {
		return nil
	}
	return &tm
}

type defaultLogObjectMarshaller struct {
	t *typeMarshaller
	v reflect.Value
}

func (v defaultLogObjectMarshaller) MarshalLogObject(output insolar.LogObjectWriter, collector insolar.LogObjectMetricCollector) string {
	return v.t.printFields(v.v, output, collector)
}

func (v defaultLogObjectMarshaller) MarshalMutedLogObject(collector insolar.LogObjectMetricCollector) {
	if collector == nil {
		return
	}
	v.t.reportFields(v.v, collector)
}

type fieldMarshallerFunc func(value reflect.Value, writer insolar.LogObjectWriter, collector insolar.LogObjectMetricCollector)
type fieldReportFunc func(value reflect.Value, collector insolar.LogObjectMetricCollector)
type fieldMarshallerMsgFunc func(value reflect.Value, collector insolar.LogObjectMetricCollector) string

type typeMarshaller struct {
	fields          []fieldMarshallerFunc
	reporters       []fieldReportFunc
	msgField        fieldMarshallerMsgFunc
	printNeedsAddr  bool
	reportNeedsAddr bool
}

type fieldOutputFunc func(insolar.LogObjectWriter, string, interface{})

type fieldDesc struct {
	reflect.StructField
	getFn    fieldValueGetterFunc
	index    int
	outputFn fieldOutputFunc
	reportFn FieldReporterFunc
}

func (p *typeMarshaller) getFieldsOf(t reflect.Type, baseOffset uintptr, getReporterFn func(reflect.Type) FieldReporterFunc) bool {
	n := t.NumField()
	var msgGetter fieldDesc
	valueGetters := make([]fieldDesc, 0, n)

	for i := 0; i < n; i++ {
		tf := t.Field(i)
		fieldName := tf.Name

		k := tf.Type.Kind()
		valueGetterFactory := fieldValueGetters[k]
		if valueGetterFactory == nil {
			continue
		}
		unexported := len(tf.PkgPath) != 0

		needsAddr := false
		fd := fieldDesc{StructField: tf, index: i}

		fd.reportFn = getReporterFn(fd.Type)
		tagType, fmtStr := singleTag(fd.Tag)

		if tagType == fmtTagText && (tf.Anonymous || fieldName == "_") {
			// a special case - we can take `txt` of an anonymous as a message text
			fieldName = "Message"
		} else if fd.reportFn == nil && (tf.Anonymous || fieldName == "" || fieldName[0] == '_') {
			continue
		}
		outputFn, optional, needsValue := outputOfField(tagType, fmtStr)

		msgField := false
		if outputFn == nil {
			if fd.reportFn == nil {
				continue
			}
		} else {
			switch fieldName {
			case "msg", "Msg", "message", "Message":
				msgField = true
			}
		}

		fd.outputFn = outputFn
		if needsValue || fd.reportFn != nil {
			needsAddr, fd.getFn = valueGetterFactory(unexported, tf.Type, optional)
		} else {
			fd.getFn = func(value reflect.Value) (v interface{}, isZero bool) {
				return nil, false
			}
		}

		if msgField {
			msgGetter = fd
		} else {
			valueGetters = append(valueGetters, fd)
		}

		if needsAddr {
			p.printNeedsAddr = true
			if fd.reportFn != nil {
				p.reportNeedsAddr = true
			}
		}
	}

	if p.reportNeedsAddr && !p.printNeedsAddr {
		panic("illegal state")
	}

	if len(valueGetters) == 0 && msgGetter.getFn == nil {
		return false
	}

	p.fields = make([]fieldMarshallerFunc, 0, len(valueGetters))

	for _, fd := range valueGetters {
		fieldGetter := getFieldGetter(fd.index, fd.StructField, p.printNeedsAddr, baseOffset)

		if fd.outputFn != nil {
			printFn := printOfField(fd, fieldGetter)
			p.fields = append(p.fields, printFn)
		}

		if fd.reportFn != nil {
			reportFieldGetter := fieldGetter
			if p.reportNeedsAddr != p.printNeedsAddr {
				reportFieldGetter = getFieldGetter(fd.index, fd.StructField, p.reportNeedsAddr, baseOffset)
			}

			reportFn := reportOfField(fd, reportFieldGetter)
			p.reporters = append(p.reporters, reportFn)
		}
	}

	if msgGetter.getFn == nil {
		p.msgField = nil
		return true
	}

	fieldGetter := getFieldGetter(msgGetter.index, msgGetter.StructField, p.printNeedsAddr, baseOffset)
	p.msgField = messageOfField(msgGetter, fieldGetter)

	return true
}

func outputOfField(tagType fmtTagType, fmtStr string) (fn fieldOutputFunc, optional bool, needsValue bool) {
	switch tagType {
	case fmtTagFormatValueOpt:
		return func(writer insolar.LogObjectWriter, fieldName string, v interface{}) {
			s := fmt.Sprintf(fmtStr, v)
			writer.AddStrField(fieldName, s)
		}, true, true
	case fmtTagFormatValue:
		return func(writer insolar.LogObjectWriter, fieldName string, v interface{}) {
			s := fmt.Sprintf(fmtStr, v)
			writer.AddStrField(fieldName, s)
		}, false, true
	case fmtTagFormatRawOpt:
		return func(writer insolar.LogObjectWriter, fieldName string, v interface{}) {
			s := fmt.Sprintf(fmtStr, v)
			writer.AddRawJSON(fieldName, []byte(s))
		}, true, true
	case fmtTagFormatRaw:
		return func(writer insolar.LogObjectWriter, fieldName string, v interface{}) {
			s := fmt.Sprintf(fmtStr, v)
			writer.AddRawJSON(fieldName, []byte(s))
		}, false, true
	case fmtTagSkip:
		return nil, false, false
	case fmtTagText:
		return func(writer insolar.LogObjectWriter, fieldName string, _ interface{}) {
			writer.AddStrField(fieldName, fmtStr)
		}, false, false
	case fmtTagOptional:
		return insolar.LogObjectWriter.AddField, true, true
	default:
		return insolar.LogObjectWriter.AddField, false, true
	}
}

type stringCapturer struct {
	v string
}

func (p *stringCapturer) AddStrField(_ string, v string) {
	p.v = v
}

func (p *stringCapturer) AddRawJSON(key string, b []byte) {
	p.v = string(b)
}

func (p *stringCapturer) AddField(key string, v interface{}) {
	p.v = fmt.Sprintf("%s", v)
}

func messageOfField(fd fieldDesc, fieldGetter func(reflect.Value) reflect.Value) func(obj reflect.Value, collector insolar.LogObjectMetricCollector) string {
	valueGetter := fd.getFn
	valueOutput := fd.outputFn

	if fd.reportFn == nil {
		return func(obj reflect.Value, _ insolar.LogObjectMetricCollector) string {
			f := fieldGetter(obj)
			v, isZero := valueGetter(f)
			if isZero {
				return ""
			}
			sc := stringCapturer{}
			valueOutput(&sc, "", v)
			return sc.v
		}
	}

	reportFn := fd.reportFn
	fieldName := fd.Name

	return func(obj reflect.Value, collector insolar.LogObjectMetricCollector) string {
		f := fieldGetter(obj)
		v, isZero := valueGetter(f)
		if collector != nil {
			reportFn(collector, fieldName, v)
		}
		if isZero {
			return ""
		}
		sc := stringCapturer{}
		valueOutput(&sc, "", v)
		return sc.v
	}
}

func printOfField(fd fieldDesc, fieldGetter func(reflect.Value) reflect.Value) fieldMarshallerFunc {
	valueGetter := fd.getFn
	valueOutput := fd.outputFn
	fieldName := fd.Name

	if fd.reportFn == nil {
		return func(obj reflect.Value, writer insolar.LogObjectWriter, _ insolar.LogObjectMetricCollector) {
			f := fieldGetter(obj)
			if v, isZero := valueGetter(f); !isZero {
				valueOutput(writer, fieldName, v)
			}
		}
	}

	reportFn := fd.reportFn
	return func(obj reflect.Value, writer insolar.LogObjectWriter, collector insolar.LogObjectMetricCollector) {
		f := fieldGetter(obj)
		v, isZero := valueGetter(f)
		if collector != nil {
			reportFn(collector, fieldName, v)
		}
		if !isZero {
			valueOutput(writer, fieldName, v)
		}
	}
}

func reportOfField(fd fieldDesc, fieldGetter func(reflect.Value) reflect.Value) fieldReportFunc {
	fieldName := fd.Name
	valueGetter := fd.getFn
	reportFn := fd.reportFn

	return func(obj reflect.Value, collector insolar.LogObjectMetricCollector) {
		f := fieldGetter(obj)
		v, _ := valueGetter(f)
		reportFn(collector, fieldName, v)
	}
}

func (p *typeMarshaller) prepareValue(value reflect.Value) reflect.Value {
	return p._prepareValue(value, p.printNeedsAddr)
}

func (p *typeMarshaller) _prepareValue(value reflect.Value, needsAddr bool) reflect.Value {
	if !needsAddr || value.CanAddr() {
		return value
	}
	valueCopy := reflect.New(value.Type()).Elem()
	valueCopy.Set(value)
	return valueCopy
}

func (p *typeMarshaller) printFields(value reflect.Value, writer insolar.LogObjectWriter, collector insolar.LogObjectMetricCollector) string {
	value = p._prepareValue(value, p.printNeedsAddr) // double check

	for _, fn := range p.fields {
		fn(value, writer, collector)
	}
	if p.msgField == nil {
		return ""
	}
	return p.msgField(value, collector)
}

func (p *typeMarshaller) reportFields(value reflect.Value, collector insolar.LogObjectMetricCollector) {
	if len(p.reporters) == 0 {
		return
	}

	value = p._prepareValue(value, p.reportNeedsAddr) // double check

	for _, fn := range p.reporters {
		fn(value, collector)
	}
}

type fmtTagType uint8

const (
	fmtTagDefault fmtTagType = iota
	fmtTagOptional

	fmtTagText
	fmtTagSkip // + opt

	fmtTagFormatRaw
	fmtTagFormatRawOpt // + opt

	fmtTagFormatValue
	fmtTagFormatValueOpt // + opt
)

func singleTag(tag reflect.StructTag) (fmtTagType, string) {
	tagType := fmtTagDefault
	if _, v, ok := ParseStructTag(tag, func(name, _ string) bool {
		switch name {
		case "fmt+opt", "opt+fmt":
			tagType = fmtTagFormatValueOpt
		case "fmt":
			tagType = fmtTagFormatValue
		case "raw+opt", "opt+raw":
			tagType = fmtTagFormatRawOpt
		case "raw":
			tagType = fmtTagFormatRaw
		case "skip":
			tagType = fmtTagSkip
		case "txt":
			tagType = fmtTagText
		case "opt":
			tagType = fmtTagOptional
		default:
			return false
		}
		return true
	}); ok {
		return tagType, v
	}
	return fmtTagDefault, ""
}
