package oasa_sync_api

import (
	"github.com/oulisnikos/oasaPlugin/oasa-sync-decode"
	"github.com/oulisnikos/oasaPlugin/oasa-sync-utils"
	"github.com/oulisnikos/oasaPlugin/oasa-sync-web"
	"os"
)

const FILE_PATH = "C:\\oasa-telematics\"

func GetData(path string) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	responseStr, error := oasa_sync_web.MakeRequest(path)
	if error != nil {
		return nil, error
	}
	file, error := os.Create(FILE_PATH + path + "_data.txt")
	if error != nil {
		return nil, error
	}
	defer file.Close()
	result = oasa_sync_decode.ReadTextCharByChar(responseStr, nil, file)

	if len(result) == 0 {
		result = append(result, map[string]interface{}{
			"response": responseStr,
		})
	}
	return result, nil
}

func GetLines() ([]map[string]interface{}, error) {
	var result []map[string]interface{}

	responseStr, error := oasa_sync_web.MakeRequest("getLines")
	if error != nil {
		return nil, error
	}
	f, err := os.Create(FILE_PATH + "lines_data.txt")
	if err != nil {
		return nil, err
	}
	defer f.Close()
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
	}, f)

	//jsonStr, error := json.Marshal(result)
	//if error != nil {
	//	return "", err
	//}
	return result, nil
}

func GetRoutes() ([]map[string]interface{}, error) {
	var result []map[string]interface{}

	responseStr, error := oasa_sync_web.MakeRequest("getRoutes")
	if error != nil {
		return nil, error
	}
	result = append(result, map[string]interface{}{
		"response": responseStr,
	})
	f, err := os.Create(FILE_PATH + "routes_data.txt")
	if err != nil {
		return nil, err
	}
	defer f.Close()
	result = oasa_sync_decode.ReadTextCharByChar(responseStr, func(dataStr []string) map[string]interface{} {
		var resultRec map[string]interface{}
		if len(dataStr) == 6 {
			resultRec = map[string]interface{}{
				"RouteCode":     oasa_sync_utils.StrToInt64(dataStr[0]),
				"LineCode":      oasa_sync_utils.StrToInt64(dataStr[1]),
				"RouteDescr":    dataStr[2],
				"RouteDescrEng": dataStr[3],
				"RouteType":     oasa_sync_utils.StrToInt(dataStr[4]),
				"RouteDistance": oasa_sync_utils.StrToFloat(dataStr[5]),
			}

		}
		return resultRec
	}, f)
	//
	////jsonStr, error := json.Marshal(result)
	//if error != nil {
	//	return "", err
	//}
	return result, nil
}

func GetStops() ([]map[string]interface{}, error) {
	var result []map[string]interface{}

	responseStr, error := oasa_sync_web.MakeRequest("getStops")
	if error != nil {
		return nil, error
	}
	f, err := os.Create(FILE_PATH + "stops-data.txt")
	if err != nil {
		return nil, err
	}
	defer f.Close()
	result = append(result, map[string]interface{}{
		"response": responseStr,
	})
	result = oasa_sync_decode.ReadTextCharByChar(responseStr, func(dataStr []string) map[string]interface{} {
		var resultRec map[string]interface{}
		if len(dataStr) == 13 {
			resultRec = map[string]interface{}{
				"StopCode":      oasa_sync_utils.StrToInt64(dataStr[0]),
				"StopID":        dataStr[1],
				"StopDescr":     dataStr[2],
				"StopDescrEng":  dataStr[3],
				"StopStreet":    dataStr[4],
				"StopStreetEng": dataStr[5],
				"StopHeading":   dataStr[6],
				"StopLat":       oasa_sync_utils.StrToFloat(dataStr[7]),
				"StopLng":       oasa_sync_utils.StrToFloat(dataStr[8]),
				//"RouteStopOrder": oasa_sync_utils.StrToInt(dataStr[9]),
				"StopType":    dataStr[9],
				"StopAmea":    dataStr[10],
				"DescrStr":    dataStr[11],
				"DescrStrEng": dataStr[12],
			}

		}
		return resultRec
	}, f)
	//
	////jsonStr, error := json.Marshal(result)
	//if error != nil {
	//	return "", err
	//}
	return result, nil
}
