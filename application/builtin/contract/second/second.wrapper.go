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
	"github.com/pkg/errors"
)

const PanicIsLogicalError = false

func INS_META_INFO() []map[string]string {
	result := make([]map[string]string, 0)

	{
		info := make(map[string]string, 3)
		info["Type"] = "SagaInfo"
		info["MethodName"] = "Accept"
		info["RollbackMethodName"] = "INS_FLAG_NO_ROLLBACK_METHOD"
		result = append(result, info)
	}

	return result
}

func INSMETHOD_GetCode(object []byte, data []byte) ([]byte, []byte, error) {
	ph := common.CurrentProxyCtx
	self := new(Second)

	if len(object) == 0 {
		return nil, nil, &foundation.Error{S: "[ Fake GetCode ] ( Generated Method ) Object is nil"}
	}

	err := ph.Deserialize(object, self)
	if err != nil {
		e := &foundation.Error{S: "[ Fake GetCode ] ( Generated Method ) Can't deserialize args.Data: " + err.Error()}
		return nil, nil, e
	}

	state := []byte{}
	err = ph.Serialize(self, &state)
	if err != nil {
		return nil, nil, err
	}

	ret := []byte{}
	err = ph.Serialize([]interface{}{self.GetCode().Bytes()}, &ret)

	return state, ret, err
}

func INSMETHOD_GetPrototype(object []byte, data []byte) ([]byte, []byte, error) {
	ph := common.CurrentProxyCtx
	self := new(Second)

	if len(object) == 0 {
		return nil, nil, &foundation.Error{S: "[ Fake GetPrototype ] ( Generated Method ) Object is nil"}
	}

	err := ph.Deserialize(object, self)
	if err != nil {
		e := &foundation.Error{S: "[ Fake GetPrototype ] ( Generated Method ) Can't deserialize args.Data: " + err.Error()}
		return nil, nil, e
	}

	state := []byte{}
	err = ph.Serialize(self, &state)
	if err != nil {
		return nil, nil, err
	}

	ret := []byte{}
	err = ph.Serialize([]interface{}{self.GetPrototype().Bytes()}, &ret)

	return state, ret, err
}

func INSMETHOD_GetName(object []byte, data []byte) (newState []byte, result []byte, err error) {
	ph := common.CurrentProxyCtx
	ph.SetSystemError(nil)

	self := new(Second)

	if len(object) == 0 {
		err = &foundation.Error{S: "[ FakeGetName ] ( INSMETHOD_* ) ( Generated Method ) Object is nil"}
		return
	}

	err = ph.Deserialize(object, self)
	if err != nil {
		err = &foundation.Error{S: "[ FakeGetName ] ( INSMETHOD_* ) ( Generated Method ) Can't deserialize args.Data: " + err.Error()}
		return
	}

	args := []interface{}{}

	err = ph.Deserialize(data, &args)
	if err != nil {
		err = &foundation.Error{S: "[ FakeGetName ] ( INSMETHOD_* ) ( Generated Method ) Can't deserialize args.Arguments: " + err.Error()}
		return
	}

	var ret0 string
	var ret1 error

	serializeResults := func() error {
		return ph.Serialize(
			foundation.Result{Returns: []interface{}{ret0, ret1}},
			&result,
		)
	}

	needRecover := true
	defer func() {
		if !needRecover {
			return
		}
		if r := recover(); r != nil {
			recoveredError := errors.Wrap(errors.Errorf("%v", r), "Failed to execute method (panic)")
			recoveredError = ph.MakeErrorSerializable(recoveredError)

			if PanicIsLogicalError {
				ret1 = recoveredError

				newState = object
				err = serializeResults()
				if err == nil {
					newState = object
				}
			} else {
				err = recoveredError
			}
		}
	}()

	ret0, ret1 = self.GetName()

	needRecover = false

	if ph.GetSystemError() != nil {
		return nil, nil, ph.GetSystemError()
	}

	err = ph.Serialize(self, &newState)
	if err != nil {
		return nil, nil, err
	}

	ret1 = ph.MakeErrorSerializable(ret1)

	err = serializeResults()
	if err != nil {
		return
	}

	return
}

