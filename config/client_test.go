package config

import (
	"reflect"
	"testing"
)

func TestClientConfig_ToArgs(t *testing.T) {
	type fields struct {
		BufSize          string
		ClientPort       uint16
		CongestionCtrl   string
		ConnectTimeout   uint64
		DSCP             string
		DisableNagle     bool
		FlowBytes        string
		FlowDuration     uint64
		FlowPackets      string
		FQRate           string
		MSS              uint
		OmitLeadingSecs  uint
		PacingTimer      string
		ParallelFlows    uint
		RSAPubKeyFile    string
		ServerAddr       string
		TargetBitrate    string
		ToS              uint8
		UDP              bool
		UDP64BitCounters bool
		Username         string
		V4Only           bool
		V6Only           bool
		WindowSize       string
		ZeroCopy         bool
		CommonConfig     CommonConfig
	}
	tests := []struct {
		name    string
		fields  fields
		want    []string
		wantErr bool
	}{
		{
			"good stuff",
			fields{
				BufSize:          "1024k",
				ClientPort:       12345,
				CongestionCtrl:   "ledbat",
				ConnectTimeout:   1234567890,
				DSCP:             "AF12",
				DisableNagle:     true,
				FlowBytes:        "1M",
				FQRate:           "10M",
				MSS:              1400,
				OmitLeadingSecs:  12345,
				PacingTimer:      "1000",
				ParallelFlows:    999,
				RSAPubKeyFile:    "/path/to/rsa/pub/key",
				ServerAddr:       "iperf-server",
				TargetBitrate:    "1M",
				ToS:              0x01,
				UDP:              false,
				UDP64BitCounters: false,
				Username:         "a user",
				V4Only:           false,
				V6Only:           false,
				WindowSize:       "1K",
				ZeroCopy:         false,
				CommonConfig:     CommonConfig{},
			},
			[]string{
				"iperf3",
				"--client",
				"iperf-server",
				"--length",
				"1024k",
				"--cport",
				"12345",
				"--congestion",
				"ledbat",
				"--connect-timeout",
				"1234567890",
				"--dscp",
				"AF12",
				"--no-delay",
				"--bytes",
				"1M",
				"--fq-rate",
				"10M",
				"--get-server-output",
				"--set-mss",
				"1400",
				"--omit",
				"12345",
				"--pacing-timer",
				"1000",
				"--parallel",
				"999",
				"--rsa-public-key-path",
				"/path/to/rsa/pub/key",
				"--reverse",
				"--bitrate",
				"1M",
				"--tos",
				"1",
				"--username",
				"a user",
				"--window",
				"1K",
				// assume Commmon's defaults
				"--json",
			},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := &ClientConfig{
				BufSize:          tt.fields.BufSize,
				ClientPort:       tt.fields.ClientPort,
				CongestionCtrl:   tt.fields.CongestionCtrl,
				ConnectTimeout:   tt.fields.ConnectTimeout,
				DSCP:             tt.fields.DSCP,
				DisableNagle:     tt.fields.DisableNagle,
				FlowBytes:        tt.fields.FlowBytes,
				FlowDuration:     tt.fields.FlowDuration,
				FlowPackets:      tt.fields.FlowPackets,
				FQRate:           tt.fields.FQRate,
				MSS:              tt.fields.MSS,
				OmitLeadingSecs:  tt.fields.OmitLeadingSecs,
				PacingTimer:      tt.fields.PacingTimer,
				ParallelFlows:    tt.fields.ParallelFlows,
				RSAPubKeyFile:    tt.fields.RSAPubKeyFile,
				ServerAddr:       tt.fields.ServerAddr,
				TargetBitrate:    tt.fields.TargetBitrate,
				ToS:              tt.fields.ToS,
				UDP:              tt.fields.UDP,
				UDP64BitCounters: tt.fields.UDP64BitCounters,
				Username:         tt.fields.Username,
				V4Only:           tt.fields.V4Only,
				V6Only:           tt.fields.V6Only,
				WindowSize:       tt.fields.WindowSize,
				ZeroCopy:         tt.fields.ZeroCopy,
				CommonConfig:     tt.fields.CommonConfig,
			}
			got, err := cfg.ToArgs()
			if (err != nil) != tt.wantErr {
				t.Errorf("ClientConfig.ToArgs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClientConfig.ToArgs() = %v, want %v", got, tt.want)
			}
		})
	}
}
