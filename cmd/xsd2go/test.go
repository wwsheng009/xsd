package main

import "encoding/xml"

//
type GBAltSeqData struct {
	GBAltSeqDataName string `xml:"GBAltSeqData_name" json:"gb_alt_seq_data_name"`

	GBAltSeqDataItems []GBAltSeqItem `xml:"GBAltSeqData_items,omitempty" json:"gb_alt_seq_data_items,omitempty"`
}
type GBAltSeqDataName struct {
	GBAltSeqDataName string `xml:"GBAltSeqData_name" json:"gb_alt_seq_data_name"`
}

//
type GBAltSeqItem struct {
	GBAltSeqItemInterval *GBInterval `xml:"GBAltSeqItem_interval,omitempty" json:"gb_alt_seq_item_interval,omitempty"`

	GBAltSeqItemIsgap *GBAltSeqItemIsgap `xml:"GBAltSeqItem_isgap,omitempty" json:"gb_alt_seq_item_isgap,omitempty"`

	GBAltSeqItemGapLength int `xml:"GBAltSeqItem_gap-length,omitempty" json:"gb_alt_seq_item_gap_length,omitempty"`

	GBAltSeqItemGapType string `xml:"GBAltSeqItem_gap-type,omitempty" json:"gb_alt_seq_item_gap_type,omitempty"`

	GBAltSeqItemGapLinkage string `xml:"GBAltSeqItem_gap-linkage,omitempty" json:"gb_alt_seq_item_gap_linkage,omitempty"`

	GBAltSeqItemGapComment string `xml:"GBAltSeqItem_gap-comment,omitempty" json:"gb_alt_seq_item_gap_comment,omitempty"`

	GBAltSeqItemFirstAccn string `xml:"GBAltSeqItem_first-accn,omitempty" json:"gb_alt_seq_item_first_accn,omitempty"`

	GBAltSeqItemLastAccn string `xml:"GBAltSeqItem_last-accn,omitempty" json:"gb_alt_seq_item_last_accn,omitempty"`

	GBAltSeqItemValue string `xml:"GBAltSeqItem_value,omitempty" json:"gb_alt_seq_item_value,omitempty"`
}

//
// type GBAltSeqItemInterval struct {
//           `xml:"" json:""`

// }

//
type GBAltSeqItemIsgap struct {
	Value string `xml:"value,attr" json:"value"`

	// extended attributes
	Attrs []xml.Attr `xml:",attr,omitempty" json:"attrs,omitempty"`
}
type GBAltSeqItemIsgapValue string

// const (
//         True    GBAltSeqItemIsgapValue  = "true"
//         False   GBAltSeqItemIsgapValue  = "false"
// )

type GBAltSeqItemFirstAccn struct {
	GBAltSeqItemFirstAccn string `xml:"GBAltSeqItem_first-accn,omitempty" json:"gb_alt_seq_item_first_accn,omitempty"`
}
type GBAltSeqItemGapComment struct {
	GBAltSeqItemGapComment string `xml:"GBAltSeqItem_gap-comment,omitempty" json:"gb_alt_seq_item_gap_comment,omitempty"`
}
type GBAltSeqItemGapLength struct {
	GBAltSeqItemGapLength int `xml:"GBAltSeqItem_gap-length,omitempty" json:"gb_alt_seq_item_gap_length,omitempty"`
}
type GBAltSeqItemGapLinkage struct {
	GBAltSeqItemGapLinkage string `xml:"GBAltSeqItem_gap-linkage,omitempty" json:"gb_alt_seq_item_gap_linkage,omitempty"`
}
type GBAltSeqItemGapType struct {
	GBAltSeqItemGapType string `xml:"GBAltSeqItem_gap-type,omitempty" json:"gb_alt_seq_item_gap_type,omitempty"`
}
type GBAltSeqItemLastAccn struct {
	GBAltSeqItemLastAccn string `xml:"GBAltSeqItem_last-accn,omitempty" json:"gb_alt_seq_item_last_accn,omitempty"`
}
type GBAltSeqItemValue struct {
	GBAltSeqItemValue string `xml:"GBAltSeqItem_value,omitempty" json:"gb_alt_seq_item_value,omitempty"`
}
type GBAuthor struct {
	GBAuthor string `xml:"GBAuthor" json:"gb_author"`
}

