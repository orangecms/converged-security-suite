package test

import (
	"fmt"

	"github.com/9elements/converged-security-suite/v2/pkg/tools"
	"github.com/9elements/go-linux-lowlevel-hw/pkg/hwapi"
)

var (
	testtxtmemoryrangevalid = Test{
		Name:                    "TXT heap ranges valid",
		Required:                true,
		function:                TXTHeapSpaceValid,
		Status:                  Implemented,
		SpecificationChapter:    "B.1",
		SpecificiationTitle:     IntelTXTSpecificationTitle,
		SpecificationDocumentID: IntelTXTSpecificationDocumentID,
	}
	testtxtpublicisreserved = Test{
		Name:                    "TXT public area reserved in e820",
		Required:                true,
		function:                TXTPublicReservedInE820,
		dependencies:            []*Test{&testtxtregisterspaceaccessible},
		Status:                  Implemented,
		SpecificationChapter:    "5.5.3 Intel TXT Public Space",
		SpecificiationTitle:     IntelTXTBGSBIOSSpecificationTitle,
		SpecificationDocumentID: IntelTXTBGSBIOSSpecificationDocumentID,
	}
	testtxtprivateisreserved = Test{
		Name:                    "TXT private area reserved in e820",
		Required:                true,
		function:                TXTPrivateReservedInE820,
		dependencies:            []*Test{&testtxtregisterspaceaccessible},
		Status:                  Implemented,
		SpecificationChapter:    "5.5.2 Intel TXT Private Space",
		SpecificiationTitle:     IntelTXTBGSBIOSSpecificationTitle,
		SpecificationDocumentID: IntelTXTBGSBIOSSpecificationDocumentID,
	}
	testmemoryisreserved = Test{
		Name:                    "TXT memory reserved in e820",
		Required:                true,
		function:                TXTReservedInE820,
		dependencies:            []*Test{&testtxtmemoryrangevalid},
		Status:                  Implemented,
		SpecificationChapter:    "5.5.4 Intel TPM Decode Area",
		SpecificiationTitle:     IntelTXTBGSBIOSSpecificationTitle,
		SpecificationDocumentID: IntelTXTBGSBIOSSpecificationDocumentID,
	}
	testtpmdecodereserved = Test{
		Name:                    "MMIO TPMDecode space reserved in e820",
		Required:                true,
		function:                TXTTPMDecodeSpaceIn820,
		Status:                  Implemented,
		SpecificationChapter:    "5.5.4 TPM Decode Area",
		SpecificiationTitle:     IntelTXTBGSBIOSSpecificationTitle,
		SpecificationDocumentID: IntelTXTBGSBIOSSpecificationDocumentID,
	}
	testtxtmemoryisdpr = Test{
		Name:                    "TXT memory in a DMA protected range",
		Required:                true,
		function:                TXTMemoryIsDPR,
		dependencies:            []*Test{&testtxtmemoryrangevalid},
		Status:                  Implemented,
		SpecificationChapter:    "1.11.1 DMA Protected Range (DPR)",
		SpecificiationTitle:     IntelTXTSpecificationTitle,
		SpecificationDocumentID: IntelTXTSpecificationDocumentID,
	}
	testtxtdprislocked = Test{
		Name:                    "TXT DPR register locked",
		Required:                true,
		function:                TXTDPRisLock,
		Status:                  Implemented,
		SpecificationChapter:    "1.11.1 DMA Protected Range (DPR)",
		SpecificiationTitle:     IntelTXTSpecificationTitle,
		SpecificationDocumentID: IntelTXTSpecificationDocumentID,
	}
	// Internal precondition to check if the platform registers are known
	testSupportsHostbridge = Test{
		Name:     "Hostbridge is supported",
		Required: true,
		function: HostbridgeIsSupported,
		Status:   Implemented,
	}
	testhostbridgeDPRcorrect = Test{
		Name:                    "CPU DPR equals hostbridge DPR",
		Required:                false,
		function:                HostbridgeDPRCorrect,
		Status:                  Implemented,
		SpecificationChapter:    "B 1.15 TXT.DPR – DMA Protected Range",
		SpecificiationTitle:     IntelTXTSpecificationTitle,
		SpecificationDocumentID: IntelTXTSpecificationDocumentID,
		dependencies:            []*Test{&testhassmrr, &testSupportsHostbridge},
	}
	testhostbridgeDPRislocked = Test{
		Name:                    "CPU hostbridge DPR register locked",
		Required:                true,
		function:                HostbridgeDPRisLocked,
		dependencies:            []*Test{&testhostbridgeDPRcorrect, &testSupportsHostbridge},
		Status:                  Implemented,
		SpecificationChapter:    "B 1.15 TXT.DPR – DMA Protected Range",
		SpecificiationTitle:     IntelTXTSpecificationTitle,
		SpecificationDocumentID: IntelTXTSpecificationDocumentID,
	}
	testsinitintxt = Test{
		Name:                    "TXT region contains SINIT ACM",
		Required:                false,
		function:                SINITInTXT,
		Status:                  Implemented,
		SpecificationChapter:    "B 1.10 TXT.SINIT.BASE – SINIT Base Address",
		SpecificiationTitle:     IntelTXTSpecificationTitle,
		SpecificationDocumentID: IntelTXTSpecificationDocumentID,
	}
	testsinitmatcheschipset = Test{
		Name:                    "SINIT ACM matches chipset",
		Required:                true,
		function:                SINITMatchesChipset,
		dependencies:            []*Test{&testsinitintxt},
		Status:                  Implemented,
		SpecificationChapter:    "2.2.3.1 Matching an AC Module to the Platform",
		SpecificiationTitle:     IntelTXTSpecificationTitle,
		SpecificationDocumentID: IntelTXTSpecificationDocumentID,
	}
	testsinitmatchescpu = Test{
		Name:                    "SINIT ACM matches CPU",
		Required:                true,
		function:                SINITMatchesCPU,
		dependencies:            []*Test{&testsinitintxt},
		Status:                  Implemented,
		SpecificationChapter:    "2.2.3.1 Matching an AC Module to the Platform",
		SpecificiationTitle:     IntelTXTSpecificationTitle,
		SpecificationDocumentID: IntelTXTSpecificationDocumentID,
	}
	testnosiniterrors = Test{
		Name:     "SINIT ACM startup successful",
		Required: false,
		function: NoSINITErrors,
		Status:   Implemented,
	}
	testbiosdataregionpresent = Test{
		Name:                    "BIOS DATA REGION present",
		Required:                true,
		function:                BIOSDATAREGIONPresent,
		Status:                  Implemented,
		SpecificationChapter:    "C.2 BIOS Data Format",
		SpecificiationTitle:     IntelTXTSpecificationTitle,
		SpecificationDocumentID: IntelTXTSpecificationDocumentID,
	}
	testbiosdataregionvalid = Test{
		Name:                    "BIOS DATA REGION valid",
		Required:                true,
		function:                BIOSDATAREGIONValid,
		dependencies:            []*Test{&testbiosdataregionpresent},
		Status:                  Implemented,
		SpecificationChapter:    "C.2 BIOS Data Format",
		SpecificiationTitle:     IntelTXTSpecificationTitle,
		SpecificationDocumentID: IntelTXTSpecificationDocumentID,
	}
	testhasmtrr = Test{
		Name:                    "CPU supports MTRRs",
		Required:                true,
		function:                HasMTRR,
		Status:                  Implemented,
		SpecificationChapter:    "2.2.5.1 MTRR Setup Prior to GETSEC[SENTER] Execution",
		SpecificiationTitle:     IntelTXTSpecificationTitle,
		SpecificationDocumentID: IntelTXTSpecificationDocumentID,
	}
	testhassmrr = Test{
		Name:         "CPU supports SMRRs",
		Required:     true,
		function:     HasSMRR,
		dependencies: []*Test{&testservermodetext},
		Status:       Implemented,
	}
	testvalidsmrr = Test{
		Name:         "SMRR covers SMM memory",
		Required:     true,
		function:     ValidSMRR,
		dependencies: []*Test{&testhassmrr, &testSupportsHostbridge},
		Status:       Implemented,
	}
	testactivesmrr = Test{
		Name:         "SMRR protection active",
		Required:     true,
		function:     ActiveSMRR,
		dependencies: []*Test{&testhassmrr},
		Status:       Implemented,
	}
	testactiveiommu = Test{
		Name:                    "IOMMU/VT-d active",
		Required:                false,
		function:                ActiveIOMMU,
		Status:                  Implemented,
		SpecificationChapter:    "1.11.2 Protected Memory Regions (PMRs)",
		SpecificiationTitle:     IntelTXTSpecificationTitle,
		SpecificationDocumentID: IntelTXTSpecificationDocumentID,
	}
	testservermodetext = Test{
		Name:     "TXT server mode enabled",
		Required: false,
		function: ServerModeTXT,
		Status:   Implemented,
	}

	// TestsMemory exposes the slice for memory related txt tests
	TestsMemory = [...]*Test{
		&testtxtmemoryrangevalid,
		&testtxtpublicisreserved,
		&testtxtprivateisreserved,
		&testmemoryisreserved,
		&testtpmdecodereserved,
		&testtxtmemoryisdpr,
		&testtxtdprislocked,
		&testhostbridgeDPRcorrect,
		&testhostbridgeDPRislocked,
		&testsinitintxt,
		&testsinitmatcheschipset,
		&testsinitmatchescpu,
		&testnosiniterrors,
		&testbiosdataregionpresent,
		&testbiosdataregionvalid,
		&testhasmtrr,
		&testhassmrr,
		&testvalidsmrr,
		&testactivesmrr,
		&testactiveiommu,
		&testservermodetext,
	}
)

