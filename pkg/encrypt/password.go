// pkg/encrypt/password.go
package encrypt

import (
    "golang.org/x/crypto/bcrypt"
)

// Password 密码加密接口
type Password interface {
    Hash(password string) (string, error)
    Compare(hashedPassword, password string) error
}

// BCryptPassword bcrypt实现
type BCryptPassword struct {
    cost int
}

// NewBCryptPassword 创建实例
func NewBCryptPassword(cost int) Password {
    if cost == 0 {
        cost = bcrypt.DefaultCost
    }
    return &BCryptPassword{cost: cost}
}

// Hash 加密密码
func (b *BCryptPassword) Hash(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), b.cost)
    return string(bytes), err
}

// Compare 比较密码
func (b *BCryptPassword) Compare(hashedPassword, password string) error {
    return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}