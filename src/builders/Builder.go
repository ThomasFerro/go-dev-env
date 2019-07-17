package builders

// ArtifactPath The path to the builded artifact
type ArtifactPath string

// Builder An artifact builder
type Builder interface {
	Build(contextPath string) (ArtifactPath, error)
}
