/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "E2AP-IEs"
 * 	found in "../../../../api/e2ap/v2beta1/e2ap_v2.0.1_Feb3_21.asn"
 * 	`asn1c -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER -no-gen-OER -D.`
 */

#ifndef	_E2nodeComponentNGeNB_DU_ID_H_
#define	_E2nodeComponentNGeNB_DU_ID_H_


#include "asn_application.h"

/* Including external dependencies */
#include "NGENB-DU-ID.h"
#include "constr_SEQUENCE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* E2nodeComponentNGeNB-DU-ID */
typedef struct E2nodeComponentNGeNB_DU_ID {
	NGENB_DU_ID_t	 ngENB_DU_ID;
	/*
	 * This type is extensible,
	 * possible extensions are below.
	 */
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} E2nodeComponentNGeNB_DU_ID_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_E2nodeComponentNGeNB_DU_ID;
extern asn_SEQUENCE_specifics_t asn_SPC_E2nodeComponentNGeNB_DU_ID_specs_1;
extern asn_TYPE_member_t asn_MBR_E2nodeComponentNGeNB_DU_ID_1[1];

#ifdef __cplusplus
}
#endif

#endif	/* _E2nodeComponentNGeNB_DU_ID_H_ */
#include "asn_internal.h"
