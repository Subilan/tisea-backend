package utils

import "fmt"
import "time"
import "github.com/golang-jwt/jwt"

func GenerateTokenFor(data map[string]interface{}) (string, error) {
	var cfg = GetConfiguration()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.MapClaims{
		"iss": "tisea",
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Duration(cfg.JwtExpiration) * time.Minute).Unix(),
		"data": data,
	})

	signed, err := token.SignedString([]byte(cfg.JwtPrivateKey))

	if err != nil {
		return "", err
	}

	return signed, nil
}

// 尝试解析给定的 token 字符串。若 token 可以被正确解析，error 为 nil，并会返回解析后的 token 对象。
func Parse(tokenString string) (*jwt.Token, error) {
	var cfg = GetConfiguration()
	return jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if alg := t.Method.Alg(); alg != "HS256" {
			return nil, fmt.Errorf("Invalid signing method %v", alg)
		}

		return []byte(cfg.JwtPrivateKey), nil
	})
}

// 检查给定的 token 是否有效。此方法检查：1. 是否可以被正常解析；2. 是否有效（token.Valid）
func CheckToken(tokenString string) error {
	token, err := Parse(tokenString)

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("Token is invalid")
	}

	return nil
}

// 尝试从给定的 token 中解出所包含的 object 数据
func Extract(tokenString string) (map[string]interface{}, error) {
	check := CheckToken(tokenString)
	if check != nil {
		return nil, check
	}
	
	parse, parseErr := Parse(tokenString)

	if parseErr != nil {
		return nil, parseErr
	}

	claims, ok := parse.Claims.(jwt.MapClaims)

	if !ok {
		return nil, fmt.Errorf("Invalid claim type")
	}

	return claims["data"].(map[string]interface{}), nil
}