var (
	biosdata tools.TXTBiosData
)

//nolint
const (
	//Heapsize from newer spec - Document 575623
	minHeapSize  = uint32(0xF0000)
	minSinitSize = uint32(0x10000)
	//Heapsize reduced for legacy spec - Document 558294
	legacyMinHeapSize = uint32(0xE0000)
)

// TXTHeapSpaceValid checks if the registers indicates the correct sizes
func TXTHeapSpaceValid(txtAPI hwapi.LowLevelHardwareInterfaces, config *tools.Configuration) (bool, error, error) {
	buf, err := tools.FetchTXTRegs(txtAPI)
	if err != nil {
		return false, nil, err
	}

	regs, err := tools.ParseTXTRegs(buf)
	if err != nil {
		return false, nil, err
	}

	if uint64(regs.HeapBase) >= FourGiB {
		return false, fmt.Errorf("HeapBase > 4Gib"), nil
	}

	if uint64(regs.HeapBase+regs.HeapSize) >= FourGiB {
		return false, fmt.Errorf("HeapBase + HeapSize >= 4Gib"), nil
	}

	//TODO: Validate against minHeapSize once legacy detection is implemented

	//This checks for legacy heap size - Document 558294
	if regs.HeapSize < legacyMinHeapSize {
		return false, fmt.Errorf("heap must be at least %v", legacyMinHeapSize), nil

	}

	if uint64(regs.SinitBase) >= FourGiB {
		return false, fmt.Errorf("SinitBase >= 4Gib"), nil
	}
	if uint64(regs.SinitBase)&0xfff > 0 {
		return false, fmt.Errorf("SinitBase must be 4 KiB aligned"), nil
	}
	if uint64(regs.SinitBase+regs.SinitSize) >= FourGiB {
		return false, fmt.Errorf("SinitBase + SinitSize >= 4Gib"), nil
	}

	if regs.SinitSize < minSinitSize {
		return false, fmt.Errorf("SINIT must be at least %v", minSinitSize), nil
	}

	if uint64(regs.MleJoin) >= FourGiB {
		return false, fmt.Errorf("MleJoin >= 4Gib"), nil
	}

	/* Document Number: 558294  5.5.6.2 SINIT Memory Region */
	if regs.SinitBase >= regs.HeapBase {
		return false, fmt.Errorf("SINIT region must be below Heap region"), nil
	}
	if regs.SinitBase > 0 && regs.SinitBase+regs.SinitSize != regs.HeapBase {
		return false, fmt.Errorf("SINIT region must end at start of Heap region"), nil
	}

	return true, nil, nil
}

