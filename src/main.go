package main

import (
	github "github.com/pulumi/pulumi-github/sdk/v4/go/github"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	// configure repository with best practices
	pulumi.Run(func(ctx *pulumi.Context) error {
		repository, err := github.NewRepository(ctx, "demo-repo", &github.RepositoryArgs{
			Description: pulumi.String("Generated from automated test"),
			Visibility:  pulumi.String("private"),
		})
		if err != nil {
			return err
		}

		ctx.Export("repositiory", repository)
		return nil
	})
}