func INSMETHOD_DoNothing(object []byte, data []byte) (newState []byte, result []byte, err error) {
	ph := common.CurrentProxyCtx
	ph.SetSystemError(nil)

	self := new(Second)

	if len(object) == 0 {
		err = &foundation.Error{S: "[ FakeDoNothing ] ( INSMETHOD_* ) ( Generated Method ) Object is nil"}
		return
	}

	err = ph.Deserialize(object, self)
	if err != nil {
		err = &foundation.Error{S: "[ FakeDoNothing ] ( INSMETHOD_* ) ( Generated Method ) Can't deserialize args.Data: " + err.Error()}
		return
	}

	args := []interface{}{}

	err = ph.Deserialize(data, &args)
	if err != nil {
		err = &foundation.Error{S: "[ FakeDoNothing ] ( INSMETHOD_* ) ( Generated Method ) Can't deserialize args.Arguments: " + err.Error()}
		return
	}

	var ret0 error

	serializeResults := func() error {
		return ph.Serialize(
			foundation.Result{Returns: []interface{}{ret0}},
			&result,
		)
	}

	needRecover := true
	defer func() {
		if !needRecover {
			return
		}
		if r := recover(); r != nil {
			recoveredError := errors.Wrap(errors.Errorf("%v", r), "Failed to execute method (panic)")
			recoveredError = ph.MakeErrorSerializable(recoveredError)

			if PanicIsLogicalError {
				ret0 = recoveredError

				newState = object
				err = serializeResults()
				if err == nil {
					newState = object
				}
			} else {
				err = recoveredError
			}
		}
	}()

	ret0 = self.DoNothing()

	needRecover = false

	if ph.GetSystemError() != nil {
		return nil, nil, ph.GetSystemError()
	}

	err = ph.Serialize(self, &newState)
	if err != nil {
		return nil, nil, err
	}

	ret0 = ph.MakeErrorSerializable(ret0)

	err = serializeResults()
	if err != nil {
		return
	}

	return
}

func INSMETHOD_Get(object []byte, data []byte) (newState []byte, result []byte, err error) {
	ph := common.CurrentProxyCtx
	ph.SetSystemError(nil)

	self := new(Second)

	if len(object) == 0 {
		err = &foundation.Error{S: "[ FakeGet ] ( INSMETHOD_* ) ( Generated Method ) Object is nil"}
		return
	}

	err = ph.Deserialize(object, self)
	if err != nil {
		err = &foundation.Error{S: "[ FakeGet ] ( INSMETHOD_* ) ( Generated Method ) Can't deserialize args.Data: " + err.Error()}
		return
	}

	args := []interface{}{}

	err = ph.Deserialize(data, &args)
	if err != nil {
		err = &foundation.Error{S: "[ FakeGet ] ( INSMETHOD_* ) ( Generated Method ) Can't deserialize args.Arguments: " + err.Error()}
		return
	}

	var ret0 int
	var ret1 error

	serializeResults := func() error {
		return ph.Serialize(
			foundation.Result{Returns: []interface{}{ret0, ret1}},
			&result,
		)
	}

	needRecover := true
	defer func() {
		if !needRecover {
			return
		}
		if r := recover(); r != nil {
			recoveredError := errors.Wrap(errors.Errorf("%v", r), "Failed to execute method (panic)")
			recoveredError = ph.MakeErrorSerializable(recoveredError)

			if PanicIsLogicalError {
				ret1 = recoveredError

				newState = object
				err = serializeResults()
				if err == nil {
					newState = object
				}
			} else {
				err = recoveredError
			}
		}
	}()

	ret0, ret1 = self.Get()

	needRecover = false

	if ph.GetSystemError() != nil {
		return nil, nil, ph.GetSystemError()
	}

	err = ph.Serialize(self, &newState)
	if err != nil {
		return nil, nil, err
	}

	ret1 = ph.MakeErrorSerializable(ret1)

	err = serializeResults()
	if err != nil {
		return
	}

	return
}

func INSMETHOD_Hello(object []byte, data []byte) (newState []byte, result []byte, err error) {
	ph := common.CurrentProxyCtx
	ph.SetSystemError(nil)

	self := new(Second)

	if len(object) == 0 {
		err = &foundation.Error{S: "[ FakeHello ] ( INSMETHOD_* ) ( Generated Method ) Object is nil"}
		return
	}

	err = ph.Deserialize(object, self)
	if err != nil {
		err = &foundation.Error{S: "[ FakeHello ] ( INSMETHOD_* ) ( Generated Method ) Can't deserialize args.Data: " + err.Error()}
		return
	}

	args := make([]interface{}, 1)
	var args0 string
	args[0] = &args0

	err = ph.Deserialize(data, &args)
	if err != nil {
		err = &foundation.Error{S: "[ FakeHello ] ( INSMETHOD_* ) ( Generated Method ) Can't deserialize args.Arguments: " + err.Error()}
		return
	}

	var ret0 string
	var ret1 error

	serializeResults := func() error {
		return ph.Serialize(
			foundation.Result{Returns: []interface{}{ret0, ret1}},
			&result,
		)
	}

	needRecover := true
	defer func() {
		if !needRecover {
			return
		}
		if r := recover(); r != nil {
			recoveredError := errors.Wrap(errors.Errorf("%v", r), "Failed to execute method (panic)")
			recoveredError = ph.MakeErrorSerializable(recoveredError)

			if PanicIsLogicalError {
				ret1 = recoveredError

				newState = object
				err = serializeResults()
				if err == nil {
					newState = object
				}
			} else {
				err = recoveredError
			}
		}
	}()

	ret0, ret1 = self.Hello(args0)

	needRecover = false

	if ph.GetSystemError() != nil {
		return nil, nil, ph.GetSystemError()
	}

	err = ph.Serialize(self, &newState)
	if err != nil {
		return nil, nil, err
	}

	ret1 = ph.MakeErrorSerializable(ret1)

	err = serializeResults()
	if err != nil {
		return
	}

	return
}

