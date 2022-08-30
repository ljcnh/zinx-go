package znet

import (
	"bytes"
	"encoding/binary"
	"errors"
	zinx_go "github.com/ljcnh/zinx-go"
	"github.com/ljcnh/zinx-go/ziface"
)

type DataPack struct {
	HeadLen uint32
}

func NewDataPack() *DataPack {
	return &DataPack{}
}

func (dp *DataPack) GetHeadLen() uint32 {
	// Len uint32 + Id uint32
	return 8
}

func (dp *DataPack) Pack(message ziface.IMessage) ([]byte, error) {
	data := bytes.NewBuffer([]byte{})
	if err := binary.Write(data, binary.LittleEndian, message.GetMsgLen()); err != nil {
		return nil, err
	}
	if err := binary.Write(data, binary.LittleEndian, message.GetMsgId()); err != nil {
		return nil, err
	}
	if err := binary.Write(data, binary.LittleEndian, message.GetData()); err != nil {
		return nil, err
	}
	return data.Bytes(), nil
}

func (dp *DataPack) UnPack(data []byte) (ziface.IMessage, error) {
	buf := bytes.NewReader(data)

	msg := &Message{}

	if err := binary.Read(buf, binary.LittleEndian, &msg.DataLen); err != nil {
		return nil, err
	}
	if err := binary.Read(buf, binary.LittleEndian, &msg.Id); err != nil {
		return nil, err
	}

	if zinx_go.GlobalObject.MaxPackageSize > 0 && msg.DataLen > zinx_go.GlobalObject.MaxPackageSize {
		return nil, errors.New("too Large msg data recv")
	}
	return msg, nil
}
