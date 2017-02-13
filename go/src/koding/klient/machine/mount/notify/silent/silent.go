package silent

import "koding/klient/machine/mount/notify"

// SilentBuilder is a factory for Silent notification objects.
type SilentBuilder struct{}

// Build satisfies notify.Builder interface. It produces Silent objects. Build
// options are not used.
func (SilentBuilder) Build(_ *notify.BuildOpts) (notify.Notifier, error) {
	return Silent{}, nil
}

// Silent is a notification object that doesn't produce any notifications. This
// means that this type should be used only in tests which don't care about
// file system notifications.
type Silent struct{}

// Close satisfies notify.Notifier interface. It does nothing.
func (Silent) Close() {}