func INSMETHOD_GetPayload(object []byte, data []byte) (newState []byte, result []byte, err error) {
	ph := common.CurrentProxyCtx
	ph.SetSystemError(nil)

	self := new(Second)

	if len(object) == 0 {
		err = &foundation.Error{S: "[ FakeGetPayload ] ( INSMETHOD_* ) ( Generated Method ) Object is nil"}
		return
	}

	err = ph.Deserialize(object, self)
	if err != nil {
		err = &foundation.Error{S: "[ FakeGetPayload ] ( INSMETHOD_* ) ( Generated Method ) Can't deserialize args.Data: " + err.Error()}
		return
	}

	args := []interface{}{}

	err = ph.Deserialize(data, &args)
	if err != nil {
		err = &foundation.Error{S: "[ FakeGetPayload ] ( INSMETHOD_* ) ( Generated Method ) Can't deserialize args.Arguments: " + err.Error()}
		return
	}

	var ret0 Payload
	var ret1 error

	serializeResults := func() error {
		return ph.Serialize(
			foundation.Result{Returns: []interface{}{ret0, ret1}},
			&result,
		)
	}

	needRecover := true
	defer func() {
		if !needRecover {
			return
		}
		if r := recover(); r != nil {
			recoveredError := errors.Wrap(errors.Errorf("%v", r), "Failed to execute method (panic)")
			recoveredError = ph.MakeErrorSerializable(recoveredError)

			if PanicIsLogicalError {
				ret1 = recoveredError

				newState = object
				err = serializeResults()
				if err == nil {
					newState = object
				}
			} else {
				err = recoveredError
			}
		}
	}()

	ret0, ret1 = self.GetPayload()

	needRecover = false

	if ph.GetSystemError() != nil {
		return nil, nil, ph.GetSystemError()
	}

	err = ph.Serialize(self, &newState)
	if err != nil {
		return nil, nil, err
	}

	ret1 = ph.MakeErrorSerializable(ret1)

	err = serializeResults()
	if err != nil {
		return
	}

	return
}

func INSMETHOD_SetPayload(object []byte, data []byte) (newState []byte, result []byte, err error) {
	ph := common.CurrentProxyCtx
	ph.SetSystemError(nil)

	self := new(Second)

	if len(object) == 0 {
		err = &foundation.Error{S: "[ FakeSetPayload ] ( INSMETHOD_* ) ( Generated Method ) Object is nil"}
		return
	}

	err = ph.Deserialize(object, self)
	if err != nil {
		err = &foundation.Error{S: "[ FakeSetPayload ] ( INSMETHOD_* ) ( Generated Method ) Can't deserialize args.Data: " + err.Error()}
		return
	}

	args := make([]interface{}, 1)
	var args0 Payload
	args[0] = &args0

	err = ph.Deserialize(data, &args)
	if err != nil {
		err = &foundation.Error{S: "[ FakeSetPayload ] ( INSMETHOD_* ) ( Generated Method ) Can't deserialize args.Arguments: " + err.Error()}
		return
	}

	var ret0 error

	serializeResults := func() error {
		return ph.Serialize(
			foundation.Result{Returns: []interface{}{ret0}},
			&result,
		)
	}

	needRecover := true
	defer func() {
		if !needRecover {
			return
		}
		if r := recover(); r != nil {
			recoveredError := errors.Wrap(errors.Errorf("%v", r), "Failed to execute method (panic)")
			recoveredError = ph.MakeErrorSerializable(recoveredError)

			if PanicIsLogicalError {
				ret0 = recoveredError

				newState = object
				err = serializeResults()
				if err == nil {
					newState = object
				}
			} else {
				err = recoveredError
			}
		}
	}()

	ret0 = self.SetPayload(args0)

	needRecover = false

	if ph.GetSystemError() != nil {
		return nil, nil, ph.GetSystemError()
	}

	err = ph.Serialize(self, &newState)
	if err != nil {
		return nil, nil, err
	}

	ret0 = ph.MakeErrorSerializable(ret0)

	err = serializeResults()
	if err != nil {
		return
	}

	return
}

