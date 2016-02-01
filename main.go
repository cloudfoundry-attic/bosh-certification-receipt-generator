package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/cloudfoundry-incubator/bosh-certification-receipt-generator/artifact"
	"github.com/cloudfoundry-incubator/bosh-certification-receipt-generator/receipt"
)

// set by the build tool
var version string

type ReleaseSlice []string

func (rs *ReleaseSlice) String() string {
	return fmt.Sprintf("%s", *rs)
}

func (rs *ReleaseSlice) Set(val string) error {
	*rs = append(*rs, val)
	return nil
}

func main() {
	var releaseArgs ReleaseSlice
	var stemcellArg string

	showVer := flag.Bool("v", false, "show version information and exit")

	flag.Var(&releaseArgs, "release", "release name/version (multiple)")
	flag.StringVar(&stemcellArg, "stemcell", "", "stemcell name/version (single)")
	flag.Parse()

	if *showVer {
		fmt.Printf("version %s\n", version)
		os.Exit(0)
	}

	if len(releaseArgs) == 0 {
		exitWithUsage("must specify at least one release")
	}

	if stemcellArg == "" {
		exitWithUsage("must specify a stemcell")
	}

	var releases []artifact.Artifact
	for i := range releaseArgs {
		release, err := artifact.New(releaseArgs[i])
		if err != nil {
			exitWithUsage(err.Error())
		}
		releases = append(releases, release)
	}

	stemcell, err := artifact.New(stemcellArg)
	if err != nil {
		exitWithUsage(err.Error())
	}

	r, err := receipt.New(releases, stemcell)
	if err != nil {
		exitWithUsage(err.Error())
	}

	prettyJSON, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		exitWithUsage(fmt.Sprintf("generating JSON: %s", err))
	}

	fmt.Println(string(prettyJSON))
}

func exitWithUsage(msg string) {
	fmt.Println(msg)
	flag.PrintDefaults()
	os.Exit(2)
}
