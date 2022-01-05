package sfexpresscitygo

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

type Signer struct {
	devId  string
	devKey string
}

func NewSigner(devId, devKey string) *Signer {
	return &Signer{
		devId:  devId,
		devKey: devKey,
	}
}

func (s *Signer) Sign(json string) string {
	// $post_data = json_encode($params);
	// $sign_char = $post_data . "&{$dev_id}&{$dev_key}";
	// $sign      = base64_encode(MD5($sign_char)); // 注：md5出来的结果是32位小写16进制字符串，$sign 的最终结果末尾包含等号=
	payload := fmt.Sprintf("%s&%s&%s", json, s.devId, s.devKey)
	h := md5.New()
	h.Write([]byte(payload))
	hash := hex.EncodeToString(h.Sum(nil))
	return base64.StdEncoding.EncodeToString([]byte(hash))
}
