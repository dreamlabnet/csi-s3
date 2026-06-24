package driver_test

import (
	"log"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/dreamlabnet/csi-s3/pkg/driver"

	"github.com/kubernetes-csi/csi-test/pkg/sanity"
)

// Ensure IDGen is initialized for all test configs
func initSanityConfig(cfg *sanity.Config) *sanity.Config {
	if cfg.IDGen == nil {
		cfg.IDGen = &sanity.DefaultIDGenerator{}
	}
	return cfg
}

var _ = Describe("S3Driver", func() {

	Context("s3fs", func() {
		socket := "/tmp/csi-s3fs.sock"
		csiEndpoint := "unix://" + socket
		if err := os.Remove(socket); err != nil && !os.IsNotExist(err) {
			Expect(err).NotTo(HaveOccurred())
		}
		driver, err := driver.New("test-node", csiEndpoint)
		if err != nil {
			log.Fatal(err)
		}
		go driver.Run()

		Describe("CSI sanity", func() {
			sanityCfg := initSanityConfig(&sanity.Config{
				TargetPath:  os.TempDir() + "/s3fs-target",
				StagingPath: os.TempDir() + "/s3fs-staging",
				Address:     csiEndpoint,
				SecretsFile: "../../test/secret.yaml",
				TestVolumeParameters: map[string]string{
					"mounter": "s3fs",
					"bucket":  "testbucket0",
				},
			})
			sanity.GinkgoTest(sanityCfg)
		})
	})

	Context("s3fs-no-bucket", func() {
		socket := "/tmp/csi-s3fs-no-bucket.sock"
		csiEndpoint := "unix://" + socket
		if err := os.Remove(socket); err != nil && !os.IsNotExist(err) {
			Expect(err).NotTo(HaveOccurred())
		}
		driver, err := driver.New("test-node", csiEndpoint)
		if err != nil {
			log.Fatal(err)
		}
		go driver.Run()

		Describe("CSI sanity", func() {
			sanityCfg := initSanityConfig(&sanity.Config{
				TargetPath:  os.TempDir() + "/s3fs-no-bucket-target",
				StagingPath: os.TempDir() + "/s3fs-no-bucket-staging",
				Address:     csiEndpoint,
				SecretsFile: "../../test/secret.yaml",
				TestVolumeParameters: map[string]string{
					"mounter": "s3fs",
				},
			})
			sanity.GinkgoTest(sanityCfg)
		})
	})

	/*
		Context("rclone", func() {
			socket := "/tmp/csi-rclone.sock"
			csiEndpoint := "unix://" + socket

			if err := os.Remove(socket); err != nil && !os.IsNotExist(err) {
				Expect(err).NotTo(HaveOccurred())
			}
			driver, err := driver.New("test-node", csiEndpoint)
			if err != nil {
				log.Fatal(err)
			}
			go driver.Run()

			Describe("CSI sanity", func() {
				sanityCfg := &sanity.Config{
					TargetPath:  os.TempDir() + "/rclone-target",
					StagingPath: os.TempDir() + "/rclone-staging",
					Address:     csiEndpoint,
					SecretsFile: "../../test/secret.yaml",
					TestVolumeParameters: map[string]string{
						"mounter": "rclone",
						"bucket":  "testbucket3",
					},
				}
				sanity.GinkgoTest(sanityCfg)
			})
		})
	*/
})