//
type GBComment struct {
	GBCommentType string `xml:"GBComment_type,omitempty" json:"gb_comment_type,omitempty"`

	GBCommentParagraphs []GBCommentParagraph `xml:"GBComment_paragraphs" json:"gb_comment_paragraphs"`
}
type GBCommentParagraph struct {
	GBCommentParagraph string `xml:"GBCommentParagraph" json:"gb_comment_paragraph"`
}
type GBCommentType struct {
	GBCommentType string `xml:"GBComment_type,omitempty" json:"gb_comment_type,omitempty"`
}

//
type GBFeature struct {
	GBFeatureKey string `xml:"GBFeature_key" json:"gb_feature_key"`

	GBFeatureLocation string `xml:"GBFeature_location" json:"gb_feature_location"`

	GBFeatureIntervals []GBInterval `xml:"GBFeature_intervals,omitempty" json:"gb_feature_intervals,omitempty"`

	GBFeatureOperator string `xml:"GBFeature_operator,omitempty" json:"gb_feature_operator,omitempty"`

	GBFeaturePartial5 *GBFeaturePartial5 `xml:"GBFeature_partial5,omitempty" json:"gb_feature_partial5,omitempty"`

	GBFeaturePartial3 *GBFeaturePartial3 `xml:"GBFeature_partial3,omitempty" json:"gb_feature_partial3,omitempty"`

	GBFeatureQuals []GBQualifier `xml:"GBFeature_quals,omitempty" json:"gb_feature_quals,omitempty"`

	GBFeatureXrefs []GBXref `xml:"GBFeature_xrefs,omitempty" json:"gb_feature_xrefs,omitempty"`
}

//
type GBFeaturePartial3 struct {
	Value string `xml:"value,attr" json:"value"`

	// extended attributes
	Attrs []xml.Attr `xml:",attr,omitempty" json:"attrs,omitempty"`
}
type GBFeaturePartial3Value string

// const (
//         True    GBFeaturePartial3Value  = "true"
//         False   GBFeaturePartial3Value  = "false"
// )

//
type GBFeaturePartial5 struct {
	Value string `xml:"value,attr" json:"value"`

	// extended attributes
	Attrs []xml.Attr `xml:",attr,omitempty" json:"attrs,omitempty"`
}
type GBFeaturePartial5Value string

// const (
//         True    GBFeaturePartial5Value  = "true"
//         False   GBFeaturePartial5Value  = "false"
// )

//
type GBFeatureSet struct {
	GBFeatureSetAnnotSource string `xml:"GBFeatureSet_annot-source,omitempty" json:"gb_feature_set_annot_source,omitempty"`

	GBFeatureSetFeatures []GBFeature `xml:"GBFeatureSet_features" json:"gb_feature_set_features"`
}
type GBFeatureSetAnnotSource struct {
	GBFeatureSetAnnotSource string `xml:"GBFeatureSet_annot-source,omitempty" json:"gb_feature_set_annot_source,omitempty"`
}
type GBFeatureKey struct {
	GBFeatureKey string `xml:"GBFeature_key" json:"gb_feature_key"`
}
type GBFeatureLocation struct {
	GBFeatureLocation string `xml:"GBFeature_location" json:"gb_feature_location"`
}
type GBFeatureOperator struct {
	GBFeatureOperator string `xml:"GBFeature_operator,omitempty" json:"gb_feature_operator,omitempty"`
}

