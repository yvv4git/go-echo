package helpers

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterIPv4(t *testing.T) {
	type args struct {
		addresses []net.Addr
	}

	testCases := []struct {
		name       string
		args       args
		wantErr    bool
		wantResult string
	}{
		{
			name: "CASE-1",
			args: args{
				addresses: []net.Addr{
					&net.IPAddr{
						IP:   []byte{127, 0, 0, 1},
						Zone: "",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "CASE-2",
			args: args{
				addresses: []net.Addr{
					&net.IPAddr{
						IP:   []byte{127, 0, 0, 1},
						Zone: "",
					},
					&net.IPAddr{
						IP:   []byte{192, 168, 0, 106},
						Zone: "",
					},
				},
			},
			wantErr:    false,
			wantResult: "192.168.0.106",
		},
		{
			name: "CASE-3",
			args: args{
				addresses: []net.Addr{
					&net.IPAddr{
						IP:   []byte{127, 0, 0, 1},
						Zone: "",
					},
					&net.IPAddr{
						IP:   []byte{192, 168, 0, 106},
						Zone: "",
					},
					&net.IPAddr{
						IP:   []byte{192, 168, 2, 27},
						Zone: "",
					},
				},
			},
			wantErr:    false,
			wantResult: "192.168.0.106",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := FilterIPv4(tc.args.addresses)
			if (err != nil) != tc.wantErr {
				t.Errorf("have error: %v \n", err)
			}

			assert.Equal(t, tc.wantResult, result)
		})
	}
}