func INSMETHOD_GetPayloadString(object []byte, data []byte) (newState []byte, result []byte, err error) {
	ph := common.CurrentProxyCtx
	ph.SetSystemError(nil)

	self := new(Second)

	if len(object) == 0 {
		err = &foundation.Error{S: "[ FakeGetPayloadString ] ( INSMETHOD_* ) ( Generated Method ) Object is nil"}
		return
	}

	err = ph.Deserialize(object, self)
	if err != nil {
		err = &foundation.Error{S: "[ FakeGetPayloadString ] ( INSMETHOD_* ) ( Generated Method ) Can't deserialize args.Data: " + err.Error()}
		return
	}

	args := []interface{}{}

	err = ph.Deserialize(data, &args)
	if err != nil {
		err = &foundation.Error{S: "[ FakeGetPayloadString ] ( INSMETHOD_* ) ( Generated Method ) Can't deserialize args.Arguments: " + err.Error()}
		return
	}

	var ret0 string
	var ret1 error

	serializeResults := func() error {
		return ph.Serialize(
			foundation.Result{Returns: []interface{}{ret0, ret1}},
			&result,
		)
	}

	needRecover := true
	defer func() {
		if !needRecover {
			return
		}
		if r := recover(); r != nil {
			recoveredError := errors.Wrap(errors.Errorf("%v", r), "Failed to execute method (panic)")
			recoveredError = ph.MakeErrorSerializable(recoveredError)

			if PanicIsLogicalError {
				ret1 = recoveredError

				newState = object
				err = serializeResults()
				if err == nil {
					newState = object
				}
			} else {
				err = recoveredError
			}
		}
	}()

	ret0, ret1 = self.GetPayloadString()

	needRecover = false

	if ph.GetSystemError() != nil {
		return nil, nil, ph.GetSystemError()
	}

	err = ph.Serialize(self, &newState)
	if err != nil {
		return nil, nil, err
	}

	ret1 = ph.MakeErrorSerializable(ret1)

	err = serializeResults()
	if err != nil {
		return
	}

	return
}

func INSMETHOD_GetBalance(object []byte, data []byte) (newState []byte, result []byte, err error) {
	ph := common.CurrentProxyCtx
	ph.SetSystemError(nil)

	self := new(Second)

	if len(object) == 0 {
		err = &foundation.Error{S: "[ FakeGetBalance ] ( INSMETHOD_* ) ( Generated Method ) Object is nil"}
		return
	}

	err = ph.Deserialize(object, self)
	if err != nil {
		err = &foundation.Error{S: "[ FakeGetBalance ] ( INSMETHOD_* ) ( Generated Method ) Can't deserialize args.Data: " + err.Error()}
		return
	}

	args := []interface{}{}

	err = ph.Deserialize(data, &args)
	if err != nil {
		err = &foundation.Error{S: "[ FakeGetBalance ] ( INSMETHOD_* ) ( Generated Method ) Can't deserialize args.Arguments: " + err.Error()}
		return
	}

	var ret0 int
	var ret1 error

	serializeResults := func() error {
		return ph.Serialize(
			foundation.Result{Returns: []interface{}{ret0, ret1}},
			&result,
		)
	}

	needRecover := true
	defer func() {
		if !needRecover {
			return
		}
		if r := recover(); r != nil {
			recoveredError := errors.Wrap(errors.Errorf("%v", r), "Failed to execute method (panic)")
			recoveredError = ph.MakeErrorSerializable(recoveredError)

			if PanicIsLogicalError {
				ret1 = recoveredError

				newState = object
				err = serializeResults()
				if err == nil {
					newState = object
				}
			} else {
				err = recoveredError
			}
		}
	}()

	ret0, ret1 = self.GetBalance()

	needRecover = false

	if ph.GetSystemError() != nil {
		return nil, nil, ph.GetSystemError()
	}

	err = ph.Serialize(self, &newState)
	if err != nil {
		return nil, nil, err
	}

	ret1 = ph.MakeErrorSerializable(ret1)

	err = serializeResults()
	if err != nil {
		return
	}

	return
}

