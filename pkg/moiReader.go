package pkg

import (
	"encoding/binary"
	"os"
)

type MoiFile struct {
	Version       uint16
	Filesize      uint32
	Year          uint16
	Month         uint8
	Day           uint8
	Hour          uint8
	Minutes       uint8
	Seconds       uint16
	VideoDuration uint16
}

func ReadMoiFile(filepath string) (*MoiFile, error) {
	moiFile := MoiFile{}
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	binary.Read(f, binary.BigEndian, &moiFile)
	return &moiFile, nil
}
