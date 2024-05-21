package utils

import "fmt"

func Contains(stringList []string, st string) bool {
	for _, s := range stringList {
		if s == st {
			return true
		}
	}
	return false
}

func AppendMaps(map1, map2 map[string]int) map[string]int {
	// Create a new map to avoid modifying the original map1
	mergedMap := make(map[string]int)

	// Copy all entries from map1 to the mergedMap
	for key, value := range map1 {
		mergedMap[key] = value
	}

	// Copy all entries from map2 to the mergedMap, overwriting any duplicates
	for key, value := range map2 {
		mergedMap[key] = value
	}

	return mergedMap
}

func GetKeys(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}

func StrListToGql(list []string) string {
	str := "["
	for index, item := range list {
		str += fmt.Sprintf("\"%s\"", item)
		if index != len(list)-1 {
			str += ","
		}
	}
	return str + "]"

}
