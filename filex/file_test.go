package filex

import (
	"fmt"
	"os/user"
	"path"
	"testing"
	"time"
)

// dummy test
func TestSelfPath(t *testing.T) {
	SelfPath()
}

func TestRealPath(t *testing.T) {
	fp, err := RealPath("/tmp")
	if err != nil {
		t.Error("error, RealPath", err)
	}
	if fp != "/tmp" {
		t.Error("error, RealPath expected:/tmp, got", fp)
	}
	//fp, err = RealPath("./")
	_, err = RealPath("./")
	if err != nil {
		t.Error("error, RealPath", err)
	}
	//if !strings.HasSuffix(fp, "github.com/researchlab/gdk/file") {
	//	t.Error("error, RealPath expected suffix:github.com/researchlab/gdk/file, got", fp)
	//}
}

func TestSelfDir(t *testing.T) {
	SelfDir()
}

func TestName(t *testing.T) {
	fp := "/tmp/test.txt"
	name := Name(fp)
	if name != "test.txt" {
		t.Error("error, Name expected:test.txt, got:", name)
	}
}

func TestDir(t *testing.T) {
	fp := "/tmp/test.txt"
	dir := Dir(fp)
	if dir != "/tmp" {
		t.Error("error, Dir expected:/tmp, got:", dir)
	}
}

func TestInsureDir(t *testing.T) {
	if err := InsureDir("/tmp"); err != nil {
		t.Error("error, InsureDir", err)
	}
	if err := InsureDir(path.Join("/tmp", fmt.Sprintf("%v", time.Now().UnixNano()))); err != nil {
		t.Error("error, InsureDir", err)
	}
}

func TestEnsureDirRW(t *testing.T) {
	user, err := user.Current()
	if err != nil {
		t.Error("error, CurrentUser")
	}
	if user.Name == "root" {
		return
	}
	tmpDir := "/tmp/test_ensure_dir/abc"
	if err := EnsureDirRW(tmpDir); err != nil {
		t.Error("error, EnsureDirRW", err)
	}
}

func TestCreate(t *testing.T) {
	user, err := user.Current()
	if err != nil {
		t.Error("error, CurrentUser")
	}
	if user.Name == "root" {
		return
	}
	tmpFile := "/tmp/create.txt"
	f, err := Create(tmpFile)
	if err != nil {
		t.Error("error, Create", err)
		return
	}
	f.Close()
}

func TestRemove(t *testing.T) {
	user, err := user.Current()
	if err != nil {
		t.Error("error, Current")
	}
	if user.Name == "root" {
		return
	}
	tmpFile := "/tmp/create.txt"
	f, err := Create(tmpFile)
	if err != nil {
		t.Error("error, Create", err)
	}
	f.Close()

	if err := Remove(tmpFile); err != nil {
		t.Error("error, Remove", err)
	}
}

func TestExt(t *testing.T) {
	fp := "/tmp/test.txt"
	if v := Ext(fp); v != ".txt" {
		t.Error("error, Ext expected:txt, got", v)
	}
}

func TestRename(t *testing.T) {
	user, err := user.Current()
	if err != nil {
		t.Error("error, Current")
	}
	if user.Name == "root" {
		return
	}
	tmpFile := "/tmp/create.txt"
	f, err := Create(tmpFile)
	if err != nil {
		t.Error("error, Create", err)
	}
	f.Close()
	if err := Rename(tmpFile, "/tmp/rename.txt"); err != nil {
		t.Error("error, Rename", err)
	}
}

func TestIsFile(t *testing.T) {
	fp := fmt.Sprintf("/tmp/%v.txt", time.Now().UnixNano())
	if IsFile(fp) {
		t.Error("error, IsFile")
	}
	user, err := user.Current()
	if err != nil {
		t.Error("error, Current")
	}
	if user.Name == "root" {
		return
	}
	f, err := Create(fp)
	if err != nil {
		t.Error("error, Create", err)
	}
	f.Close()
	if !IsFile(fp) {
		t.Error("error, IsFile")
	}
}

func TestIsExist(t *testing.T) {
	if !IsExist("/tmp") {
		t.Error("error, IsExist")
	}
}