func INSMETHOD_Accept(object []byte, data []byte) (newState []byte, result []byte, err error) {
	ph := common.CurrentProxyCtx
	ph.SetSystemError(nil)

	self := new(Second)

	if len(object) == 0 {
		err = &foundation.Error{S: "[ FakeAccept ] ( INSMETHOD_* ) ( Generated Method ) Object is nil"}
		return
	}

	err = ph.Deserialize(object, self)
	if err != nil {
		err = &foundation.Error{S: "[ FakeAccept ] ( INSMETHOD_* ) ( Generated Method ) Can't deserialize args.Data: " + err.Error()}
		return
	}

	args := make([]interface{}, 1)
	var args0 int
	args[0] = &args0

	err = ph.Deserialize(data, &args)
	if err != nil {
		err = &foundation.Error{S: "[ FakeAccept ] ( INSMETHOD_* ) ( Generated Method ) Can't deserialize args.Arguments: " + err.Error()}
		return
	}

	var ret0 error

	serializeResults := func() error {
		return ph.Serialize(
			foundation.Result{Returns: []interface{}{ret0}},
			&result,
		)
	}

	needRecover := true
	defer func() {
		if !needRecover {
			return
		}
		if r := recover(); r != nil {
			recoveredError := errors.Wrap(errors.Errorf("%v", r), "Failed to execute method (panic)")
			recoveredError = ph.MakeErrorSerializable(recoveredError)

			if PanicIsLogicalError {
				ret0 = recoveredError

				newState = object
				err = serializeResults()
				if err == nil {
					newState = object
				}
			} else {
				err = recoveredError
			}
		}
	}()

	ret0 = self.Accept(args0)

	needRecover = false

	if ph.GetSystemError() != nil {
		return nil, nil, ph.GetSystemError()
	}

	err = ph.Serialize(self, &newState)
	if err != nil {
		return nil, nil, err
	}

	ret0 = ph.MakeErrorSerializable(ret0)

	err = serializeResults()
	if err != nil {
		return
	}

	return
}

func INSMETHOD_AnError(object []byte, data []byte) (newState []byte, result []byte, err error) {
	ph := common.CurrentProxyCtx
	ph.SetSystemError(nil)

	self := new(Second)

	if len(object) == 0 {
		err = &foundation.Error{S: "[ FakeAnError ] ( INSMETHOD_* ) ( Generated Method ) Object is nil"}
		return
	}

	err = ph.Deserialize(object, self)
	if err != nil {
		err = &foundation.Error{S: "[ FakeAnError ] ( INSMETHOD_* ) ( Generated Method ) Can't deserialize args.Data: " + err.Error()}
		return
	}

	args := []interface{}{}

	err = ph.Deserialize(data, &args)
	if err != nil {
		err = &foundation.Error{S: "[ FakeAnError ] ( INSMETHOD_* ) ( Generated Method ) Can't deserialize args.Arguments: " + err.Error()}
		return
	}

	var ret0 error

	serializeResults := func() error {
		return ph.Serialize(
			foundation.Result{Returns: []interface{}{ret0}},
			&result,
		)
	}

	needRecover := true
	defer func() {
		if !needRecover {
			return
		}
		if r := recover(); r != nil {
			recoveredError := errors.Wrap(errors.Errorf("%v", r), "Failed to execute method (panic)")
			recoveredError = ph.MakeErrorSerializable(recoveredError)

			if PanicIsLogicalError {
				ret0 = recoveredError

				newState = object
				err = serializeResults()
				if err == nil {
					newState = object
				}
			} else {
				err = recoveredError
			}
		}
	}()

	ret0 = self.AnError()

	needRecover = false

	if ph.GetSystemError() != nil {
		return nil, nil, ph.GetSystemError()
	}

	err = ph.Serialize(self, &newState)
	if err != nil {
		return nil, nil, err
	}

	ret0 = ph.MakeErrorSerializable(ret0)

	err = serializeResults()
	if err != nil {
		return
	}

	return
}

func INSMETHOD_NoError(object []byte, data []byte) (newState []byte, result []byte, err error) {
	ph := common.CurrentProxyCtx
	ph.SetSystemError(nil)

	self := new(Second)

	if len(object) == 0 {
		err = &foundation.Error{S: "[ FakeNoError ] ( INSMETHOD_* ) ( Generated Method ) Object is nil"}
		return
	}

	err = ph.Deserialize(object, self)
	if err != nil {
		err = &foundation.Error{S: "[ FakeNoError ] ( INSMETHOD_* ) ( Generated Method ) Can't deserialize args.Data: " + err.Error()}
		return
	}

	args := []interface{}{}

	err = ph.Deserialize(data, &args)
	if err != nil {
		err = &foundation.Error{S: "[ FakeNoError ] ( INSMETHOD_* ) ( Generated Method ) Can't deserialize args.Arguments: " + err.Error()}
		return
	}

	var ret0 error

	serializeResults := func() error {
		return ph.Serialize(
			foundation.Result{Returns: []interface{}{ret0}},
			&result,
		)
	}

	needRecover := true
	defer func() {
		if !needRecover {
			return
		}
		if r := recover(); r != nil {
			recoveredError := errors.Wrap(errors.Errorf("%v", r), "Failed to execute method (panic)")
			recoveredError = ph.MakeErrorSerializable(recoveredError)

			if PanicIsLogicalError {
				ret0 = recoveredError

				newState = object
				err = serializeResults()
				if err == nil {
					newState = object
				}
			} else {
				err = recoveredError
			}
		}
	}()

	ret0 = self.NoError()

	needRecover = false

	if ph.GetSystemError() != nil {
		return nil, nil, ph.GetSystemError()
	}

	err = ph.Serialize(self, &newState)
	if err != nil {
		return nil, nil, err
	}

	ret0 = ph.MakeErrorSerializable(ret0)

	err = serializeResults()
	if err != nil {
		return
	}

	return
}

