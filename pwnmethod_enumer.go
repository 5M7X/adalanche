// Code generated by "enumer -type=PwnMethod -trimprefix=Pwn -json"; DO NOT EDIT.

package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

const _PwnMethodName = "CreateUserCreateGroupCreateComputerCreateAnyObjectDeleteChildrenTargetDeleteObjectInheritsSecurityACLContainsDenyResetPasswordOwnsGenericAllWriteAllWritePropertyAllWriteExtendedAllTakeOwnershipWriteDACLWriteSPNWriteValidatedSPNWriteAllowedToActAddMemberAddMemberGroupAttrAddSelfMemberReadMSAPasswordHasMSAWriteKeyCredentialLinkWriteAttributeSecurityGUIDSIDHistoryEqualityAllExtendedRightsDCReplicationGetChangesDCReplicationSyncronizeDSReplicationGetChangesAllReadLAPSPasswordMemberOfGroupHasSPNHasSPNNoPreauthAdminSDHolderOverwriteACLComputerAffectedByGPOGPOMachineConfigPartOfGPOGPOUserConfigPartOfGPOLocalAdminRightsLocalRDPRightsLocalDCOMRightsScheduledTaskOnUNCPathPwdMachineScriptWriteAltSecurityIdentitiesWriteProfilePathWriteScriptPath"
const _PwnMethodLowerName = "createusercreategroupcreatecomputercreateanyobjectdeletechildrentargetdeleteobjectinheritssecurityaclcontainsdenyresetpasswordownsgenericallwriteallwritepropertyallwriteextendedalltakeownershipwritedaclwritespnwritevalidatedspnwriteallowedtoactaddmemberaddmembergroupattraddselfmemberreadmsapasswordhasmsawritekeycredentiallinkwriteattributesecurityguidsidhistoryequalityallextendedrightsdcreplicationgetchangesdcreplicationsyncronizedsreplicationgetchangesallreadlapspasswordmemberofgrouphasspnhasspnnopreauthadminsdholderoverwriteaclcomputeraffectedbygpogpomachineconfigpartofgpogpouserconfigpartofgpolocaladminrightslocalrdprightslocaldcomrightsscheduledtaskonuncpathpwdmachinescriptwritealtsecurityidentitieswriteprofilepathwritescriptpath"

var _PwnMethodMap = map[PwnMethod]string{
	2:               _PwnMethodName[0:10],
	4:               _PwnMethodName[10:21],
	8:               _PwnMethodName[21:35],
	16:              _PwnMethodName[35:50],
	32:              _PwnMethodName[50:70],
	64:              _PwnMethodName[70:82],
	128:             _PwnMethodName[82:98],
	256:             _PwnMethodName[98:113],
	512:             _PwnMethodName[113:126],
	1024:            _PwnMethodName[126:130],
	2048:            _PwnMethodName[130:140],
	4096:            _PwnMethodName[140:148],
	8192:            _PwnMethodName[148:164],
	16384:           _PwnMethodName[164:180],
	32768:           _PwnMethodName[180:193],
	65536:           _PwnMethodName[193:202],
	131072:          _PwnMethodName[202:210],
	262144:          _PwnMethodName[210:227],
	524288:          _PwnMethodName[227:244],
	1048576:         _PwnMethodName[244:253],
	2097152:         _PwnMethodName[253:271],
	4194304:         _PwnMethodName[271:284],
	8388608:         _PwnMethodName[284:299],
	16777216:        _PwnMethodName[299:305],
	33554432:        _PwnMethodName[305:327],
	67108864:        _PwnMethodName[327:353],
	134217728:       _PwnMethodName[353:371],
	268435456:       _PwnMethodName[371:388],
	536870912:       _PwnMethodName[388:411],
	1073741824:      _PwnMethodName[411:434],
	2147483648:      _PwnMethodName[434:460],
	4294967296:      _PwnMethodName[460:476],
	8589934592:      _PwnMethodName[476:489],
	17179869184:     _PwnMethodName[489:495],
	34359738368:     _PwnMethodName[495:510],
	68719476736:     _PwnMethodName[510:535],
	137438953472:    _PwnMethodName[535:556],
	274877906944:    _PwnMethodName[556:581],
	549755813888:    _PwnMethodName[581:603],
	1099511627776:   _PwnMethodName[603:619],
	2199023255552:   _PwnMethodName[619:633],
	4398046511104:   _PwnMethodName[633:648],
	8796093022208:   _PwnMethodName[648:670],
	17592186044416:  _PwnMethodName[670:686],
	35184372088832:  _PwnMethodName[686:712],
	70368744177664:  _PwnMethodName[712:728],
	140737488355328: _PwnMethodName[728:743],
}

