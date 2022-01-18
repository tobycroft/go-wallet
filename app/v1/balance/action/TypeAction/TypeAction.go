package TypeAction

import "main.go/app/v1/balance/model/TypeModel"

func App_map_type() (name map[int64]interface{}, direction map[int64]interface{}) {
	types := TypeModel.Api_select()
	name = map[int64]interface{}{}
	direction = map[int64]interface{}{}
	for _, data := range types {
		name[data["id"].(int64)] = data["name"]
		direction[data["id"].(int64)] = data["direction"]
	}
	return
}

func App_select_in_out_exchange() (in []interface{}, out []interface{}, exchange []interface{}) {
	types := TypeModel.Api_select()
	for _, data := range types {
		if data["direction"].(string) == "in" {
			in = append(in, data["id"])
			continue
		}
		if data["direction"].(string) == "out" {
			out = append(out, data["id"])
			continue
		}
		if data["direction"].(string) == "exchange" {
			exchange = append(exchange, data["id"])
		}
	}
	return
}
