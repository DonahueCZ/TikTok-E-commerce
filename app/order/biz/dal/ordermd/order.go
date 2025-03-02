package ordermd

import (
	"errors"
	"fmt"
	"github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/kitex_gen/order_service"
	"google.golang.org/protobuf/encoding/protojson"
)

type Order struct {
	Id      int64 `gorm:"column:id;primaryKey;autoIncrement" redis:"id"`
	UserId  int64 `gorm:"column:user_id;not null" redis:"user_id"`
	GoodsId int64 `gorm:"column:goods_id;not null" redis:"goods_id"`

	// Status 0: 未支付 1: 已支付 -1: 已取消
	Status        int32  `gorm:"column:status;not null;default:0" redis:"status"`
	CreateTime    int64  `gorm:"column:create_time;not null" redis:"create_time"`
	GoodsCount    int32  `gorm:"column:goods_count;not null" redis:"goods_count"`
	Cost          int32  `gorm:"column:cost;not null" redis:"cost"`
	AddresseeInfo string `gorm:"column:addressee_info;not null" redis:"addressee_info"`
}

type Tabler interface {
	TableName() string
}

var (
	ExpireTime = int64(60 * 60 * 24)
	IsPaid     = int32(1)
	IsNotPaid  = int32(0)
	IsExpire   = int32(-1)
)

func (o *Order) TableName() string {
	return "orders"
}

func (o *Order) GetCacheKey() string {
	return fmt.Sprintf("%s:%d", o.TableName(), o.Id)
}

func AddresseeInfo2Str(addresseeInfo *order_service.AddresseeInfo) (string, error) {
	if addresseeInfo == nil {
		return "", errors.New("addresseeInfo is nil")
	}
	jsonStr, err := protojson.Marshal(addresseeInfo)
	if err != nil {
		return "", err
	}
	return string(jsonStr), nil
}

func Str2AddresseeInfo(str string) (*order_service.AddresseeInfo, error) {
	if str == "" {
		return nil, errors.New("str is empty")
	}
	addresseeInfo := &order_service.AddresseeInfo{}
	err := protojson.Unmarshal([]byte(str), addresseeInfo)
	if err != nil {
		return nil, err
	}
	return addresseeInfo, nil
}
