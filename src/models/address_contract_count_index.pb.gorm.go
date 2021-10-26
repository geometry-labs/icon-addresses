// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: address_contract_count_index.proto

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

type AddressContractCountIndexORM struct {
	PublicKey string `gorm:"primary_key"`
}

// TableName overrides the default tablename generated by GORM
func (AddressContractCountIndexORM) TableName() string {
	return "address_contract_count_indices"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *AddressContractCountIndex) ToORM(ctx context.Context) (AddressContractCountIndexORM, error) {
	to := AddressContractCountIndexORM{}
	var err error
	if prehook, ok := interface{}(m).(AddressContractCountIndexWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.PublicKey = m.PublicKey
	if posthook, ok := interface{}(m).(AddressContractCountIndexWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *AddressContractCountIndexORM) ToPB(ctx context.Context) (AddressContractCountIndex, error) {
	to := AddressContractCountIndex{}
	var err error
	if prehook, ok := interface{}(m).(AddressContractCountIndexWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.PublicKey = m.PublicKey
	if posthook, ok := interface{}(m).(AddressContractCountIndexWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type AddressContractCountIndex the arg will be the target, the caller the one being converted from

// AddressContractCountIndexBeforeToORM called before default ToORM code
type AddressContractCountIndexWithBeforeToORM interface {
	BeforeToORM(context.Context, *AddressContractCountIndexORM) error
}

// AddressContractCountIndexAfterToORM called after default ToORM code
type AddressContractCountIndexWithAfterToORM interface {
	AfterToORM(context.Context, *AddressContractCountIndexORM) error
}

// AddressContractCountIndexBeforeToPB called before default ToPB code
type AddressContractCountIndexWithBeforeToPB interface {
	BeforeToPB(context.Context, *AddressContractCountIndex) error
}

// AddressContractCountIndexAfterToPB called after default ToPB code
type AddressContractCountIndexWithAfterToPB interface {
	AfterToPB(context.Context, *AddressContractCountIndex) error
}

// DefaultCreateAddressContractCountIndex executes a basic gorm create call
func DefaultCreateAddressContractCountIndex(ctx context.Context, in *AddressContractCountIndex, db *gorm1.DB) (*AddressContractCountIndex, error) {
	if in == nil {
		return nil, errors1.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(AddressContractCountIndexORMWithBeforeCreate_); ok {
		if db, err = hook.BeforeCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(AddressContractCountIndexORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

type AddressContractCountIndexORMWithBeforeCreate_ interface {
	BeforeCreate_(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type AddressContractCountIndexORMWithAfterCreate_ interface {
	AfterCreate_(context.Context, *gorm1.DB) error
}

// DefaultApplyFieldMaskAddressContractCountIndex patches an pbObject with patcher according to a field mask.
func DefaultApplyFieldMaskAddressContractCountIndex(ctx context.Context, patchee *AddressContractCountIndex, patcher *AddressContractCountIndex, updateMask *field_mask1.FieldMask, prefix string, db *gorm1.DB) (*AddressContractCountIndex, error) {
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
	}
	if err != nil {
		return nil, err
	}
	return patchee, nil
}

// DefaultListAddressContractCountIndex executes a gorm list call
func DefaultListAddressContractCountIndex(ctx context.Context, db *gorm1.DB) ([]*AddressContractCountIndex, error) {
	in := AddressContractCountIndex{}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(AddressContractCountIndexORMWithBeforeListApplyQuery); ok {
		if db, err = hook.BeforeListApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	db, err = gorm2.ApplyCollectionOperators(ctx, db, &AddressContractCountIndexORM{}, &AddressContractCountIndex{}, nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(AddressContractCountIndexORMWithBeforeListFind); ok {
		if db, err = hook.BeforeListFind(ctx, db); err != nil {
			return nil, err
		}
	}
	db = db.Where(&ormObj)
	db = db.Order("public_key")
	ormResponse := []AddressContractCountIndexORM{}
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(AddressContractCountIndexORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*AddressContractCountIndex{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

type AddressContractCountIndexORMWithBeforeListApplyQuery interface {
	BeforeListApplyQuery(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type AddressContractCountIndexORMWithBeforeListFind interface {
	BeforeListFind(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type AddressContractCountIndexORMWithAfterListFind interface {
	AfterListFind(context.Context, *gorm1.DB, *[]AddressContractCountIndexORM) error
}