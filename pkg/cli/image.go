package cli

func resolveImageRegistry(imageRegistry string, imageRepository string) string {
	if imageRepository != "" {
		return imageRepository
	}
	return imageRegistry
}
