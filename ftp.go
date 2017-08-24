package foldersync

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/dutchcoders/goftp"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type FTP struct {
	host     string
	port     int
	user     string
	password string
}

func (ftp *FTP) NewFTP(host string, port int, user string, password string) error {
	ftp.host = host
	ftp.port = port
	ftp.user = user
	ftp.password = password
	return nil
}

func (ftp *FTP) FolderSync(local string, remote string, flag string) (bool, error) {

	var err error
	var gftp *goftp.FTP

	//check
	local = filepath.ToSlash(local)
	remote = filepath.ToSlash(remote + "/")

	if ftp.needUpdate(local, remote, flag) {

		err = os.RemoveAll(local)
		if err != nil {
			fmt.Println(err.Error())
		}

		if rst, err := ftp.loginFTP(goftp); rst == false {
			return false, err
		}
		defer gftp.Close()

		// Download each file into local memory, and calculate it's sha256 hash
		err = gftp.Walk(remote, func(dir string, info os.FileMode, err error) error {
			_, err = gftp.Retr(dir, func(r io.Reader) error {

				fmt.Printf("File in remote: %q\n", dir)

				rel, err := filepath.Rel(remote, dir)
				if err != nil {
					fmt.Println(err.Error())
				}
				fmt.Printf("Relative path in remote: %q\n", rel)

				loc := filepath.Join(local, rel)
				fmt.Printf("Full path in Local: %q\n", rel)

				folder := filepath.Dir(loc)
				fmt.Printf("Folder of path: %q\n", folder)

				err = os.MkdirAll(folder, os.ModePerm)
				if err != nil {
					fmt.Println(err.Error())
				}

				file, err := os.Create(loc)
				if err != nil {
					fmt.Println(err.Error())
				}

				if _, err = io.Copy(file, r); err != nil {
					fmt.Println(err.Error())
					return err
				}
				if err != nil {
					fmt.Println(err.Error())
				}

				return err
			})

			return err
		})

		if err != nil {
			fmt.Println(err.Error())
		}
	}
	return true, nil
}

func (ftp *FTP) needUpdate(local string, remote string, flag string) (bool, error) {
	needUpdate := false

	if b, _ := pathExists(local); b == false {
		needUpdate = true
	} else {
		localFlag := filepath.Join(local, flag)
		remoteFlag := filepath.Join(remote, flag)

		fmt.Println(fmt.Sprintf("metadata.xml:local(%s),remote(%s)\n", localFlag, remoteFlag))

		lmd5, _ := ftp.localFileMd5(localFlag)
		rmd5, _ := ftp.remoteFileMd5(remoteFlag)

		fmt.Println(fmt.Sprintf("MD5 of metadata.xml:local(%s),remote(%s)\n", lmd5, rmd5))

		if !((strings.Compare(lmd5, rmd5) == 0) && (strings.Compare(lmd5, "") != 0)) {
			fmt.Println("metadata.xml in both local and remote is different")
			needUpdate = true
		}
	}

	return needUpdate, nil
}

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func (ftp *FTP) localFileMd5(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := md5.New()
	if _, err = io.Copy(hash, file); err != nil {
		return "", err
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}

func (ftp *FTP) remoteFileMd5(path string) (string, error) {

	var err error
	var gftp *goftp.FTP

	if rst, err := ftp.loginFTP(goftp); rst == false {
		return "", err
	}
	defer gftp.Close()

	var hash = ""
	_, err = gftp.Retr(path, func(r io.Reader) error {
		var md5er = md5.New()
		if _, err = io.Copy(md5er, r); err != nil {
			return err
		}
		hash = hex.EncodeToString(md5er.Sum(nil))
		return err
	})

	return hash, nil
}

func (ftp *FTP) loginFTP(gftp *goftp.FTP) (bool, error) {
	var err error

	url := fmt.Sprintf("%s:%d", ftp.host, ftp.port)
	if gftp, err = goftp.Connect(url); err != nil {
		return false, err
	}
	defer gftp.Close()

	// Username / password authentication
	if err = gftp.Login(ftp.user, ftp.password); err != nil {
		return false, err
	}

	return true, nil
}
