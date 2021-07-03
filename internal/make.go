package main

import (
	"github.com/storozhukBM/build"
	"github.com/storozhukBM/downloader"
)

const golangCiLinterVersion = "1.24.0"

var b = build.NewBuild(build.BuildOptions{})

var commands = []build.Command{{
	Name: `verify`, Body: testDownloaderAndRunRunLinter,
}}

func main() {
	b.Register(commands)
	b.BuildFromOsArgsAndExit()
}

func testDownloaderAndRunRunLinter() {
	urlTemplate := "https://github.com/golangci/golangci-lint/releases/download/v{version}/{fileName}"
	linter, downloadErr := downloader.DownloadExecutable(downloader.DownloadExecutableOptions{
		ExecutableName:           "golangci-lint",
		Version:                  golangCiLinterVersion,
		FileNameTemplate:         "golangci-lint-{version}-{os}-{arch}.{osArchiveType}",
		ReleaseBinaryUrlTemplate: urlTemplate,
		ChecksumFilePath:         "./internal/checksumFile.txt",
		DestinationDirectory:     "bin/linters/",
		BinaryPathInsideTemplate: "golangci-lint-{version}-{os}-{arch}/{executableName}{executableExtension}",
		InfoPrinter:              b.Info,
		WarnPrinter:              b.Warn,
	})
	if downloadErr != nil {
		b.AddError(downloadErr)
		return
	}
	b.Run(linter, `run`)
}
