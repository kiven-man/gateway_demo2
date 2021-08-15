package unpack

import (
	"encoding/binary"
	"io"
)

const Msg_header = "12345678"

//编码
func Encode(byteBuffer io.Writer, content string) error {
	//byteBuffer 缓冲区
	//binary.BigEndian 大端序
	//content

	//格式：Msg_header(8) + conten_len(4) + content
	//写入header
	if err := binary.Write(byteBuffer, binary.BigEndian, []byte(Msg_header)); err != nil {
		return err
	}
	//写入content_len
	clen := int32(len([]byte(content)))
	if err := binary.Write(byteBuffer, binary.BigEndian, clen); err != nil {
		return err
	}
	//写入content
	if err := binary.Write(byteBuffer, binary.BigEndian, []byte(content)); err != nil {
		return err
	}
	return nil
}

//解码
func Decode(byteBuffer io.Reader) (bodyBuf []byte, err error) {
	//1.读取header
	msgBuf := make([]byte, len(Msg_header))
	if _, err := io.ReadFull(byteBuffer, msgBuf); err != nil {
		return nil, err
	}
	//2.读取content_len
	lenBuf := make([]byte, 4)
	if _, err := io.ReadFull(byteBuffer, lenBuf); err != nil {
		return nil, err
	}
	length := binary.BigEndian.Uint32(lenBuf)
	//3.读取content
	bodyBuf = make([]byte, length)
	if _, err := io.ReadFull(byteBuffer, bodyBuf); err != nil {
		return nil, err
	}
	return bodyBuf, nil
}
