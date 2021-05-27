package zip

import (
	"archive/zip"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

// archive/zip
// zip相关

// 压缩单个文件
func TestZipFile(t *testing.T) {
	// 准备文件
	srcFileName := "tmp.a.txt"
	err := os.WriteFile(srcFileName, []byte("hello world."), os.ModePerm)
	assert.Nil(t, err)

	// 创建准备写入的zip文件
	zipFileName := "tmp.a.zip"
	zipFile, err := os.Create(zipFileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer zipFile.Close()

	// 通过io.Writer创建zip.Writer
	zipWriter := zip.NewWriter(zipFile)
	defer func() {
		if err := zipWriter.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	// 获取写入源文件信息
	fileInfo, err := os.Stat(srcFileName)
	if err != nil {
		log.Fatalln(err)
	}

	// 根据源文件信息创建压缩文件头
	fileHeader, err := zip.FileInfoHeader(fileInfo)
	if err != nil {
		log.Fatalln(err)
	}

	// 设置压缩文件名（对于多层目录结构这一步需要设置）
	// fileHeader.Name = srcFileName

	// 通过压缩文件信息创建真正的执行写入的io.Writer
	w, err := zipWriter.CreateHeader(fileHeader)
	assert.Nil(t, err)

	// 打开要写入的文件
	sf, err := os.Open(srcFileName)
	assert.Nil(t, err)
	defer sf.Close()

	// 拷贝写入数据
	n, err := io.Copy(w, sf)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("写入成功字节数:", n) // 写入成功字节数: 12
}

// 压缩目录
func TestZipDir(t *testing.T) {
	// 准备文件
	// tmp
	// ├── a
	// ├── b
	// └── c
	//    └── d
	_ = os.Mkdir("tmp", 0755)
	_ = os.WriteFile("tmp/a", []byte("aaaaa"), 0755)
	_ = os.WriteFile("tmp/b", []byte("bbbbb"), 0755)
	_ = os.Mkdir("tmp/c", 0755)
	_ = os.WriteFile("tmp/c/d", []byte("ddddd"), 0755)

	zipFile, err := os.Create("tmp.zip")
	if err != nil {
		log.Fatalln(err)
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer func() {
		if err := zipWriter.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	// 注意"tmp"路径应该是相对路径,不是绝对路径
	err = filepath.Walk("tmp", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		fileHeader, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		fileHeader.Name = path
		if info.IsDir() {
			fileHeader.Name += "/"
		}

		w, err := zipWriter.CreateHeader(fileHeader)
		if err != nil {
			return err
		}

		if !info.Mode().IsRegular() {
			return nil
		}

		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()

		n, err := io.Copy(w, f)
		if err != nil {
			return err
		}

		fmt.Printf("成功压缩文件:%s, 写入字节数:%d\n", path, n)
		return nil
	})
	if err != nil {
		log.Fatalln(err)
	}
	// output:
	// 成功压缩文件:tmp/a, 写入字节数:5
	// 成功压缩文件:tmp/b, 写入字节数:5
	// 成功压缩文件:tmp/c/d, 写入字节数:5
}
