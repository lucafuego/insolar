// Copyright 2020 Insolar Network Ltd.
// All rights reserved.
// This material is licensed under the Insolar License version 1.0,
// available at https://github.com/insolar/insolar/blob/master/LICENSE.md.

// Code generated by insgocc. DO NOT EDIT.
// source template in logicrunner/preprocessor/templates

package second

import (
	"github.com/insolar/insolar/insolar"
	"github.com/insolar/insolar/logicrunner/builtin/foundation"
	"github.com/insolar/insolar/logicrunner/common"
)

type Payload struct {
	Int int
	Str string
}

// PrototypeReference to prototype of this contract
// error checking hides in generator
var PrototypeReference, _ = insolar.NewObjectReferenceFromString("insolar:0AAABAnMNN5fPUhNWiXpSMacCG0MKI38i2h3anCnmDa8")

// Second holds proxy type
type Second struct {
	Reference insolar.Reference
	Prototype insolar.Reference
	Code      insolar.Reference
}

// ContractConstructorHolder holds logic with object construction
type ContractConstructorHolder struct {
	constructorName string
	argsSerialized  []byte
}

// AsChild saves object as child
func (r *ContractConstructorHolder) AsChild(objRef insolar.Reference) (*Second, error) {
	ret, err := common.CurrentProxyCtx.SaveAsChild(objRef, *PrototypeReference, r.constructorName, r.argsSerialized)
	if err != nil {
		return nil, err
	}

	var ref insolar.Reference
	var constructorError *foundation.Error
	resultContainer := foundation.Result{
		Returns: []interface{}{&ref, &constructorError},
	}
	err = common.CurrentProxyCtx.Deserialize(ret, &resultContainer)
	if err != nil {
		return nil, err
	}

	if resultContainer.Error != nil {
		return nil, resultContainer.Error
	}

	if constructorError != nil {
		return nil, constructorError
	}

	return &Second{Reference: ref}, nil
}

// GetObject returns proxy object
func GetObject(ref insolar.Reference) *Second {
	if !ref.IsObjectReference() {
		return nil
	}
	return &Second{Reference: ref}
}

// GetPrototype returns reference to the prototype
func GetPrototype() insolar.Reference {
	return *PrototypeReference
}

// New is constructor
func New() *ContractConstructorHolder {
	var args [0]interface{}

	var argsSerialized []byte
	err := common.CurrentProxyCtx.Serialize(args, &argsSerialized)
	if err != nil {
		panic(err)
	}

	return &ContractConstructorHolder{constructorName: "New", argsSerialized: argsSerialized}
}

// NewWithOne is constructor
func NewWithOne(oneNumber int) *ContractConstructorHolder {
	var args [1]interface{}
	args[0] = oneNumber

	var argsSerialized []byte
	err := common.CurrentProxyCtx.Serialize(args, &argsSerialized)
	if err != nil {
		panic(err)
	}

	return &ContractConstructorHolder{constructorName: "NewWithOne", argsSerialized: argsSerialized}
}

// NewWithX is constructor
func NewWithX() *ContractConstructorHolder {
	var args [0]interface{}

	var argsSerialized []byte
	err := common.CurrentProxyCtx.Serialize(args, &argsSerialized)
	if err != nil {
		panic(err)
	}

	return &ContractConstructorHolder{constructorName: "NewWithX", argsSerialized: argsSerialized}
}

// NewSaga is constructor
func NewSaga() *ContractConstructorHolder {
	var args [0]interface{}

	var argsSerialized []byte
	err := common.CurrentProxyCtx.Serialize(args, &argsSerialized)
	if err != nil {
		panic(err)
	}

	return &ContractConstructorHolder{constructorName: "NewSaga", argsSerialized: argsSerialized}
}

// NewNil is constructor
func NewNil() *ContractConstructorHolder {
	var args [0]interface{}

	var argsSerialized []byte
	err := common.CurrentProxyCtx.Serialize(args, &argsSerialized)
	if err != nil {
		panic(err)
	}

	return &ContractConstructorHolder{constructorName: "NewNil", argsSerialized: argsSerialized}
}