//
type GBInterval struct {
	GBIntervalFrom int `xml:"GBInterval_from,omitempty" json:"gb_interval_from,omitempty"`

	GBIntervalTo int `xml:"GBInterval_to,omitempty" json:"gb_interval_to,omitempty"`

	GBIntervalPoint int `xml:"GBInterval_point,omitempty" json:"gb_interval_point,omitempty"`

	GBIntervalIscomp *GBIntervalIscomp `xml:"GBInterval_iscomp,omitempty" json:"gb_interval_iscomp,omitempty"`

	GBIntervalInterbp *GBIntervalInterbp `xml:"GBInterval_interbp,omitempty" json:"gb_interval_interbp,omitempty"`

	GBIntervalAccession string `xml:"GBInterval_accession" json:"gb_interval_accession"`
}

//
type GBIntervalInterbp struct {
	Value string `xml:"value,attr" json:"value"`

	// extended attributes
	Attrs []xml.Attr `xml:",attr,omitempty" json:"attrs,omitempty"`
}
type GBIntervalInterbpValue string

const (
	True  GBIntervalInterbpValue = "true"
	False GBIntervalInterbpValue = "false"
)

//
type GBIntervalIscomp struct {
	Value string `xml:"value,attr" json:"value"`

	// extended attributes
	Attrs []xml.Attr `xml:",attr,omitempty" json:"attrs,omitempty"`
}
type GBIntervalIscompValue string

// const (
//         True    GBIntervalIscompValue   = "true"
//         False   GBIntervalIscompValue   = "false"
// )

type GBIntervalAccession struct {
	GBIntervalAccession string `xml:"GBInterval_accession" json:"gb_interval_accession"`
}
type GBIntervalFrom struct {
	GBIntervalFrom int `xml:"GBInterval_from,omitempty" json:"gb_interval_from,omitempty"`
}
type GBIntervalPoint struct {
	GBIntervalPoint int `xml:"GBInterval_point,omitempty" json:"gb_interval_point,omitempty"`
}
type GBIntervalTo struct {
	GBIntervalTo int `xml:"GBInterval_to,omitempty" json:"gb_interval_to,omitempty"`
}
type GBKeyword struct {
	GBKeyword string `xml:"GBKeyword" json:"gb_keyword"`
}

//
type GBQualifier struct {
	GBQualifierName string `xml:"GBQualifier_name" json:"gb_qualifier_name"`

	GBQualifierValue string `xml:"GBQualifier_value,omitempty" json:"gb_qualifier_value,omitempty"`
}
type GBQualifierName struct {
	GBQualifierName string `xml:"GBQualifier_name" json:"gb_qualifier_name"`
}
type GBQualifierValue struct {
	GBQualifierValue string `xml:"GBQualifier_value,omitempty" json:"gb_qualifier_value,omitempty"`
}

//
type GBReference struct {
	GBReferenceReference string `xml:"GBReference_reference" json:"gb_reference_reference"`

	GBReferencePosition string `xml:"GBReference_position,omitempty" json:"gb_reference_position,omitempty"`

	GBReferenceAuthors []GBAuthor `xml:"GBReference_authors,omitempty" json:"gb_reference_authors,omitempty"`

	GBReferenceConsortium string `xml:"GBReference_consortium,omitempty" json:"gb_reference_consortium,omitempty"`

	GBReferenceTitle string `xml:"GBReference_title,omitempty" json:"gb_reference_title,omitempty"`

	GBReferenceJournal string `xml:"GBReference_journal" json:"gb_reference_journal"`

	GBReferenceXrefs []GBXref `xml:"GBReference_xref,omitempty" json:"gb_reference_xrefs,omitempty"`

	GBReferencePubmed int `xml:"GBReference_pubmed,omitempty" json:"gb_reference_pubmed,omitempty"`

	GBReferenceRemark string `xml:"GBReference_remark,omitempty" json:"gb_reference_remark,omitempty"`
}
type GBReferenceConsortium struct {
	GBReferenceConsortium string `xml:"GBReference_consortium,omitempty" json:"gb_reference_consortium,omitempty"`
}
type GBReferenceJournal struct {
	GBReferenceJournal string `xml:"GBReference_journal" json:"gb_reference_journal"`
}
type GBReferencePosition struct {
	GBReferencePosition string `xml:"GBReference_position,omitempty" json:"gb_reference_position,omitempty"`
}
type GBReferencePubmed struct {
	GBReferencePubmed int `xml:"GBReference_pubmed,omitempty" json:"gb_reference_pubmed,omitempty"`
}
type GBReferenceReference struct {
	GBReferenceReference string `xml:"GBReference_reference" json:"gb_reference_reference"`
}
type GBReferenceRemark struct {
	GBReferenceRemark string `xml:"GBReference_remark,omitempty" json:"gb_reference_remark,omitempty"`
}
type GBReferenceTitle struct {
	GBReferenceTitle string `xml:"GBReference_title,omitempty" json:"gb_reference_title,omitempty"`
}
type GBSecondaryAccn struct {
	GBSecondaryAccn string `xml:"GBSecondary-accn" json:"gb_secondary_accn"`
}

