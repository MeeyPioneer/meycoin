/*
 * @file
 * @copyright defined in meycoin/LICENSE.txt
 */

package network

import (
	"net"
	"reflect"
	"testing"
)

func TestResolveHostDomain(t *testing.T) {
	type args struct {
		domainName string
	}
	tests := []struct {
		name    string
		args    args
		exist   bool
		wantErr bool
	}{
		{"TSucc", args{"www.google.com"}, true, false},
		{"TNowhere", args{"not.in.my.meycoin.io"}, false, true},
		{"TWrongName", args{"!#@doigjw"}, false, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ResolveHostDomain(tt.args.domainName)
			if (err != nil) != tt.wantErr {
				t.Errorf("ResolveHostDomain() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (len(got) > 0) != tt.exist {
				t.Errorf("ResolveHostDomain() = %v, want %v", got, tt.exist)
			}
		})
	}
}

func TestResolveHostDomainLocal(t *testing.T) {
	t.Skip("skip env dependent test")
	type args struct {
		domainName string
	}
	tests := []struct {
		name    string
		args    args
		want    []net.IP
		wantErr bool
	}{
		{"TPrivate", args{"devuntu31"}, []net.IP{net.ParseIP("192.168.0.215")}, false},
		{"TPrivate", args{"devuntu31.blocko.io"}, []net.IP{net.ParseIP("192.168.0.215")}, false},
		{"TPrivate", args{"devuntu31ss"}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ResolveHostDomain(tt.args.domainName)
			if (err != nil) != tt.wantErr {
				t.Errorf("ResolveHostDomain() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ResolveHostDomain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseAddress(t *testing.T) {
	tests := []struct {
		name string

		in string

		wantErr  bool
		wantHost string
		wantPort string
	}{
		{"TIP4", "211.34.56.78", false, "211.34.56.78", ""},
		{"TIP6", "fe80::dcbf:beff:fe87:e30a", false, "fe80::dcbf:beff:fe87:e30a", ""},
		{"TIP6_2", "::ffff:192.0.1.2", false, "::ffff:192.0.1.2", ""},
		{"TFQDN", "iparkmac.meycoin.io", false, "iparkmac.meycoin.io", ""},
		{"TIP4WithPort", "211.34.56.78:1234", true, "211.34.56.78", "1234"},
		{"TIP6WithPort", "[fe80::dcbf:beff:fe87:e30a]:1234", true, "fe80::dcbf:beff:fe87:e30a", "1234"},
		{"TFQDNWithPort", "iparkmac.meycoin.io:1234", true, "iparkmac.meycoin.io", "1234"},
		// TODO: test cases
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := CheckAddress(test.in)
			if (err != nil) != test.wantErr {
				t.Errorf("CheckAddress() error = %v, wantErr %v", err, test.wantErr)
				return
			}
			if !test.wantErr {
				if got != test.wantHost {
					t.Errorf("CheckAddress() = host %v, wantHost %v", got, test.wantHost)
				}
			}
		})
	}
}

func TestCheckAddressType(t *testing.T) {
	tests := []struct {
		name string

		in string

		want AddressType
	}{
		{"TIP4", "211.34.56.78", AddressTypeIP},
		{"TIP6", "fe80::dcbf:beff:fe87:e30a", AddressTypeIP},
		{"TIP6_2", "::ffff:192.0.1.2", AddressTypeIP},
		{"TFQDN", "iparkmac.meycoin.io", AddressTypeFQDN},
		{"TFQDN_2", "3com.com", AddressTypeFQDN},
		{"TWrongDN", "3com!.com", AddressTypeError},
		{"TIP4withPort", "211.34.56.78:1234", AddressTypeError},
		{"TIP6withPort", "[fe80::dcbf:beff:fe87:e30a]:1234", AddressTypeError},
		{"TFQDNwithPort", "iparkmac.meycoin.io:1234", AddressTypeError},
		// TODO: test cases
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := CheckAddressType(test.in)
			if got != test.want {
				t.Errorf("CheckAddressType() = type %v, wantType %v", got, test.want)
			}
		})
	}
}