// TXTPublicReservedInE820 checks if TXTPublic area is marked reserved in e820 map
func TXTPublicReservedInE820(txtAPI hwapi.LowLevelHardwareInterfaces, config *tools.Configuration) (bool, error, error) {
	res, err := txtAPI.IsReservedInE820(uint64(tools.TxtPublicSpace), uint64(tools.TxtPublicSpace+tools.TxtPublicSpaceSize))
	if err != nil {
		return false, nil, err
	}
	if !res {
		return false, fmt.Errorf("TXTPublic area is not marked reserved in e820 map"), nil
	}
	return true, nil, nil

}

// TXTPrivateReservedInE820 checks if TXTPrivate area is marked reserved in e820 map
func TXTPrivateReservedInE820(txtAPI hwapi.LowLevelHardwareInterfaces, config *tools.Configuration) (bool, error, error) {
	res, err := txtAPI.IsReservedInE820(uint64(tools.TxtPrivateSpace), uint64(tools.TxtPrivateSpace+tools.TxtPrivateSpaceSize))
	if err != nil {
		return false, nil, err
	}
	if !res {
		return false, fmt.Errorf("TXTPrivate area is not marked reserved in e820 map"), nil
	}
	return true, nil, nil
}

// TXTReservedInE820 checks if the HEAP/MSEG/SINIT TXT regions are marked reserved in e820 map.
func TXTReservedInE820(txtAPI hwapi.LowLevelHardwareInterfaces, config *tools.Configuration) (bool, error, error) {
	buf, err := tools.FetchTXTRegs(txtAPI)
	if err != nil {
		return false, nil, err
	}
	regs, err := tools.ParseTXTRegs(buf)
	if err != nil {
		return false, nil, err
	}

	heapReserved, err := txtAPI.IsReservedInE820(uint64(regs.HeapBase), uint64(regs.HeapBase+regs.HeapSize))
	if err != nil {
		return false, nil, err
	}

	sinitReserved, err := txtAPI.IsReservedInE820(uint64(regs.SinitBase), uint64(regs.SinitBase+regs.SinitSize))
	if err != nil {
		return false, nil, err
	}

	if heapReserved && sinitReserved {
		return true, nil, nil
	}
	return false, fmt.Errorf("HEAP/;SEG/SINIT TXT regions are NOT marked as reserved. HeapReserve: %v - SINITReserved: %v", heapReserved, sinitReserved), nil
}