//
type GBSeq struct {
	GBSeqLocus string `xml:"GBSeq_locus,omitempty" json:"gb_seq_locus,omitempty"`

	GBSeqLength int `xml:"GBSeq_length" json:"gb_seq_length"`

	GBSeqStrandedness string `xml:"GBSeq_strandedness,omitempty" json:"gb_seq_strandedness,omitempty"`

	GBSeqMoltype string `xml:"GBSeq_moltype" json:"gb_seq_moltype"`

	GBSeqTopology string `xml:"GBSeq_topology,omitempty" json:"gb_seq_topology,omitempty"`

	GBSeqDivision string `xml:"GBSeq_division,omitempty" json:"gb_seq_division,omitempty"`

	GBSeqUpdateDate string `xml:"GBSeq_update-date,omitempty" json:"gb_seq_update_date,omitempty"`

	GBSeqCreateDate string `xml:"GBSeq_create-date,omitempty" json:"gb_seq_create_date,omitempty"`

	GBSeqUpdateRelease string `xml:"GBSeq_update-release,omitempty" json:"gb_seq_update_release,omitempty"`

	GBSeqCreateRelease string `xml:"GBSeq_create-release,omitempty" json:"gb_seq_create_release,omitempty"`

	GBSeqDefinition string `xml:"GBSeq_definition,omitempty" json:"gb_seq_definition,omitempty"`

	GBSeqPrimaryAccession string `xml:"GBSeq_primary-accession,omitempty" json:"gb_seq_primary_accession,omitempty"`

	GBSeqEntryVersion string `xml:"GBSeq_entry-version,omitempty" json:"gb_seq_entry_version,omitempty"`

	GBSeqAccessionVersion string `xml:"GBSeq_accession-version,omitempty" json:"gb_seq_accession_version,omitempty"`

	GBSeqOtherSeqids []GBSeqid `xml:"GBSeq_other-seqids,omitempty" json:"gb_seq_other_seqids,omitempty"`

	GBSeqSecondaryAccessions []GBSecondaryAccn `xml:"GBSeq_secondary-accessions,omitempty" json:"gb_seq_secondary_accessions,omitempty"`

	GBSeqProject string `xml:"GBSeq_project,omitempty" json:"gb_seq_project,omitempty"`

	GBSeqKeywords []GBKeyword `xml:"GBSeq_keywords,omitempty" json:"gb_seq_keywords,omitempty"`

	GBSeqSegment string `xml:"GBSeq_segment,omitempty" json:"gb_seq_segment,omitempty"`

	GBSeqSource string `xml:"GBSeq_source,omitempty" json:"gb_seq_source,omitempty"`

	GBSeqOrganism string `xml:"GBSeq_organism,omitempty" json:"gb_seq_organism,omitempty"`

	GBSeqTaxonomy string `xml:"GBSeq_taxonomy,omitempty" json:"gb_seq_taxonomy,omitempty"`

	GBSeqReferences []GBReference `xml:"GBSeq_references,omitempty" json:"gb_seq_references,omitempty"`

	GBSeqComment string `xml:"GBSeq_comment,omitempty" json:"gb_seq_comment,omitempty"`

	GBSeqCommentSets []GBComment `xml:"GBSeq_comment-set,omitempty" json:"gb_seq_comment_sets,omitempty"`

	GBSeqStrucComments []GBStrucComment `xml:"GBSeq_struc-comments,omitempty" json:"gb_seq_struc_comments,omitempty"`

	GBSeqPrimary string `xml:"GBSeq_primary,omitempty" json:"gb_seq_primary,omitempty"`

	GBSeqSourceDb string `xml:"GBSeq_source-db,omitempty" json:"gb_seq_source_db,omitempty"`

	GBSeqDatabaseReference string `xml:"GBSeq_database-reference,omitempty" json:"gb_seq_database_reference,omitempty"`

	GBSeqFeatureTables []GBFeature `xml:"GBSeq_feature-table,omitempty" json:"gb_seq_feature_tables,omitempty"`

	GBSeqFeatureSets []GBFeatureSet `xml:"GBSeq_feature-set,omitempty" json:"gb_seq_feature_sets,omitempty"`

	GBSeqSequence string `xml:"GBSeq_sequence,omitempty" json:"gb_seq_sequence,omitempty"`

	GBSeqContig string `xml:"GBSeq_contig,omitempty" json:"gb_seq_contig,omitempty"`

	GBSeqAltSeqs []GBAltSeqData `xml:"GBSeq_alt-seq,omitempty" json:"gb_seq_alt_seqs,omitempty"`

	GBSeqXrefs []GBXref `xml:"GBSeq_xrefs,omitempty" json:"gb_seq_xrefs,omitempty"`
}
type GBSeqAccessionVersion struct {
	GBSeqAccessionVersion string `xml:"GBSeq_accession-version,omitempty" json:"gb_seq_accession_version,omitempty"`
}
type GBSeqComment struct {
	GBSeqComment string `xml:"GBSeq_comment,omitempty" json:"gb_seq_comment,omitempty"`
}
type GBSeqContig struct {
	GBSeqContig string `xml:"GBSeq_contig,omitempty" json:"gb_seq_contig,omitempty"`
}
type GBSeqCreateDate struct {
	GBSeqCreateDate string `xml:"GBSeq_create-date,omitempty" json:"gb_seq_create_date,omitempty"`
}
type GBSeqCreateRelease struct {
	GBSeqCreateRelease string `xml:"GBSeq_create-release,omitempty" json:"gb_seq_create_release,omitempty"`
}
type GBSeqDatabaseReference struct {
	GBSeqDatabaseReference string `xml:"GBSeq_database-reference,omitempty" json:"gb_seq_database_reference,omitempty"`
}
type GBSeqDefinition struct {
	GBSeqDefinition string `xml:"GBSeq_definition,omitempty" json:"gb_seq_definition,omitempty"`
}
type GBSeqDivision struct {
	GBSeqDivision string `xml:"GBSeq_division,omitempty" json:"gb_seq_division,omitempty"`
}
type GBSeqEntryVersion struct {
	GBSeqEntryVersion string `xml:"GBSeq_entry-version,omitempty" json:"gb_seq_entry_version,omitempty"`
}
type GBSeqLength struct {
	GBSeqLength int `xml:"GBSeq_length" json:"gb_seq_length"`
}
type GBSeqLocus struct {
	GBSeqLocus string `xml:"GBSeq_locus,omitempty" json:"gb_seq_locus,omitempty"`
}
type GBSeqMoltype struct {
	GBSeqMoltype string `xml:"GBSeq_moltype" json:"gb_seq_moltype"`
}
type GBSeqOrganism struct {
	GBSeqOrganism string `xml:"GBSeq_organism,omitempty" json:"gb_seq_organism,omitempty"`
}
type GBSeqPrimary struct {
	GBSeqPrimary string `xml:"GBSeq_primary,omitempty" json:"gb_seq_primary,omitempty"`
}
type GBSeqPrimaryAccession struct {
	GBSeqPrimaryAccession string `xml:"GBSeq_primary-accession,omitempty" json:"gb_seq_primary_accession,omitempty"`
}
type GBSeqProject struct {
	GBSeqProject string `xml:"GBSeq_project,omitempty" json:"gb_seq_project,omitempty"`
}
type GBSeqSegment struct {
	GBSeqSegment string `xml:"GBSeq_segment,omitempty" json:"gb_seq_segment,omitempty"`
}
type GBSeqSequence struct {
	GBSeqSequence string `xml:"GBSeq_sequence,omitempty" json:"gb_seq_sequence,omitempty"`
}
type GBSeqSource struct {
	GBSeqSource string `xml:"GBSeq_source,omitempty" json:"gb_seq_source,omitempty"`
}
type GBSeqSourceDb struct {
	GBSeqSourceDb string `xml:"GBSeq_source-db,omitempty" json:"gb_seq_source_db,omitempty"`
}
type GBSeqStrandedness struct {
	GBSeqStrandedness string `xml:"GBSeq_strandedness,omitempty" json:"gb_seq_strandedness,omitempty"`
}
type GBSeqTaxonomy struct {
	GBSeqTaxonomy string `xml:"GBSeq_taxonomy,omitempty" json:"gb_seq_taxonomy,omitempty"`
}
type GBSeqTopology struct {
	GBSeqTopology string `xml:"GBSeq_topology,omitempty" json:"gb_seq_topology,omitempty"`
}
type GBSeqUpdateDate struct {
	GBSeqUpdateDate string `xml:"GBSeq_update-date,omitempty" json:"gb_seq_update_date,omitempty"`
}
type GBSeqUpdateRelease struct {
	GBSeqUpdateRelease string `xml:"GBSeq_update-release,omitempty" json:"gb_seq_update_release,omitempty"`
}
type GBSeqid struct {
	GBSeqid string `xml:"GBSeqid" json:"gb_seqid"`
}

