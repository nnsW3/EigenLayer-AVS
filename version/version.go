package version

var (
	// Version can also be set through tag release at build time
	semver = "1.0.0"
)

// Get return the version. Note that we're injecting this at build time when we tag release
func Get() string {
	return semver
}