// TXTTPMDecodeSpaceIn820 checks if TPMDecode area is marked as reserved in e820 map
func TXTTPMDecodeSpaceIn820(txtAPI hwapi.LowLevelHardwareInterfaces, config *tools.Configuration) (bool, error, error) {
	res, err := txtAPI.IsReservedInE820(uint64(tools.TxtTPMDecode), uint64(tools.TxtTPMDecode+tools.TxtTPMDecodeSize-1))
	if err != nil {
		return false, nil, err
	}
	if !res {
		return false, fmt.Errorf("TPMDecode area not reserved in e820"), nil
	}
	return true, nil, nil
}

// TXTMemoryIsDPR checks if the TXT DPR protects TXT memory.
func TXTMemoryIsDPR(txtAPI hwapi.LowLevelHardwareInterfaces, config *tools.Configuration) (bool, error, error) {
	buf, err := tools.FetchTXTRegs(txtAPI)
	if err != nil {
		return false, nil, err
	}
	regs, err := tools.ParseTXTRegs(buf)
	if err != nil {
		return false, nil, err
	}

	var dprBase uint32
	var dprSize uint32
	var dprLimit uint32

	dprSize = uint32(regs.Dpr.Size) * 1024 * 1024
	dprLimit = uint32(regs.Dpr.Top+1) * 1024 * 1024
	dprBase = dprLimit - dprSize

	/* Chapter 5.5.6.1 DMA Protection Memory Region */
	if dprSize < 3*1024*1024 {
		return false, fmt.Errorf("DPR region is smaller than 3 MiB"), nil
	}
	if dprBase > regs.HeapBase {
		return false, fmt.Errorf("TXT Heap region not covered by DPR region"), nil
	}
	if regs.SinitBase > 0 && dprBase > regs.SinitBase {
		return false, fmt.Errorf("TXT Sinit region not covered by DPR region"), nil
	}
	if dprLimit < regs.HeapBase+regs.HeapSize {
		return false, fmt.Errorf("TXT Heap region not covered by DPR region"), nil
	}
	if regs.SinitBase > 0 && dprLimit < regs.SinitBase+regs.SinitSize {
		return false, fmt.Errorf("TXT Sinit region not covered by DPR region"), nil
	}
	/* Chapter 5.5.6.3 Intel TXT Heap Memory Region */
	if dprLimit != regs.HeapBase+regs.HeapSize {
		return false, fmt.Errorf("TXT heap region must end at top of DPR region"), nil
	}
	/* Chapter 5.5.6.1 DMA Protection Memory Region */
	if dprBase > dprLimit-2*1024*1024-regs.HeapSize-regs.SinitSize {
		return false, fmt.Errorf("MLE region in DPR region is less than 2 MiB"), nil
	}

	return true, nil, nil
}

