package gdk

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
)

type Ltx struct {
	name  string `json:"name"`
	logid int64  `json:"logId"`
	count int    `json:"count"`
}

const LTX_DEFAULT_NAME = "ltx"

func NewLtx() *Ltx {
	return &Ltx{
		name:  LTX_DEFAULT_NAME,
		logid: time.Now().UnixNano(),
	}
}

func (p *Ltx) Set(name string) *Ltx {
	p.name = name
	return p
}
func (p *Ltx) step() int {
	p.count++
	return p.count
}

func (p *Ltx) format(format string) string {
	return fmt.Sprintf("%s log%d %d %s", p.name, p.logid, p.step(), format)
}

func (p *Ltx) Warnf(format string, args ...interface{}) {
	log.Warnf(p.format(format), args...)
}

func (p *Ltx) Errorf(format string, args ...interface{}) {
	log.Errorf(p.format(format), args...)
}
