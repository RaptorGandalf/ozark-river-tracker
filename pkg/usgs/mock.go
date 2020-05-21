package usgs

import (
	"bytes"

	gock "gopkg.in/h2non/gock.v1"
)

type Mock struct{ Host string }

func NewMock() *Mock {
	return &Mock{Host: "https://waterservices.usgs.gov"}
}

func (m *Mock) Start() {
	gock.New(m.Host).
		Persist().
		Get("nwis/iv").
		Reply(200).
		Body(bytes.NewBufferString(response))
}

const (
	response = `
	{
		"name":"ns720:timeSeriesResponseType",
		"declaredType":"org.cuahsi.waterml.TimeSeriesResponseType",
		"scope":"javax.xml.bind.JAXBElement$GlobalScope",
		"value":{
		   "queryInfo":{
			  "queryURL":"http://waterservices.usgs.gov/nwis/iv/format=JSON&sites=07064533&parameterCd=00065",
			  "criteria":{
				 "locationParam":"[ALL:07064533]",
				 "variableParam":"[00065]",
				 "parameter":[
	 
				 ]
			  },
			  "note":[]
		   },
		   "timeSeries":[
			  {
				 "sourceInfo":{},
				 "variable":{
					"variableCode":[
					   {
						  "value":"00065",
						  "network":"NWIS",
						  "vocabulary":"NWIS:UnitValues",
						  "variableID":45807202,
						  "default":true
					   }
					],
					"variableName":"Gage height, ft",
					"variableDescription":"Gage height, feet",
					"valueType":"Derived Value",
					"unit":{},
					"options":{},
					"note":[
	 
					],
					"noDataValue":-999999.0,
					"variableProperty":[
	 
					],
					"oid":"45807202"
				 },
				 "values":[],
				 "name":"USGS:07064533:00065:00000"
			  }
		   ]
		},
		"nil":false,
		"globalScope":true,
		"typeSubstituted":false
	 }
	`
)
