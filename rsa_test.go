//Package crypto implements IPIN encryption as per EBS
package ipin

import (
	"testing"

	"github.com/google/uuid"
)

func TestEncrypt(t *testing.T) {

	id, _ := uuid.NewRandom()
	type args struct {
		pubkey string
		pin    string
		uuid   string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"testing enc", args{uuid: id.String(), pubkey: "MFwwDQYJKoZIhvcNAQEBBQADSwAwSAJBANx4gKYSMv3CrWWsxdPfxDxFvl+Is/0kc1dvMI1yNWDXI3AgdI4127KMUOv7gmwZ6SnRsHX/KAM0IPRe0+Sa0vMCAwEAAQ==", pin: "0009"}, "3232", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := Encrypt(tt.args.pubkey, tt.args.pin, tt.args.uuid)

			if got != tt.want {
				t.Errorf("Encrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}
