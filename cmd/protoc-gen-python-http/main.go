package main

import (
	"flag"
	"google.golang.org/protobuf/compiler/protogen"
)

func main() {
	var (
		flags flag.FlagSet
	)

	gen := protogen.Options{
		ParamFunc: flags.Set,
		ImportRewriteFunc: func(path protogen.GoImportPath) protogen.GoImportPath {
			return path
		},
	}

	gen.Run(func(plugin *protogen.Plugin) error {
		plugin.SupportedFeatures = 0

		for _, file := range plugin.Files {
			if !file.Generate {
				continue
			}

			GenerateFile(plugin, file)
		}

		return nil
	})
}
