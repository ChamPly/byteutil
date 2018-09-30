byteutil
===
byteutil is golang byte utility

## Quick Start
##### Install
``` bash
$ go get github.com/champly/byteutil
```

#### Slice Byte to String
``` go
str := SliceByteToString([]byte{65, 1, 3, 2, 123, 59})
fmt.Println(str)
```

#### Num to Byte
``` go
Int8ToSliceByte(int8(1))
Int16ToSliceByte(int16(1))
Int32ToSliceByte(int32(1))
Int64ToSliceByte(int64(1))
Uint16ToSliceByte(uint16(1))
Uint32ToSliceByte(uint32(1))
Uint64ToSliceByte(uint64(1))
```

#### String to Slice Byte
``` go
input := "饕餮大餐！@#¥%……&*（）™ "
result := StringToSliceByte(input)
fmt.Println(string(result))
```