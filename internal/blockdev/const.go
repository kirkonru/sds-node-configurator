package blockdev

const (
	DRBDName                  = "drbd"
	DeckhouseEmptyBlockDevice = "DeckhouseEmptyBlockDevice"
	APIVersion                = "storage.deckhouse.io/v1"
	AppName                   = "linstor-pools-importer"
)

var (
	//lsblkCommand = []string{"lsblk", "--json", "-lpf "}
	echoCommand = []string{"echo", "{  \"blockdevices\": [   {     \"name\": \"/dev/loop0\",\n      \"mountpoint\": \"/snap/core20/1587\",\n      \"partuuid\": null,\n      \"hotplug\": false,\n      \"model\": null,\n      \"serial\": null,\n      \"size\": \"62M\",\n      \"type\": \"loop\",\n      \"wwn\": null,\n      \"kname\": \"/dev/loop0\",\n      \"pkname\": null\n    },{\n      \"name\": \"/dev/loop1\",\n      \"mountpoint\": \"/snap/lxd/22923\",\n      \"partuuid\": null,\n      \"hotplug\": false,\n      \"model\": null,\n      \"serial\": null,\n      \"size\": \"79.9M\",\n      \"type\": \"loop\",\n      \"wwn\": null,\n      \"kname\": \"/dev/loop1\",\n      \"pkname\": null\n    },{\n      \"name\": \"/dev/loop2\",\n      \"mountpoint\": \"/snap/core18/2785\",\n      \"partuuid\": null,\n      \"hotplug\": false,\n      \"model\": null,\n      \"serial\": null,\n      \"size\": \"55.7M\",\n      \"type\": \"loop\",\n      \"wwn\": null,\n      \"kname\": \"/dev/loop2\",\n      \"pkname\": null\n    },{\n      \"name\": \"/dev/loop3\",\n      \"mountpoint\": \"/snap/snapd/19457\",\n      \"partuuid\": null,\n      \"hotplug\": false,\n      \"model\": null,\n      \"serial\": null,\n      \"size\": \"53.3M\",\n      \"type\": \"loop\",\n      \"wwn\": null,\n      \"kname\": \"/dev/loop3\",\n      \"pkname\": null\n    },{\n      \"name\": \"/dev/loop4\",\n      \"mountpoint\": \"/snap/core20/1950\",\n      \"partuuid\": null,\n      \"hotplug\": false,\n      \"model\": null,\n      \"serial\": null,\n      \"size\": \"63.4M\",\n      \"type\": \"loop\",\n      \"wwn\": null,\n      \"kname\": \"/dev/loop4\",\n      \"pkname\": null\n    },{\n      \"name\": \"/dev/loop5\",\n      \"mountpoint\": \"/snap/lxd/24322\",\n      \"partuuid\": null,\n      \"hotplug\": false,\n      \"model\": null,\n      \"serial\": null,\n      \"size\": \"111.9M\",\n      \"type\": \"loop\",\n      \"wwn\": null,\n      \"kname\": \"/dev/loop5\",\n      \"pkname\": null\n    },{\n      \"name\": \"/dev/loop6\",\n      \"mountpoint\": \"/snap/helm/372\",\n      \"partuuid\": null,\n      \"hotplug\": false,\n      \"model\": null,\n      \"serial\": null,\n      \"size\": \"11.2M\",\n      \"type\": \"loop\",\n      \"wwn\": null,\n      \"kname\": \"/dev/loop6\",\n      \"pkname\": null\n    },{\n      \"name\": \"/dev/vda\",\n      \"mountpoint\": null,\n      \"partuuid\": null,\n      \"hotplug\": false,\n      \"model\": null,\n      \"serial\": \"e7ccf5f5-f1f1-450e-8\",\n      \"size\": \"25G\",\n      \"type\": \"disk\",\n      \"wwn\": null,\n      \"kname\": \"/dev/vda\",\n      \"pkname\": null\n    },{\n      \"name\": \"/dev/vda1\",\n      \"mountpoint\": \"/\",\n      \"partuuid\": \"222de6d2-728f-4405-956e-48920db7c4c4\",\n      \"hotplug\": false,\n      \"model\": null,\n      \"serial\": null,\n      \"size\": \"24.9G\",\n      \"type\": \"part\",\n      \"wwn\": null,\n      \"kname\": \"/dev/vda1\",\n      \"pkname\": \"/dev/vda\"\n    },{\n      \"name\": \"/dev/vda14\",\n      \"mountpoint\": null,\n      \"partuuid\": \"1b81f7d3-904e-4008-8292-8723d49a8418\",\n      \"hotplug\": false,\n      \"model\": null,\n      \"serial\": null,\n      \"size\": \"4M\",\n      \"type\": \"part\",\n      \"wwn\": null,\n      \"kname\": \"/dev/vda14\",\n      \"pkname\": \"/dev/vda\"\n    },{\n      \"name\": \"/dev/vda15\",\n      \"mountpoint\": \"/boot/efi\",\n      \"partuuid\": \"c20b055b-76dc-4718-9ce7-5f713e3612c7\",\n      \"hotplug\": false,\n      \"model\": null,\n      \"serial\": null,\n      \"size\": \"106M\",\n      \"type\": \"part\",\n      \"wwn\": null,\n      \"kname\": \"/dev/vda15\",\n      \"pkname\": \"/dev/vda\"\n    },{\n      \"name\": \"/dev/vdb\",\n      \"mountpoint\": \"/mnt/kubernetes-data\",\n      \"partuuid\": null,\n      \"hotplug\": false,\n      \"model\": null,\n      \"serial\": \"62083394-8003-4b6e-a\",\n      \"size\": \"12G\",\n      \"type\": \"disk\",\n      \"wwn\": null,\n      \"kname\": \"/dev/vdb\",\n      \"pkname\": null\n    }\n  ]\n}"}
)
