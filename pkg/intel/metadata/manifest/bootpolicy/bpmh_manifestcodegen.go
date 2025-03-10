//go:build !manifestcodegen
// +build !manifestcodegen

// Code generated by "menifestcodegen". DO NOT EDIT.
// To reproduce: go run github.com/9elements/converged-security-suite/v2/pkg/intel/metadata/manifest/common/manifestcodegen/cmd/manifestcodegen github.com/9elements/converged-security-suite/v2/pkg/intel/metadata/manifest/bootpolicy

package bootpolicy

import (
	"encoding/binary"
	"fmt"
	"io"
	"strings"

	"github.com/9elements/converged-security-suite/v2/pkg/intel/metadata/manifest"
	"github.com/9elements/converged-security-suite/v2/pkg/intel/metadata/manifest/common/pretty"
)

var (
	// Just to avoid errors in "import" above in case if it wasn't used below
	_ = binary.LittleEndian
	_ = (fmt.Stringer)(nil)
	_ = (io.Reader)(nil)
	_ = pretty.Header
	_ = strings.Join
	_ = manifest.StructInfo{}
)

// NewBPMH returns a new instance of BPMH with
// all default values set.
func NewBPMH() *BPMH {
	s := &BPMH{}
	copy(s.StructInfo.ID[:], []byte(StructureIDBPMH))
	s.StructInfo.Version = 0x23
	s.Rehash()
	return s
}

// Validate (recursively) checks the structure if there are any unexpected
// values. It returns an error if so.
func (s *BPMH) Validate() error {
	// See tag "require"
	for idx := range s.Reserved0 {
		if s.Reserved0[idx] != 0 {
			return fmt.Errorf("'Reserved0[%d]' is expected to be 0, but it is %v", idx, s.Reserved0[idx])
		}
	}

	return nil
}

// StructureIDBPMH is the StructureID (in terms of
// the document #575623) of element 'BPMH'.
const StructureIDBPMH = "__ACBP__"

// GetStructInfo returns current value of StructInfo of the structure.
//
// StructInfo is a set of standard fields with presented in any element
// ("element" in terms of document #575623).
func (s *BPMH) GetStructInfo() manifest.StructInfo {
	return s.StructInfo
}

// SetStructInfo sets new value of StructInfo to the structure.
//
// StructInfo is a set of standard fields with presented in any element
// ("element" in terms of document #575623).
func (s *BPMH) SetStructInfo(newStructInfo manifest.StructInfo) {
	s.StructInfo = newStructInfo
}

// ReadFrom reads the BPMH from 'r' in format defined in the document #575623.
func (s *BPMH) ReadFrom(r io.Reader) (int64, error) {
	var totalN int64

	err := binary.Read(r, binary.LittleEndian, &s.StructInfo)
	if err != nil {
		return totalN, fmt.Errorf("unable to read structure info at %d: %w", totalN, err)
	}
	totalN += int64(binary.Size(s.StructInfo))

	n, err := s.ReadDataFrom(r)
	if err != nil {
		return totalN, fmt.Errorf("unable to read data: %w", err)
	}
	totalN += n

	return totalN, nil
}

// ReadDataFrom reads the BPMH from 'r' excluding StructInfo,
// in format defined in the document #575623.
func (s *BPMH) ReadDataFrom(r io.Reader) (int64, error) {
	totalN := int64(0)

	// StructInfo (ManifestFieldType: structInfo)
	{
		// ReadDataFrom does not read Struct, use ReadFrom for that.
	}

	// KeySignatureOffset (ManifestFieldType: endValue)
	{
		n, err := 2, binary.Read(r, binary.LittleEndian, &s.KeySignatureOffset)
		if err != nil {
			return totalN, fmt.Errorf("unable to read field 'KeySignatureOffset': %w", err)
		}
		totalN += int64(n)
	}

	// BPMRevision (ManifestFieldType: endValue)
	{
		n, err := 1, binary.Read(r, binary.LittleEndian, &s.BPMRevision)
		if err != nil {
			return totalN, fmt.Errorf("unable to read field 'BPMRevision': %w", err)
		}
		totalN += int64(n)
	}

	// BPMSVN (ManifestFieldType: endValue)
	{
		n, err := 1, binary.Read(r, binary.LittleEndian, &s.BPMSVN)
		if err != nil {
			return totalN, fmt.Errorf("unable to read field 'BPMSVN': %w", err)
		}
		totalN += int64(n)
	}

	// ACMSVNAuth (ManifestFieldType: endValue)
	{
		n, err := 1, binary.Read(r, binary.LittleEndian, &s.ACMSVNAuth)
		if err != nil {
			return totalN, fmt.Errorf("unable to read field 'ACMSVNAuth': %w", err)
		}
		totalN += int64(n)
	}

	// Reserved0 (ManifestFieldType: arrayStatic)
	{
		n, err := 1, binary.Read(r, binary.LittleEndian, s.Reserved0[:])
		if err != nil {
			return totalN, fmt.Errorf("unable to read field 'Reserved0': %w", err)
		}
		totalN += int64(n)
	}

	// NEMDataStack (ManifestFieldType: endValue)
	{
		n, err := 2, binary.Read(r, binary.LittleEndian, &s.NEMDataStack)
		if err != nil {
			return totalN, fmt.Errorf("unable to read field 'NEMDataStack': %w", err)
		}
		totalN += int64(n)
	}

	return totalN, nil
}

