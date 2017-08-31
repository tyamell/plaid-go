package plaid

import "fmt"

// Version represents the version of the Plaid SDK.
type Version struct {
	// Major and minor version.
	Version float32

	// Increment this for bug releases
	Revision int

	// Suffix is the suffix used in the Plaid version string.
	// It will be blank for release versions.
	Suffix string
}

// CurrentVersion is the default version of the SDK without build information
var CurrentVersion = Version{
	Version:  0.1,
	Revision: 0,
	Suffix:   "-DEV",
}

func (t Version) String() string {
	if t.Revision > 0 {
		return fmt.Sprintf("%.2f.%d%s", t.Version, t.Revision, t.Suffix)
	}
	return fmt.Sprintf("%.2f%s", t.Version, t.Suffix)
}

// ReleaseVersion represents the release version.
func (t Version) ReleaseVersion() Version {
	t.Suffix = ""
	return t
}
