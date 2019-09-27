// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import mock "github.com/stretchr/testify/mock"

import v1beta1 "github.com/kubernetes-sigs/service-catalog/pkg/apis/servicecatalog/v1beta1"

// serviceClassGetter is an autogenerated mock type for the serviceClassGetter type
type serviceClassGetter struct {
	mock.Mock
}

// Find provides a mock function with given fields: name, namespace
func (_m *serviceClassGetter) Find(name string, namespace string) (*v1beta1.ServiceClass, error) {
	ret := _m.Called(name, namespace)

	var r0 *v1beta1.ServiceClass
	if rf, ok := ret.Get(0).(func(string, string) *v1beta1.ServiceClass); ok {
		r0 = rf(name, namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1beta1.ServiceClass)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(name, namespace)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByExternalName provides a mock function with given fields: externalName, namespace
func (_m *serviceClassGetter) FindByExternalName(externalName string, namespace string) (*v1beta1.ServiceClass, error) {
	ret := _m.Called(externalName, namespace)

	var r0 *v1beta1.ServiceClass
	if rf, ok := ret.Get(0).(func(string, string) *v1beta1.ServiceClass); ok {
		r0 = rf(externalName, namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1beta1.ServiceClass)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(externalName, namespace)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
