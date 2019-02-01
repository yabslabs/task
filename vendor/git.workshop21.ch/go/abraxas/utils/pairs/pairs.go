/*
 * Copyright (c) 2017 VRSG | Verwaltungsrechenzentrum AG, St.Gallen
 * All Rights Reserved.
 */

package pairs

import (
	"fmt"
)

func PairsString(kv ...string) map[string]string { // nolint: golint
	if len(kv)%2 == 1 {
		panic(fmt.Sprintf("Pairs got the odd number of input pairs for metadata: %d", len(kv)))
	}

	v := map[string]string{}
	var key string
	for i, s := range kv {
		if i%2 == 0 {
			key = s
			continue
		}

		v[key] = s
	}
	return v
}

func Pairs(kv ...interface{}) map[string]interface{} {
	if len(kv)%2 == 1 {
		panic(fmt.Sprintf("Pairs got the odd number of input pairs for metadata: %d", len(kv)))
	}

	v := map[string]interface{}{}
	var key string
	for i, s := range kv {
		if i%2 == 0 {
			key = fmt.Sprint(s)
			continue
		}

		v[key] = s
	}
	return v
}
