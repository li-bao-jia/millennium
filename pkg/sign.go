package pkg

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

func Sign(p map[string]interface{}, appSecret string) string {
	// 将字典序列化为 JSON 字符串
	jsonBytes, _ := json.Marshal(p)

	jsonStr := string(jsonBytes)

	// 将 JSON 字符串转换为字符数组
	jsonArr := strings.Split(jsonStr, "")

	// 对字符数组进行排序
	sort.Strings(jsonArr)

	// 将排序后的字符数组转换为字符串并添加 appSecret
	sortedStr := strings.Join(jsonArr, "") + appSecret

	// 进行 MD5 加密并返回结果
	sum := md5.Sum([]byte(sortedStr))

	return fmt.Sprintf("%x", sum)
}
