package bytesToInteger

import (
	"testing"
)

func TestWriter(t *testing.T) {
	var tint8 int8 = -100
	//int16ToBytes
	b := Int8ToBytes(tint8)
	tint82 := BytesToInt8(b)
	if tint8 != tint82 {
		t.Fatalf("unexpected value obtained; got %b want %b", tint8, tint82)
	}
	var tuint8 uint8 = 100
	//int16ToBytes
	b = Uint8ToBytes(tuint8)
	tuint82 := BytesToUint8(b)
	if tuint82 != tuint82 {
		t.Fatalf("unexpected value obtained; got %b want %b", tuint8, tuint82)
	}

	var tint16 int16 = -1000
	//int16ToBytes
	b = Int16ToBytes(tint16)
	tInt162 := BytesToInt16(b)
	if tint16 != tInt162 {
		t.Fatalf("unexpected value obtained; got %b want %b", tint16, tInt162)
	}
	//int16ToBytesBigEndian
	b = Int16ToBytesBigEndian(tint16)
	tInt162 = BytesToInt16BigEndian(b)
	if tint16 != tInt162 {
		t.Fatalf("unexpected value obtained; got %b want %b", tint16, tInt162)
	}

	var tuint16 uint16 = 1000
	//uint16ToBytes
	b = Uint16ToBytes(tuint16)
	tuint162 := BytesToUint16(b)
	if tuint16 != tuint162 {
		t.Fatalf("unexpected value obtained; got %b want %b", tuint16, tuint162)
	}
	b = Uint16ToBytesBigEndian(tuint16)
	tuint162 = BytesToUint16BigEndian(b)
	if tuint16 != tuint162 {
		t.Fatalf("unexpected value obtained; got %b want %b", tuint16, tuint162)
	}

	var tInt32 int32 = -1000
	//int32ToBytes
	b = Int32ToBytes(tInt32)
	tInt322 := BytesToInt32(b)
	if tInt32 != tInt322 {
		t.Fatalf("unexpected value obtained; got %b want %b", tInt32, tInt322)
	}
	//int32ToBytesBigEndian
	b = Int32ToBytesBigEndian(tInt32)
	tInt322 = BytesToInt32BigEndian(b)
	if tInt32 != tInt322 {
		t.Fatalf("unexpected value obtained; got %b want %b", tInt32, tInt322)
	}

	//uint32ToBytes
	var tuint32 uint32 = 1000
	b = Uint32ToBytes(tuint32)
	tuint322 := BytesToUint32(b)
	if tuint32 != tuint322 {
		t.Fatalf("unexpected value obtained; got %b want %b", tuint32, tuint322)
	}
	//uint32ToBytesBigEndian
	b = Uint32ToBytesBigEndian(tuint32)
	tuint322 = BytesToUint32BigEndian(b)
	if tuint32 != tuint322 {
		t.Fatalf("unexpected value obtained; got %b want %b", tuint32, tuint322)
	}

	var tInt64 int64 = -1000
	//int64ToBytes
	b = Int64ToBytes(tInt64)
	tInt642 := BytesToInt64(b)
	if tInt64 != tInt642 {
		t.Fatalf("unexpected value obtained; got %b want %b", tInt64, tInt642)
	}
	//int64ToBytesBigEndian
	b = Int64ToBytesBigEndian(tInt64)
	tInt642 = BytesToInt64BigEndian(b)
	if tInt64 != tInt642 {
		t.Fatalf("unexpected value obtained; got %b want %b", tInt64, tInt642)
	}

	var tuint64 uint64 = 1000
	//uint64ToBytes
	b = Uint64ToBytes(tuint64)
	tuint62 := BytesToUint64(b)
	if tuint64 != tuint62 {
		t.Fatalf("unexpected value obtained; got %b want %b", tuint64, tuint62)
	}
	//uint64ToBytesBigEndian
	b = Uint64ToBytesBigEndian(tuint64)
	tuint62 = BytesToUint64BigEndian(b)
	if tuint64 != tuint62 {
		t.Fatalf("unexpected value obtained; got %b want %b", tuint64, tuint62)
	}
}