// RehashRecursive calls Rehash (see below) recursively.
func (s *BPMH) RehashRecursive() {
	s.StructInfo.Rehash()
	s.Rehash()
}

// Rehash sets values which are calculated automatically depending on the rest
// data. It is usually about the total size field of an element.
func (s *BPMH) Rehash() {
	s.Variable0 = 0x20
	s.ElementSize = uint16(s.TotalSize())
}

// WriteTo writes the BPMH into 'w' in format defined in
// the document #575623.
func (s *BPMH) WriteTo(w io.Writer) (int64, error) {
	totalN := int64(0)
	s.Rehash()

	// StructInfo (ManifestFieldType: structInfo)
	{
		n, err := s.StructInfo.WriteTo(w)
		if err != nil {
			return totalN, fmt.Errorf("unable to write field 'StructInfo': %w", err)
		}
		totalN += int64(n)
	}

	// KeySignatureOffset (ManifestFieldType: endValue)
	{
		n, err := 2, binary.Write(w, binary.LittleEndian, &s.KeySignatureOffset)
		if err != nil {
			return totalN, fmt.Errorf("unable to write field 'KeySignatureOffset': %w", err)
		}
		totalN += int64(n)
	}

	// BPMRevision (ManifestFieldType: endValue)
	{
		n, err := 1, binary.Write(w, binary.LittleEndian, &s.BPMRevision)
		if err != nil {
			return totalN, fmt.Errorf("unable to write field 'BPMRevision': %w", err)
		}
		totalN += int64(n)
	}

	// BPMSVN (ManifestFieldType: endValue)
	{
		n, err := 1, binary.Write(w, binary.LittleEndian, &s.BPMSVN)
		if err != nil {
			return totalN, fmt.Errorf("unable to write field 'BPMSVN': %w", err)
		}
		totalN += int64(n)
	}

	// ACMSVNAuth (ManifestFieldType: endValue)
	{
		n, err := 1, binary.Write(w, binary.LittleEndian, &s.ACMSVNAuth)
		if err != nil {
			return totalN, fmt.Errorf("unable to write field 'ACMSVNAuth': %w", err)
		}
		totalN += int64(n)
	}

	// Reserved0 (ManifestFieldType: arrayStatic)
	{
		n, err := 1, binary.Write(w, binary.LittleEndian, s.Reserved0[:])
		if err != nil {
			return totalN, fmt.Errorf("unable to write field 'Reserved0': %w", err)
		}
		totalN += int64(n)
	}

	// NEMDataStack (ManifestFieldType: endValue)
	{
		n, err := 2, binary.Write(w, binary.LittleEndian, &s.NEMDataStack)
		if err != nil {
			return totalN, fmt.Errorf("unable to write field 'NEMDataStack': %w", err)
		}
		totalN += int64(n)
	}

	return totalN, nil
}

// StructInfoSize returns the size in bytes of the value of field StructInfo
func (s *BPMH) StructInfoTotalSize() uint64 {
	return s.StructInfo.TotalSize()
}

// KeySignatureOffsetSize returns the size in bytes of the value of field KeySignatureOffset
func (s *BPMH) KeySignatureOffsetTotalSize() uint64 {
	return 2
}

// BPMRevisionSize returns the size in bytes of the value of field BPMRevision
func (s *BPMH) BPMRevisionTotalSize() uint64 {
	return 1
}

// BPMSVNSize returns the size in bytes of the value of field BPMSVN
func (s *BPMH) BPMSVNTotalSize() uint64 {
	return 1
}

// ACMSVNAuthSize returns the size in bytes of the value of field ACMSVNAuth
func (s *BPMH) ACMSVNAuthTotalSize() uint64 {
	return 1
}

// Reserved0Size returns the size in bytes of the value of field Reserved0
func (s *BPMH) Reserved0TotalSize() uint64 {
	return 1
}

// NEMDataStackSize returns the size in bytes of the value of field NEMDataStack
func (s *BPMH) NEMDataStackTotalSize() uint64 {
	return 2
}

