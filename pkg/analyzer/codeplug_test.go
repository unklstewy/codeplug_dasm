package analyzer

import (
	"testing"

	cptesting "github.com/unklstewy/codeplug_dasm/pkg/analyzer/testing"
)

func TestIdentifyCodeplugType(t *testing.T) {
	tests := []struct {
		name     string
		header   []byte
		expected CodeplugType
	}{
		{
			name:     "Baofeng DMR codeplug",
			header:   []byte{0x42, 0x46, 0x44, 0x43, 0x50}, // "BFDCP"
			expected: CodeplugTypeBaofengDMR,
		},
		{
			name:     "TYT MD-380 codeplug",
			header:   []byte{0x44, 0x4D, 0x52, 0x33, 0x38}, // "DMR38"
			expected: CodeplugTypeTYTMD380,
		},
		{
			name:     "Unknown codeplug",
			header:   []byte{0x00, 0x01, 0x02, 0x03, 0x04},
			expected: CodeplugTypeUnknown,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IdentifyCodeplugType(tt.header)

			if result != tt.expected {
				t.Errorf("IdentifyCodeplugType() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestExtractChannelInfo(t *testing.T) {
	// This test is a placeholder - we need actual codeplug samples
	t.Skip("Requires real codeplug samples")

	analyzer := NewCodeplugAnalyzer()

	// Sample codeplug data
	sampleData := cptesting.LoadTestCodeplug(t, "samples/codeplug_sample.bin")

	channels, err := analyzer.ExtractChannels(sampleData)
	if err != nil {
		t.Fatalf("ExtractChannels failed: %v", err)
	}

	if len(channels) == 0 {
		t.Error("Expected to extract channel information, got none")
	}
}