func (i PwnMethod) String() string {
	if str, ok := _PwnMethodMap[i]; ok {
		return str
	}
	return fmt.Sprintf("PwnMethod(%d)", i)
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _PwnMethodNoOp() {
	var x [1]struct{}
	_ = x[PwnCreateUser-(2)]
	_ = x[PwnCreateGroup-(4)]
	_ = x[PwnCreateComputer-(8)]
	_ = x[PwnCreateAnyObject-(16)]
	_ = x[PwnDeleteChildrenTarget-(32)]
	_ = x[PwnDeleteObject-(64)]
	_ = x[PwnInheritsSecurity-(128)]
	_ = x[PwnACLContainsDeny-(256)]
	_ = x[PwnResetPassword-(512)]
	_ = x[PwnOwns-(1024)]
	_ = x[PwnGenericAll-(2048)]
	_ = x[PwnWriteAll-(4096)]
	_ = x[PwnWritePropertyAll-(8192)]
	_ = x[PwnWriteExtendedAll-(16384)]
	_ = x[PwnTakeOwnership-(32768)]
	_ = x[PwnWriteDACL-(65536)]
	_ = x[PwnWriteSPN-(131072)]
	_ = x[PwnWriteValidatedSPN-(262144)]
	_ = x[PwnWriteAllowedToAct-(524288)]
	_ = x[PwnAddMember-(1048576)]
	_ = x[PwnAddMemberGroupAttr-(2097152)]
	_ = x[PwnAddSelfMember-(4194304)]
	_ = x[PwnReadMSAPassword-(8388608)]
	_ = x[PwnHasMSA-(16777216)]
	_ = x[PwnWriteKeyCredentialLink-(33554432)]
	_ = x[PwnWriteAttributeSecurityGUID-(67108864)]
	_ = x[PwnSIDHistoryEquality-(134217728)]
	_ = x[PwnAllExtendedRights-(268435456)]
	_ = x[PwnDCReplicationGetChanges-(536870912)]
	_ = x[PwnDCReplicationSyncronize-(1073741824)]
	_ = x[PwnDSReplicationGetChangesAll-(2147483648)]
	_ = x[PwnReadLAPSPassword-(4294967296)]
	_ = x[PwnMemberOfGroup-(8589934592)]
	_ = x[PwnHasSPN-(17179869184)]
	_ = x[PwnHasSPNNoPreauth-(34359738368)]
	_ = x[PwnAdminSDHolderOverwriteACL-(68719476736)]
	_ = x[PwnComputerAffectedByGPO-(137438953472)]
	_ = x[PwnGPOMachineConfigPartOfGPO-(274877906944)]
	_ = x[PwnGPOUserConfigPartOfGPO-(549755813888)]
	_ = x[PwnLocalAdminRights-(1099511627776)]
	_ = x[PwnLocalRDPRights-(2199023255552)]
	_ = x[PwnLocalDCOMRights-(4398046511104)]
	_ = x[PwnScheduledTaskOnUNCPath-(8796093022208)]
	_ = x[PwdMachineScript-(17592186044416)]
	_ = x[PwnWriteAltSecurityIdentities-(35184372088832)]
	_ = x[PwnWriteProfilePath-(70368744177664)]
	_ = x[PwnWriteScriptPath-(140737488355328)]
}

var _PwnMethodValues = []PwnMethod{PwnCreateUser, PwnCreateGroup, PwnCreateComputer, PwnCreateAnyObject, PwnDeleteChildrenTarget, PwnDeleteObject, PwnInheritsSecurity, PwnACLContainsDeny, PwnResetPassword, PwnOwns, PwnGenericAll, PwnWriteAll, PwnWritePropertyAll, PwnWriteExtendedAll, PwnTakeOwnership, PwnWriteDACL, PwnWriteSPN, PwnWriteValidatedSPN, PwnWriteAllowedToAct, PwnAddMember, PwnAddMemberGroupAttr, PwnAddSelfMember, PwnReadMSAPassword, PwnHasMSA, PwnWriteKeyCredentialLink, PwnWriteAttributeSecurityGUID, PwnSIDHistoryEquality, PwnAllExtendedRights, PwnDCReplicationGetChanges, PwnDCReplicationSyncronize, PwnDSReplicationGetChangesAll, PwnReadLAPSPassword, PwnMemberOfGroup, PwnHasSPN, PwnHasSPNNoPreauth, PwnAdminSDHolderOverwriteACL, PwnComputerAffectedByGPO, PwnGPOMachineConfigPartOfGPO, PwnGPOUserConfigPartOfGPO, PwnLocalAdminRights, PwnLocalRDPRights, PwnLocalDCOMRights, PwnScheduledTaskOnUNCPath, PwdMachineScript, PwnWriteAltSecurityIdentities, PwnWriteProfilePath, PwnWriteScriptPath}

var _PwnMethodNameToValueMap = map[string]PwnMethod{
	_PwnMethodName[0:10]:         PwnCreateUser,
	_PwnMethodLowerName[0:10]:    PwnCreateUser,
	_PwnMethodName[10:21]:        PwnCreateGroup,
	_PwnMethodLowerName[10:21]:   PwnCreateGroup,
	_PwnMethodName[21:35]:        PwnCreateComputer,
	_PwnMethodLowerName[21:35]:   PwnCreateComputer,
	_PwnMethodName[35:50]:        PwnCreateAnyObject,
	_PwnMethodLowerName[35:50]:   PwnCreateAnyObject,
	_PwnMethodName[50:70]:        PwnDeleteChildrenTarget,
	_PwnMethodLowerName[50:70]:   PwnDeleteChildrenTarget,
	_PwnMethodName[70:82]:        PwnDeleteObject,
	_PwnMethodLowerName[70:82]:   PwnDeleteObject,
	_PwnMethodName[82:98]:        PwnInheritsSecurity,
	_PwnMethodLowerName[82:98]:   PwnInheritsSecurity,
	_PwnMethodName[98:113]:       PwnACLContainsDeny,
	_PwnMethodLowerName[98:113]:  PwnACLContainsDeny,
	_PwnMethodName[113:126]:      PwnResetPassword,
	_PwnMethodLowerName[113:126]: PwnResetPassword,
	_PwnMethodName[126:130]:      PwnOwns,
	_PwnMethodLowerName[126:130]: PwnOwns,
	_PwnMethodName[130:140]:      PwnGenericAll,
	_PwnMethodLowerName[130:140]: PwnGenericAll,
	_PwnMethodName[140:148]:      PwnWriteAll,
	_PwnMethodLowerName[140:148]: PwnWriteAll,
	_PwnMethodName[148:164]:      PwnWritePropertyAll,
	_PwnMethodLowerName[148:164]: PwnWritePropertyAll,
	_PwnMethodName[164:180]:      PwnWriteExtendedAll,
	_PwnMethodLowerName[164:180]: PwnWriteExtendedAll,
	_PwnMethodName[180:193]:      PwnTakeOwnership,
	_PwnMethodLowerName[180:193]: PwnTakeOwnership,
	_PwnMethodName[193:202]:      PwnWriteDACL,
	_PwnMethodLowerName[193:202]: PwnWriteDACL,
	_PwnMethodName[202:210]:      PwnWriteSPN,
	_PwnMethodLowerName[202:210]: PwnWriteSPN,
	_PwnMethodName[210:227]:      PwnWriteValidatedSPN,
	_PwnMethodLowerName[210:227]: PwnWriteValidatedSPN,
	_PwnMethodName[227:244]:      PwnWriteAllowedToAct,
	_PwnMethodLowerName[227:244]: PwnWriteAllowedToAct,
	_PwnMethodName[244:253]:      PwnAddMember,
	_PwnMethodLowerName[244:253]: PwnAddMember,
	_PwnMethodName[253:271]:      PwnAddMemberGroupAttr,
	_PwnMethodLowerName[253:271]: PwnAddMemberGroupAttr,
	_PwnMethodName[271:284]:      PwnAddSelfMember,
	_PwnMethodLowerName[271:284]: PwnAddSelfMember,
	_PwnMethodName[284:299]:      PwnReadMSAPassword,
	_PwnMethodLowerName[284:299]: PwnReadMSAPassword,
	_PwnMethodName[299:305]:      PwnHasMSA,
	_PwnMethodLowerName[299:305]: PwnHasMSA,
	_PwnMethodName[305:327]:      PwnWriteKeyCredentialLink,
	_PwnMethodLowerName[305:327]: PwnWriteKeyCredentialLink,
	_PwnMethodName[327:353]:      PwnWriteAttributeSecurityGUID,
	_PwnMethodLowerName[327:353]: PwnWriteAttributeSecurityGUID,
	_PwnMethodName[353:371]:      PwnSIDHistoryEquality,
	_PwnMethodLowerName[353:371]: PwnSIDHistoryEquality,
	_PwnMethodName[371:388]:      PwnAllExtendedRights,
	_PwnMethodLowerName[371:388]: PwnAllExtendedRights,
	_PwnMethodName[388:411]:      PwnDCReplicationGetChanges,
	_PwnMethodLowerName[388:411]: PwnDCReplicationGetChanges,
	_PwnMethodName[411:434]:      PwnDCReplicationSyncronize,
	_PwnMethodLowerName[411:434]: PwnDCReplicationSyncronize,
	_PwnMethodName[434:460]:      PwnDSReplicationGetChangesAll,
	_PwnMethodLowerName[434:460]: PwnDSReplicationGetChangesAll,
	_PwnMethodName[460:476]:      PwnReadLAPSPassword,
	_PwnMethodLowerName[460:476]: PwnReadLAPSPassword,
	_PwnMethodName[476:489]:      PwnMemberOfGroup,
	_PwnMethodLowerName[476:489]: PwnMemberOfGroup,
	_PwnMethodName[489:495]:      PwnHasSPN,
	_PwnMethodLowerName[489:495]: PwnHasSPN,
	_PwnMethodName[495:510]:      PwnHasSPNNoPreauth,
	_PwnMethodLowerName[495:510]: PwnHasSPNNoPreauth,
	_PwnMethodName[510:535]:      PwnAdminSDHolderOverwriteACL,
	_PwnMethodLowerName[510:535]: PwnAdminSDHolderOverwriteACL,
	_PwnMethodName[535:556]:      PwnComputerAffectedByGPO,
	_PwnMethodLowerName[535:556]: PwnComputerAffectedByGPO,
	_PwnMethodName[556:581]:      PwnGPOMachineConfigPartOfGPO,
	_PwnMethodLowerName[556:581]: PwnGPOMachineConfigPartOfGPO,
	_PwnMethodName[581:603]:      PwnGPOUserConfigPartOfGPO,
	_PwnMethodLowerName[581:603]: PwnGPOUserConfigPartOfGPO,
	_PwnMethodName[603:619]:      PwnLocalAdminRights,
	_PwnMethodLowerName[603:619]: PwnLocalAdminRights,
	_PwnMethodName[619:633]:      PwnLocalRDPRights,
	_PwnMethodLowerName[619:633]: PwnLocalRDPRights,
	_PwnMethodName[633:648]:      PwnLocalDCOMRights,
	_PwnMethodLowerName[633:648]: PwnLocalDCOMRights,
	_PwnMethodName[648:670]:      PwnScheduledTaskOnUNCPath,
	_PwnMethodLowerName[648:670]: PwnScheduledTaskOnUNCPath,
	_PwnMethodName[670:686]:      PwdMachineScript,
	_PwnMethodLowerName[670:686]: PwdMachineScript,
	_PwnMethodName[686:712]:      PwnWriteAltSecurityIdentities,
	_PwnMethodLowerName[686:712]: PwnWriteAltSecurityIdentities,
	_PwnMethodName[712:728]:      PwnWriteProfilePath,
	_PwnMethodLowerName[712:728]: PwnWriteProfilePath,
	_PwnMethodName[728:743]:      PwnWriteScriptPath,
	_PwnMethodLowerName[728:743]: PwnWriteScriptPath,
}

var _PwnMethodNames = []string{
	_PwnMethodName[0:10],
	_PwnMethodName[10:21],
	_PwnMethodName[21:35],
	_PwnMethodName[35:50],
	_PwnMethodName[50:70],
	_PwnMethodName[70:82],
	_PwnMethodName[82:98],
	_PwnMethodName[98:113],
	_PwnMethodName[113:126],
	_PwnMethodName[126:130],
	_PwnMethodName[130:140],
	_PwnMethodName[140:148],
	_PwnMethodName[148:164],
	_PwnMethodName[164:180],
	_PwnMethodName[180:193],
	_PwnMethodName[193:202],
	_PwnMethodName[202:210],
	_PwnMethodName[210:227],
	_PwnMethodName[227:244],
	_PwnMethodName[244:253],
	_PwnMethodName[253:271],
	_PwnMethodName[271:284],
	_PwnMethodName[284:299],
	_PwnMethodName[299:305],
	_PwnMethodName[305:327],
	_PwnMethodName[327:353],
	_PwnMethodName[353:371],
	_PwnMethodName[371:388],
	_PwnMethodName[388:411],
	_PwnMethodName[411:434],
	_PwnMethodName[434:460],
	_PwnMethodName[460:476],
	_PwnMethodName[476:489],
	_PwnMethodName[489:495],
	_PwnMethodName[495:510],
	_PwnMethodName[510:535],
	_PwnMethodName[535:556],
	_PwnMethodName[556:581],
	_PwnMethodName[581:603],
	_PwnMethodName[603:619],
	_PwnMethodName[619:633],
	_PwnMethodName[633:648],
	_PwnMethodName[648:670],
	_PwnMethodName[670:686],
	_PwnMethodName[686:712],
	_PwnMethodName[712:728],
	_PwnMethodName[728:743],
}

// PwnMethodString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func PwnMethodString(s string) (PwnMethod, error) {
	if val, ok := _PwnMethodNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _PwnMethodNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to PwnMethod values", s)
}

// PwnMethodValues returns all values of the enum
func PwnMethodValues() []PwnMethod {
	return _PwnMethodValues
}

// PwnMethodStrings returns a slice of all String values of the enum
func PwnMethodStrings() []string {
	strs := make([]string, len(_PwnMethodNames))
	copy(strs, _PwnMethodNames)
	return strs
}

// IsAPwnMethod returns "true" if the value is listed in the enum definition. "false" otherwise
func (i PwnMethod) IsAPwnMethod() bool {
	_, ok := _PwnMethodMap[i]
	return ok
}

// MarshalJSON implements the json.Marshaler interface for PwnMethod
func (i PwnMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for PwnMethod
func (i *PwnMethod) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("PwnMethod should be a string, got %s", data)
	}

	var err error
	*i, err = PwnMethodString(s)
	return err
}