func TestSearchFile(t *testing.T) {
	user, err := user.Current()
	if err != nil {
		t.Error("error, user.Current")
	}
	if user.Name == "root" {
		return
	}
	tmpFile := fmt.Sprintf("%v.txt", time.Now().UnixNano())
	fp, err := SearchFile(tmpFile, "/tmp")
	if err == nil {
		t.Error("error, SearchFile")
	}
	tmpFilePath := path.Join("/tmp", tmpFile)
	Create(tmpFilePath)
	fp, err = SearchFile(tmpFile, "/tmp")
	if err != nil {
		t.Error("error, SearchFile", err)
	}
	if fp != tmpFilePath {
		t.Error("error, SearchFile, expected:", tmpFile, " got:", fp)
	}
}

func TestFileMTime(t *testing.T) {
	user, err := user.Current()
	if err != nil {
		t.Error("error, Current")
	}
	if user.Name == "root" {
		return
	}

	fp := fmt.Sprintf("/tmp/%v.txt", time.Now().UnixNano())
	_, err = FileMTime(fp)
	if err == nil {
		t.Error("error, FileMTime")
	}
	f, err := Create(fp)
	if err != nil {
		t.Error("error, Create", err)
	}
	f.Close()
	_, err = FileMTime(fp)
	if err != nil {
		t.Error("error, FileMTime", err)
	}
}

func TestFileSize(t *testing.T) {
	user, err := user.Current()
	if err != nil {
		t.Error("error, Current")
	}
	if user.Name == "root" {
		return
	}

	fp := fmt.Sprintf("/tmp/%v.txt", time.Now().UnixNano())
	_, err = FileSize(fp)
	if err == nil {
		t.Error("error, FileSize")
	}
	f, err := Create(fp)
	if err != nil {
		t.Error("error, Create", err)
	}
	f.Close()
	_, err = FileSize(fp)
	if err != nil {
		t.Error("error, FileSize", err)
	}
}

func TestDirsUnder(t *testing.T) {
	_, err := DirsUnder(fmt.Sprintf("/tmp/%v", time.Now().UnixNano()))
	if err != nil {
		t.Error("error, DirsUnder", err)
	}

	ret, err := DirsUnder("/tmp")
	if err != nil {
		t.Error("error, DirsUnder", err)
	}
	for _, v := range ret {
		if !IsExist(path.Join("/tmp", v)) {
			t.Error("error, DirsUnder ", v, " not exist")
		}
	}
}

func TestFilesUnder(t *testing.T) {
	_, err := FilesUnder(fmt.Sprintf("/tmp/%v", time.Now().UnixNano()))
	if err != nil {
		t.Error("error, FilesUnder", err)
	}

	ret, err := FilesUnder("/tmp")
	if err != nil {
		t.Error("error, FilesUnder", err)
	}
	for _, v := range ret {
		if !IsExist(path.Join("/tmp", v)) {
			t.Error("error, FilesUnder ", v, " not exist")
		}
	}
}

func TestSearchFileWithAffix(t *testing.T) {
	user, err := user.Current()
	if err != nil {
		t.Error("error, Current")
	}
	if user.Name == "root" {
		return
	}

	_, exist := SearchFileWithAffix("/tmxxx", "", "")
	if exist {
		t.Error("error, SearchFileWithAffix")
	}
	fp := fmt.Sprintf("/tmp/test-%v.txt", time.Now().UnixNano())
	_, err = FileSize(fp)
	if err == nil {
		t.Error("error, FileSize")
	}
	f, err := Create(fp)
	if err != nil {
		t.Error("error, Create", err)
	}
	f.Close()
	prefix := Name(fp)
	v, exist := SearchFileWithAffix("/tmp", prefix, ".txt")
	if !exist {
		t.Error("error, SearchFileWithAffix")
	}
	if path.Join("/tmp", v) != fp {
		t.Error("error, SearchFileWithAffix, expected:", fp, " got:/tmp/", v)
	}

	_, exist = SearchFileWithAffix("/tmp", "", "")
	if exist {
		t.Error("error, SearchFileWithAffix")
	}
	_, exist = SearchFileWithAffix("/tmp", "/tmp", ".txt")

	if exist {
		t.Error("error, SearchFileWithAffix")
	}
	_path := fmt.Sprintf("/tmp/%v", time.Now().UnixNano())
	Create(_path)
	_, exist = SearchFileWithAffix(_path, "", "")
	if exist {
		t.Error("error, SearchFileWithAffix")
	}
}