//
type GBStrucComment struct {
	GBStrucCommentName string `xml:"GBStrucComment_name,omitempty" json:"gb_struc_comment_name,omitempty"`

	GBStrucCommentItems []GBStrucCommentItem `xml:"GBStrucComment_items" json:"gb_struc_comment_items"`
}

//
type GBStrucCommentItem struct {
	GBStrucCommentItemTag string `xml:"GBStrucCommentItem_tag,omitempty" json:"gb_struc_comment_item_tag,omitempty"`

	GBStrucCommentItemValue string `xml:"GBStrucCommentItem_value,omitempty" json:"gb_struc_comment_item_value,omitempty"`

	GBStrucCommentItemUrl string `xml:"GBStrucCommentItem_url,omitempty" json:"gb_struc_comment_item_url,omitempty"`
}
type GBStrucCommentItemTag struct {
	GBStrucCommentItemTag string `xml:"GBStrucCommentItem_tag,omitempty" json:"gb_struc_comment_item_tag,omitempty"`
}
type GBStrucCommentItemUrl struct {
	GBStrucCommentItemUrl string `xml:"GBStrucCommentItem_url,omitempty" json:"gb_struc_comment_item_url,omitempty"`
}
type GBStrucCommentItemValue struct {
	GBStrucCommentItemValue string `xml:"GBStrucCommentItem_value,omitempty" json:"gb_struc_comment_item_value,omitempty"`
}
type GBStrucCommentName struct {
	GBStrucCommentName string `xml:"GBStrucComment_name,omitempty" json:"gb_struc_comment_name,omitempty"`
}

//
type GBXref struct {
	GBXrefDbname string `xml:"GBXref_dbname" json:"gb_xref_dbname"`

	GBXrefID string `xml:"GBXref_id" json:"gb_xref_id"`
}
type GBXrefDbname struct {
	GBXrefDbname string `xml:"GBXref_dbname" json:"gb_xref_dbname"`
}
type GBXrefID struct {
	GBXrefID string `xml:"GBXref_id" json:"gb_xref_id"`
}