// StructInfoOffset returns the offset in bytes of field StructInfo
func (s *BPMH) StructInfoOffset() uint64 {
	return 0
}

// KeySignatureOffsetOffset returns the offset in bytes of field KeySignatureOffset
func (s *BPMH) KeySignatureOffsetOffset() uint64 {
	return s.StructInfoOffset() + s.StructInfoTotalSize()
}

// BPMRevisionOffset returns the offset in bytes of field BPMRevision
func (s *BPMH) BPMRevisionOffset() uint64 {
	return s.KeySignatureOffsetOffset() + s.KeySignatureOffsetTotalSize()
}

// BPMSVNOffset returns the offset in bytes of field BPMSVN
func (s *BPMH) BPMSVNOffset() uint64 {
	return s.BPMRevisionOffset() + s.BPMRevisionTotalSize()
}

// ACMSVNAuthOffset returns the offset in bytes of field ACMSVNAuth
func (s *BPMH) ACMSVNAuthOffset() uint64 {
	return s.BPMSVNOffset() + s.BPMSVNTotalSize()
}

// Reserved0Offset returns the offset in bytes of field Reserved0
func (s *BPMH) Reserved0Offset() uint64 {
	return s.ACMSVNAuthOffset() + s.ACMSVNAuthTotalSize()
}

// NEMDataStackOffset returns the offset in bytes of field NEMDataStack
func (s *BPMH) NEMDataStackOffset() uint64 {
	return s.Reserved0Offset() + s.Reserved0TotalSize()
}

// Size returns the total size of the BPMH.
func (s *BPMH) TotalSize() uint64 {
	if s == nil {
		return 0
	}

	var size uint64
	size += s.StructInfoTotalSize()
	size += s.KeySignatureOffsetTotalSize()
	size += s.BPMRevisionTotalSize()
	size += s.BPMSVNTotalSize()
	size += s.ACMSVNAuthTotalSize()
	size += s.Reserved0TotalSize()
	size += s.NEMDataStackTotalSize()
	return size
}

// PrettyString returns the content of the structure in an easy-to-read format.
func (s *BPMH) PrettyString(depth uint, withHeader bool, opts ...pretty.Option) string {
	var lines []string
	if withHeader {
		lines = append(lines, pretty.Header(depth, "BPMH", s))
	}
	if s == nil {
		return strings.Join(lines, "\n")
	}
	// ManifestFieldType is structInfo
	lines = append(lines, pretty.SubValue(depth+1, "Struct Info", "", &s.StructInfo, opts...)...)
	// ManifestFieldType is endValue
	lines = append(lines, pretty.SubValue(depth+1, "Key Signature Offset", "", &s.KeySignatureOffset, opts...)...)
	// ManifestFieldType is endValue
	lines = append(lines, pretty.SubValue(depth+1, "BPM Revision", "", &s.BPMRevision, opts...)...)
	// ManifestFieldType is endValue
	lines = append(lines, pretty.SubValue(depth+1, "BPM SVN", "", &s.BPMSVN, opts...)...)
	// ManifestFieldType is endValue
	lines = append(lines, pretty.SubValue(depth+1, "ACM SVN Auth", "", &s.ACMSVNAuth, opts...)...)
	// ManifestFieldType is arrayStatic
	lines = append(lines, pretty.SubValue(depth+1, "Reserved 0", "", &s.Reserved0, opts...)...)
	// ManifestFieldType is endValue
	lines = append(lines, pretty.SubValue(depth+1, "NEM Data Stack", "", &s.NEMDataStack, opts...)...)
	if depth < 2 {
		lines = append(lines, "")
	}
	return strings.Join(lines, "\n")
}

// PrettyString returns the bits of the flags in an easy-to-read format.
func (v Size4K) PrettyString(depth uint, withHeader bool, opts ...pretty.Option) string {
	var lines []string
	if withHeader {
		lines = append(lines, pretty.Header(depth, "Size 4 K", v))
	}
	lines = append(lines, pretty.SubValue(depth+1, "In Bytes", "", v.InBytes(), opts...)...)
	return strings.Join(lines, "\n")
}

// TotalSize returns the total size measured through binary.Size.
func (v Size4K) TotalSize() uint64 {
	return uint64(binary.Size(v))
}

// WriteTo writes the Size4K into 'w' in binary format.
func (v Size4K) WriteTo(w io.Writer) (int64, error) {
	return int64(v.TotalSize()), binary.Write(w, binary.LittleEndian, v)
}

// ReadFrom reads the Size4K from 'r' in binary format.
func (v Size4K) ReadFrom(r io.Reader) (int64, error) {
	return int64(v.TotalSize()), binary.Read(r, binary.LittleEndian, v)
}