// TXTDPRisLock checks the TXTRegister in memory about the status of DPR if it's locked.
func TXTDPRisLock(txtAPI hwapi.LowLevelHardwareInterfaces, config *tools.Configuration) (bool, error, error) {
	buf, err := tools.FetchTXTRegs(txtAPI)
	if err != nil {
		return false, nil, err
	}
	regs, err := tools.ParseTXTRegs(buf)
	if err != nil {
		return false, nil, err
	}

	if !regs.Dpr.Lock {
		return false, fmt.Errorf("TXTDPR is not locked"), nil
	}
	return true, nil, nil
}

// HostbridgeIsSupported checks if the suite supports the hostbridge
func HostbridgeIsSupported(txtAPI hwapi.LowLevelHardwareInterfaces, config *tools.Configuration) (bool, error, error) {

	_, _, err := txtAPI.ReadHostBridgeTseg()
	if err != nil {
		return false, nil, err
	}
	return true, nil, nil
}

// HostbridgeDPRCorrect checks if TXT DPR equals PCI Hostbridge DPR
func HostbridgeDPRCorrect(txtAPI hwapi.LowLevelHardwareInterfaces, config *tools.Configuration) (bool, error, error) {
	buf, err := tools.FetchTXTRegs(txtAPI)
	if err != nil {
		return false, fmt.Errorf("cannot read DPR registers: %s", err), nil
	}
	regs, err := tools.ParseTXTRegs(buf)
	if err != nil {
		return false, fmt.Errorf("cannot parse DPR registers: %s", err), nil
	}

	hostbridgeDpr, err := txtAPI.ReadHostBridgeDPR()
	// No need to validate hostbridge register, already done for TXT DPR
	// Just make sure they match.
	if err != nil {
		return false, fmt.Errorf("unable to read the hostbridge DPR: %w", err), nil
	}

	if hostbridgeDpr.Top != regs.Dpr.Top {
		return false, fmt.Errorf("hostbridge DPR Top doesn't match TXT DPR Top"), nil
	}

	if hostbridgeDpr.Size != regs.Dpr.Size {
		return false, fmt.Errorf("hostbridge DPR Size doesn't match TXT DPR Size"), nil
	}

	return true, nil, nil
}

