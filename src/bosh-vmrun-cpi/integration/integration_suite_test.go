package integration_test

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"text/template"

	"github.com/mholt/archiver"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

func TestIntegration(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Integration Suite")
}

var ExtractedStemcellTempDir string
var CpiConfigPath string
var vmStoreDir string

func GexecCommandWithStdin(commandBin string, commandArgs ...string) (*gexec.Session, io.WriteCloser) {
	command := exec.Command(commandBin, commandArgs...)
	stdin, err := command.StdinPipe()
	Expect(err).ToNot(HaveOccurred())

	session, err := gexec.Start(command, GinkgoWriter, os.Stderr)
	Expect(err).ToNot(HaveOccurred())

	return session, stdin
}

func extractStemcell() string {
	stemcellFile := "../../../ci/deploy-test/state/stemcell.tgz"

	stemcellTempDir, err := ioutil.TempDir("", "stemcell-")
	Expect(err).ToNot(HaveOccurred())

	err = archiver.TarGz.Open(stemcellFile, stemcellTempDir)
	Expect(err).ToNot(HaveOccurred())

	return stemcellTempDir
}

var configTemplate, _ = template.New("parse").Parse(`{
	"cloud": {
		"plugin": "vsphere",
		"properties": {
			"vmrun": {
				"vm_store_path": "{{.VmStorePath}}",
				"vmrun_bin_path": "{{.VmrunBinPath}}",
				"ovftool_bin_path": "{{.OvftoolBinPath}}"
			},
			"vcenters": [
			{
				"host": "{{.EsxiHost}}",
				"user": "{{.EsxiUser}}",
				"password": "{{.EsxiPassword}}",
				"datacenters": [
				{
					"name": "{{.EsxiDatacenter}}",
					"vm_folder": "BOSH_VMs",
					"template_folder": "BOSH_Templates",
					"disk_path": "bosh_disks",
					"datastore_pattern": "{{.EsxiDatastore}}"
				}
				]
			}
			],
			"agent": {
				"ntp": [
				],
				"blobstore": {
					"provider": "local",
					"options": {
						"blobstore_path": "/var/vcap/micro_bosh/data/cache"
					}
				},
				"mbus": "https://mbus:p2an3m7idfm6vmqp3w74@0.0.0.0:6868"
			}
		}
	}
}`)

func generateCPIConfig() (string, string) {
	vmStoreTempDir, err := ioutil.TempDir("", "vm-store-path-")
	Expect(err).ToNot(HaveOccurred())

	var configValues = struct {
		VmStorePath    string
		VmrunBinPath   string
		OvftoolBinPath string
		EsxiHost       string
		EsxiUser       string
		EsxiPassword   string
		EsxiDatacenter string
		EsxiDatastore  string
	}{
		VmStorePath:    vmStoreTempDir,
		VmrunBinPath:   os.Getenv("VMRUN_BIN_PATH"),
		OvftoolBinPath: os.Getenv("OVFTOOL_BIN_PATH"),
		EsxiHost:       os.Getenv("VCENTER_HOST"),
		EsxiUser:       os.Getenv("VCENTER_USER"),
		EsxiPassword:   os.Getenv("VCENTER_PASSWORD"),
		EsxiDatacenter: os.Getenv("VCENTER_DATACENTER"),
		EsxiDatastore:  os.Getenv("VCENTER_DATASTORE"),
	}

	configFile, err := ioutil.TempFile("", "config")
	Expect(err).ToNot(HaveOccurred())

	var configContent bytes.Buffer
	configTemplate.Execute(&configContent, configValues)

	configFile.WriteString(configContent.String())
	configPath, err := filepath.Abs(configFile.Name())
	Expect(err).ToNot(HaveOccurred())

	return configPath, vmStoreTempDir
}

var _ = BeforeSuite(func() {
	ExtractedStemcellTempDir = extractStemcell()
	CpiConfigPath, vmStoreDir = generateCPIConfig()
})

var _ = AfterSuite(func() {
	os.RemoveAll(ExtractedStemcellTempDir)
	os.RemoveAll(CpiConfigPath)
	//os.RemoveAll(vmStoreDir)
	gexec.CleanupBuildArtifacts()
})