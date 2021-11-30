module github.com/orangecms/converged-security-suite

go 1.16

require (
	github.com/9elements/go-linux-lowlevel-hw v0.0.0-20210922135846-40d2937646a8
	github.com/alecthomas/kong v0.2.11
	github.com/creasty/defaults v1.5.1
	github.com/digitalocean/go-smbios v0.0.0-20180907143718-390a4f403a8e
	github.com/dustin/go-humanize v1.0.0
	github.com/edsrzf/mmap-go v1.0.0
	github.com/fatih/camelcase v1.0.0
	github.com/fearful-symmetry/gomsr v0.0.1
	github.com/golang-collections/go-datastructures v0.0.0-20150211160725-59788d5eb259
	github.com/google/go-attestation v0.2.2
	github.com/google/go-tpm v0.3.3-0.20210120190357-1ff48daca32f
	github.com/google/uuid v1.2.0
	github.com/linuxboot/cbfs v0.0.0-20210504130259-7e6ab4ccb5aa
	github.com/linuxboot/fiano v1.0.1
	github.com/logrusorgru/aurora v2.0.3+incompatible
	github.com/steakknife/hamming v0.0.0-20180906055917-c99c65617cd3
	github.com/stretchr/testify v1.6.1
	github.com/tidwall/pretty v1.0.2
	github.com/tjfoc/gmsm v1.4.0
	github.com/ulikunitz/xz v0.5.10
	github.com/xaionaro-go/gosrc v0.0.0-20201124181305-3fdf8476a735
	github.com/xaionaro-go/unsafetools v0.0.0-20200202162159-021b112c4d30
	golang.org/x/crypto v0.0.0-20201117144127-c1f2f97bffc9
)

require github.com/9elements/converged-security-suite/v2 v2.6.0

replace github.com/9elements/converged-security-suite/v2 => ./

replace github.com/linuxboot/cbfs => github.com/orangecms/cbfs v0.0.0-20211130150506-eec4ac140592
