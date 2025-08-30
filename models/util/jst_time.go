package util

import (
	"fmt"
	"time"
)

/** 時刻関係共通メソッド*/
type JSTTime struct {
	time.Time
}

// JSONマーシャル時の出力形式をカスタマイズ
func (jt JSTTime) MarshalJSON() ([]byte, error) {
	// 日本時間に変換
	jst := jt.In(time.FixedZone("Asia/Tokyo", 9*60*60))

	// ISO8601形式 +09:00 付きでフォーマット
	formatted := fmt.Sprintf(`"%s"`, jst.Format("2006-01-02T15:04:05+09:00"))
	return []byte(formatted), nil
}
