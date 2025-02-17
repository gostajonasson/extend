package main

import (
	"context"
	"net/http"

	"go.einride.tech/sage/sg"
	"go.einride.tech/sage/tools/sgconvco"
	"go.einride.tech/sage/tools/sggit"
	"go.einride.tech/sage/tools/sggolangcilint"
	"go.einride.tech/sage/tools/sgmarkdownfmt"
	"go.einride.tech/sage/tools/sgyamlfmt"
)

func main() {
	sg.GenerateMakefiles(
		sg.Makefile{
			Path:          sg.FromGitRoot("Makefile"),
			DefaultTarget: Default,
		},

		sg.Makefile{
			Path:          sg.FromGitRoot("proto", "Makefile"),
			Namespace:     Proto{},
			DefaultTarget: Proto.Default,
		},
	)
}

func Default(ctx context.Context) error {
	sg.Deps(ctx, ConvcoCheck, FormatMarkdown, FormatYaml)
	sg.Deps(ctx, Proto.Default)
	sg.Deps(ctx, GoModTidy)
	sg.Deps(ctx, GoLint)
	sg.Deps(ctx, GitVerifyNoDiff)
	return nil
}

func GoModTidy(ctx context.Context) error {
	sg.Logger(ctx).Println("tidying Go module files...")
	command := sg.Command(ctx, "go", "mod", "tidy", "-v")
	command.Dir = sg.FromGitRoot()
	return command.Run()
}

func ConvcoCheck(ctx context.Context) error {
	sg.Logger(ctx).Println("checking git commits...")
	return sgconvco.Command(ctx, "check", "origin/master..HEAD").Run()
}

func FormatMarkdown(ctx context.Context) error {
	sg.Logger(ctx).Println("formatting Markdown files...")
	return sgmarkdownfmt.Command(ctx, "-w", ".").Run()
}

func FormatYaml(ctx context.Context) error {
	sg.Logger(ctx).Println("formatting YAML files...")
	return sgyamlfmt.Run(ctx)
}

func GitVerifyNoDiff(ctx context.Context) error {
	sg.Logger(ctx).Println("verifying that git has no diff...")
	return sggit.VerifyNoDiff(ctx)
}

func GoLint(ctx context.Context) error {
	sg.Logger(ctx).Println("linting Go files...")
	return sggolangcilint.CommandInDirectory(
		ctx,
		sg.FromGitRoot(),
	).Run()
}

func ServePDF(ctx context.Context) error {
	const addr = ":8080"
	httpMux := http.NewServeMux()
	httpMux.Handle("/openapiv2/", http.StripPrefix("/openapiv2", http.FileServer(http.Dir("./openapiv2"))))
	httpMux.Handle("/", http.FileServer(http.Dir("./pdf")))
	sg.Logger(ctx).Printf("listening on %s", addr)
	return (&http.Server{
		Addr:    addr,
		Handler: httpMux,
	}).ListenAndServe()
}