func INSMETHOD_ReturnNil(object []byte, data []byte) (newState []byte, result []byte, err error) {
	ph := common.CurrentProxyCtx
	ph.SetSystemError(nil)

	self := new(Second)

	if len(object) == 0 {
		err = &foundation.Error{S: "[ FakeReturnNil ] ( INSMETHOD_* ) ( Generated Method ) Object is nil"}
		return
	}

	err = ph.Deserialize(object, self)
	if err != nil {
		err = &foundation.Error{S: "[ FakeReturnNil ] ( INSMETHOD_* ) ( Generated Method ) Can't deserialize args.Data: " + err.Error()}
		return
	}

	args := []interface{}{}

	err = ph.Deserialize(data, &args)
	if err != nil {
		err = &foundation.Error{S: "[ FakeReturnNil ] ( INSMETHOD_* ) ( Generated Method ) Can't deserialize args.Arguments: " + err.Error()}
		return
	}

	var ret0 *string
	var ret1 error

	serializeResults := func() error {
		return ph.Serialize(
			foundation.Result{Returns: []interface{}{ret0, ret1}},
			&result,
		)
	}

	needRecover := true
	defer func() {
		if !needRecover {
			return
		}
		if r := recover(); r != nil {
			recoveredError := errors.Wrap(errors.Errorf("%v", r), "Failed to execute method (panic)")
			recoveredError = ph.MakeErrorSerializable(recoveredError)

			if PanicIsLogicalError {
				ret1 = recoveredError

				newState = object
				err = serializeResults()
				if err == nil {
					newState = object
				}
			} else {
				err = recoveredError
			}
		}
	}()

	ret0, ret1 = self.ReturnNil()

	needRecover = false

	if ph.GetSystemError() != nil {
		return nil, nil, ph.GetSystemError()
	}

	err = ph.Serialize(self, &newState)
	if err != nil {
		return nil, nil, err
	}

	ret1 = ph.MakeErrorSerializable(ret1)

	err = serializeResults()
	if err != nil {
		return
	}

	return
}

func INSCONSTRUCTOR_New(ref insolar.Reference, data []byte) (state []byte, result []byte, err error) {
	ph := common.CurrentProxyCtx
	ph.SetSystemError(nil)

	args := []interface{}{}

	err = ph.Deserialize(data, &args)
	if err != nil {
		err = &foundation.Error{S: "[ FakeNew ] ( INSCONSTRUCTOR_* ) ( Generated Method ) Can't deserialize args.Arguments: " + err.Error()}
		return
	}

	var ret0 *Second
	var ret1 error

	serializeResults := func() error {
		return ph.Serialize(
			foundation.Result{Returns: []interface{}{ref, ret1}},
			&result,
		)
	}

	needRecover := true
	defer func() {
		if !needRecover {
			return
		}
		if r := recover(); r != nil {
			recoveredError := errors.Wrap(errors.Errorf("%v", r), "Failed to execute constructor (panic)")
			recoveredError = ph.MakeErrorSerializable(recoveredError)

			if PanicIsLogicalError {
				ret1 = recoveredError

				err = serializeResults()
				if err == nil {
					state = data
				}
			} else {
				err = recoveredError
			}
		}
	}()

	ret0, ret1 = New()

	needRecover = false

	ret1 = ph.MakeErrorSerializable(ret1)
	if ret0 == nil && ret1 == nil {
		ret1 = &foundation.Error{S: "constructor returned nil"}
	}

	if ph.GetSystemError() != nil {
		err = ph.GetSystemError()
		return
	}

	err = serializeResults()
	if err != nil {
		return
	}

	if ret1 != nil {
		// logical error, the result should be registered with type RequestSideEffectNone
		state = nil
		return
	}

	err = ph.Serialize(ret0, &state)
	if err != nil {
		return
	}

	return
}

