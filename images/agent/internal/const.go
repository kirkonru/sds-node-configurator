/*
Copyright 2023 Flant JSC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package internal

const (
	TypePart                  = "part"
	DRBDName                  = "/dev/drbd"
	LoopDeviceType            = "loop"
	LVMDeviceType             = "lvm"
	LVMFSType                 = "LVM2_member"
	SdsNodeConfigurator       = "storage.deckhouse.io/sds-node-configurator"
	LVMVGHealthOperational    = "Operational"
	LVMVGHealthNonOperational = "NonOperational"
)

var (
	AllowedFSTypes       = [...]string{LVMFSType}
	InvalidDeviceTypes   = [...]string{LoopDeviceType, LVMDeviceType}
	BlockDeviceValidSize = "1G"
	ResizeDelta          = "32Mi"
	Finalizers           = []string{SdsNodeConfigurator}
	LVMTags              = []string{"storage.deckhouse.io/enabled=true", "linstor-"}
)
