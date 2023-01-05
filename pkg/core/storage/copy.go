package storage

import (
	"github.com/vmmgr/node/pkg/core/request"
	"github.com/vmmgr/node/pkg/core/tool"
	"io"
	"os"
	"time"
)

type File struct {
	uuid string
}

type Progress struct {
	total int64
	size  int64
}

func (p *Progress) Write(data []byte) (int, error) {
	n := len(data)
	p.size += int64(n)

	return n, nil
}

func (b *Base) FileCopy(srcFile, dstFile string) error {
	tool.ExportLog("---Copy disk image\nsrc/dest = " + srcFile + "_" + dstFile)
	src, err := os.Open(srcFile)
	if err != nil {
		return tool.ExportLog("Error: open error")
	}
	defer src.Close()
	file, err := src.Stat()
	if err != nil {
		return tool.ExportLog("Error: file gateway error")
	}

	dst, err := os.Create(dstFile)
	if err != nil {
		return tool.ExportLog("Error: file create")
	}
	defer dst.Close()

	p := Progress{total: file.Size()}

	go func() {
		for {
			if p.size != p.total {
				<-time.NewTimer(200 * time.Microsecond).C
				request.SendServer(b.Controller, b.UUID, uint((float64(p.size)/float64(p.total))*100), "", err)
			} else {
				request.SendServer(b.Controller, b.UUID, uint((float64(p.size)/float64(p.total))*100), "success", err)
				tool.ExportLog("FileCopy end")
				return
			}
		}
	}()

	_, err = io.Copy(dst, io.TeeReader(src, &p))
	if err != nil {
		return tool.ExportLog("Error: file node error")
	}

	return nil
}
