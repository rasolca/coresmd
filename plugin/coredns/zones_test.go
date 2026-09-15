// SPDX-FileCopyrightText: © 2026 OpenCHAMI a Series of LF Projects, LLC
//
// SPDX-License-Identifier: MIT

package plugin

import "testing"

func TestZoneAllowsBMCType(t *testing.T) {
	tests := []struct {
		name     string
		zone     Zone
		compType string
		want     bool
	}{
		{
			name:     "default allows NodeBMC",
			zone:     Zone{Name: "cluster.local"},
			compType: "NodeBMC",
			want:     true,
		},
		{
			name:     "default denies RouterBMC",
			zone:     Zone{Name: "cluster.local"},
			compType: "RouterBMC",
			want:     false,
		},
		{
			name:     "default denies ChassisBMC",
			zone:     Zone{Name: "cluster.local"},
			compType: "ChassisBMC",
			want:     false,
		},
		{
			name:     "explicit RouterBMC allowed",
			zone:     Zone{Name: "cluster.local", BMCTypes: []string{"NodeBMC", "RouterBMC"}},
			compType: "RouterBMC",
			want:     true,
		},
		{
			name:     "explicit ChassisBMC allowed",
			zone:     Zone{Name: "cluster.local", BMCTypes: []string{"NodeBMC", "ChassisBMC"}},
			compType: "ChassisBMC",
			want:     true,
		},
		{
			name:     "explicit NodeBMC still allowed",
			zone:     Zone{Name: "cluster.local", BMCTypes: []string{"NodeBMC", "RouterBMC"}},
			compType: "NodeBMC",
			want:     true,
		},
		{
			name:     "explicit list denies unlisted type",
			zone:     Zone{Name: "cluster.local", BMCTypes: []string{"RouterBMC"}},
			compType: "NodeBMC",
			want:     false,
		},
		{
			name:     "Node type is never allowed as BMC",
			zone:     Zone{Name: "cluster.local", BMCTypes: []string{"NodeBMC", "RouterBMC", "ChassisBMC"}},
			compType: "Node",
			want:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.zone.AllowsBMCType(tt.compType); got != tt.want {
				t.Errorf("Zone.AllowsBMCType(%q) = %v, want %v", tt.compType, got, tt.want)
			}
		})
	}
}
