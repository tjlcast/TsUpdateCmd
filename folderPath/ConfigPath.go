package folderPath

import (
	"strings"
)

const SFNODESPATH = "C:\\SfDevCluster\\Data\\_App\\"
const PATH_SCRIPT = `\\_Node_\d\\Ipam_App\d\\IpamPolicyServicePkg.Code.1.0.0.0\\wwwroot\\scripts`

type ConfigPath struct {
	localProjTsPath		Entry
	clusterProjTsPath 	Entry
}

func Parse(localProjTsPath string, clusterProjPath string) *ConfigPath {
	cp := &ConfigPath{}
	cp.parseLocalProjPath(localProjTsPath)
	cp.parseClusterProjPathRegex(clusterProjPath)
	return cp
}

func (self *ConfigPath)CopyFromLocal2Cluster(fileName string) error {
	bytes, _, err := self.localProjTsPath.readFile(fileName)
	if err != nil {
		return err
	}
	_, err = self.clusterProjTsPath.addFile(bytes, fileName)
	if err != nil {
		return err
	}
	return nil
}

func (self *ConfigPath) parseLocalProjPath(localProjTsPath string) {
	self.localProjTsPath = newEntry(localProjTsPath)
}

func (self *ConfigPath) parseClusterProjPathRegex(cluterProjPath string) {
	cluterProjPath = strings.Replace(cluterProjPath, `\`, `\\`, -1)
	entryPath := cluterProjPath + PATH_SCRIPT
	self.clusterProjTsPath = newRegexEntry(cluterProjPath, entryPath)
}

func (self *ConfigPath) String() string {
	strs := make([]string, 2)
	strs[0] = self.localProjTsPath.String()
	strs[1] = self.clusterProjTsPath.String()
	return strings.Join(strs, pathListSeparator)
}
