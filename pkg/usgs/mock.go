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
			  "queryURL":"http://waterservices.usgs.gov/nwis/iv/format=JSON&sites=07067000&parameterCd=00065,00060,00010",
			  "criteria":{
				 "locationParam":"[ALL:07067000]",
				 "variableParam":"[00065, 00060, 00010]",
				 "parameter":[
	 
				 ]
			  },
			  "note":[
				 {
					"value":"[ALL:07067000]",
					"title":"filter:sites"
				 },
				 {
					"value":"[mode=LATEST, modifiedSince=null]",
					"title":"filter:timeRange"
				 },
				 {
					"value":"methodIds=[ALL]",
					"title":"filter:methodId"
				 },
				 {
					"value":"2020-06-10T00:08:23.525Z",
					"title":"requestDT"
				 },
				 {
					"value":"7fe92a40-aaae-11ea-948a-6cae8b6642f6",
					"title":"requestId"
				 },
				 {
					"value":"Provisional data are subject to revision. Go to http://waterdata.usgs.gov/nwis/help/?provisional for more information.",
					"title":"disclaimer"
				 },
				 {
					"value":"caas01",
					"title":"server"
				 }
			  ]
		   },
		   "timeSeries":[
			  {
				 "sourceInfo":{
					"siteName":"Current River at Van Buren, MO",
					"siteCode":[
					   {
						  "value":"07067000",
						  "network":"NWIS",
						  "agencyCode":"USGS"
					   }
					],
					"timeZoneInfo":{
					   "defaultTimeZone":{
						  "zoneOffset":"-06:00",
						  "zoneAbbreviation":"CST"
					   },
					   "daylightSavingsTimeZone":{
						  "zoneOffset":"-05:00",
						  "zoneAbbreviation":"CDT"
					   },
					   "siteUsesDaylightSavingsTime":true
					},
					"geoLocation":{
					   "geogLocation":{
						  "srs":"EPSG:4326",
						  "latitude":36.99138889,
						  "longitude":-91.0135
					   },
					   "localSiteXY":[
	 
					   ]
					},
					"note":[
	 
					],
					"siteType":[
	 
					],
					"siteProperty":[
					   {
						  "value":"ST",
						  "name":"siteTypeCd"
					   },
					   {
						  "value":"11010008",
						  "name":"hucCd"
					   },
					   {
						  "value":"29",
						  "name":"stateCd"
					   },
					   {
						  "value":"29035",
						  "name":"countyCd"
					   }
					]
				 },
				 "variable":{
					"variableCode":[
					   {
						  "value":"00060",
						  "network":"NWIS",
						  "vocabulary":"NWIS:UnitValues",
						  "variableID":45807197,
						  "default":true
					   }
					],
					"variableName":"Streamflow, ft&#179;/s",
					"variableDescription":"Discharge, cubic feet per second",
					"valueType":"Derived Value",
					"unit":{
					   "unitCode":"ft3/s"
					},
					"options":{
					   "option":[
						  {
							 "name":"Statistic",
							 "optionCode":"00000"
						  }
					   ]
					},
					"note":[
	 
					],
					"noDataValue":-999999.0,
					"variableProperty":[
	 
					],
					"oid":"45807197"
				 },
				 "values":[
					{
					   "value":[
						  {
							 "value":"7760",
							 "qualifiers":[
								"P"
							 ],
							 "dateTime":"2020-06-09T18:30:00.000-05:00"
						  }
					   ],
					   "qualifier":[
						  {
							 "qualifierCode":"P",
							 "qualifierDescription":"Provisional data subject to revision.",
							 "qualifierID":0,
							 "network":"NWIS",
							 "vocabulary":"uv_rmk_cd"
						  }
					   ],
					   "qualityControlLevel":[
	 
					   ],
					   "method":[
						  {
							 "methodDescription":"",
							 "methodID":77827
						  }
					   ],
					   "source":[
	 
					   ],
					   "offset":[
	 
					   ],
					   "sample":[
	 
					   ],
					   "censorCode":[
	 
					   ]
					}
				 ],
				 "name":"USGS:07067000:00060:00000"
			  },
			  {
				 "sourceInfo":{
					"siteName":"Current River at Van Buren, MO",
					"siteCode":[
					   {
						  "value":"07067000",
						  "network":"NWIS",
						  "agencyCode":"USGS"
					   }
					],
					"timeZoneInfo":{
					   "defaultTimeZone":{
						  "zoneOffset":"-06:00",
						  "zoneAbbreviation":"CST"
					   },
					   "daylightSavingsTimeZone":{
						  "zoneOffset":"-05:00",
						  "zoneAbbreviation":"CDT"
					   },
					   "siteUsesDaylightSavingsTime":true
					},
					"geoLocation":{
					   "geogLocation":{
						  "srs":"EPSG:4326",
						  "latitude":36.9,
						  "longitude":-91.0
					   },
					   "localSiteXY":[
	 
					   ]
					},
					"note":[
	 
					],
					"siteType":[
	 
					],
					"siteProperty":[
					   {
						  "value":"ST",
						  "name":"siteTypeCd"
					   },
					   {
						  "value":"11010008",
						  "name":"hucCd"
					   },
					   {
						  "value":"29",
						  "name":"stateCd"
					   },
					   {
						  "value":"29035",
						  "name":"countyCd"
					   }
					]
				 },
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
					"unit":{
					   "unitCode":"ft"
					},
					"options":{
					   "option":[
						  {
							 "name":"Statistic",
							 "optionCode":"00000"
						  }
					   ]
					},
					"note":[
	 
					],
					"noDataValue":-999999.0,
					"variableProperty":[
	 
					],
					"oid":"45807202"
				 },
				 "values":[
					{
					   "value":[
						  {
							 "value":"6.98",
							 "qualifiers":[
								"P"
							 ],
							 "dateTime":"2020-06-09T18:30:00.000-05:00"
						  }
					   ],
					   "qualifier":[
						  {
							 "qualifierCode":"P",
							 "qualifierDescription":"Provisional data subject to revision.",
							 "qualifierID":0,
							 "network":"NWIS",
							 "vocabulary":"uv_rmk_cd"
						  }
					   ],
					   "qualityControlLevel":[
	 
					   ],
					   "method":[
						  {
							 "methodDescription":"",
							 "methodID":77826
						  }
					   ],
					   "source":[
	 
					   ],
					   "offset":[
	 
					   ],
					   "sample":[
	 
					   ],
					   "censorCode":[
	 
					   ]
					}
				 ],
				 "name":"USGS:07067000:00065:00000"
			  }
		   ]
		},
		"nil":false,
		"globalScope":true,
		"typeSubstituted":false
	 }
	`
)