// NewWithErr is constructor
func NewWithErr() *ContractConstructorHolder {
	var args [0]interface{}

	var argsSerialized []byte
	err := common.CurrentProxyCtx.Serialize(args, &argsSerialized)
	if err != nil {
		panic(err)
	}

	return &ContractConstructorHolder{constructorName: "NewWithErr", argsSerialized: argsSerialized}
}

// GetReference returns reference of the object
func (r *Second) GetReference() insolar.Reference {
	return r.Reference
}

// GetPrototype returns reference to the code
func (r *Second) GetPrototype() (insolar.Reference, error) {
	if r.Prototype.IsEmpty() {
		ret := [2]interface{}{}
		var ret0 insolar.Reference
		ret[0] = &ret0
		var ret1 *foundation.Error
		ret[1] = &ret1

		res, err := common.CurrentProxyCtx.RouteCall(r.Reference, false, false, "GetPrototype", make([]byte, 0), *PrototypeReference)
		if err != nil {
			return ret0, err
		}

		err = common.CurrentProxyCtx.Deserialize(res, &ret)
		if err != nil {
			return ret0, err
		}

		if ret1 != nil {
			return ret0, ret1
		}

		r.Prototype = ret0
	}

	return r.Prototype, nil

}

// GetCode returns reference to the code
func (r *Second) GetCode() (insolar.Reference, error) {
	if r.Code.IsEmpty() {
		ret := [2]interface{}{}
		var ret0 insolar.Reference
		ret[0] = &ret0
		var ret1 *foundation.Error
		ret[1] = &ret1

		res, err := common.CurrentProxyCtx.RouteCall(r.Reference, false, false, "GetCode", make([]byte, 0), *PrototypeReference)
		if err != nil {
			return ret0, err
		}

		err = common.CurrentProxyCtx.Deserialize(res, &ret)
		if err != nil {
			return ret0, err
		}

		if ret1 != nil {
			return ret0, ret1
		}

		r.Code = ret0
	}

	return r.Code, nil
}

// GetName is proxy generated method
func (r *Second) GetName() (string, error) {
	var args [0]interface{}

	var argsSerialized []byte

	ret := make([]interface{}, 2)
	var ret0 string
	ret[0] = &ret0
	var ret1 *foundation.Error
	ret[1] = &ret1

	err := common.CurrentProxyCtx.Serialize(args, &argsSerialized)
	if err != nil {
		return ret0, err
	}

	res, err := common.CurrentProxyCtx.RouteCall(r.Reference, false, false, "GetName", argsSerialized, *PrototypeReference)
	if err != nil {
		return ret0, err
	}

	resultContainer := foundation.Result{
		Returns: ret,
	}
	err = common.CurrentProxyCtx.Deserialize(res, &resultContainer)
	if err != nil {
		return ret0, err
	}
	if resultContainer.Error != nil {
		err = resultContainer.Error
		return ret0, err
	}
	if ret1 != nil {
		return ret0, ret1
	}
	return ret0, nil
}

// GetNameAsImmutable is proxy generated method
func (r *Second) GetNameAsImmutable() (string, error) {
	var args [0]interface{}

	var argsSerialized []byte

	ret := make([]interface{}, 2)
	var ret0 string
	ret[0] = &ret0
	var ret1 *foundation.Error
	ret[1] = &ret1

	err := common.CurrentProxyCtx.Serialize(args, &argsSerialized)
	if err != nil {
		return ret0, err
	}

	res, err := common.CurrentProxyCtx.RouteCall(r.Reference, true, false, "GetName", argsSerialized, *PrototypeReference)
	if err != nil {
		return ret0, err
	}

	resultContainer := foundation.Result{
		Returns: ret,
	}
	err = common.CurrentProxyCtx.Deserialize(res, &resultContainer)
	if err != nil {
		return ret0, err
	}
	if resultContainer.Error != nil {
		err = resultContainer.Error
		return ret0, err
	}
	if ret1 != nil {
		return ret0, ret1
	}
	return ret0, nil
}

