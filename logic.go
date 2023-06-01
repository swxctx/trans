package trans

import "fmt"

func (trans *Trans) reload() error {
	// read file
	dataList, err := ReadCsv(trans.Config.FilePath)
	if err != nil {
		return err
	}

	// reload data logic
	for i, data := range dataList {
		// i == 0 表头
		if i == 0 {
			trans.langCodes = data
			trans.sourceLanIndex = trans.getSourceLanIndex()
			if trans.sourceLanIndex < 0 {
				return fmt.Errorf("trans: reload err, no source language")
			}
			continue
		}

		// insert lan
		trans.insertData(data)
	}

	return nil
}

// insertData
func (trans *Trans) insertData(dataList []string) {
	// other lan map
	other := make(map[string]string)
	for i, d := range dataList {
		other[trans.langCodes[i]] = d
	}

	// insert main lan
	trans.Data[dataList[trans.sourceLanIndex]] = other
}

// getSourceLanIndex
func (trans *Trans) getSourceLanIndex() int {
	for i, d := range trans.langCodes {
		if d == trans.Config.SourceLan {
			return i
		}
	}
	return -1
}

// translate
func (trans *Trans) translate(lanCode, text string) string {
	return trans.Data[text][lanCode]
}
