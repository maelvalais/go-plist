package plist

import (
	"strings"
	"testing"
)

var InvalidBplists []string = []string{
	"bplist00\xdc\x01\x02\x03\x04\x05\x06\a\b\t\n\v\f\r\x10\x11\x12\x13\x14\x15\x16\x19\x15#$_\x10\x11CFPlugInFactories_\x10\x12CFBundleIdentifier_\x10\x1bCFPlugInDynamicRegistration_\x10\x1dCFBundleInfoDictionaryVersion_\x10\x0fCFBundleVersion_\x10\x12CFBundleExecutable_\x10\x16CFPlugInUnloadFunction]CFPlugInTypes_\x10\x15CFBundleDocumentTypes_\x10\x1fCFPlugInDynamicRegisterFunction_\x10\x19CFBundleDevelopmentRegion\\CFBundleName\xd1\x0e\x0f_\x10$9B6E3B91-772C-4DA9-9EFC-06B51846737E_\x10\x1dMetadataImporterPluginFactory_\x10\x19com.adiumX.adiumXimporterRNOS6.0S1.0_\x10\x16AdiumSpotlightImporterP\xd1\x17\x18_\x10$8B08C4BF-415B-11D8-B3F9-0003936726FC\xa1\x0e\xa1\x1a\xd3\x1b\x1c\x1d\x1e\x1f _\x10\x10CFBundleTypeRole_\x10\x0fLSTypeIsPackage_\x10\x12LSItemContentTypesZMDImporter\t\xa2!\"_\x10\x12com.adiumx.htmllog_\x10\x11com.adiumx.xmllogWEnglish_\x10\x16AdiumSpotlightImporter\x00\b\x00J\x00h\x00\x88\x00\x9a\x00\xaf\x00\xc8\x00\xd6\x00\xee\x01\x10\x01\x019\x01<\x01c\x01\x83\x01\x9f\x01\xa2\x01\xa6\x01\xaa\x01\xc3\x01\xc7\x01\xee\x01\x01\xf2\x01\xf9\x02\f\x02\x1e\x023>\x02\x01\x00\x00\x00\x00\x00\x00\x00%\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x02\x8c",
	"bplist00\xdc\x01\x02\x03\x04\x05\x06\a\b\t\n\v\f\r\x10\x11\x12\x13\x14\x15\x16\x19\x15#$_\x10\x11CFPlugInFactories_\x10\x12CFBundleIdentifier_\x10\x1bCFPlugInDynamicRegistration_\x10\x1dCFBundleInfoDictionaryVersion_\x10\x0fCFBundleVersion_\x10\x12CFBundleExecutable_\x10\x16CFPlugInUnloadFunction]CFPlugInTypes_\x10\x15CFBundleDocumentTypes_\x10\x1fCFPlugInDynamicRegisterFunction_\x10\x19CFBundleDevelopmentRegion\\CFBundleName\xd1\x0e\x0f_\x10$9B6E3B91-772C-4DA9-9EFC-06B51846737E_\x10\x1dMetadataImporterPluginFactory_\x10\x19com.adiumX.adiumXimporterRNOS6.0S1.0_\x10\x16AdiumSpotlightImporterP\xd1\x17\x18_\x10$8B08C4BF-415B-11D8-B3F9-0003936726FC\xa1\x0e\xa1\x1a\xd3\x1b\x1c\x1d\x1e\x1f _\x10\x10CFBundleTypeRole_\x10\x0fLSTypeIsPackage_\x10\x12LSItemContentTypesZMDImporter\t\xa2!\"_\x10\x12com.adiumx.htmllog_\x10\x11com.adiumx.xmllogWEnglish_\x10\x16AdiumSpotlightImporter\x00\b\x00!\x005\x00J\x00h\x00\x88\x00\x9a\x00\xaf\x00\xc8\x00\xd6\x00\xee\x01\x10\x01\x019\x01<\x01c\x01\x83\x01\x9f\x01\xa2\x01\xa6\x01\xaa\x01\xc3\x01\xc7\x01\xee\x01\xf0\x02\x01\x00\x00\x00\x00\x00\x00\x00%\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x02\x8c",
	"bplist00\xdc\x01\x02PlugInDynamicRegistration_\x10\x1dCFBundleInfoDictionaryVersion_\x10\x0fCFBundleVersion_\x10\x12CFBundleExecutable_\x10\x16CFPlugInUnloadFunction]CFPlugInTypes_\x10\x15CFBundleDocumentTypes_\x10\x1fCFPlugInDynamicRegisterFunction_\x10\x19CFBundleDevelopmentRegion\\CFBundleName\xd1\x0e\x0f_\x10$9B6E3B91-772C-4DA9-9EFC-06B51846737E_\x10\x1dMetadataImporterPluginFactory_\x10\x19com.adiumX.adiumXimporterRNOS6.0S1.0_\x10\x16AdiumSpotlightImporterP\xd1\x17\x18_\x10$8B08C4BF-415B-11D8-B3F9-0003936726FC\xa1\x0e\xa1\x1a\xd3\x1b\x1c\x1d\x1e\x1f _\x10\x10CFBundleTypeRole_\x10\x0fLSTypeIsPackage_\x10\x12LSItemContentTypesZMDImporter\t\xa2!\"_\x10\x12com.adiumx.htmllog_\x10\x11com.adiumx.xmllogWEnglish_\x10\x16AdiumSpotlightImporter\x00\b\x00!\x005\x00J\x00h\x00\x88\x00\x9a\x00\xaf\x00\xc8\x00\xd6\x00\xee\x01\x10\x01,\x019\x01<\x01c\x01\x83\x01\x9f\x01\xa2\x01\xa6\x01\xaa\x01\xc3\x01\xc4\x01\xc7\x01\xee\x01\xf0\x01\xf2\x01\xf9\x02\f\x02\x1e\x023\x02>\x02?\x02B\x02W\x02k\x02s\x02\x01\x00\x00\x00\x00\x00\x00\x00%\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x02\x8c",
	"bplist00000000000000000000000000",
	"bplist00\xdc\r\x10\x11\x12\x13\x14\x15\x16\x19\x15#$_\x10\x11CFPlugInFactories_\x10\x12CFBundleIdentifier_\x10\x1bCFPlugInDynamicRegistration_\x10\x1dCFBundleInfoDictionaryVersion_\x10\x0fCFBundleVersion_\x10\x12CFBundleExecutable_\x10\x16CFPlugInUnloadFunction]CFPlugInTypes_\x10\x15CFBundleDocumentTypes_\x10\x1fCFPlugInDynamicRegisterFunction_\x10\x19CFBundleDevelopmentRegion\\CFBundleName\xd1\x0e\x0f_\x10$9B6E3B91-772C-4DA9-9EFC-06B51846737E_\x10\x1dMetadataImporterPluginFactory_\x10\x19com.adiumX.adiumXimporterRNOS6.0S1.0_\x10\x16AdiumSpotlightImporterP\xd1\x17\x18_\x10$8B08C4BF-415B-11D8-B3F9-0003936726FC\xa1\x0e\xa1\x1a\xd3\x1b\x1c\x1d\x1e\x1f _\x10\x10CFBundleTypeRole_\x10\x0fLSTypeIsPackage_\x10\x12LSItemContentTypesZMDImporter\t\xa2!\"_\x10\x12com.adiumx.htmllog_\x10\x11com.adiumx.xmllogWEnglish_\x10\x16AdiumSpotlightImporter\x00\b\x00J\x00h\x00\x88\x00\x9a\x00\xaf\x00\xc8\x00\xd6\x02\x01\x00\x00\x00\x00\x00\x00\x00%\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x02\x8c",
	"bplist00\xdc\x01\x02\x03\x04\x05\x06\a\b\t\n\v\f\r\x10\x11\x12\x13\x14\x15\x16\x19\x15#$_\x10\x11CFPlugInFactories_\x10\x12CFBundleIdentifier_\x10\x1bCFPlugInDynamicRegistration_\x10\x1dCFBundleInfoDictionaryVersion_\x10\x0fCFBundleVersion_\x10\x12CFBundleExecutable_\x10\x16CFPlugInUnloadFunction]CFPlugInTypes_\x10\x15CFBundleDocumentTypes_\x10\x1fCFPlugInDynamicRegisterFunction_\x10\x19CFBundleDevelopmentRegion\\CFBundleName\xd1\x0e\x0f_\x10$9B6E3B91-772C-4DA9-9EFC-06B51846737E_\x10\x1dMetadataImporterPluginFactory_\x10\x19com.adiumX.adiumXimporterRNOS6.0S1.0_\x10\x16AdiumSpotlightImporterP\xd1\x17\x18_\x10$8B08C4BF-415B-11D8-B3F9-0003936726FC\xa1\x0e\xa1\x1a\xd3\x1b\x1c\x1d\x1e\x1f _\x10\x10CFBundleTypeRole_\x10\x0fLSTypeIsPackage_\x10\x12LSItemContentTypesZMDImporter\t\xa2!\"_\x10\x12com.adiumx.htmllog_\x10\x11com.adiumx.xmllogWEnglish_\x10\x16AdiumSpotlightImporter\x00\b\x00!\x005\x00J\x00h\x00\x88\x00\x9a\x00\xaf\x00\xc8\x00\xd6\x00\xee\x01\x10\x01\x019\x01<\x01c\x01\x83\x01\x9f\x01\xa2\x01\xa6\x01\xaa\x01\xc3\x01\xc7\x01\xee\x01\x01\xf2\x01\xf9\x02\f\x02\x1e\x023\x02>\x02\x01\x00\x00\x00\x00\x00\x00\x00%\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x02\x8c",
	"bplist00\xdc\x01\x02\x03\x04\x05\x06\a\b\t\n\v\f\r\x10\x11\x12\x13\x14\x15\x16\x19\x15#$_\x10\x11CFPlugInFactories_\x10\x12CFBundleIdentifier_\x10\x1bCFPlugInDynamicRegistration_\x10\x1dCFBundleInfoDictionaryVersion_\x10\x0fCFBundleVersion_\x10\x12CFBundleExecutable_\x10\x16CFPlugInUnloadFunction]CFPlugInTypes_\x10\x15CFBundleDocumentTypes_\x10\x1fCFPlugInDynamicRegisterFunction_\x10\x19CFBundleDevelopmentRegion\\CFBundleName\xd1\x0e\x0f_\x10$9B6E3B91-772C-4DA9-9EFC-06B51846737E_\x10\x1dMetadataImporterPluginFactory_\x10\x19com.adiumX.adiumXimporterRNOS6.0S1.0_\x10\x16AdiumSpotlightImporterP\xd1\x17\x18_\x10$8B08C4BF-415B-11D8-B3F9-0003936726FC\xa1\x0e\xa1\x1a\xd3\x1b\x1c\x1d\x1e\x1f _\x10\x10CFBundleTypeRole_\x10\x0fLSTypeIsPackage_\x10\x12LSItemContentTypesZMDImporter\t\xa2!\"_\x10\x12com.adiumx.htmllog_\x10\x11com.adiumx.xmllogWEnglish_\x10\x16AdiumSpotlightImporter\x00\b5\x02\x01\x00\x00\x00\x00\x00\x00\x00%\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x02\x8c",
	"bplist00\xd60000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000\b000\x01\x01\x00\x00\x00\x00\x00\x00\x00\x1d\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\xa6",
	"bplist00\xd6\x0100000\a0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000\xa2\x1b\x1c00000000000000000000000000000000\b\x8300\x01\x01\x00\x00\x00\x00\x00\x00\x00\x1d\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\xa6",
	"bplist00\xd6\x0100000\a000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000\b000\x01\x01\x00\x00\x00\x00\x00\x00\x00\x1d\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\xa6",
	"bplist00\xd5\x01\x02\x03\x04\x05\x06\a\b\t\n000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000\x1300000000\b\xa6\x01\x01\x00\x00\x00\x00\x00\x00\x00\v\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\xaf",
}

func TestInvalidBinaryPlists(t *testing.T) {
	for _, data := range InvalidBplists {
		var obj interface{}
		buf := strings.NewReader(data)
		err := NewDecoder(buf).Decode(&obj)
		if err == nil {
			t.Fatal("invalid plist failed to throw error")
		} else {
			t.Log(err)
		}
	}
}
