package main

import (
	"archive/zip"
	"bytes"
	"compress/flate"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/cloudfoundry-attic/jibber_jabber"
	"github.com/hellodword/archiver"
	"github.com/hellodword/gopkid"
)

var userLocale string

var (
	encoding string
)

func init() {

	// detect the locale
	userLocale, _ = jibber_jabber.DetectIETF()

	if userLocale == "zh-CN" {
		flag.StringVar(&encoding, "encoding", "utf-8", "编码名称, 例如 utf-8、gbk、gb18030, 默认为 utf-8")
	} else {
		flag.StringVar(&encoding, "encoding", "utf-8", "encoding name, utf-8 by default")
	}
}

func main() {

	if len(os.Args) >= 2 &&
		(os.Args[1] == "-h" || os.Args[1] == "--help" || os.Args[1] == "help") {
		fmt.Println(usageString())
		os.Exit(0)
	}

	if len(os.Args) < 2 {
		fatal(usageString())
	}

	flag.Parse()

	apkPath := flag.Arg(0)

	if !exists(apkPath) {
		fatal(apkPath, "not exist!")
	}

	z := archiver.Zip{
		CompressionLevel:       flate.DefaultCompression,
		MkdirAll:               true,
		SelectiveCompression:   true,
		ContinueOnError:        false,
		OverwriteExisting:      false,
		ImplicitTopLevelFolder: false,
	}

	m := gopkid.GetInfo()

	z.Walk(apkPath, func(f archiver.File) error {
		zfh, ok := f.Header.(zip.FileHeader)
		s := strings.ToLower(zfh.Name)
		if ok {

			if strings.HasSuffix(s, ".png") {

			} else if strings.HasSuffix(s, ".mp3") {

			} else if strings.HasSuffix(s, ".js") {

			} else if strings.HasSuffix(s, ".otf") {

			} else if strings.HasSuffix(s, ".ttf") {

			} else if strings.HasSuffix(s, ".woff") {

			} else if strings.HasSuffix(s, ".css") {

			} else if strings.HasSuffix(s, ".json") {

			} else if strings.HasSuffix(s, ".html") {

			} else if strings.HasSuffix(s, ".dex") {

			} else if strings.HasSuffix(s, ".webp") {

			} else if strings.HasSuffix(s, ".ogg") {

			} else if strings.HasSuffix(s, ".txt") {

			} else {
				for k, v := range m {
					for _, p := range v {
						if strings.Contains(s, strings.ToLower(p)) {
							fmt.Fprintln(os.Stdout, "\t", k, zfh.Name)
							break
						}

					}
				}
			}

		}
		return nil
	})

}

func exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func fatal(v ...interface{}) {
	fmt.Fprintln(os.Stderr, v...)
	os.Exit(1)
}

func usageString() string {

	buf := new(bytes.Buffer)

	if userLocale == "zh-CN" {
		fmt.Println("Locale:", userLocale)
		buf.WriteString(usageCN)
	} else {
		buf.WriteString(usage)
	}

	flag.CommandLine.SetOutput(buf)
	flag.CommandLine.PrintDefaults()
	return buf.String()
}

const usage = `Usage: pkid {help} [apk file path] ([--encoding ...])
  help
    Display this help text. Also -h or --help.

 Detect the reinforce of apk file. 

`

const usageCN = `用法: pkid {help} [apk 文件路径] ([--encoding ...])
  help
    展示帮助信息. -h 或者 --help 也能达到相同效果.

 探测 apk 文件的加固.  

`
