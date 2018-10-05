package version

import (
	"fmt"
)

const (
	STABLE Status = "Stable"
	BETA   Status = "Beta"
	ALPHA  Status = "Alpha"
)

type Version struct {
	Major  int
	Minor  int
	Build  int
	Status Status
	Date   string
}

type Status string

func newStatus() Status {
	return ALPHA
}

func New(major int, minor int, build int, status Status, date string) Version {
	return Version{
		Major:  major,
		Minor:  minor,
		Build:  build,
		Status: status,
		Date:   date,
	}
}

func (v Version) ToString() string {
	return fmt.Sprintf("%d.%d.%d %s (%s)", v.Major, v.Minor, v.Build, v.Status, v.Date)
}