// DoNothing is proxy generated method
func (r *Second) DoNothing() error {
	var args [0]interface{}

	var argsSerialized []byte

	ret := make([]interface{}, 1)
	var ret0 *foundation.Error
	ret[0] = &ret0

	err := common.CurrentProxyCtx.Serialize(args, &argsSerialized)
	if err != nil {
		return err
	}

	res, err := common.CurrentProxyCtx.RouteCall(r.Reference, false, false, "DoNothing", argsSerialized, *PrototypeReference)
	if err != nil {
		return err
	}

	resultContainer := foundation.Result{
		Returns: ret,
	}
	err = common.CurrentProxyCtx.Deserialize(res, &resultContainer)
	if err != nil {
		return err
	}
	if resultContainer.Error != nil {
		err = resultContainer.Error
		return err
	}
	if ret0 != nil {
		return ret0
	}
	return nil
}

// DoNothingAsImmutable is proxy generated method
func (r *Second) DoNothingAsImmutable() error {
	var args [0]interface{}

	var argsSerialized []byte

	ret := make([]interface{}, 1)
	var ret0 *foundation.Error
	ret[0] = &ret0

	err := common.CurrentProxyCtx.Serialize(args, &argsSerialized)
	if err != nil {
		return err
	}

	res, err := common.CurrentProxyCtx.RouteCall(r.Reference, true, false, "DoNothing", argsSerialized, *PrototypeReference)
	if err != nil {
		return err
	}

	resultContainer := foundation.Result{
		Returns: ret,
	}
	err = common.CurrentProxyCtx.Deserialize(res, &resultContainer)
	if err != nil {
		return err
	}
	if resultContainer.Error != nil {
		err = resultContainer.Error
		return err
	}
	if ret0 != nil {
		return ret0
	}
	return nil
}

// Get is proxy generated method
func (r *Second) Get() (int, error) {
	var args [0]interface{}

	var argsSerialized []byte

	ret := make([]interface{}, 2)
	var ret0 int
	ret[0] = &ret0
	var ret1 *foundation.Error
	ret[1] = &ret1

	err := common.CurrentProxyCtx.Serialize(args, &argsSerialized)
	if err != nil {
		return ret0, err
	}

	res, err := common.CurrentProxyCtx.RouteCall(r.Reference, false, false, "Get", argsSerialized, *PrototypeReference)
	if err != nil {
		return ret0, err
	}

	resultContainer := foundation.Result{
		Returns: ret,
	}
	err = common.CurrentProxyCtx.Deserialize(res, &resultContainer)
	if err != nil {
		return ret0, err
	}
	if resultContainer.Error != nil {
		err = resultContainer.Error
		return ret0, err
	}
	if ret1 != nil {
		return ret0, ret1
	}
	return ret0, nil
}

// GetAsImmutable is proxy generated method
func (r *Second) GetAsImmutable() (int, error) {
	var args [0]interface{}

	var argsSerialized []byte

	ret := make([]interface{}, 2)
	var ret0 int
	ret[0] = &ret0
	var ret1 *foundation.Error
	ret[1] = &ret1

	err := common.CurrentProxyCtx.Serialize(args, &argsSerialized)
	if err != nil {
		return ret0, err
	}

	res, err := common.CurrentProxyCtx.RouteCall(r.Reference, true, false, "Get", argsSerialized, *PrototypeReference)
	if err != nil {
		return ret0, err
	}

	resultContainer := foundation.Result{
		Returns: ret,
	}
	err = common.CurrentProxyCtx.Deserialize(res, &resultContainer)
	if err != nil {
		return ret0, err
	}
	if resultContainer.Error != nil {
		err = resultContainer.Error
		return ret0, err
	}
	if ret1 != nil {
		return ret0, ret1
	}
	return ret0, nil
}