func INSCONSTRUCTOR_NewWithOne(ref insolar.Reference, data []byte) (state []byte, result []byte, err error) {
	ph := common.CurrentProxyCtx
	ph.SetSystemError(nil)

	args := make([]interface{}, 1)
	var args0 int
	args[0] = &args0

	err = ph.Deserialize(data, &args)
	if err != nil {
		err = &foundation.Error{S: "[ FakeNewWithOne ] ( INSCONSTRUCTOR_* ) ( Generated Method ) Can't deserialize args.Arguments: " + err.Error()}
		return
	}

	var ret0 *Second
	var ret1 error

	serializeResults := func() error {
		return ph.Serialize(
			foundation.Result{Returns: []interface{}{ref, ret1}},
			&result,
		)
	}

	needRecover := true
	defer func() {
		if !needRecover {
			return
		}
		if r := recover(); r != nil {
			recoveredError := errors.Wrap(errors.Errorf("%v", r), "Failed to execute constructor (panic)")
			recoveredError = ph.MakeErrorSerializable(recoveredError)

			if PanicIsLogicalError {
				ret1 = recoveredError

				err = serializeResults()
				if err == nil {
					state = data
				}
			} else {
				err = recoveredError
			}
		}
	}()

	ret0, ret1 = NewWithOne(args0)

	needRecover = false

	ret1 = ph.MakeErrorSerializable(ret1)
	if ret0 == nil && ret1 == nil {
		ret1 = &foundation.Error{S: "constructor returned nil"}
	}

	if ph.GetSystemError() != nil {
		err = ph.GetSystemError()
		return
	}

	err = serializeResults()
	if err != nil {
		return
	}

	if ret1 != nil {
		// logical error, the result should be registered with type RequestSideEffectNone
		state = nil
		return
	}

	err = ph.Serialize(ret0, &state)
	if err != nil {
		return
	}

	return
}

func INSCONSTRUCTOR_NewWithX(ref insolar.Reference, data []byte) (state []byte, result []byte, err error) {
	ph := common.CurrentProxyCtx
	ph.SetSystemError(nil)

	args := []interface{}{}

	err = ph.Deserialize(data, &args)
	if err != nil {
		err = &foundation.Error{S: "[ FakeNewWithX ] ( INSCONSTRUCTOR_* ) ( Generated Method ) Can't deserialize args.Arguments: " + err.Error()}
		return
	}

	var ret0 *Second
	var ret1 error

	serializeResults := func() error {
		return ph.Serialize(
			foundation.Result{Returns: []interface{}{ref, ret1}},
			&result,
		)
	}

	needRecover := true
	defer func() {
		if !needRecover {
			return
		}
		if r := recover(); r != nil {
			recoveredError := errors.Wrap(errors.Errorf("%v", r), "Failed to execute constructor (panic)")
			recoveredError = ph.MakeErrorSerializable(recoveredError)

			if PanicIsLogicalError {
				ret1 = recoveredError

				err = serializeResults()
				if err == nil {
					state = data
				}
			} else {
				err = recoveredError
			}
		}
	}()

	ret0, ret1 = NewWithX()

	needRecover = false

	ret1 = ph.MakeErrorSerializable(ret1)
	if ret0 == nil && ret1 == nil {
		ret1 = &foundation.Error{S: "constructor returned nil"}
	}

	if ph.GetSystemError() != nil {
		err = ph.GetSystemError()
		return
	}

	err = serializeResults()
	if err != nil {
		return
	}

	if ret1 != nil {
		// logical error, the result should be registered with type RequestSideEffectNone
		state = nil
		return
	}

	err = ph.Serialize(ret0, &state)
	if err != nil {
		return
	}

	return
}

func INSCONSTRUCTOR_NewSaga(ref insolar.Reference, data []byte) (state []byte, result []byte, err error) {
	ph := common.CurrentProxyCtx
	ph.SetSystemError(nil)

	args := []interface{}{}

	err = ph.Deserialize(data, &args)
	if err != nil {
		err = &foundation.Error{S: "[ FakeNewSaga ] ( INSCONSTRUCTOR_* ) ( Generated Method ) Can't deserialize args.Arguments: " + err.Error()}
		return
	}

	var ret0 *Second
	var ret1 error

	serializeResults := func() error {
		return ph.Serialize(
			foundation.Result{Returns: []interface{}{ref, ret1}},
			&result,
		)
	}

	needRecover := true
	defer func() {
		if !needRecover {
			return
		}
		if r := recover(); r != nil {
			recoveredError := errors.Wrap(errors.Errorf("%v", r), "Failed to execute constructor (panic)")
			recoveredError = ph.MakeErrorSerializable(recoveredError)

			if PanicIsLogicalError {
				ret1 = recoveredError

				err = serializeResults()
				if err == nil {
					state = data
				}
			} else {
				err = recoveredError
			}
		}
	}()

	ret0, ret1 = NewSaga()

	needRecover = false

	ret1 = ph.MakeErrorSerializable(ret1)
	if ret0 == nil && ret1 == nil {
		ret1 = &foundation.Error{S: "constructor returned nil"}
	}

	if ph.GetSystemError() != nil {
		err = ph.GetSystemError()
		return
	}

	err = serializeResults()
	if err != nil {
		return
	}

	if ret1 != nil {
		// logical error, the result should be registered with type RequestSideEffectNone
		state = nil
		return
	}

	err = ph.Serialize(ret0, &state)
	if err != nil {
		return
	}

	return
}

