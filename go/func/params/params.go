package params

import (
	"time"
)

const (
	CommonCart = "common"
)

type cartExts struct {
	cartType string
	ttl      time.Duration
}

type DemoCart struct {
	UserID string
	ItemID string
	Sku    int64
	Ext    cartExts
}

var DefaultExt = cartExts{
	cartType: CommonCart,       // 默认是普通购物车类型
	ttl:      time.Minute * 60, // 默认 60min 过期
}

func NewExt() *cartExts {
	return &DefaultExt
}

func (c *cartExts) WithCartType(cartType string) {
	c.cartType = cartType
}

func (c *cartExts) WithTTL(d time.Duration) {
	c.ttl = d
}

func NewCart(userID string, Sku int64, ext *cartExts) *DemoCart {
	c := &DemoCart{
		UserID: userID,
		Sku:    Sku,
		Ext:    *ext, // 设置默认值
	}
	return c
}

type CartOption func(c *cartExts)

func NewCart1(userID string, Sku int64, cartOptions ...CartOption) *DemoCart {

	option := DefaultExt

	for _, fn := range cartOptions {
		fn(&option)
	}

	return &DemoCart{
		UserID: userID,
		Sku:    Sku,
		Ext:    option, // 设置默认值
	}
}

func WithCartType1(cartType string) CartOption {
	return func(c *cartExts) {
		c.cartType = cartType
	}
}

func WithTTL1(ttl time.Duration) CartOption {
	return func(c *cartExts) {
		c.ttl = ttl
	}
}