// Hello is proxy generated method
func (r *Second) Hello(s string) (string, error) {
	var args [1]interface{}
	args[0] = s

	var argsSerialized []byte

	ret := make([]interface{}, 2)
	var ret0 string
	ret[0] = &ret0
	var ret1 *foundation.Error
	ret[1] = &ret1

	err := common.CurrentProxyCtx.Serialize(args, &argsSerialized)
	if err != nil {
		return ret0, err
	}

	res, err := common.CurrentProxyCtx.RouteCall(r.Reference, false, false, "Hello", argsSerialized, *PrototypeReference)
	if err != nil {
		return ret0, err
	}

	resultContainer := foundation.Result{
		Returns: ret,
	}
	err = common.CurrentProxyCtx.Deserialize(res, &resultContainer)
	if err != nil {
		return ret0, err
	}
	if resultContainer.Error != nil {
		err = resultContainer.Error
		return ret0, err
	}
	if ret1 != nil {
		return ret0, ret1
	}
	return ret0, nil
}

// HelloAsImmutable is proxy generated method
func (r *Second) HelloAsImmutable(s string) (string, error) {
	var args [1]interface{}
	args[0] = s

	var argsSerialized []byte

	ret := make([]interface{}, 2)
	var ret0 string
	ret[0] = &ret0
	var ret1 *foundation.Error
	ret[1] = &ret1

	err := common.CurrentProxyCtx.Serialize(args, &argsSerialized)
	if err != nil {
		return ret0, err
	}

	res, err := common.CurrentProxyCtx.RouteCall(r.Reference, true, false, "Hello", argsSerialized, *PrototypeReference)
	if err != nil {
		return ret0, err
	}

	resultContainer := foundation.Result{
		Returns: ret,
	}
	err = common.CurrentProxyCtx.Deserialize(res, &resultContainer)
	if err != nil {
		return ret0, err
	}
	if resultContainer.Error != nil {
		err = resultContainer.Error
		return ret0, err
	}
	if ret1 != nil {
		return ret0, ret1
	}
	return ret0, nil
}

// GetPayload is proxy generated method
func (r *Second) GetPayload() (Payload, error) {
	var args [0]interface{}

	var argsSerialized []byte

	ret := make([]interface{}, 2)
	var ret0 Payload
	ret[0] = &ret0
	var ret1 *foundation.Error
	ret[1] = &ret1

	err := common.CurrentProxyCtx.Serialize(args, &argsSerialized)
	if err != nil {
		return ret0, err
	}

	res, err := common.CurrentProxyCtx.RouteCall(r.Reference, false, false, "GetPayload", argsSerialized, *PrototypeReference)
	if err != nil {
		return ret0, err
	}

	resultContainer := foundation.Result{
		Returns: ret,
	}
	err = common.CurrentProxyCtx.Deserialize(res, &resultContainer)
	if err != nil {
		return ret0, err
	}
	if resultContainer.Error != nil {
		err = resultContainer.Error
		return ret0, err
	}
	if ret1 != nil {
		return ret0, ret1
	}
	return ret0, nil
}

// GetPayloadAsImmutable is proxy generated method
func (r *Second) GetPayloadAsImmutable() (Payload, error) {
	var args [0]interface{}

	var argsSerialized []byte

	ret := make([]interface{}, 2)
	var ret0 Payload
	ret[0] = &ret0
	var ret1 *foundation.Error
	ret[1] = &ret1

	err := common.CurrentProxyCtx.Serialize(args, &argsSerialized)
	if err != nil {
		return ret0, err
	}

	res, err := common.CurrentProxyCtx.RouteCall(r.Reference, true, false, "GetPayload", argsSerialized, *PrototypeReference)
	if err != nil {
		return ret0, err
	}

	resultContainer := foundation.Result{
		Returns: ret,
	}
	err = common.CurrentProxyCtx.Deserialize(res, &resultContainer)
	if err != nil {
		return ret0, err
	}
	if resultContainer.Error != nil {
		err = resultContainer.Error
		return ret0, err
	}
	if ret1 != nil {
		return ret0, ret1
	}
	return ret0, nil
}

