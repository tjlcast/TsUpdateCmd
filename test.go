package main

import (
	"path/filepath"
	"fmt"
)

func main() {
	pattern 	:= `C:\\SfDevCluster\\Data\\_App\\_Node_[0-9]\\Ipam_App[0-9]\\IpamPolicyServicePkg.Code.1.0.0.0\\wwwroot\\scripts`
	fmt.Println(pattern)
	const path		= `C:\SfDevCluster\Data\_App\_Node_0\Ipam_App0\IpamPolicyServicePkg.Code.1.0.0.0\wwwroot\scripts`
	//const path		= "C:\\SfDevCluster\\Data\\_App\\_Node_0\\Ipam_App0\\IpamPolicyServicePkg.Code.1.0.0.0\\wwwroot\\scripts"
		matched, _ := filepath.Match(pattern, path)
	fmt.Println(matched)

	pattern1	:= `c:\\a\\b[0-9]\\c`
	path1	:= "c:\\a\\b1\\c"

	result, _ := filepath.Match(pattern1, path1)
	fmt.Println(result)
}
