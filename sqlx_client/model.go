package sqlx_client

import "strconv"

func LoginByDeviceName(deviceName string) (map[string]interface{},error){
	resMap:=map[string]interface{}{}
	rows,err:=db.Queryx("select secret_key as ProductSecret,product_id as ProductId," +
		"device_name as DeviceName,device_id as DeviceId from device_detail where device_name=?",deviceName)
	if rows!=nil{
		for rows.Next(){
			err=rows.MapScan(resMap)
		}
	}
	//mapscan返回的是[]uint8，需要你自己转换类型
	if resMap["ProductSecret"]!=nil{
		t:=string(resMap["ProductSecret"].([]uint8))
		resMap["ProductSecret"]=t
	}
	if resMap["ProductId"]!=nil{
		t:=string(resMap["ProductId"].([]uint8))
		t1,_:=strconv.Atoi(t)
		resMap["ProductId"]=t1
	}
	if resMap["DeviceName"]!=nil{
		t:=string(resMap["DeviceName"].([]uint8))
		resMap["DeviceName"]=t
	}
	if resMap["DeviceId"]!=nil{
		t:=string(resMap["DeviceId"].([]uint8))
		resMap["DeviceId"]=t
	}
	return resMap,err
}