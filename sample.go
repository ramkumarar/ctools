import "github.com/r3labs/diff"
n := map[string]string{"name": "ram sdlhfsfh sdlhfshdfhsdlfh  sdfhslhflshf sfhsdlf slhflshf slhflshfl shlfhslhflshdfhsdlfhosfoahosh fsfh sdh", "age": "42", "gender": "M", "city": "chennai"}
	m := map[string]string{"name": "kokila sdhshfd sfhlshfd shfshlf shflshdflhslfhlsf sjf;sjfspfjs;df;sjf sjfsjdf ssfshfshdlfhsl flsdfpsdfshkldfhlshdf lshfhsfsf sfamk'fksf'js fsd", "age": "7", "gender": "F", "city": "chennai"}

	changelogs, _ := diff.Diff(n, m)

	var data [][]string

	for _, changelog := range changelogs {
		var row []string

		row = append(row, strings.Join(changelog.Path, ""))
		row = append(row, fmt.Sprintf("%s", changelog.From))
		row = append(row, fmt.Sprintf("%s", changelog.To))
		data = append(data, row)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeader([]string{"Key", "Source", "Destination"})
	table.AppendBulk(data)
	table.SetRowLine(true)
	table.Render()

func getConfigMapFromUcdObj(configs []string) map[string]string {
	configMap := make(map[string]string)
	configSize := len(configs)
	keySize := configSize / 2

	for i := 0; i < keySize; i++ {
		configMap[configs[i]] = configs[keySize+i]
	}
	return configMap
}