// HostbridgeDPRisLocked checks if the Hostbridge DPR is marked as locked
func HostbridgeDPRisLocked(txtAPI hwapi.LowLevelHardwareInterfaces, config *tools.Configuration) (bool, error, error) {
	hostbridgeDpr, err := txtAPI.ReadHostBridgeDPR()

	if err != nil {
		return false, nil, err
	}

	if !hostbridgeDpr.Lock {
		return false, fmt.Errorf("hostbridge DPR isn't locked"), nil
	}
	return true, nil, nil
}

// SINITInTXT checks the TXT region on containing a valid SINIT ACM.
func SINITInTXT(txtAPI hwapi.LowLevelHardwareInterfaces, config *tools.Configuration) (bool, error, error) {
	buf, err := tools.FetchTXTRegs(txtAPI)
	if err != nil {
		return false, nil, err
	}
	regs, err := tools.ParseTXTRegs(buf)
	if err != nil {
		return false, nil, err
	}

	sinitBuf := make([]byte, regs.SinitSize)
	err = txtAPI.ReadPhysBuf(int64(regs.SinitBase), sinitBuf)
	if err != nil {
		return false, nil, err
	}

	acm, _, _, _, err, internalerr := tools.ParseACM(sinitBuf)
	if internalerr != nil {
		return false, nil, internalerr
	}
	if err != nil {
		return false, err, nil
	}
	if acm == nil {
		return false, fmt.Errorf("ACM is nil"), nil
	}

	if acm.Header.ModuleType != 2 {
		return false, fmt.Errorf("SINIT in TXT: ACM ModuleType not 2"), nil
	}
	return true, nil, nil

}

// SINITMatchesChipset checks if the SINIT ACM matches the chipset.
func SINITMatchesChipset(txtAPI hwapi.LowLevelHardwareInterfaces, config *tools.Configuration) (bool, error, error) {
	buf, err := tools.FetchTXTRegs(txtAPI)
	if err != nil {
		return false, nil, err
	}
	regs, err := tools.ParseTXTRegs(buf)
	if err != nil {
		return false, nil, err
	}

	acm, chps, _, _, err, internalerr := sinitACM(txtAPI, regs)
	if internalerr != nil {
		return false, nil, internalerr
	}
	if err != nil {
		return false, err, nil
	}
	if chps == nil {
		return false, fmt.Errorf("CHPS is nil"), nil
	}

	for _, ch := range chps.IDList {
		a := ch.VendorID == regs.Vid
		b := ch.DeviceID == regs.Did

		if a && b {
			if acm.Header.Flags&1 != 0 {
				if ch.RevisionID&regs.Rid == regs.Rid {
					return true, nil, nil
				}
			} else {
				if ch.RevisionID == regs.Rid {
					return true, nil, nil
				}
			}
		}
	}

	return false, fmt.Errorf("SINIT doesn't match chipset"), nil
}

// SINITMatchesCPU checks if the SINITACM matches the CPU
func SINITMatchesCPU(txtAPI hwapi.LowLevelHardwareInterfaces, config *tools.Configuration) (bool, error, error) {
	buf, err := tools.FetchTXTRegs(txtAPI)
	if err != nil {
		return false, nil, err
	}
	regs, err := tools.ParseTXTRegs(buf)
	if err != nil {
		return false, nil, err
	}

	_, _, cpus, _, err, internalerr := sinitACM(txtAPI, regs)
	if internalerr != nil {
		return false, nil, internalerr
	}
	if err != nil {
		return false, err, nil
	}

	// IA32_PLATFORM_ID
	platform, err := txtAPI.IA32PlatformID()
	if err != nil {
		return false, nil, err
	}

	fms := txtAPI.CPUSignature()

	for _, cpu := range cpus.IDList {
		a := fms&cpu.FMSMask == cpu.FMS
		b := platform&cpu.PlatformMask == cpu.PlatformID

		if a && b {
			return true, nil, nil
		}
	}

	return false, fmt.Errorf("CPU signature not found in SINIT processor ID list"), nil
}

