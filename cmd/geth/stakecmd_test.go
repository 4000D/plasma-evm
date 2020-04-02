package main

import (
	"flag"
	"testing"

	"github.com/Onther-Tech/plasma-evm/log"
	"github.com/Onther-Tech/plasma-evm/params"
	"github.com/mattn/go-colorable"
)

func init() {
	loglevel := flag.Int("loglevel", 4, "verbosity of logs")

	log.PrintOrigins(true)
	log.Root().SetHandler(log.LvlFilterHandler(log.Lvl(*loglevel), log.StreamHandler(colorable.NewColorableStderr(), log.TerminalFormat(true))))
}

func TestParseInt(t *testing.T) {
	testParseInt("Test 1", "0.01")
	testParseInt("Test 1.2", "0.5")
	testParseInt("Test 2", ".01")
	testParseInt("Test 3", "1.")
	testParseInt("Test 4", "1")
}

func testParseInt(caption, str string) {
	b := parseFloatString(str, params.WTONDecimals)
	f := params.ToRayFloat64(b)
	log.Info(caption, "str", str, "big", b, "float", f)
}
