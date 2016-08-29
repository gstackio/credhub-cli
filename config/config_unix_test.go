// +build !windows

package config_test

import (
	"os"

	"path"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pivotal-cf/credhub-cli/config"
)

var _ = Describe("Config", func() {
	var cfg config.Config

	BeforeEach(func() {
		cfg = config.Config{
			ApiURL:  "http://api.example.com",
			AuthURL: "http://auth.example.com",
		}
	})

	It("set appropriate permissions for persisted files", func() {
		config.WriteConfig(cfg)

		parentStat, _ := os.Stat(path.Dir(config.ConfigPath()))
		Expect(parentStat.Mode().String()).To(Equal("drwxr-xr-x"))

		fileStat, _ := os.Stat(config.ConfigPath())
		Expect(fileStat.Mode().String()).To(Equal("-rw-------"))
	})
})