// NoSINITErrors checks if the SINITACM was executed without any errors
func NoSINITErrors(txtAPI hwapi.LowLevelHardwareInterfaces, config *tools.Configuration) (bool, error, error) {
	buf, err := tools.FetchTXTRegs(txtAPI)
	if err != nil {
		return false, nil, err
	}
	regs, err := tools.ParseTXTRegs(buf)
	if err != nil {
		return false, nil, err
	}

	if regs.ErrorCodeRaw != 0xc0000001 {
		return false, fmt.Errorf("SINIT Error detected"), nil
	}
	return true, nil, nil
}

// BIOSDATAREGIONPresent checks is the BIOSDATA Region is present in TXT Register Space
func BIOSDATAREGIONPresent(txtAPI hwapi.LowLevelHardwareInterfaces, config *tools.Configuration) (bool, error, error) {
	buf, err := tools.FetchTXTRegs(txtAPI)
	if err != nil {
		return false, nil, err
	}
	regs, err := tools.ParseTXTRegs(buf)
	if err != nil {
		return false, nil, err
	}

	txtHeap := make([]byte, regs.HeapSize)
	err = txtAPI.ReadPhysBuf(int64(regs.HeapBase), txtHeap)
	if err != nil {
		return false, nil, err
	}

	biosdata, err = tools.ParseBIOSDataRegion(txtHeap)
	if err != nil {
		return false, nil, err
	}

	return true, nil, nil
}

// BIOSDATAREGIONValid checks if the BIOSDATA Region in TXT Register Space is valid
func BIOSDATAREGIONValid(txtAPI hwapi.LowLevelHardwareInterfaces, config *tools.Configuration) (bool, error, error) {
	if biosdata.Version < 2 {
		return false, fmt.Errorf("BIOS DATA regions version < 2 are not supperted"), nil
	}

	if biosdata.BiosSinitSize < 8 {
		return false, fmt.Errorf("BIOS DATA region is too small"), nil
	}

	if biosdata.NumLogProcs == 0 {
		return false, fmt.Errorf("BIOS DATA region corrupted"), nil
	}
	return true, nil, nil
}

// HasMTRR checks if MTRR is supported by CPU
func HasMTRR(txtAPI hwapi.LowLevelHardwareInterfaces, config *tools.Configuration) (bool, error, error) {
	if !txtAPI.HasMTRR() {
		return false, fmt.Errorf("CPU does not have MTRR"), nil
	}
	return true, nil, nil
}

// HasSMRR checks if SMRR is supported
func HasSMRR(txtAPI hwapi.LowLevelHardwareInterfaces, config *tools.Configuration) (bool, error, error) {
	ret, err := txtAPI.HasSMRR()
	if err != nil {
		return false, nil, err
	}
	if !ret {
		return false, fmt.Errorf("CPU has no SMRR"), nil
	}
	return true, nil, nil
}

