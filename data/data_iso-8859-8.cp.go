// This file is automatically generated by generate-charset-data.
// Do not hand-edit.

package data

import (
	"io"
	"io/ioutil"
	"strings"

	"github.com/ezoic/go-charset/charset"
)

func init() {
	charset.RegisterDataFile("iso-8859-8.cp", func() (io.ReadCloser, error) {
		r := strings.NewReader("\x00\x01\x02\x03\x04\x05\x06\a\b\t\n\v\f\r\x0e\x0f\x10\x11\x12\x13\x14\x15\x16\x17\x18\x19\x1a\x1b\x1c\x1d\x1e\x1f !\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~\x7f\u0080\u0081\u0082\u0083\u0084\u0085\u0086\u0087\u0088\u0089\u008a\u008b\u008c\u008d\u008e\u008f\u0090\u0091\u0092\u0093\u0094\u0095\u0096\u0097\u0098\u0099\u009a\u009b\u009c\u009d\u009e\u009f\u00a0�¢£¤¥¦§¨©×«¬\u00ad®‾°±²³´µ¶·¸¹÷»¼½¾��������������������������������‗אבגדהוזחטיךכלםמןנסעףפץצקרשת�����")
		return ioutil.NopCloser(r), nil
	})
}
