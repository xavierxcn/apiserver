package main

import (
	"context"
	"fmt"

	"github.com/containers/buildah"
	is "github.com/containers/image/v5/storage"
	"github.com/containers/storage"
	"github.com/containers/storage/pkg/unshare"
)

func main() {
	if buildah.InitReexec() {
		return
	}
	unshare.MaybeReexecUsingUserNamespace(false)

	buildStoreOptions, err := storage.DefaultStoreOptions(unshare.IsRootless(), unshare.GetRootlessUID())
	if err != nil {
		panic(err)
	}

	buildStore, err := storage.GetStore(buildStoreOptions)
	if err != nil {
		panic(err)
	}
	defer buildStore.Shutdown(false)

	builderOpts := buildah.BuilderOptions{
		Args: map[string]string{
			"-f": "Dockerfile",
		},
	}

	builder, err := buildah.NewBuilder(context.TODO(), buildStore, builderOpts)
	if err != nil {
		panic(err)
	}
	defer builder.Delete()

	// buildah.BuildUs

	imageRef, err := is.Transport.ParseStoreReference(buildStore, "docker.io/myusername/my-image")
	if err != nil {
		panic(err)
	}

	imageId, _, _, err := builder.Commit(context.TODO(), imageRef, buildah.CommitOptions{})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Image built! %s\n", imageId)
}
