package cli

import "testing"

func TestResolveImageRegistry(t *testing.T) {
	t.Run("uses image registry by default", func(t *testing.T) {
		if got, want := resolveImageRegistry("registry.example.com", ""), "registry.example.com"; got != want {
			t.Fatalf("resolveImageRegistry() = %q, want %q", got, want)
		}
	})

	t.Run("supports the legacy image repository flag", func(t *testing.T) {
		if got, want := resolveImageRegistry("registry.example.com", "legacy.example.com"), "legacy.example.com"; got != want {
			t.Fatalf("resolveImageRegistry() = %q, want %q", got, want)
		}
	})
}
