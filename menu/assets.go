// Code generated by compile.go. DO NOT EDIT.
//go:generate go run ./assets/compile.go

package menu

import (
	"errors"
	"text/template"
)

var Assets = map[string]string{
	"bird_core.conf": "log stderr all;\nprotocol device {\n    scan time 60;\n}\nprotocol bfd {\n    interface \"*\" {\n       min rx interval 400 ms;\n       min tx interval 400 ms;\n    };\n}\nprotocol static defaultgw {\n    ipv4;\n    route 0.0.0.0/0 via {{.Network.Endpoints.Host.IP}};\n}\nprotocol kernel {\n    merge paths;\n    ipv4 {\n        export all;\n    };\n}\ntemplate bgp bgpcore {\n    local as {{.Network.ASNCore}};\n    bfd;\n\n    ipv4 {\n        import all;\n        export all;\n        next hop self;\n    };\n}\n{{$asnSpine := .Network.ASNSpine -}}\n{{range $spineIdx, $spine :=  .Spines -}}\nprotocol bgp '{{$spine.Name}}' from bgpcore {\n    neighbor {{$spine.CoreAddress.IP}} as {{$asnSpine}};\n}\n{{end -}}\n",
	"bird_rack-tor1.conf": "{{$rackIdx := .RackIdx -}}\n{{$self := index .Args.Racks $rackIdx -}}\nlog stderr all;\nprotocol device {\n    scan time 60;\n}\nprotocol direct direct1 {\n    ipv4;\n    interface \"{{$self.ToR1.NodeInterface}}\";\n}\nprotocol bfd {\n    interface \"*\" {\n       min rx interval 400 ms;\n       min tx interval 400 ms;\n    };\n}\nprotocol kernel {\n    merge paths;\n    ipv4 {\n        export filter {\n            if source = RTS_DEVICE then reject;\n            accept;\n        };\n    };\n}\n{{$asnSpine := .Args.Network.ASNSpine -}}\n{{range $spine := .Args.Spines -}}\nprotocol bgp '{{$spine.Name}}' {\n    local as {{$self.ASN}};\n    neighbor {{($spine.ToR1Address $rackIdx).IP}} as {{$asnSpine}};\n    bfd;\n\n    ipv4 {\n        import all;\n        export all;\n    };\n}\n{{end -}}\ntemplate bgp bgpnode {\n    local as {{$self.ASN}};\n    direct;\n    rr client;\n    bfd;\n    passive;\n    error wait time 5,20;\n\n    ipv4 {\n        import all;\n        export filter {\n                if proto = \"direct1\" then reject;\n                accept;\n        };\n    };\n}\nprotocol bgp 'boot-{{$rackIdx}}' from bgpnode {\n    neighbor {{$self.BootNode.Node1Address.IP}} as {{$self.ASN}};\n}\n{{range $cs := $self.CSList -}}\nprotocol bgp '{{$self.Name}}-{{$cs.Name}}' from bgpnode {\n    neighbor {{$cs.Node1Address.IP}} as {{$self.ASN}};\n}\n{{end -}}\n{{range $ss := $self.SSList -}}\nprotocol bgp '{{$self.Name}}-{{$ss.Name}}' from bgpnode {\n    neighbor {{$ss.Node1Address.IP}} as {{$self.ASN}};\n}\n{{end -}}\n",
	"bird_rack-tor2.conf": "{{$rackIdx := .RackIdx -}}\n{{$self := index .Args.Racks $rackIdx -}}\nlog stderr all;\nprotocol device {\n    scan time 60;\n}\nprotocol direct direct1 {\n    ipv4;\n    interface \"{{$self.ToR2.NodeInterface}}\";\n}\nprotocol bfd {\n    interface \"*\" {\n       min rx interval 400 ms;\n       min tx interval 400 ms;\n    };\n}\nprotocol kernel {\n    merge paths;\n    ipv4 {\n        export filter {\n            if source = RTS_DEVICE then reject;\n            accept;\n        };\n    };\n}\n{{$asnSpine := .Args.Network.ASNSpine -}}\n{{range $spine := .Args.Spines -}}\nprotocol bgp '{{$spine.Name}}' {\n    local as {{$self.ASN}};\n    neighbor {{($spine.ToR2Address $rackIdx).IP}} as {{$asnSpine}};\n    bfd;\n\n    ipv4 {\n        import all;\n        export all;\n    };\n}\n{{end -}}\ntemplate bgp bgpnode {\n    local as {{$self.ASN}};\n    direct;\n    rr client;\n    bfd;\n    passive;\n    error wait time 5,20;\n\n    ipv4 {\n        import all;\n        export filter {\n                if proto = \"direct1\" then reject;\n                accept;\n        };\n    };\n}\nprotocol bgp 'boot-{{$rackIdx}}' from bgpnode {\n    neighbor {{$self.BootNode.Node2Address.IP}} as {{$self.ASN}};\n}\n{{range $cs := $self.CSList -}}\nprotocol bgp '{{$self.Name}}-{{$cs.Name}}' from bgpnode {\n    neighbor {{$cs.Node2Address.IP}} as {{$self.ASN}};\n}\n{{end -}}\n{{range $ss := $self.SSList -}}\nprotocol bgp '{{$self.Name}}-{{$ss.Name}}' from bgpnode {\n    neighbor {{$ss.Node2Address.IP}} as {{$self.ASN}};\n}\n{{end -}}\n",
	"bird_spine.conf": "{{$spineIdx := .SpineIdx -}}\n{{$self := index .Args.Spines $spineIdx -}}\nlog stderr all;\nprotocol device {\n    scan time 60;\n}\nprotocol bfd {\n    interface \"*\" {\n       min rx interval 400 ms;\n       min tx interval 400 ms;\n    };\n}\nprotocol kernel {\n    merge paths;\n    ipv4 {\n        export all;\n    };\n}\ntemplate bgp bgptor {\n    local as {{.Args.Network.ASNSpine}};\n    bfd;\n\n    ipv4 {\n        import all;\n        export all;\n        next hop self;\n    };\n}\n{{range $rack := .Args.Racks -}}\nprotocol bgp '{{$rack.Name}}-tor1' from bgptor {\n    neighbor {{(index $rack.ToR1.SpineAddresses $spineIdx).IP}} as {{$rack.ASN}};\n}\nprotocol bgp '{{$rack.Name}}-tor2' from bgptor {\n    neighbor {{(index $rack.ToR2.SpineAddresses $spineIdx).IP}} as {{$rack.ASN}};\n}\n{{end -}}\nipv4 table outertab;\nprotocol static myroutes {\n    ipv4 {\n        table outertab;\n    };\n    # LoadBalancer\n    route {{.Args.Network.Exposed.LoadBalancer}} via {{(index .Args.Core.SpineAddresses $spineIdx).IP}};\n    # Bastion\n    route {{.Args.Network.Exposed.Bastion}} via {{(index .Args.Core.SpineAddresses $spineIdx).IP}};\n    # Ingress\n    route {{.Args.Network.Exposed.Ingress}} via {{(index .Args.Core.SpineAddresses $spineIdx).IP}};\n    # Global\n    route {{.Args.Network.Exposed.Global}} via {{(index .Args.Core.SpineAddresses $spineIdx).IP}};\n}\n\nprotocol bgp 'core' {\n    local as {{.Args.Network.ASNSpine}};\n    neighbor {{(index .Args.Core.SpineAddresses $spineIdx).IP}} as {{.Args.Network.ASNCore}};\n    bfd;\n\n    ipv4 {\n        table outertab;\n        import all;\n        export all;\n        next hop self;\n    };\n}\n\nprotocol pipe outerroutes {\n    table master4;\n    peer table outertab;\n    import filter {\n        if proto = \"myroutes\" then reject;\n        accept;\n    };\n    export none;\n}\n",
	"chrony.conf": "# ntp servers\nserver 216.239.35.12 iburst\nserver 216.239.35.4 iburst\n\n# Allow chronyd to make a rapid measurement of the system clock error at boot time\ninitstepslew 0.1 216.239.35.12 216.239.35.4\n\n# Record the rate at which the system clock gains/losses time.\ndriftfile /var/lib/chrony/drift\n\n# Allow the system clock to be stepped in the first three updates\n# if its offset is larger than 1 second.\nmakestep 1.0 3\n\n# Enable kernel synchronization of the real-time clock (RTC).\nrtcsync\n\n# Allow NTP client access from local network.\nallow 10.0.0.0/8\n\n# Ignore leap second; ajdust by slewing\nleapsecmode slew\nmaxslewrate 1000\nsmoothtime 400 0.001 leaponly\n\n# mlockall\nlock_all\n\n# set highest scheduling priority\nsched_priority 99\n",
	"setup-default-gateway": "#!/bin/sh\n\nip route add default via {{.IP}}\n",
	"setup-iptables": "#!/bin/sh\n\n# eth0 -> internet\n# eth1 -> BMC\niptables -t nat -A POSTROUTING -o eth0 -s {{.}} -j MASQUERADE\niptables -t nat -A POSTROUTING -o eth1 -j MASQUERADE\n",
	"squid.conf": "# Only allow cachemgr access from localhost\nhttp_access allow manager localhost\nhttp_access deny manager\nhttp_access allow all\nhttp_port 3128\naccess_log stdio:/var/log/squid/access.log\npid_filename \"none\"\ncache_dir aufs /var/spool/squid 50000 16 256\ncoredump_dir /var/spool/squid\ncache_mem 200 MB\nmaximum_object_size_in_memory 100 MB\nmaximum_object_size 100 MB\ndetect_broken_pconn on\nforwarded_for delete\nhttpd_suppress_version_string on\n",
}

func GetAssetTemplate(name string) (*template.Template, error) {
	data, ok := Assets[name]
	if !ok {
		return nil, errors.New("no such asset: " + name)
	}
	return template.New(name).Parse(data)
}
