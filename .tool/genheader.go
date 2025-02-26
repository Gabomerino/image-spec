// Copyright 2017 The Linux Foundation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the 
var headerTemplate = template.Must(template.New("gen").Parse(`<title>image-spec {{.Version}}</title>
<base href="https://raw.githubusercontent.com/opencontainers/image-spec/{{.Branch}}/">`))

type Obj struct {
	Version string
	Branch  string
}

func main() {
	obj := Obj{
		Version: specs.Version,
		Branch:  specs.Version,
	}
	if strings.HasSuffix(specs.Version, "-dev") {
		cmd := exec.Command("git", "log", "-1", `--pretty=%H`, "HEAD")
		var out bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		obj.Branch = strings.Trim(out.String(), " \n\r")
	}
	headerTemplate.Execute(os.Stdout, obj)
}