func INSCONSTRUCTOR_NewNil(ref insolar.Reference, data []byte) (state []byte, result []byte, err error) {
	ph := common.CurrentProxyCtx
	ph.SetSystemError(nil)

	args := []interface{}{}

	err = ph.Deserialize(data, &args)
	if err != nil {
		err = &foundation.Error{S: "[ FakeNewNil ] ( INSCONSTRUCTOR_* ) ( Generated Method ) Can't deserialize args.Arguments: " + err.Error()}
		return
	}

	var ret0 *Second
	var ret1 error

	serializeResults := func() error {
		return ph.Serialize(
			foundation.Result{Returns: []interface{}{ref, ret1}},
			&result,
		)
	}

	needRecover := true
	defer func() {
		if !needRecover {
			return
		}
		if r := recover(); r != nil {
			recoveredError := errors.Wrap(errors.Errorf("%v", r), "Failed to execute constructor (panic)")
			recoveredError = ph.MakeErrorSerializable(recoveredError)

			if PanicIsLogicalError {
				ret1 = recoveredError

				err = serializeResults()
				if err == nil {
					state = data
				}
			} else {
				err = recoveredError
			}
		}
	}()

	ret0, ret1 = NewNil()

	needRecover = false

	ret1 = ph.MakeErrorSerializable(ret1)
	if ret0 == nil && ret1 == nil {
		ret1 = &foundation.Error{S: "constructor returned nil"}
	}

	if ph.GetSystemError() != nil {
		err = ph.GetSystemError()
		return
	}

	err = serializeResults()
	if err != nil {
		return
	}

	if ret1 != nil {
		// logical error, the result should be registered with type RequestSideEffectNone
		state = nil
		return
	}

	err = ph.Serialize(ret0, &state)
	if err != nil {
		return
	}

	return
}

func INSCONSTRUCTOR_NewWithErr(ref insolar.Reference, data []byte) (state []byte, result []byte, err error) {
	ph := common.CurrentProxyCtx
	ph.SetSystemError(nil)

	args := []interface{}{}

	err = ph.Deserialize(data, &args)
	if err != nil {
		err = &foundation.Error{S: "[ FakeNewWithErr ] ( INSCONSTRUCTOR_* ) ( Generated Method ) Can't deserialize args.Arguments: " + err.Error()}
		return
	}

	var ret0 *Second
	var ret1 error

	serializeResults := func() error {
		return ph.Serialize(
			foundation.Result{Returns: []interface{}{ref, ret1}},
			&result,
		)
	}

	needRecover := true
	defer func() {
		if !needRecover {
			return
		}
		if r := recover(); r != nil {
			recoveredError := errors.Wrap(errors.Errorf("%v", r), "Failed to execute constructor (panic)")
			recoveredError = ph.MakeErrorSerializable(recoveredError)

			if PanicIsLogicalError {
				ret1 = recoveredError

				err = serializeResults()
				if err == nil {
					state = data
				}
			} else {
				err = recoveredError
			}
		}
	}()

	ret0, ret1 = NewWithErr()

	needRecover = false

	ret1 = ph.MakeErrorSerializable(ret1)
	if ret0 == nil && ret1 == nil {
		ret1 = &foundation.Error{S: "constructor returned nil"}
	}

	if ph.GetSystemError() != nil {
		err = ph.GetSystemError()
		return
	}

	err = serializeResults()
	if err != nil {
		return
	}

	if ret1 != nil {
		// logical error, the result should be registered with type RequestSideEffectNone
		state = nil
		return
	}

	err = ph.Serialize(ret0, &state)
	if err != nil {
		return
	}

	return
}

func Initialize() insolar.ContractWrapper {
	return insolar.ContractWrapper{
		GetCode:      INSMETHOD_GetCode,
		GetPrototype: INSMETHOD_GetPrototype,
		Methods: insolar.ContractMethods{
			"GetName":          INSMETHOD_GetName,
			"DoNothing":        INSMETHOD_DoNothing,
			"Get":              INSMETHOD_Get,
			"Hello":            INSMETHOD_Hello,
			"GetPayload":       INSMETHOD_GetPayload,
			"SetPayload":       INSMETHOD_SetPayload,
			"GetPayloadString": INSMETHOD_GetPayloadString,
			"GetBalance":       INSMETHOD_GetBalance,
			"Accept":           INSMETHOD_Accept,
			"AnError":          INSMETHOD_AnError,
			"NoError":          INSMETHOD_NoError,
			"ReturnNil":        INSMETHOD_ReturnNil,
		},
		Constructors: insolar.ContractConstructors{
			"New":        INSCONSTRUCTOR_New,
			"NewWithOne": INSCONSTRUCTOR_NewWithOne,
			"NewWithX":   INSCONSTRUCTOR_NewWithX,
			"NewSaga":    INSCONSTRUCTOR_NewSaga,
			"NewNil":     INSCONSTRUCTOR_NewNil,
			"NewWithErr": INSCONSTRUCTOR_NewWithErr,
		},
	}
}
