package dnsmasq

import _ "embed"

//go:embed scripts/dnsmasq.conf.gotmpl
var configFile string

//go:embed scripts/dnsmasq.service.gotmpl
var unitFile string

//go:embed scripts/aro-dnsmasq-pre.sh.gotmpl
var preScriptFile string
