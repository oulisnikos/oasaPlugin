package oasa_sync_api

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"github.com/oulisnikos/oasaPlugin/oasa-sync-decode"
	"github.com/oulisnikos/oasaPlugin/oasa-sync-utils"
	"github.com/oulisnikos/oasaPlugin/oasa-sync-web"
)

func GetLines() ([]map[string]interface{}, error) {
	var result []map[string]interface{}

	response, err := oasa_sync_web.GetRequest("http://telematics.oasa.gr/api/?act=getLines", map[string]string{
		"Accept-Encoding": "gzip, deflate"})
	if err != nil {
		return nil, err
	}

	reader, err := gzip.NewReader(response.Body)

	if err != nil {
		fmt.Printf(err.Error())
		return nil, err
	} else {
		defer reader.Close()

		buf := new(bytes.Buffer)
		buf.ReadFrom(reader)
		responseStr := buf.String()
		result = oasa_sync_decode.ReadTextCharByChar(responseStr, func(dataStr []string) map[string]interface{} {
			var resultRec map[string]interface{}
			if len(dataStr) == 18 {
				resultRec = map[string]interface{}{
					"ID":       oasa_sync_utils.StrToInt64(dataStr[0]),
					"CODE":     dataStr[1],
					"DESCR":    dataStr[2],
					"DESCRENG": dataStr[3],
					"NUM1":     oasa_sync_utils.StrToInt(dataStr[4]),
					"NUM2":     oasa_sync_utils.StrToFloat(dataStr[5]),
					"NUM3":     oasa_sync_utils.StrToInt(dataStr[6]),
					"NUM4":     oasa_sync_utils.StrToFloat(dataStr[7]),
					"NUM5":     oasa_sync_utils.StrToInt(dataStr[8]),
					"NUM6":     oasa_sync_utils.StrToFloat(dataStr[9]),
					"NUM7":     oasa_sync_utils.StrToInt(dataStr[10]),
					"NUM8":     oasa_sync_utils.StrToFloat(dataStr[11]),
					"NUM9":     oasa_sync_utils.StrToInt(dataStr[12]),
					"NUM10":    oasa_sync_utils.StrToFloat(dataStr[13]),
					"NUM11":    oasa_sync_utils.StrToInt(dataStr[14]),
					"NUM12":    oasa_sync_utils.StrToFloat(dataStr[15]),
					"NUM13":    oasa_sync_utils.StrToInt(dataStr[16]),
					"NUM14":    oasa_sync_utils.StrToFloat(dataStr[17]),
				}

			}
			return resultRec
		})
	}

	//jsonStr, error := json.Marshal(result)
	//if error != nil {
	//	return "", err
	//}
	return result, nil
}
