/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "E2AP-IEs"
 * 	found in "../../../../api/e2ap/v2beta1/e2ap_v2.0.1_Feb3_21.asn"
 * 	`asn1c -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER -no-gen-OER -D.`
 */

#include "TNLinformation.h"

static int
memb_tnlAddress_constraint_1(const asn_TYPE_descriptor_t *td, const void *sptr,
			asn_app_constraint_failed_f *ctfailcb, void *app_key) {
	const BIT_STRING_t *st = (const BIT_STRING_t *)sptr;
	size_t size;
	
	if(!sptr) {
		ASN__CTFAIL(app_key, td, sptr,
			"%s: value not given (%s:%d)",
			td->name, __FILE__, __LINE__);
		return -1;
	}
	
	if(st->size > 0) {
		/* Size in bits */
		size = 8 * st->size - (st->bits_unused & 0x07);
	} else {
		size = 0;
	}
	
	if((size >= 1 && size <= 160)) {
		/* Constraint check succeeded */
		return 0;
	} else {
		ASN__CTFAIL(app_key, td, sptr,
			"%s: constraint failed (%s:%d)",
			td->name, __FILE__, __LINE__);
		return -1;
	}
}

static int
memb_tnlPort_constraint_1(const asn_TYPE_descriptor_t *td, const void *sptr,
			asn_app_constraint_failed_f *ctfailcb, void *app_key) {
	const BIT_STRING_t *st = (const BIT_STRING_t *)sptr;
	size_t size;
	
	if(!sptr) {
		ASN__CTFAIL(app_key, td, sptr,
			"%s: value not given (%s:%d)",
			td->name, __FILE__, __LINE__);
		return -1;
	}
	
	if(st->size > 0) {
		/* Size in bits */
		size = 8 * st->size - (st->bits_unused & 0x07);
	} else {
		size = 0;
	}
	
	if((size == 16)) {
		/* Constraint check succeeded */
		return 0;
	} else {
		ASN__CTFAIL(app_key, td, sptr,
			"%s: constraint failed (%s:%d)",
			td->name, __FILE__, __LINE__);
		return -1;
	}
}

static asn_per_constraints_t asn_PER_memb_tnlAddress_constr_2 CC_NOTUSED = {
	{ APC_UNCONSTRAINED,	-1, -1,  0,  0 },
	{ APC_CONSTRAINED | APC_EXTENSIBLE,  8,  8,  1,  160 }	/* (SIZE(1..160,...)) */,
	0, 0	/* No PER value map */
};
static asn_per_constraints_t asn_PER_memb_tnlPort_constr_3 CC_NOTUSED = {
	{ APC_UNCONSTRAINED,	-1, -1,  0,  0 },
	{ APC_CONSTRAINED,	 0,  0,  16,  16 }	/* (SIZE(16..16)) */,
	0, 0	/* No PER value map */
};
asn_TYPE_member_t asn_MBR_TNLinformation_1[] = {
	{ ATF_NOFLAGS, 0, offsetof(struct TNLinformation, tnlAddress),
		(ASN_TAG_CLASS_CONTEXT | (0 << 2)),
		-1,	/* IMPLICIT tag at current level */
		&asn_DEF_BIT_STRING,
		0,
		{ 0, &asn_PER_memb_tnlAddress_constr_2,  memb_tnlAddress_constraint_1 },
		0, 0, /* No default value */
		"tnlAddress"
		},
	{ ATF_POINTER, 1, offsetof(struct TNLinformation, tnlPort),
		(ASN_TAG_CLASS_CONTEXT | (1 << 2)),
		-1,	/* IMPLICIT tag at current level */
		&asn_DEF_BIT_STRING,
		0,
		{ 0, &asn_PER_memb_tnlPort_constr_3,  memb_tnlPort_constraint_1 },
		0, 0, /* No default value */
		"tnlPort"
		},
};
static const int asn_MAP_TNLinformation_oms_1[] = { 1 };
static const ber_tlv_tag_t asn_DEF_TNLinformation_tags_1[] = {
	(ASN_TAG_CLASS_UNIVERSAL | (16 << 2))
};
static const asn_TYPE_tag2member_t asn_MAP_TNLinformation_tag2el_1[] = {
    { (ASN_TAG_CLASS_CONTEXT | (0 << 2)), 0, 0, 0 }, /* tnlAddress */
    { (ASN_TAG_CLASS_CONTEXT | (1 << 2)), 1, 0, 0 } /* tnlPort */
};
asn_SEQUENCE_specifics_t asn_SPC_TNLinformation_specs_1 = {
	sizeof(struct TNLinformation),
	offsetof(struct TNLinformation, _asn_ctx),
	asn_MAP_TNLinformation_tag2el_1,
	2,	/* Count of tags in the map */
	asn_MAP_TNLinformation_oms_1,	/* Optional members */
	1, 0,	/* Root/Additions */
	2,	/* First extension addition */
};
asn_TYPE_descriptor_t asn_DEF_TNLinformation = {
	"TNLinformation",
	"TNLinformation",
	&asn_OP_SEQUENCE,
	asn_DEF_TNLinformation_tags_1,
	sizeof(asn_DEF_TNLinformation_tags_1)
		/sizeof(asn_DEF_TNLinformation_tags_1[0]), /* 1 */
	asn_DEF_TNLinformation_tags_1,	/* Same as above */
	sizeof(asn_DEF_TNLinformation_tags_1)
		/sizeof(asn_DEF_TNLinformation_tags_1[0]), /* 1 */
	{ 0, 0, SEQUENCE_constraint },
	asn_MBR_TNLinformation_1,
	2,	/* Elements count */
	&asn_SPC_TNLinformation_specs_1	/* Additional specs */
};

