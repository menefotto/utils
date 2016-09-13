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
	"io"
	"net"
	"net/http"
	"os"
	"path"
	"sync"
	"time"

	"github.com/sonic/lib/errors"
	"github.com/sonic/lib/utils/miscutils"
)

var Silent bool = false

func DownloadMulti(baseurl, saveto string, pkgs []string) chan error {
	errchan := make(chan error)
	var wg sync.WaitGroup

	for _, pkg := range pkgs {
		wg.Add(1)
		go func(baseurl, pkgname string, wg sync.WaitGroup) {
			err := DownloadSingle(baseurl, saveto, pkgname)
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

func DownloadSingle(baseurl, saveto, pkgname string) error {
	client := clientInit()

	resp, err := client.Get(baseurl + pkgname)
	defer resp.Body.Close()
	if err != nil {
		return errors.Wrap(err)()
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf(resp.Status)
	}

	f, err := os.Create(path.Join(saveto, pkgname))
	if err != nil {
		return errors.Wrap(err)()
	}
	defer f.Close()

	return copy(resp.Body, f, resp.ContentLength, pkgname)
}

func copy(src io.Reader, dst io.Writer, srcsize int64, pkgname string) error {
	var (
		bufferSize int64 = 4096
		total      int64 = 0
	)

	buffer := make([]byte, bufferSize)
	body := io.LimitReader(src, srcsize)

	percent := srcsize / 100

	for {
		nreads, err := body.Read(buffer)
		if nreads > 0 {
			if err != nil && err != io.EOF {
				return errors.Wrap(err)()
			}

			total += int64(nreads)

			_, err = dst.Write(buffer[:nreads])
			if err != nil && err != io.EOF {
				return errors.Wrap(err)()
			}

			if !Silent {
				message := miscutils.ProgressMsgBuild("Downloading " + pkgname)
				miscutils.ProgressPrinter(message, total, percent)
			}

			if total == srcsize {
				return nil
			}
		}

	}
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
