package generator

func getFieldOrder(data DockerfileData) []string {
	fieldsOrder := make([]string, 0)

	// Extract field names in the order they appear in the input YAML
	if data.From != "" {
		fieldsOrder = append(fieldsOrder, "From")
	}
	if data.Label != "" {
		fieldsOrder = append(fieldsOrder, "Label")
	}
	if data.Maintainer != "" {
		fieldsOrder = append(fieldsOrder, "Maintainer")
	}
	if data.User != "" {
		fieldsOrder = append(fieldsOrder, "User")
	}
	if data.Volume != "" {
		fieldsOrder = append(fieldsOrder, "Volume")
	}
	if data.OnBuild != "" {
		fieldsOrder = append(fieldsOrder, "OnBuild")
	}
	if data.StopSignal != "" {
		fieldsOrder = append(fieldsOrder, "StopSignal")
	}
	if data.Healthcheck != "" {
		fieldsOrder = append(fieldsOrder, "Healthcheck")
	}
	if data.Workdir != "" {
		fieldsOrder = append(fieldsOrder, "Workdir")
	}
	if len(data.Copy) > 0 {
		fieldsOrder = append(fieldsOrder, "Copy")
		if len(data.Run) > 0 {
			fieldsOrder = append(fieldsOrder, "Run")
		}
	}
	if len(data.CopyCmd) > 0 {
		fieldsOrder = append(fieldsOrder, "CopyCmd")
	}
	if data.Arg != "" {
		fieldsOrder = append(fieldsOrder, "Arg")
	}
	if data.Env != "" {
		fieldsOrder = append(fieldsOrder, "Env")
	}
	if len(data.RunCmd) > 0 {
		fieldsOrder = append(fieldsOrder, "RunCmd")
	}
	if data.Env1 != "" {
		fieldsOrder = append(fieldsOrder, "Env1")
	}
	if data.Env2 != "" {
		fieldsOrder = append(fieldsOrder, "Env2")
	}

	return fieldsOrder
}
