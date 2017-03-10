package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"stringutil"
)

type FieldMismatch struct {
	expected, found int
}

type fileContent struct {
	col1 string
	col2 string
}

type folderContentRecord struct {
	subFolderContext  string
	subFolderFullPath string
}

type foldercontents []folderContentRecord

type fileContents []fileContent

func (e *FieldMismatch) Error() string {
	return "CSV line fields mismatch. Expected " + strconv.Itoa(e.expected) + " found " + strconv.Itoa(e.found)
}

type UnsupportedType struct {
	Type string
}

func (e *UnsupportedType) Error() string {
	return "Unsupported type: " + e.Type
}

type fileMetadataRecord struct {
	Filename     string
	FullFileName string
	Datetime     string
	Size         string
}

type fileMetadata []fileMetadataRecord

func getParentDirectory(dir string, file2 string) string {
	var parentDir string

	//fmt.Printf("input: %q\n\tdir: %q\n\tfile: %q\n", file, dir, file2)
	for _, t := range strings.Split(dir, "\\") {
		if t != "" {
			parentDir = t
			//fmt.Println(parentDir)
		}
	}
	//fmt.Printf("dir: %q\n\tfile: %q\n", parentDir, file2)
	return parentDir
}

func getFileList(searchDir string) []string {
	fileList := []string{}
	err := filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		fileList = append(fileList, path)
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
	return fileList
}

func readLogFile(file string) (fileMetadata, map[uint64]string) {
	var m map[uint64]string
	m = make(map[uint64]string)
	var myint uint64

	//var inputFileName []string
	var vFileMetadataRecord fileMetadataRecord
	var vfileMetadata fileMetadata

	fmt.Println(file)
	logfile, err := os.Open(file)
	if err != nil {
		fmt.Println("Error", err)
		return vfileMetadata, m
	}
	defer logfile.Close()

	reader := csv.NewReader(logfile)

	for {
		err := Unmarshal(reader, &vFileMetadataRecord)
		if err == io.EOF {
			fmt.Println("Error unmashelling")
			break
		}
		if err != nil {
			fmt.Println("Error unmashelling")
			panic(err)

		}

		myint = hashFolderContents(vFileMetadataRecord)

		m[myint] = vFileMetadataRecord.Filename
		vfileMetadata = append(vfileMetadata, vFileMetadataRecord)
	}

	return vfileMetadata, m

}

func hashFolderContents(vFileMetadataRecord fileMetadataRecord) uint64 {
	var myint uint64
	var stringToHash string
	stringToHash = vFileMetadataRecord.Filename
	stringToHash = vFileMetadataRecord.FullFileName
	stringToHash += vFileMetadataRecord.Datetime
	stringToHash += vFileMetadataRecord.Size

	myint = stringutil.HashThisString(stringToHash)
	return myint
}

func Unmarshal(reader *csv.Reader, v interface{}) error {
	record, err := reader.Read()
	reader.Comma = ','
	if err != nil {
		return err
	}
	s := reflect.ValueOf(v).Elem()
	if s.NumField() != len(record) {
		return &FieldMismatch{s.NumField(), len(record)}
	}
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		switch f.Type().String() {
		case "string":
			f.SetString(record[i])
		case "int":
			ival, err := strconv.ParseInt(record[i], 10, 0)
			if err != nil {
				return err
			}
			f.SetInt(ival)
		default:
			return &UnsupportedType{f.Type().String()}
		}
	}
	return nil
}

func writeFileLog(vLogFile string, data [][]string) {
	file, err := os.OpenFile(vLogFile, os.O_APPEND, 0666)
	checkError("Cannot create file", err)
	defer file.Close()
	writer := csv.NewWriter(file)
	for _, record := range data {
		err := writer.Write(record)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	}
	defer writer.Flush()
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

func getTopDirectory(folderpath string) foldercontents {
	var fullpath string
	files, _ := ioutil.ReadDir(folderpath)
	var vfoldercontents foldercontents
	for _, f := range files {
		fullpath = folderpath
		var vfolderContentRecord folderContentRecord
		if f.IsDir() {
			vfolderContentRecord.subFolderContext = f.Name()
			fullpath += "/"
			fullpath += vfolderContentRecord.subFolderContext
			fmt.Println("Printing folder: ", fullpath)
			vfolderContentRecord.subFolderFullPath = fullpath
			vfoldercontents = append(vfoldercontents, vfolderContentRecord)
		}
	}
	return vfoldercontents
}

func main() {
	searchDir := "c:/temp/research"
	fmt.Println("*******************")
	vfoldercontents := getTopDirectory(searchDir)
	var folders folderContentRecord
	for _, folders = range vfoldercontents {
		fmt.Println("+++++++++++++++++++++++++++++++++++")
		traversContextFolders(folders.subFolderFullPath)
		fmt.Println("+++++++++++++++++++++++++++++++++++")
	}
	fmt.Println("*******************")
	//var mLogFile map[string]uint64

}

func traversContextFolders(searchDir string) {
	var contextDir string
	//	var vfileMetadata fileMetadata
	var xfileMetadata fileMetadata
	var xfileMetadataRecord fileMetadataRecord
	var vLogFile string
	var m map[uint64]string

	m = make(map[uint64]string)
	var myint uint64

	vLogFile = searchDir
	vLogFile += "/log/log.csv"

	m = consumingLogFile(vLogFile)

	for _, file := range getFileList(searchDir) {
		fmt.Println(file)
		dir, file2 := filepath.Split(file)
		fi, err := os.Stat(file)

		if err != nil {
			fmt.Println(err)
			return
		}

		switch mode := fi.Mode(); {
		case mode.IsDir():
			if file2 == "log" {
				contextDir = getParentDirectory(dir, file2)
				fmt.Println("print:", contextDir)
			}
		case mode.IsRegular():
			if file2 != "log.csv" { /*
					fmt.Println("reading log file")

				} else {*/
				fmt.Println("Director: ", dir)
				fmt.Println("File: ", file2)

				xfileMetadataRecord.Filename = file2
				xfileMetadataRecord.FullFileName = file
				xfileMetadataRecord.Size = strconv.FormatInt(fi.Size(), 10)
				xfileMetadataRecord.Datetime = fi.ModTime().String()

				myint = hashFolderContents(xfileMetadataRecord)
				_, present := m[myint]
				if present {
					fmt.Println("file exists")
				} else {
					fmt.Println(file2, "Does NOT exist")
					xfileMetadata = append(xfileMetadata, xfileMetadataRecord)
				}
			}

		} //end of case statement

	}

	var vfileContents [][]string
	for _, vrecord := range xfileMetadata {

		vfileContents = append(vfileContents, createFileContents(vrecord))

	}
	//		fmt.Println(vLogFile)
	if vLogFile != "" {
		writeFileLog(vLogFile, vfileContents)
	}

}

func createFileContents(vrecord fileMetadataRecord) []string {

	strFileContents := []string{vrecord.Filename, vrecord.FullFileName, vrecord.Datetime, vrecord.Size}
	return strFileContents
}

func consumingLogFile(file string) map[uint64]string {

	//			fmt.Println(file)
	vfileMetadata, m := readLogFile(file)

	fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
	for _, vrecord := range vfileMetadata {
		if vrecord.Filename != "" {
			fmt.Println("Name:", vrecord.Filename, "Time:", vrecord.Datetime, "Size: ", vrecord.Size)

		}
	}
	fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

	return m

}