// SetPayload is proxy generated method
func (r *Second) SetPayload(P Payload) error {
	var args [1]interface{}
	args[0] = P

	var argsSerialized []byte

	ret := make([]interface{}, 1)
	var ret0 *foundation.Error
	ret[0] = &ret0

	err := common.CurrentProxyCtx.Serialize(args, &argsSerialized)
	if err != nil {
		return err
	}

	res, err := common.CurrentProxyCtx.RouteCall(r.Reference, false, false, "SetPayload", argsSerialized, *PrototypeReference)
	if err != nil {
		return err
	}

	resultContainer := foundation.Result{
		Returns: ret,
	}
	err = common.CurrentProxyCtx.Deserialize(res, &resultContainer)
	if err != nil {
		return err
	}
	if resultContainer.Error != nil {
		err = resultContainer.Error
		return err
	}
	if ret0 != nil {
		return ret0
	}
	return nil
}

// SetPayloadAsImmutable is proxy generated method
func (r *Second) SetPayloadAsImmutable(P Payload) error {
	var args [1]interface{}
	args[0] = P

	var argsSerialized []byte

	ret := make([]interface{}, 1)
	var ret0 *foundation.Error
	ret[0] = &ret0

	err := common.CurrentProxyCtx.Serialize(args, &argsSerialized)
	if err != nil {
		return err
	}

	res, err := common.CurrentProxyCtx.RouteCall(r.Reference, true, false, "SetPayload", argsSerialized, *PrototypeReference)
	if err != nil {
		return err
	}

	resultContainer := foundation.Result{
		Returns: ret,
	}
	err = common.CurrentProxyCtx.Deserialize(res, &resultContainer)
	if err != nil {
		return err
	}
	if resultContainer.Error != nil {
		err = resultContainer.Error
		return err
	}
	if ret0 != nil {
		return ret0
	}
	return nil
}

// GetPayloadString is proxy generated method
func (r *Second) GetPayloadString() (string, error) {
	var args [0]interface{}

	var argsSerialized []byte

	ret := make([]interface{}, 2)
	var ret0 string
	ret[0] = &ret0
	var ret1 *foundation.Error
	ret[1] = &ret1

	err := common.CurrentProxyCtx.Serialize(args, &argsSerialized)
	if err != nil {
		return ret0, err
	}

	res, err := common.CurrentProxyCtx.RouteCall(r.Reference, false, false, "GetPayloadString", argsSerialized, *PrototypeReference)
	if err != nil {
		return ret0, err
	}

	resultContainer := foundation.Result{
		Returns: ret,
	}
	err = common.CurrentProxyCtx.Deserialize(res, &resultContainer)
	if err != nil {
		return ret0, err
	}
	if resultContainer.Error != nil {
		err = resultContainer.Error
		return ret0, err
	}
	if ret1 != nil {
		return ret0, ret1
	}
	return ret0, nil
}

// GetPayloadStringAsImmutable is proxy generated method
func (r *Second) GetPayloadStringAsImmutable() (string, error) {
	var args [0]interface{}

	var argsSerialized []byte

	ret := make([]interface{}, 2)
	var ret0 string
	ret[0] = &ret0
	var ret1 *foundation.Error
	ret[1] = &ret1

	err := common.CurrentProxyCtx.Serialize(args, &argsSerialized)
	if err != nil {
		return ret0, err
	}

	res, err := common.CurrentProxyCtx.RouteCall(r.Reference, true, false, "GetPayloadString", argsSerialized, *PrototypeReference)
	if err != nil {
		return ret0, err
	}

	resultContainer := foundation.Result{
		Returns: ret,
	}
	err = common.CurrentProxyCtx.Deserialize(res, &resultContainer)
	if err != nil {
		return ret0, err
	}
	if resultContainer.Error != nil {
		err = resultContainer.Error
		return ret0, err
	}
	if ret1 != nil {
		return ret0, ret1
	}
	return ret0, nil
}

