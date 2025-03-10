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

// NewReserved returns a new instance of Reserved with
// all default values set.
func NewReserved() *Reserved {
	s := &Reserved{}
	copy(s.StructInfo.ID[:], []byte(StructureIDReserved))
	s.StructInfo.Version = 0x21
	s.Rehash()
	return s
}

// Validate (recursively) checks the structure if there are any unexpected
// values. It returns an error if so.
func (s *Reserved) Validate() error {

	return nil
}

// StructureIDReserved is the StructureID (in terms of
// the document #575623) of element 'Reserved'.
const StructureIDReserved = "__PFRS__"

// GetStructInfo returns current value of StructInfo of the structure.
//
// StructInfo is a set of standard fields with presented in any element
// ("element" in terms of document #575623).
func (s *Reserved) GetStructInfo() manifest.StructInfo {
	return s.StructInfo
}

// SetStructInfo sets new value of StructInfo to the structure.
//
// StructInfo is a set of standard fields with presented in any element
// ("element" in terms of document #575623).
func (s *Reserved) SetStructInfo(newStructInfo manifest.StructInfo) {
	s.StructInfo = newStructInfo
}

// ReadFrom reads the Reserved from 'r' in format defined in the document #575623.
func (s *Reserved) ReadFrom(r io.Reader) (int64, error) {
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

// ReadDataFrom reads the Reserved from 'r' excluding StructInfo,
// in format defined in the document #575623.
func (s *Reserved) ReadDataFrom(r io.Reader) (int64, error) {
	totalN := int64(0)

	// StructInfo (ManifestFieldType: structInfo)
	{
		// ReadDataFrom does not read Struct, use ReadFrom for that.
	}

	// ReservedData (ManifestFieldType: arrayStatic)
	{
		n, err := 32, binary.Read(r, binary.LittleEndian, s.ReservedData[:])
		if err != nil {
			return totalN, fmt.Errorf("unable to read field 'ReservedData': %w", err)
		}
		totalN += int64(n)
	}

	return totalN, nil
}

// RehashRecursive calls Rehash (see below) recursively.
func (s *Reserved) RehashRecursive() {
	s.StructInfo.Rehash()
	s.Rehash()
}

// Rehash sets values which are calculated automatically depending on the rest
// data. It is usually about the total size field of an element.
func (s *Reserved) Rehash() {
	s.Variable0 = 0
	s.ElementSize = uint16(s.TotalSize())
}

// WriteTo writes the Reserved into 'w' in format defined in
// the document #575623.
func (s *Reserved) WriteTo(w io.Writer) (int64, error) {
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

	// ReservedData (ManifestFieldType: arrayStatic)
	{
		n, err := 32, binary.Write(w, binary.LittleEndian, s.ReservedData[:])
		if err != nil {
			return totalN, fmt.Errorf("unable to write field 'ReservedData': %w", err)
		}
		totalN += int64(n)
	}

	return totalN, nil
}

// StructInfoSize returns the size in bytes of the value of field StructInfo
func (s *Reserved) StructInfoTotalSize() uint64 {
	return s.StructInfo.TotalSize()
}

// ReservedDataSize returns the size in bytes of the value of field ReservedData
func (s *Reserved) ReservedDataTotalSize() uint64 {
	return 32
}

// StructInfoOffset returns the offset in bytes of field StructInfo
func (s *Reserved) StructInfoOffset() uint64 {
	return 0
}

// ReservedDataOffset returns the offset in bytes of field ReservedData
func (s *Reserved) ReservedDataOffset() uint64 {
	return s.StructInfoOffset() + s.StructInfoTotalSize()
}

// Size returns the total size of the Reserved.
func (s *Reserved) TotalSize() uint64 {
	if s == nil {
		return 0
	}

	var size uint64
	size += s.StructInfoTotalSize()
	size += s.ReservedDataTotalSize()
	return size
}

// PrettyString returns the content of the structure in an easy-to-read format.
func (s *Reserved) PrettyString(depth uint, withHeader bool, opts ...pretty.Option) string {
	var lines []string
	if withHeader {
		lines = append(lines, pretty.Header(depth, "Reserved", s))
	}
	if s == nil {
		return strings.Join(lines, "\n")
	}
	// ManifestFieldType is structInfo
	lines = append(lines, pretty.SubValue(depth+1, "Struct Info", "", &s.StructInfo, opts...)...)
	// ManifestFieldType is arrayStatic
	lines = append(lines, pretty.SubValue(depth+1, "Reserved Data", "", &s.ReservedData, opts...)...)
	if depth < 2 {
		lines = append(lines, "")
	}
	return strings.Join(lines, "\n")
}
