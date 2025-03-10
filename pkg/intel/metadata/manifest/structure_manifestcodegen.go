//go:build !manifestcodegen
// +build !manifestcodegen

// Code generated by "menifestcodegen". DO NOT EDIT.
// To reproduce: go run github.com/9elements/converged-security-suite/v2/pkg/intel/metadata/manifest/common/manifestcodegen/cmd/manifestcodegen github.com/9elements/converged-security-suite/v2/pkg/intel/metadata/manifest

package manifest

import (
	"encoding/binary"
	"fmt"
	"io"
	"strings"

	"github.com/9elements/converged-security-suite/v2/pkg/intel/metadata/manifest/common/pretty"
)

var (
	// Just to avoid errors in "import" above in case if it wasn't used below
	_ = binary.LittleEndian
	_ = (fmt.Stringer)(nil)
	_ = (io.Reader)(nil)
	_ = pretty.Header
	_ = strings.Join
)

// NewStructInfo returns a new instance of StructInfo with
// all default values set.
func NewStructInfo() *StructInfo {
	s := &StructInfo{}
	s.Rehash()
	return s
}

// Validate (recursively) checks the structure if there are any unexpected
// values. It returns an error if so.
func (s *StructInfo) Validate() error {

	return nil
}

// ReadFrom reads the StructInfo from 'r' in format defined in the document #575623.
func (s *StructInfo) ReadFrom(r io.Reader) (int64, error) {
	totalN := int64(0)

	// ID (ManifestFieldType: arrayStatic)
	{
		n, err := 8, binary.Read(r, binary.LittleEndian, s.ID[:])
		if err != nil {
			return totalN, fmt.Errorf("unable to read field 'ID': %w", err)
		}
		totalN += int64(n)
	}

	// Version (ManifestFieldType: endValue)
	{
		n, err := 1, binary.Read(r, binary.LittleEndian, &s.Version)
		if err != nil {
			return totalN, fmt.Errorf("unable to read field 'Version': %w", err)
		}
		totalN += int64(n)
	}

	// Variable0 (ManifestFieldType: endValue)
	{
		n, err := 1, binary.Read(r, binary.LittleEndian, &s.Variable0)
		if err != nil {
			return totalN, fmt.Errorf("unable to read field 'Variable0': %w", err)
		}
		totalN += int64(n)
	}

	// ElementSize (ManifestFieldType: endValue)
	{
		n, err := 2, binary.Read(r, binary.LittleEndian, &s.ElementSize)
		if err != nil {
			return totalN, fmt.Errorf("unable to read field 'ElementSize': %w", err)
		}
		totalN += int64(n)
	}

	return totalN, nil
}

// RehashRecursive calls Rehash (see below) recursively.
func (s *StructInfo) RehashRecursive() {
	s.Rehash()
}

// Rehash sets values which are calculated automatically depending on the rest
// data. It is usually about the total size field of an element.
func (s *StructInfo) Rehash() {
}

// WriteTo writes the StructInfo into 'w' in format defined in
// the document #575623.
func (s *StructInfo) WriteTo(w io.Writer) (int64, error) {
	totalN := int64(0)
	s.Rehash()

	// ID (ManifestFieldType: arrayStatic)
	{
		n, err := 8, binary.Write(w, binary.LittleEndian, s.ID[:])
		if err != nil {
			return totalN, fmt.Errorf("unable to write field 'ID': %w", err)
		}
		totalN += int64(n)
	}

	// Version (ManifestFieldType: endValue)
	{
		n, err := 1, binary.Write(w, binary.LittleEndian, &s.Version)
		if err != nil {
			return totalN, fmt.Errorf("unable to write field 'Version': %w", err)
		}
		totalN += int64(n)
	}

	// Variable0 (ManifestFieldType: endValue)
	{
		n, err := 1, binary.Write(w, binary.LittleEndian, &s.Variable0)
		if err != nil {
			return totalN, fmt.Errorf("unable to write field 'Variable0': %w", err)
		}
		totalN += int64(n)
	}

	// ElementSize (ManifestFieldType: endValue)
	{
		n, err := 2, binary.Write(w, binary.LittleEndian, &s.ElementSize)
		if err != nil {
			return totalN, fmt.Errorf("unable to write field 'ElementSize': %w", err)
		}
		totalN += int64(n)
	}

	return totalN, nil
}

// IDSize returns the size in bytes of the value of field ID
func (s *StructInfo) IDTotalSize() uint64 {
	return 8
}

// VersionSize returns the size in bytes of the value of field Version
func (s *StructInfo) VersionTotalSize() uint64 {
	return 1
}

// Variable0Size returns the size in bytes of the value of field Variable0
func (s *StructInfo) Variable0TotalSize() uint64 {
	return 1
}

// ElementSizeSize returns the size in bytes of the value of field ElementSize
func (s *StructInfo) ElementSizeTotalSize() uint64 {
	return 2
}

// IDOffset returns the offset in bytes of field ID
func (s *StructInfo) IDOffset() uint64 {
	return 0
}

// VersionOffset returns the offset in bytes of field Version
func (s *StructInfo) VersionOffset() uint64 {
	return s.IDOffset() + s.IDTotalSize()
}

// Variable0Offset returns the offset in bytes of field Variable0
func (s *StructInfo) Variable0Offset() uint64 {
	return s.VersionOffset() + s.VersionTotalSize()
}

// ElementSizeOffset returns the offset in bytes of field ElementSize
func (s *StructInfo) ElementSizeOffset() uint64 {
	return s.Variable0Offset() + s.Variable0TotalSize()
}

// Size returns the total size of the StructInfo.
func (s *StructInfo) TotalSize() uint64 {
	if s == nil {
		return 0
	}

	var size uint64
	size += s.IDTotalSize()
	size += s.VersionTotalSize()
	size += s.Variable0TotalSize()
	size += s.ElementSizeTotalSize()
	return size
}

// PrettyString returns the content of the structure in an easy-to-read format.
func (s *StructInfo) PrettyString(depth uint, withHeader bool, opts ...pretty.Option) string {
	var lines []string
	if withHeader {
		lines = append(lines, pretty.Header(depth, "Struct Info", s))
	}
	if s == nil {
		return strings.Join(lines, "\n")
	}
	// ManifestFieldType is arrayStatic
	lines = append(lines, pretty.SubValue(depth+1, "ID", "", &s.ID, opts...)...)
	// ManifestFieldType is endValue
	lines = append(lines, pretty.SubValue(depth+1, "Version", "", &s.Version, opts...)...)
	// ManifestFieldType is endValue
	lines = append(lines, pretty.SubValue(depth+1, "Variable 0", "", &s.Variable0, opts...)...)
	// ManifestFieldType is endValue
	lines = append(lines, pretty.SubValue(depth+1, "Element Size", "", &s.ElementSize, opts...)...)
	if depth < 2 {
		lines = append(lines, "")
	}
	return strings.Join(lines, "\n")
}
