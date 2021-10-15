// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: address_token.proto

package models

import (
	context "context"
	fmt "fmt"
	
	_ "github.com/infobloxopen/protoc-gen-gorm/options"
	math "math"

	gorm2 "github.com/infobloxopen/atlas-app-toolkit/gorm"
	errors1 "github.com/infobloxopen/protoc-gen-gorm/errors"
	gorm1 "github.com/jinzhu/gorm"
	field_mask1 "google.golang.org/genproto/protobuf/field_mask"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = fmt.Errorf
var _ = math.Inf

type AddressTokenORM struct {
	PublicKey            string `gorm:"primary_key"`
	TokenContractAddress string `gorm:"primary_key"`
}

// TableName overrides the default tablename generated by GORM
func (AddressTokenORM) TableName() string {
	return "address_tokens"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *AddressToken) ToORM(ctx context.Context) (AddressTokenORM, error) {
	to := AddressTokenORM{}
	var err error
	if prehook, ok := interface{}(m).(AddressTokenWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.PublicKey = m.PublicKey
	to.TokenContractAddress = m.TokenContractAddress
	if posthook, ok := interface{}(m).(AddressTokenWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *AddressTokenORM) ToPB(ctx context.Context) (AddressToken, error) {
	to := AddressToken{}
	var err error
	if prehook, ok := interface{}(m).(AddressTokenWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.PublicKey = m.PublicKey
	to.TokenContractAddress = m.TokenContractAddress
	if posthook, ok := interface{}(m).(AddressTokenWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type AddressToken the arg will be the target, the caller the one being converted from

// AddressTokenBeforeToORM called before default ToORM code
type AddressTokenWithBeforeToORM interface {
	BeforeToORM(context.Context, *AddressTokenORM) error
}

// AddressTokenAfterToORM called after default ToORM code
type AddressTokenWithAfterToORM interface {
	AfterToORM(context.Context, *AddressTokenORM) error
}

// AddressTokenBeforeToPB called before default ToPB code
type AddressTokenWithBeforeToPB interface {
	BeforeToPB(context.Context, *AddressToken) error
}

// AddressTokenAfterToPB called after default ToPB code
type AddressTokenWithAfterToPB interface {
	AfterToPB(context.Context, *AddressToken) error
}

// DefaultCreateAddressToken executes a basic gorm create call
func DefaultCreateAddressToken(ctx context.Context, in *AddressToken, db *gorm1.DB) (*AddressToken, error) {
	if in == nil {
		return nil, errors1.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(AddressTokenORMWithBeforeCreate_); ok {
		if db, err = hook.BeforeCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(AddressTokenORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

type AddressTokenORMWithBeforeCreate_ interface {
	BeforeCreate_(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type AddressTokenORMWithAfterCreate_ interface {
	AfterCreate_(context.Context, *gorm1.DB) error
}

// DefaultApplyFieldMaskAddressToken patches an pbObject with patcher according to a field mask.
func DefaultApplyFieldMaskAddressToken(ctx context.Context, patchee *AddressToken, patcher *AddressToken, updateMask *field_mask1.FieldMask, prefix string, db *gorm1.DB) (*AddressToken, error) {
	if patcher == nil {
		return nil, nil
	} else if patchee == nil {
		return nil, errors1.NilArgumentError
	}
	var err error
	for _, f := range updateMask.Paths {
		if f == prefix+"PublicKey" {
			patchee.PublicKey = patcher.PublicKey
			continue
		}
		if f == prefix+"TokenContractAddress" {
			patchee.TokenContractAddress = patcher.TokenContractAddress
			continue
		}
	}
	if err != nil {
		return nil, err
	}
	return patchee, nil
}

// DefaultListAddressToken executes a gorm list call
func DefaultListAddressToken(ctx context.Context, db *gorm1.DB) ([]*AddressToken, error) {
	in := AddressToken{}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(AddressTokenORMWithBeforeListApplyQuery); ok {
		if db, err = hook.BeforeListApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	db, err = gorm2.ApplyCollectionOperators(ctx, db, &AddressTokenORM{}, &AddressToken{}, nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(AddressTokenORMWithBeforeListFind); ok {
		if db, err = hook.BeforeListFind(ctx, db); err != nil {
			return nil, err
		}
	}
	db = db.Where(&ormObj)
	db = db.Order("public_key")
	ormResponse := []AddressTokenORM{}
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(AddressTokenORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*AddressToken{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

type AddressTokenORMWithBeforeListApplyQuery interface {
	BeforeListApplyQuery(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type AddressTokenORMWithBeforeListFind interface {
	BeforeListFind(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type AddressTokenORMWithAfterListFind interface {
	AfterListFind(context.Context, *gorm1.DB, *[]AddressTokenORM) error
}
