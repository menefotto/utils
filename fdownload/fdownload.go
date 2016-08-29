// downlaoder package provides a single function Download
// which perform multiple downloads given a base url and
// multiple resources to download from that base url,
// downloads are done cuncurently, for each resource
// to download a new http.client is created as well as a
// new goroutune and returns channel of errors if any
// otherwise nil is returned in case of success (non error)

package fdownload

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"sync"
	"time"
)

func DownloadMulti(baseurl string, pkgs []string) chan error {
	errchan := make(chan error)
	var wg sync.WaitGroup

	for _, pkg := range pkgs {
		wg.Add(1)
		go func(baseurl, pkgname string, wg sync.WaitGroup) {
			err := DownloadSingle(baseurl, pkgname)
			if err != nil {
				errchan <- fmt.Errorf("name: %v,%v\n", pkg, err)
			} else {
				errchan <- nil
			}
			wg.Done()
		}(baseurl, pkg, wg)
	}
	return errchan
}

func DownloadSingle(baseurl, pkgname string) error {
	client := clientInit()
	resp, err := client.Get(baseurl + pkgname)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf(resp.Status)
	}

	data, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(pkgname, data, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func clientInit() http.Client {
	transport := &http.Transport{
		Dial: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).Dial,
	}
	return http.Client{Transport: transport}

}
