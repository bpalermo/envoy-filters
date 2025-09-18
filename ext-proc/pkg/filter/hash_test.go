package filter

import (
	"testing"
)

func TestNewFilter(t *testing.T) {
	f := NewFilter()
	if f == nil {
		t.Fatal("NewFilter() returned nil")
	}
	if f.re == nil {
		t.Fatal("NewFilter() did not initialize regex")
	}
}

func TestFilter_Hash(t *testing.T) {
	tests := []struct {
		name        string
		partitionID string
		want        *string
		wantNil     bool
	}{
		{
			name:        "even number",
			partitionID: "s0-456",
			want:        stringPtr("0"),
		},
		{
			name:        "odd number",
			partitionID: "s1-123",
			want:        stringPtr("1"),
		},
		{
			name:        "large even number",
			partitionID: "s5-1000000",
			want:        stringPtr("0"),
		},
		{
			name:        "large odd number",
			partitionID: "s20-999999",
			want:        stringPtr("1"),
		},
		{
			name:        "unicode letters with number",
			partitionID: "αβγ123-789",
			want:        stringPtr("1"),
		},
		{
			name:        "no match - missing hyphen",
			partitionID: "abc123456",
			wantNil:     true,
		},
		{
			name:        "no match - no number after hyphen",
			partitionID: "abc123-",
			wantNil:     true,
		},
		{
			name:        "no match - letters after hyphen",
			partitionID: "abc123-abc",
			wantNil:     true,
		},
		{
			name:        "empty string",
			partitionID: "",
			wantNil:     true,
		},
		{
			name:        "zero after hyphen",
			partitionID: "test1-0",
			want:        stringPtr("0"),
		},
	}

	f := NewFilter()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := f.Hash(tt.partitionID)

			if tt.wantNil {
				if got != nil {
					t.Errorf("Hash() = %v, want nil", *got)
				}
				return
			}

			if got == nil {
				t.Errorf("Hash() = nil, want %v", *tt.want)
				return
			}

			if *got != *tt.want {
				t.Errorf("Hash() = %v, want %v", *got, *tt.want)
			}
		})
	}
}

// Helper function to create string pointers
func stringPtr(s string) *string {
	return &s
}
