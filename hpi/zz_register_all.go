package hpi

import "fmt"

// process is function to call transfrom function by each type
// If you new transform function please register here function
// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
func process(t *Transform, tConfig interface{}) (err error) {

	// init
	tConfigMap, ok := tConfig.(map[string]interface{})
	if ok {

		// type of transform
		tfType := tConfigMap["type"].(string)

		switch tfType {

		// Register to call function by function type
		case HPI_TType_MergeString:
			_, err = t.HpiMergeString(tConfig)
			return err
		case HPI_TType_SplitString:
			_, err = t.HpiSplitString(tConfig)
			return err
		case HPI_TType_SplitStringArray:
			_, err = t.HpiSplitStringArray(tConfig)
			return err
		case HPI_TType_ReplaceString:
			_, err = t.HpiReplaceString(tConfig)
			return err
		case HPI_TType_TrimString:
			_, err = t.HpiTrimString(tConfig)
			return err
		case HPI_TType_GenUUID:
			_, err = t.HpiGenUUID(tConfig)
			return err
		case HPI_TType_GenUniqueId:
			_, err = t.HpiGenUniqueID(tConfig)
			return err
		case HPI_TType_RandomInteger:
			_, err = t.HpiRandomInteger(tConfig)
			return err
		case HPI_TType_RandomAlphabet:
			_, err = t.HpiRandomAlphabet(tConfig)
			return err
		case HPI_TType_GetCurrentTime:
			_, err = t.HpiGetCurrentTime(tConfig)
			return err
		case HPI_TType_RenameFields:
			_, err = t.HpiRenameFields(tConfig)
			return err
		case HPI_TType_DateFormat:
			_, err = t.HpiDateFormat(tConfig)
			return err
		case HPI_TType_NumberFormat:
			_, err = t.HpiNumberFormat(tConfig)
			return err
		case HPI_TType_NameFormat:
			_, err = t.HpiNameFormat(tConfig)
			return err
		case HPI_TType_IncreaseNum:
			_, err = t.HpiIncreaseNumber(tConfig)
			return err
		case HPI_TType_DecreaseNum:
			_, err = t.HpiDecreaseNumber(tConfig)
			return err
		case HPI_TType_ToNumber:
			_, err = t.HpiToNumber(tConfig)
			return err
		case HPI_TType_ToBoolean:
			_, err = t.HpiToBoolean(tConfig)
			return err
		case HPI_TType_FixValue:
			_, err = t.HpiFixValue(tConfig)
			return err
		case HPI_TType_DeleteField:
			_, err = t.HpiDeleteField(tConfig)
			return err
		case HPI_TType_GetKeysInMap:
			_, err = t.HpiGetKeysInMap(tConfig)
			return err
		case HPI_TType_GetValuesInMap:
			_, err = t.HpiGetValuesInMap(tConfig)
			return err
		case HPI_TType_CheckValue:
			_, err = t.HpiCheckValue(tConfig)
			return err
		case HPI_TType_TimeStampToDate:
			_, err = t.HpiTimeStampToDate(tConfig)
			return err
		case HPI_TType_DateToTimeStamp:
			_, err = t.HpiDateToTimeStamp(tConfig)
			return err
		case HPI_TType_GetCurrentTimeStamp:
			_, err = t.HpiGetCurrentTimeStamp(tConfig)
			return err

		case HPI_TType_ListSize:
			_, err = t.HpiListSize(tConfig)
			return err
		case HPI_TType_AddToList:
			_, err = t.HpiAddToList(tConfig)
			return err
		case HPI_TType_AddObjectToList:
			_, err = t.HpiAddObjectToList(tConfig)
			return err
		case HPI_TType_UpdateRecord:
			_, err = t.HpiUpdateRecord(tConfig)
			return err
		case HPI_TType_CopyToList:
			_, err = t.HpiCopyToList(tConfig)
			return err
		case HPI_TType_CopyFromList:
			_, err = t.HpiCopyFromList(tConfig)
			return err
		case HPI_TType_CopyFromListToObject:
			_, err = t.HpiCopyFromListToObject(tConfig)
			return err
		case HPI_TType_JoinList:
			_, err = t.HpiJoinList(tConfig)
			return err
		case HPI_TType_DeleteRecord:
			_, err = t.HpiDeleteRecord(tConfig)
			return err
		case HPI_TType_FilterRecord:
			_, err = t.HpiFilterRecord(tConfig)
			return err
		case HPI_TType_FindOneRecord:
			_, err = t.HpiFindOneRecord(tConfig)
			return err
		case HPI_TType_SumFieldsInList:
			_, err = t.HpiSumFieldsInList(tConfig)
			return err
		case HPI_TType_SplitListToObject:
			_, err = t.HpiSplitListToObject(tConfig)
			return err
		case HPI_TType_SearchValueInList:
			_, err = t.HpiSearchValueInList(tConfig)
			return err
		case HPI_TType_CalculateNumber:
			_, err = t.HpiCalculateNumber(tConfig)
			return err
		case HPI_TType_MathUtil:
			_, err = t.HpiMathUtil(tConfig)
			return err
		case HPI_TType_CompareNumber:
			_, err = t.HpiCompareNumber(tConfig)
			return err
		case HPI_TType_CompareString:
			_, err = t.HpiCompareString(tConfig)
			return err
		case HPI_TType_CompareDateTime:
			_, err = t.HpiCompareDateTime(tConfig)
			return err
		case HPI_TType_CopyObjectToSameParent:
			_, err = t.HpiCopyObjectToSameParent(tConfig)
			return err
		case HPI_TType_ValidateAndCopy:
			_, err = t.HpiValidateAndCopy(tConfig)
			return err
		case HPI_TType_SortList:
			_, err = t.HpiSortList(tConfig)
			return err
		case HPI_TType_CheckExistingField:
			_, err = t.HpiCheckExistingField(tConfig)
			return err

		case HPI_TType_Base64Encode:
			_, err = t.HpiBase64Encode(tConfig)
			return err
		case HPI_TType_Base64Decode:
			_, err = t.HpiBase64Decode(tConfig)
			return err
		case HPI_TType_Aes256GcmEncode:
			_, err = t.HpiAes256GcmEncode(tConfig)
			return err
		case HPI_TType_Aes256GcmDecode:
			_, err = t.HpiAes256GcmDecode(tConfig)
			return err

		case HPI_TType_FindValuesByRegExp:
			_, err = t.HpiFindValuesByRegExp(tConfig)
			return err
		case HPI_TType_ModifyString:
			_, err = t.HpiModifyString(tConfig)
			return err

		case HPI_TType_StringToJson:
			_, err = t.HpiStringToJson(tConfig)
			return err
		case HPI_TType_JsonToString:
			_, err = t.HpiJsonToString(tConfig)
			return err

		case HPI_TType_SsqToDma:
			_, err = t.HpiSsqToDmaDataFormat(tConfig)
			return err
		case HPI_TType_DmaToSsq:
			_, err = t.HpiDmaToSsqDataFormat(tConfig)
			return err

		}

		// Return error if transform type is missing
		return fmt.Errorf("tranfrom type='" + tfType + "' is missing")

	}

	return
}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
