package byteutil

import "testing"

func TestUtf8ToUnicode(t *testing.T) {
	t.Log(Utf8ToUnicode("你好"))
}
