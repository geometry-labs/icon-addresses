// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: block.proto

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

type BlockORM struct {
	LogCount         uint32
	Number           uint32 `gorm:"primary_key"`
	TransactionCount uint32
}

// TableName overrides the default tablename generated by GORM
func (BlockORM) TableName() string {
	return "blocks"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *Block) ToORM(ctx context.Context) (BlockORM, error) {
	to := BlockORM{}
	var err error
	if prehook, ok := interface{}(m).(BlockWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Number = m.Number
	to.TransactionCount = m.TransactionCount
	to.LogCount = m.LogCount
	if posthook, ok := interface{}(m).(BlockWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *BlockORM) ToPB(ctx context.Context) (Block, error) {
	to := Block{}
	var err error
	if prehook, ok := interface{}(m).(BlockWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Number = m.Number
	to.TransactionCount = m.TransactionCount
	to.LogCount = m.LogCount
	if posthook, ok := interface{}(m).(BlockWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type Block the arg will be the target, the caller the one being converted from

// BlockBeforeToORM called before default ToORM code
type BlockWithBeforeToORM interface {
	BeforeToORM(context.Context, *BlockORM) error
}

// BlockAfterToORM called after default ToORM code
type BlockWithAfterToORM interface {
	AfterToORM(context.Context, *BlockORM) error
}

// BlockBeforeToPB called before default ToPB code
type BlockWithBeforeToPB interface {
	BeforeToPB(context.Context, *Block) error
}

// BlockAfterToPB called after default ToPB code
type BlockWithAfterToPB interface {
	AfterToPB(context.Context, *Block) error
}

// DefaultCreateBlock executes a basic gorm create call
func DefaultCreateBlock(ctx context.Context, in *Block, db *gorm1.DB) (*Block, error) {
	if in == nil {
		return nil, errors1.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(BlockORMWithBeforeCreate_); ok {
		if db, err = hook.BeforeCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(BlockORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

type BlockORMWithBeforeCreate_ interface {
	BeforeCreate_(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type BlockORMWithAfterCreate_ interface {
	AfterCreate_(context.Context, *gorm1.DB) error
}

// DefaultApplyFieldMaskBlock patches an pbObject with patcher according to a field mask.
func DefaultApplyFieldMaskBlock(ctx context.Context, patchee *Block, patcher *Block, updateMask *field_mask1.FieldMask, prefix string, db *gorm1.DB) (*Block, error) {
	if patcher == nil {
		return nil, nil
	} else if patchee == nil {
		return nil, errors1.NilArgumentError
	}
	var err error
	for _, f := range updateMask.Paths {
		if f == prefix+"Number" {
			patchee.Number = patcher.Number
			continue
		}
		if f == prefix+"TransactionCount" {
			patchee.TransactionCount = patcher.TransactionCount
			continue
		}
		if f == prefix+"LogCount" {
			patchee.LogCount = patcher.LogCount
			continue
		}
	}
	if err != nil {
		return nil, err
	}
	return patchee, nil
}

// DefaultListBlock executes a gorm list call
func DefaultListBlock(ctx context.Context, db *gorm1.DB) ([]*Block, error) {
	in := Block{}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(BlockORMWithBeforeListApplyQuery); ok {
		if db, err = hook.BeforeListApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	db, err = gorm2.ApplyCollectionOperators(ctx, db, &BlockORM{}, &Block{}, nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(BlockORMWithBeforeListFind); ok {
		if db, err = hook.BeforeListFind(ctx, db); err != nil {
			return nil, err
		}
	}
	db = db.Where(&ormObj)
	db = db.Order("number")
	ormResponse := []BlockORM{}
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(BlockORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*Block{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

type BlockORMWithBeforeListApplyQuery interface {
	BeforeListApplyQuery(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type BlockORMWithBeforeListFind interface {
	BeforeListFind(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type BlockORMWithAfterListFind interface {
	AfterListFind(context.Context, *gorm1.DB, *[]BlockORM) error
}
