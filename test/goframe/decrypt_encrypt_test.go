package goframe

import (
	"fmt"
	"github.com/gogf/gf/crypto/gaes"
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"testing"
)

/*
@Author kim
@Description   加解密
@date 2021-2-7 16:58
*/

var IV = "my iv must be 16"                  //iv必须要16/24/32位数，非必须参数
var KEY = "5f4dcc3b5aa765d61d8327deb882cf99" //key必须是16/24/32位数
var TEST_PASSWORD = "123456"

//md5加密
func TestMD5Encrypt(t *testing.T) {
	encrypt, err := gmd5.Encrypt(TEST_PASSWORD)
	if err != nil {
		g.Dump(err)
	}
	g.Dump(encrypt)
	g.Dump("位数:", len(encrypt))
}

func TestEncrypt(t *testing.T) {
	bytes := gconv.Bytes(TEST_PASSWORD)
	encrypt, err := gaes.Encrypt(bytes, gconv.Bytes(KEY))
	if err != nil {
		g.Dump(err)
	}
	s := string(encrypt)
	g.Dump(s)
}

func TestDecrypt(t *testing.T) {
	bytes := gconv.Bytes(TEST_PASSWORD)
	encrypt, err := gaes.Encrypt(bytes, gconv.Bytes(KEY), gconv.Bytes(IV))
	if err != nil {
		g.Dump(err)
	}
	s := string(encrypt)
	fmt.Println(s == s)

	decrypt, err := gaes.Decrypt(gconv.Bytes(s), gconv.Bytes(KEY), gconv.Bytes(IV))
	if err != nil {
		g.Dump(err)
	}
	g.Dump(gconv.String(decrypt))
}
