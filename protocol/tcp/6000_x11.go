package tcp

import (
	"encoding/json"
	"fmt"
	"net"
	"time"
)

type X11Reply struct {
	Success                        uint8  `json:"success,omitempty"`
	ProtocolMajorVersion           uint16 `json:"protocol-major-version,omitempty"`
	ProtocolMinorVersion           uint16 `json:"protocol-minor-version,omitempty"`
	ReplyLength                    uint16 `json:"reply-length,omitempty"`
	ReleaseNumber                  uint32 `json:"release-number,omitempty"`
	ResourceIDBase                 uint32 `json:"resource-id-base,omitempty"`
	ResourceIDMask                 uint32 `json:"resource-id-mask,omitempty"`
	MotionBufferSize               uint32 `json:"motion-buffer-size,omitempty"`
	LengthOfVendor                 uint16 `json:"length-of-vendor,omitempty"`
	MaximumRequestLength           uint16 `json:"maximum-request-length,omitemptyh"`
	NumberOfScreensInRoots         uint8  `json:"number-of-screens-in-roots,omitempty"`
	NumberOfFormatsInPixmapFormats uint8  `json:"number-of-formats-in-pixmap-formats,omitempty"`
	ImageByteOrder                 string `json:"image-byte-order,omitempty"` // 0-LSBFirst or MSBFirst
	BitmapFormatBitOrder           string `json:"bitmap-format-bit-order,omitempty"`
	BitmapFormatScanlineUnit       uint8  `json:"bitmap-format-scanline-unit,omitempty"`
	BitmapFormatScanlinePad        uint8  `json:"bitmap-format-scanline-pad,omitempty"`
	MinKeycode                     uint8  `json:"min-keycode,omitempty"`
	MaxKeycode                     uint8  `json:"max-keycode,omitempty"`
	Vendor                         string `json:"verdor,omitempty"`
}

func GrabX11(addr string) {
	dialer := net.Dialer{
		Timeout:  time.Second * 10,
		Deadline: time.Now().Add(time.Second * 10),
	}
	conn, err := dialer.Dial("tcp", addr+":6000")
	if err != nil {
		fmt.Println(err)
		return
	}
	conn.Write([]byte("\x6c\x00\x0b\x00\x00\x00\x00\x00\x00\x00\x00\x00"))
	var buf [1024]byte
	_, err = conn.Read(buf[:])
	if err != nil {
		fmt.Println(err)
		return
	}
	if buf[0] > 0x01 {
		fmt.Println(false)
		fmt.Println(buf)
		return
	}
	var (
		reply X11Reply
	)
	switch buf[0] {
	case 0x01:
		reply.Success = buf[0]
		reply.ProtocolMajorVersion = uint16(buf[2]) | uint16(buf[3])<<8
		reply.ProtocolMinorVersion = uint16(buf[4]) | uint16(buf[5])<<8
		reply.ReplyLength = uint16(buf[6]) | uint16(buf[7])<<8
		reply.ReleaseNumber = uint32(buf[8]) | uint32(buf[9])<<8 | uint32(buf[10])<<16 | uint32(buf[11])<<24
		reply.ResourceIDBase = uint32(buf[12]) | uint32(buf[13])<<8 | uint32(buf[14])<<16 | uint32(buf[15])<<24
		reply.ResourceIDMask = uint32(buf[16]) | uint32(buf[17])<<8 | uint32(buf[18])<<16 | uint32(buf[19])<<24
		reply.MotionBufferSize = uint32(buf[20]) | uint32(buf[21])<<8 | uint32(buf[22])<<16 | uint32(buf[23])<<24
		reply.LengthOfVendor = uint16(buf[24]) | uint16(buf[25])<<8
		reply.MaximumRequestLength = uint16(buf[26]) | uint16(buf[27])<<8
		reply.NumberOfScreensInRoots = uint8(buf[28])
		reply.NumberOfFormatsInPixmapFormats = uint8(buf[29])
		if uint8(buf[30]) == 0 {
			reply.ImageByteOrder = "LSBFirst"
		} else {
			reply.ImageByteOrder = "MSBFirst"
		}
		if uint8(buf[31]) == 0 {
			reply.BitmapFormatBitOrder = "LSBFirst"
		} else {
			reply.BitmapFormatBitOrder = "MSBFirst"
		}
		reply.BitmapFormatScanlineUnit = uint8(buf[32])
		reply.BitmapFormatScanlinePad = uint8(buf[33])
		reply.MinKeycode = uint8(buf[34])
		reply.MaxKeycode = uint8(buf[35])
		reply.Vendor = string(buf[40:64])
	case 0x00:
		reply.Success = buf[0]
		reasonLen := buf[1]
		reply.ProtocolMajorVersion = uint16(buf[2]) | uint16(buf[3])<<8
		reply.ProtocolMinorVersion = uint16(buf[4]) | uint16(buf[5])<<8
		reply.ReplyLength = uint16(buf[6]) | uint16(buf[7])<<8
		reply.Vendor = string(buf[8 : 8+reasonLen])
	default:
	}

	data, err := json.Marshal(reply)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s\n", data)
}
