package param

import "os"

func AntiDummy(path string)  {
	if path =="/"{
		os.Exit(1)
	}
}
