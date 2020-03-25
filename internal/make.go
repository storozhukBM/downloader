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
		ChecksumFileContent:      golangcilintChecksumFile,
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

const golangcilintChecksumFile = `
aeaa5498682246b87d0b77ece283897348ea03d98e816760a074058bfca60b2a  golangci-lint-1.24.0-windows-amd64.zip
7e854a70d449fe77b7a91583ec88c8603eb3bf96c45d52797dc4ba3f2f278dbe  golangci-lint-1.24.0-darwin-386.tar.gz
835101fae192c3a2e7a51cb19d5ac3e1a40b0e311955e89bc21d61de78635979  golangci-lint-1.24.0-linux-armv6.tar.gz
a041a6e6a61c9ff3dbe58673af13ea00c76bcd462abede0ade645808e97cdd6d  golangci-lint-1.24.0-windows-386.zip
7cc73eb9ca02b7a766c72b913f8080401862b10e7bb90c09b085415a81f21609  golangci-lint-1.24.0-freebsd-armv6.tar.gz
537bb2186987b5e68ad4e8829230557f26087c3028eb736dea1662a851bad73d  golangci-lint-1.24.0-linux-armv7.tar.gz
8cb1bc1e63d8f0d9b71fcb10b38887e1646a6b8a120ded2e0cd7c3284528f633  golangci-lint-1.24.0-linux-mips64.tar.gz
095d3f8bf7fc431739861574d0b58d411a617df2ed5698ce5ae5ecc66d23d44d  golangci-lint-1.24.0-freebsd-armv7.tar.gz
e245df27cec3827aef9e7afbac59e92816978ee3b64f84f7b88562ff4b2ac225  golangci-lint-1.24.0-linux-arm64.tar.gz
35d6d5927e19f0577cf527f0e4441dbb37701d87e8cf729c98a510fce397fbf7  golangci-lint-1.24.0-linux-ppc64le.tar.gz
a1ed66353b8ceb575d78db3051491bce3ac1560e469a9bc87e8554486fec7dfe  golangci-lint-1.24.0-freebsd-386.tar.gz
241ca454102e909de04957ff8a5754c757cefa255758b3e1fba8a4533d19d179  golangci-lint-1.24.0-linux-amd64.tar.gz
ff488423db01a0ec8ffbe4e1d65ef1be6a8d5e6d7930cf380ce8aaf714125470  golangci-lint-1.24.0-linux-386.tar.gz
f05af56f15ebbcf77663a8955d1e39009b584ce8ea4c5583669369d80353a113  golangci-lint-1.24.0-darwin-amd64.tar.gz
b0096796c0ffcd6c350a2ec006100e7ef5f0597b43a204349d4f997273fb32a7  golangci-lint-1.24.0-freebsd-amd64.tar.gz
c9c2867380e85628813f1f7d1c3cfc6c6f7931e89bea86f567ff451b8cdb6654  golangci-lint-1.24.0-linux-mips64le.tar.gz
2feb97fa61c934aa3eba9bc104ab5dd8fb946791d58e64060e8857e800eeae0b  golangci-lint-1.24.0-linux-s390x.tar.gz
`