// GetBalance is proxy generated method
func (r *Second) GetBalance() (int, error) {
	var args [0]interface{}

	var argsSerialized []byte

	ret := make([]interface{}, 2)
	var ret0 int
	ret[0] = &ret0
	var ret1 *foundation.Error
	ret[1] = &ret1

	err := common.CurrentProxyCtx.Serialize(args, &argsSerialized)
	if err != nil {
		return ret0, err
	}

	res, err := common.CurrentProxyCtx.RouteCall(r.Reference, false, false, "GetBalance", argsSerialized, *PrototypeReference)
	if err != nil {
		return ret0, err
	}

	resultContainer := foundation.Result{
		Returns: ret,
	}
	err = common.CurrentProxyCtx.Deserialize(res, &resultContainer)
	if err != nil {
		return ret0, err
	}
	if resultContainer.Error != nil {
		err = resultContainer.Error
		return ret0, err
	}
	if ret1 != nil {
		return ret0, ret1
	}
	return ret0, nil
}

// GetBalanceAsImmutable is proxy generated method
func (r *Second) GetBalanceAsImmutable() (int, error) {
	var args [0]interface{}

	var argsSerialized []byte

	ret := make([]interface{}, 2)
	var ret0 int
	ret[0] = &ret0
	var ret1 *foundation.Error
	ret[1] = &ret1

	err := common.CurrentProxyCtx.Serialize(args, &argsSerialized)
	if err != nil {
		return ret0, err
	}

	res, err := common.CurrentProxyCtx.RouteCall(r.Reference, true, false, "GetBalance", argsSerialized, *PrototypeReference)
	if err != nil {
		return ret0, err
	}

	resultContainer := foundation.Result{
		Returns: ret,
	}
	err = common.CurrentProxyCtx.Deserialize(res, &resultContainer)
	if err != nil {
		return ret0, err
	}
	if resultContainer.Error != nil {
		err = resultContainer.Error
		return ret0, err
	}
	if ret1 != nil {
		return ret0, ret1
	}
	return ret0, nil
}

// Accept is proxy generated method
func (r *Second) Accept(amount int) error {
	var args [1]interface{}
	args[0] = amount

	var argsSerialized []byte

	ret := make([]interface{}, 1)
	var ret0 *foundation.Error
	ret[0] = &ret0

	err := common.CurrentProxyCtx.Serialize(args, &argsSerialized)
	if err != nil {
		return err
	}

	_, err = common.CurrentProxyCtx.RouteCall(r.Reference, false, true, "Accept", argsSerialized, *PrototypeReference)
	if err != nil {
		return err
	}
	return nil
}

// AnError is proxy generated method
func (r *Second) AnError() error {
	var args [0]interface{}

	var argsSerialized []byte

	ret := make([]interface{}, 1)
	var ret0 *foundation.Error
	ret[0] = &ret0

	err := common.CurrentProxyCtx.Serialize(args, &argsSerialized)
	if err != nil {
		return err
	}

	res, err := common.CurrentProxyCtx.RouteCall(r.Reference, false, false, "AnError", argsSerialized, *PrototypeReference)
	if err != nil {
		return err
	}

	resultContainer := foundation.Result{
		Returns: ret,
	}
	err = common.CurrentProxyCtx.Deserialize(res, &resultContainer)
	if err != nil {
		return err
	}
	if resultContainer.Error != nil {
		err = resultContainer.Error
		return err
	}
	if ret0 != nil {
		return ret0
	}
	return nil
}

// AnErrorAsImmutable is proxy generated method
func (r *Second) AnErrorAsImmutable() error {
	var args [0]interface{}

	var argsSerialized []byte

	ret := make([]interface{}, 1)
	var ret0 *foundation.Error
	ret[0] = &ret0

	err := common.CurrentProxyCtx.Serialize(args, &argsSerialized)
	if err != nil {
		return err
	}

	res, err := common.CurrentProxyCtx.RouteCall(r.Reference, true, false, "AnError", argsSerialized, *PrototypeReference)
	if err != nil {
		return err
	}

	resultContainer := foundation.Result{
		Returns: ret,
	}
	err = common.CurrentProxyCtx.Deserialize(res, &resultContainer)
	if err != nil {
		return err
	}
	if resultContainer.Error != nil {
		err = resultContainer.Error
		return err
	}
	if ret0 != nil {
		return ret0
	}
	return nil
}