// ValidSMRR checks if SMRR is valid
func ValidSMRR(txtAPI hwapi.LowLevelHardwareInterfaces, config *tools.Configuration) (bool, error, error) {
	smrr, err := txtAPI.GetSMRRInfo()
	if err != nil {
		return false, nil, err
	}

	if smrr.PhysMask == 0 {
		return false, fmt.Errorf("SMRR PhysMask isn't set"), nil
	}
	if smrr.PhysBase == 0 {
		return false, fmt.Errorf("SMRR PhysBase isn't set"), nil
	}

	tsegbase, tseglimit, err := txtAPI.ReadHostBridgeTseg()
	if err != nil {
		return false, nil, err
	}
	if tsegbase == 0 || tsegbase == 0xffffffff {
		return false, fmt.Errorf("TSEG base register isn't valid"), nil
	}
	if tseglimit == 0 || tseglimit == 0xffffffff {
		return false, fmt.Errorf("TSEG limit register isn't valid"), nil
	}

	if tsegbase&(^(uint32(smrr.PhysMask) << 12)) != 0 {
		return false, fmt.Errorf("TSEG base isn't aligned to SMRR Physmask"), nil
	}
	if tsegbase != (uint32(smrr.PhysBase) << 12) {
		return false, fmt.Errorf("TSEG base doesn't start at SMRR PhysBase"), nil
	}
	if tseglimit&(^(uint32(smrr.PhysMask) << 12)) != 0 {
		return false, fmt.Errorf("TSEG limit isn't aligned to SMRR Physmask"), nil
	}
	if ((tseglimit - 1) & (uint32(smrr.PhysMask) << 12)) != (uint32(smrr.PhysBase) << 12) {
		return false, fmt.Errorf("SMRR Physmask doesn't cover whole TSEG"), nil
	}

	return true, nil, nil
}

// ActiveSMRR checks if SMMR is set active
func ActiveSMRR(txtAPI hwapi.LowLevelHardwareInterfaces, config *tools.Configuration) (bool, error, error) {
	smrr, err := txtAPI.GetSMRRInfo()
	if err != nil {
		return false, nil, err
	}

	if !smrr.Active {
		return false, fmt.Errorf("SMRR not active"), nil
	}
	return true, nil, nil
}

// ActiveIOMMU checks if IOMMU is active
func ActiveIOMMU(txtAPI hwapi.LowLevelHardwareInterfaces, config *tools.Configuration) (bool, error, error) {
	smrr, err := txtAPI.GetSMRRInfo()
	if err != nil {
		return false, nil, err
	}
	smrrPhysEnd := (smrr.PhysBase | ^smrr.PhysMask) & 0xfffff
	ret, err := txtAPI.AddressRangesIsDMAProtected(smrr.PhysBase, smrrPhysEnd)
	if err != nil {
		return false, fmt.Errorf("failed to check SMRR DMA protection: %s", err), nil
	}
	if !ret {
		return false, fmt.Errorf("IOMMU does not protect SMRR (%x-%x) from DMA", smrr.PhysBase, smrrPhysEnd), nil
	}
	return true, nil, nil
}

// ServerModeTXT checks if TXT runs in Servermode
func ServerModeTXT(txtAPI hwapi.LowLevelHardwareInterfaces, config *tools.Configuration) (bool, error, error) {
	// FIXME: FindOverlapping GetSec[Parameters] ebx = 5
	// Assume yes if dependencies are satisfied
	val, err := txtAPI.HasSMRR()
	if err != nil {
		return false, nil, err
	}
	if txtAPI.HasSMX() && txtAPI.HasVMX() && val {
		return true, nil, nil
	}
	return false, fmt.Errorf("servermode not active"), nil
}

//ReleaseFusedFSBI checks if the FSBI is release fused
func ReleaseFusedFSBI(txtAPI hwapi.LowLevelHardwareInterfaces, config *tools.Configuration) (bool, error, error) {
	return false, nil, fmt.Errorf("ReleaseFusedFSBI: Unimplemented")
}

func sinitACM(txtAPI hwapi.LowLevelHardwareInterfaces, regs tools.TXTRegisterSpace) (*tools.ACM, *tools.Chipsets, *tools.Processors, *tools.TPMs, error, error) {
	sinitBuf := make([]byte, regs.SinitSize)
	err := txtAPI.ReadPhysBuf(int64(regs.SinitBase), sinitBuf)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	return tools.ParseACM(sinitBuf)
}