// NoError is proxy generated method
func (r *Second) NoError() error {
	var args [0]interface{}

	var argsSerialized []byte

	ret := make([]interface{}, 1)
	var ret0 *foundation.Error
	ret[0] = &ret0

	err := common.CurrentProxyCtx.Serialize(args, &argsSerialized)
	if err != nil {
		return err
	}

	res, err := common.CurrentProxyCtx.RouteCall(r.Reference, false, false, "NoError", argsSerialized, *PrototypeReference)
	if err != nil {
		return err
	}

	resultContainer := foundation.Result{
		Returns: ret,
	}
	err = common.CurrentProxyCtx.Deserialize(res, &resultContainer)
	if err != nil {
		return err
	}
	if resultContainer.Error != nil {
		err = resultContainer.Error
		return err
	}
	if ret0 != nil {
		return ret0
	}
	return nil
}

// NoErrorAsImmutable is proxy generated method
func (r *Second) NoErrorAsImmutable() error {
	var args [0]interface{}

	var argsSerialized []byte

	ret := make([]interface{}, 1)
	var ret0 *foundation.Error
	ret[0] = &ret0

	err := common.CurrentProxyCtx.Serialize(args, &argsSerialized)
	if err != nil {
		return err
	}

	res, err := common.CurrentProxyCtx.RouteCall(r.Reference, true, false, "NoError", argsSerialized, *PrototypeReference)
	if err != nil {
		return err
	}

	resultContainer := foundation.Result{
		Returns: ret,
	}
	err = common.CurrentProxyCtx.Deserialize(res, &resultContainer)
	if err != nil {
		return err
	}
	if resultContainer.Error != nil {
		err = resultContainer.Error
		return err
	}
	if ret0 != nil {
		return ret0
	}
	return nil
}

// ReturnNil is proxy generated method
func (r *Second) ReturnNil() (*string, error) {
	var args [0]interface{}

	var argsSerialized []byte

	ret := make([]interface{}, 2)
	var ret0 *string
	ret[0] = &ret0
	var ret1 *foundation.Error
	ret[1] = &ret1

	err := common.CurrentProxyCtx.Serialize(args, &argsSerialized)
	if err != nil {
		return ret0, err
	}

	res, err := common.CurrentProxyCtx.RouteCall(r.Reference, false, false, "ReturnNil", argsSerialized, *PrototypeReference)
	if err != nil {
		return ret0, err
	}

	resultContainer := foundation.Result{
		Returns: ret,
	}
	err = common.CurrentProxyCtx.Deserialize(res, &resultContainer)
	if err != nil {
		return ret0, err
	}
	if resultContainer.Error != nil {
		err = resultContainer.Error
		return ret0, err
	}
	if ret1 != nil {
		return ret0, ret1
	}
	return ret0, nil
}

// ReturnNilAsImmutable is proxy generated method
func (r *Second) ReturnNilAsImmutable() (*string, error) {
	var args [0]interface{}

	var argsSerialized []byte

	ret := make([]interface{}, 2)
	var ret0 *string
	ret[0] = &ret0
	var ret1 *foundation.Error
	ret[1] = &ret1

	err := common.CurrentProxyCtx.Serialize(args, &argsSerialized)
	if err != nil {
		return ret0, err
	}

	res, err := common.CurrentProxyCtx.RouteCall(r.Reference, true, false, "ReturnNil", argsSerialized, *PrototypeReference)
	if err != nil {
		return ret0, err
	}

	resultContainer := foundation.Result{
		Returns: ret,
	}
	err = common.CurrentProxyCtx.Deserialize(res, &resultContainer)
	if err != nil {
		return ret0, err
	}
	if resultContainer.Error != nil {
		err = resultContainer.Error
		return ret0, err
	}
	if ret1 != nil {
		return ret0, ret1
	}
	return ret0, nil
}
