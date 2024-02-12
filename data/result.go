package data

type TransactionResult int16

const (
	// Note: Exact number must stay stable.  This code is stored by value
	// in metadata for historic transactions.

	// 0: S Success (success)
	// Causes:
	// - Success.
	// Implications:
	// - Applied
	// - Forwarded
	TesSUCCESS TransactionResult = 0
)
const (
	// Note: Exact numbers must stay stable.  These codes are stored by
	// value in metadata for historic transactions.

	// 100 .. 255 C
	//   Claim fee only (ripple transaction with no good paths, pay to
	//   non-existent account, no path)
	//
	// Causes:
	// - Success, but does not achieve optimal result.
	// - Invalid transaction or no effect, but claim fee to use the sequence
	//   number.
	//
	// Implications:
	// - Applied
	// - Forwarded
	//
	// Only allowed as a return code of appliedTransaction when !tapRETRY.
	// Otherwise, treated as terRETRY.
	//
	// DO NOT CHANGE THESE NUMBERS: They appear in ledger meta data.
	TecCLAIM                              = TransactionResult(100)
	TecPATH_PARTIAL                       = TransactionResult(101)
	TecUNFUNDED_ADD                       = TransactionResult(102) // Unused legacy code
	TecUNFUNDED_OFFER                     = TransactionResult(103)
	TecUNFUNDED_PAYMENT                   = TransactionResult(104)
	TecFAILED_PROCESSING                  = TransactionResult(105)
	TecDIR_FULL                           = TransactionResult(121)
	TecINSUF_RESERVE_LINE                 = TransactionResult(122)
	TecINSUF_RESERVE_OFFER                = TransactionResult(123)
	TecNO_DST                             = TransactionResult(124)
	TecNO_DST_INSUF_XRP                   = TransactionResult(125)
	TecNO_LINE_INSUF_RESERVE              = TransactionResult(126)
	TecNO_LINE_REDUNDANT                  = TransactionResult(127)
	TecPATH_DRY                           = TransactionResult(128)
	TecUNFUNDED                           = TransactionResult(129)
	TecNO_ALTERNATIVE_KEY                 = TransactionResult(130)
	TecNO_REGULAR_KEY                     = TransactionResult(131)
	TecOWNERS                             = TransactionResult(132)
	TecNO_ISSUER                          = TransactionResult(133)
	TecNO_AUTH                            = TransactionResult(134)
	TecNO_LINE                            = TransactionResult(135)
	TecINSUFF_FEE                         = TransactionResult(136)
	TecFROZEN                             = TransactionResult(137)
	TecNO_TARGET                          = TransactionResult(138)
	TecNO_PERMISSION                      = TransactionResult(139)
	TecNO_ENTRY                           = TransactionResult(140)
	TecINSUFFICIENT_RESERVE               = TransactionResult(141)
	TecNEED_MASTER_KEY                    = TransactionResult(142)
	TecDST_TAG_NEEDED                     = TransactionResult(143)
	TecINTERNAL                           = TransactionResult(144)
	TecOVERSIZE                           = TransactionResult(145)
	TecCRYPTOCONDITION_ERROR              = TransactionResult(146)
	TecINVARIANT_FAILED                   = TransactionResult(147)
	TecEXPIRED                            = TransactionResult(148)
	TecDUPLICATE                          = TransactionResult(149)
	TecKILLED                             = TransactionResult(150)
	TecHAS_OBLIGATIONS                    = TransactionResult(151)
	TecTOO_SOON                           = TransactionResult(152)
	TecHOOK_REJECTED                      = TransactionResult(153)
	TecMAX_SEQUENCE_REACHED               = TransactionResult(154)
	TecNO_SUITABLE_NFTOKEN_PAGE           = TransactionResult(155)
	TecNFTOKEN_BUY_SELL_MISMATCH          = TransactionResult(156)
	TecNFTOKEN_OFFER_TYPE_MISMATCH        = TransactionResult(157)
	TecCANT_ACCEPT_OWN_NFTOKEN_OFFER      = TransactionResult(158)
	TecINSUFFICIENT_FUNDS                 = TransactionResult(159)
	TecOBJECT_NOT_FOUND                   = TransactionResult(160)
	TecINSUFFICIENT_PAYMENT               = TransactionResult(161)
	TecUNFUNDED_AMM                       = TransactionResult(162)
	TecAMM_BALANCE                        = TransactionResult(163)
	TecAMM_FAILED                         = TransactionResult(164)
	TecAMM_INVALID_TOKENS                 = TransactionResult(165)
	TecAMM_EMPTY                          = TransactionResult(166)
	TecAMM_NOT_EMPTY                      = TransactionResult(167)
	TecAMM_ACCOUNT                        = TransactionResult(168)
	TecINCOMPLETE                         = TransactionResult(169)
	TecXCHAIN_BAD_TRANSFER_ISSUE          = TransactionResult(170)
	TecXCHAIN_NO_CLAIM_ID                 = TransactionResult(171)
	TecXCHAIN_BAD_CLAIM_ID                = TransactionResult(172)
	TecXCHAIN_CLAIM_NO_QUORUM             = TransactionResult(173)
	TecXCHAIN_PROOF_UNKNOWN_KEY           = TransactionResult(174)
	TecXCHAIN_CREATE_ACCOUNT_NONXRP_ISSUE = TransactionResult(175)
	TecXCHAIN_WRONG_CHAIN                 = TransactionResult(176)
	TecXCHAIN_REWARD_MISMATCH             = TransactionResult(177)
	TecXCHAIN_NO_SIGNERS_LIST             = TransactionResult(178)
	TecXCHAIN_SENDING_ACCOUNT_MISMATCH    = TransactionResult(179)
	TecXCHAIN_INSUFF_CREATE_AMOUNT        = TransactionResult(180)
	TecXCHAIN_ACCOUNT_CREATE_PAST         = TransactionResult(181)
	TecXCHAIN_ACCOUNT_CREATE_TOO_MANY     = TransactionResult(182)
	TecXCHAIN_PAYMENT_FAILED              = TransactionResult(183)
	TecXCHAIN_SELF_COMMIT                 = TransactionResult(184)
	TecXCHAIN_BAD_PUBLIC_KEY_ACCOUNT_PAIR = TransactionResult(185)
	TecXCHAIN_CREATE_ACCOUNT_DISABLED     = TransactionResult(186)
	TecEMPTY_DID                          = TransactionResult(187)
)

const (
	// Note: Range is stable.
	// Exact numbers are used in ripple-binary-codec:
	//     https://github.com/ripple/ripple-binary-codec/blob/master/src/enums/definitions.json
	// Use tokens.

	// -399 .. -300: L Local error (transaction fee inadequate, exceeds local
	// limit) Only valid during non-consensus processing. Implications:
	// - Not forwarded
	// - No fee check
	TelLOCAL_ERROR TransactionResult = iota - 399
	TelBAD_DOMAIN
	TelBAD_PATH_COUNT
	TelBAD_PUBLIC_KEY
	TelFAILED_PROCESSING
	TelINSUF_FEE_P
	TelNO_DST_PARTIAL
	TelCAN_NOT_QUEUE
	TelCAN_NOT_QUEUE_BALANCE
	TelCAN_NOT_QUEUE_BLOCKS
	TelCAN_NOT_QUEUE_BLOCKED
	TelCAN_NOT_QUEUE_FEE
	TelCAN_NOT_QUEUE_FULL
	TelWRONG_NETWORK
	TelREQUIRES_NETWORK_ID
	TelNETWORK_ID_MAKES_TX_NON_CANONICAL
)
const (
	// Note: Range is stable.
	// Exact numbers are used in ripple-binary-codec:
	//     https://github.com/ripple/ripple-binary-codec/blob/master/src/enums/definitions.json
	// Use tokens.

	// -299 .. -200: M Malformed (bad signature)
	// Causes:
	// - Transaction corrupt.
	// Implications:
	// - Not applied
	// - Not forwarded
	// - Reject
	// - Cannot succeed in any imagined ledger.
	TemMALFORMED TransactionResult = iota - 299
	TemBAD_AMOUNT
	TemBAD_CURRENCY
	TemBAD_EXPIRATION
	TemBAD_FEE
	TemBAD_ISSUER
	TemBAD_LIMIT
	TemBAD_OFFER
	TemBAD_PATH
	TemBAD_PATH_LOOP
	TemBAD_REGKEY
	TemBAD_SEND_XRP_LIMIT
	TemBAD_SEND_XRP_MAX
	TemBAD_SEND_XRP_NO_DIRECT
	TemBAD_SEND_XRP_PARTIAL
	TemBAD_SEND_XRP_PATHS
	TemBAD_SEQUENCE
	TemBAD_SIGNATURE
	TemBAD_SRC_ACCOUNT
	TemBAD_TRANSFER_RATE
	TemDST_IS_SRC
	TemDST_NEEDED
	TemINVALID
	TemINVALID_FLAG
	TemREDUNDANT
	TemRIPPLE_EMPTY
	TemDISABLED
	TemBAD_SIGNER
	TemBAD_QUORUM
	TemBAD_WEIGHT
	TemBAD_TICK_SIZE
	TemINVALID_ACCOUNT_ID
	TemCANNOT_PREAUTH_SELF
	TemINVALID_COUNT
	TemUNCERTAIN // An internal intermediate result; should never be returned.
	TemUNKNOWN   // An internal intermediate result; should never be returned.
	TemSEQ_AND_TICKET
	TemBAD_NFTOKEN_TRANSFER_FEE
	TemBAD_AMM_TOKENS
	TemXCHAIN_EQUAL_DOOR_ACCOUNTS
	TemXCHAIN_BAD_PROOF
	TemXCHAIN_BRIDGE_BAD_ISSUES
	TemXCHAIN_BRIDGE_NONDOOR_OWNER
	TemXCHAIN_BRIDGE_BAD_MIN_ACCOUNT_CREATE_AMOUNT
	TemXCHAIN_BRIDGE_BAD_REWARD_AMOUNT
	TemEMPTY_DID
)
const (
	// Note: Range is stable.
	// Exact numbers are used in ripple-binary-codec:
	//     https://github.com/ripple/ripple-binary-codec/blob/master/src/enums/definitions.json
	// Use tokens.

	// -199 .. -100: F
	//    Failure (sequence number previously used)
	//
	// Causes:
	// - Transaction cannot succeed because of ledger state.
	// - Unexpected ledger state.
	// - C++ exception.
	//
	// Implications:
	// - Not applied
	// - Not forwarded
	// - Could succeed in an imagined ledger.
	TefFAILURE TransactionResult = iota - 199
	TefALREADY
	TefBAD_ADD_AUTH
	TefBAD_AUTH
	TefBAD_LEDGER
	TefCREATED
	TefEXCEPTION
	TefINTERNAL
	TefNO_AUTH_REQUIRED // Can't set auth if auth is not required.
	TefPAST_SEQ
	TefWRONG_PRIOR
	TefMASTER_DISABLED
	TefMAX_LEDGER
	TefBAD_SIGNATURE
	TefBAD_QUORUM
	TefNOT_MULTI_SIGNING
	TefBAD_AUTH_MASTER
	TefINVARIANT_FAILED
	TefTOO_BIG
	TefNO_TICKET
	TefNFTOKEN_IS_NOT_TRANSFERABLE
)
const (
	// Note: Range is stable.
	// Exact numbers are used in ripple-binary-codec:
	//     https://github.com/ripple/ripple-binary-codec/blob/master/src/enums/definitions.json
	// Use tokens.

	// -99 .. -1: R Retry
	//   sequence too high, no funds for txn fee, originating -account
	//   non-existent
	//
	// Cause:
	//   Prior application of another, possibly non-existent, transaction could
	//   allow this transaction to succeed.
	//
	// Implications:
	// - Not applied
	// - May be forwarded
	//   - Results indicating the txn was forwarded: terQUEUED
	//   - All others are not forwarded.
	// - Might succeed later
	// - Hold
	// - Makes hole in sequence which jams transactions.
	TerRETRY       TransactionResult = iota - 99
	TerFUNDS_SPENT                   // DEPRECATED.
	TerINSUF_FEE_B                   // Can't pay fee, therefore don't burden network.
	TerNO_ACCOUNT                    // Can't pay fee, therefore don't burden network.
	TerNO_AUTH                       // Not authorized to hold IOUs.
	TerNO_LINE                       // Internal flag.
	TerOWNERS                        // Can't succeed with non-zero owner count.
	TerPRE_SEQ                       // Can't pay fee, no point in forwarding, so don't
	// burden network.
	TerLAST       // DEPRECATED.
	TerNO_RIPPLE  // Rippling not allowed
	TerQUEUED     // Transaction is being held in TxQ until fee drops
	TerPRE_TICKET // Ticket is not yet in ledger but might be on its way
	TerNO_AMM     // AMM doesn't exist for the asset pair
)

var resultNames = map[TransactionResult]struct {
	Token string
	Human string
}{
	TesSUCCESS: {"tesSUCCESS", "The transaction was applied."},

	TecAMM_BALANCE:                        {"tecAMM_BALANCE", "AMM has invalid balance. "},
	TecAMM_INVALID_TOKENS:                 {"tecAMM_INVALID_TOKENS", "AMM invalid LP tokens. "},
	TecAMM_FAILED:                         {"tecAMM_FAILED", "AMM transaction failed. "},
	TecAMM_EMPTY:                          {"tecAMM_EMPTY", "AMM is in empty state. "},
	TecAMM_NOT_EMPTY:                      {"tecAMM_NOT_EMPTY", "AMM is not in empty state. "},
	TecAMM_ACCOUNT:                        {"tecAMM_ACCOUNT", "This operation is not allowed on an AMM Account. "},
	TecCLAIM:                              {"tecCLAIM", "Fee claimed. Sequence used. No action. "},
	TecDIR_FULL:                           {"tecDIR_FULL", "Can not add entry to full directory. "},
	TecFAILED_PROCESSING:                  {"tecFAILED_PROCESSING", "Failed to correctly process transaction. "},
	TecINSUF_RESERVE_LINE:                 {"tecINSUF_RESERVE_LINE", "Insufficient reserve to add trust line. "},
	TecINSUF_RESERVE_OFFER:                {"tecINSUF_RESERVE_OFFER", "Insufficient reserve to create offer. "},
	TecNO_DST:                             {"tecNO_DST", "Destination does not exist. Send XRP to create it. "},
	TecNO_DST_INSUF_XRP:                   {"tecNO_DST_INSUF_XRP", "Destination does not exist. Too little XRP sent to create it. "},
	TecNO_LINE_INSUF_RESERVE:              {"tecNO_LINE_INSUF_RESERVE", "No such line. Too little reserve to create it. "},
	TecNO_LINE_REDUNDANT:                  {"tecNO_LINE_REDUNDANT", "Can't set non-existent line to default. "},
	TecPATH_DRY:                           {"tecPATH_DRY", "Path could not send partial amount. "},
	TecPATH_PARTIAL:                       {"tecPATH_PARTIAL", "Path could not send full amount. "},
	TecNO_ALTERNATIVE_KEY:                 {"tecNO_ALTERNATIVE_KEY", "The operation would remove the ability to sign transactions with the account. "},
	TecNO_REGULAR_KEY:                     {"tecNO_REGULAR_KEY", "Regular key is not set. "},
	TecOVERSIZE:                           {"tecOVERSIZE", "Object exceeded serialization limits. "},
	TecUNFUNDED:                           {"tecUNFUNDED", "Not enough XRP to satisfy the reserve requirement. "},
	TecUNFUNDED_ADD:                       {"tecUNFUNDED_ADD", "DEPRECATED. "},
	TecUNFUNDED_AMM:                       {"tecUNFUNDED_AMM", "Insufficient balance to fund AMM. "},
	TecUNFUNDED_OFFER:                     {"tecUNFUNDED_OFFER", "Insufficient balance to fund created offer. "},
	TecUNFUNDED_PAYMENT:                   {"tecUNFUNDED_PAYMENT", "Insufficient XRP balance to send. "},
	TecOWNERS:                             {"tecOWNERS", "Non-zero owner count. "},
	TecNO_ISSUER:                          {"tecNO_ISSUER", "Issuer account does not exist. "},
	TecNO_AUTH:                            {"tecNO_AUTH", "Not authorized to hold asset. "},
	TecNO_LINE:                            {"tecNO_LINE", "No such line. "},
	TecINSUFF_FEE:                         {"tecINSUFF_FEE", "Insufficient balance to pay fee. "},
	TecFROZEN:                             {"tecFROZEN", "Asset is frozen. "},
	TecNO_TARGET:                          {"tecNO_TARGET", "Target account does not exist. "},
	TecNO_PERMISSION:                      {"tecNO_PERMISSION", "No permission to perform requested operation. "},
	TecNO_ENTRY:                           {"tecNO_ENTRY", "No matching entry found. "},
	TecINSUFFICIENT_RESERVE:               {"tecINSUFFICIENT_RESERVE", "Insufficient reserve to complete requested operation. "},
	TecNEED_MASTER_KEY:                    {"tecNEED_MASTER_KEY", "The operation requires the use of the Master Key. "},
	TecDST_TAG_NEEDED:                     {"tecDST_TAG_NEEDED", "A destination tag is required. "},
	TecINTERNAL:                           {"tecINTERNAL", "An internal error has occurred during processing. "},
	TecCRYPTOCONDITION_ERROR:              {"tecCRYPTOCONDITION_ERROR", "Malformed, invalid, or mismatched conditional or fulfillment. "},
	TecINVARIANT_FAILED:                   {"tecINVARIANT_FAILED", "One or more invariants for the transaction were not satisfied. "},
	TecEXPIRED:                            {"tecEXPIRED", "Expiration time is passed. "},
	TecDUPLICATE:                          {"tecDUPLICATE", "Ledger object already exists. "},
	TecKILLED:                             {"tecKILLED", "No funds transferred and no offer created. "},
	TecHAS_OBLIGATIONS:                    {"tecHAS_OBLIGATIONS", "The account cannot be deleted since it has obligations. "},
	TecTOO_SOON:                           {"tecTOO_SOON", "It is too early to attempt the requested operation. Please wait. "},
	TecMAX_SEQUENCE_REACHED:               {"tecMAX_SEQUENCE_REACHED", "The maximum sequence number was reached. "},
	TecNO_SUITABLE_NFTOKEN_PAGE:           {"tecNO_SUITABLE_NFTOKEN_PAGE", "A suitable NFToken page could not be located. "},
	TecNFTOKEN_BUY_SELL_MISMATCH:          {"tecNFTOKEN_BUY_SELL_MISMATCH", "The 'Buy' and 'Sell' NFToken offers are mismatched. "},
	TecNFTOKEN_OFFER_TYPE_MISMATCH:        {"tecNFTOKEN_OFFER_TYPE_MISMATCH", "The type of NFToken offer is incorrect. "},
	TecCANT_ACCEPT_OWN_NFTOKEN_OFFER:      {"tecCANT_ACCEPT_OWN_NFTOKEN_OFFER", "An NFToken offer cannot be claimed by its owner. "},
	TecINSUFFICIENT_FUNDS:                 {"tecINSUFFICIENT_FUNDS", "Not enough funds available to complete requested transaction. "},
	TecOBJECT_NOT_FOUND:                   {"tecOBJECT_NOT_FOUND", "A requested object could not be located. "},
	TecINSUFFICIENT_PAYMENT:               {"tecINSUFFICIENT_PAYMENT", "The payment is not sufficient. "},
	TecINCOMPLETE:                         {"tecINCOMPLETE", "Some work was completed, but more submissions required to finish. "},
	TecXCHAIN_BAD_TRANSFER_ISSUE:          {"tecXCHAIN_BAD_TRANSFER_ISSUE", "Bad xchain transfer issue. "},
	TecXCHAIN_NO_CLAIM_ID:                 {"tecXCHAIN_NO_CLAIM_ID", "No such xchain claim id. "},
	TecXCHAIN_BAD_CLAIM_ID:                {"tecXCHAIN_BAD_CLAIM_ID", "Bad xchain claim id. "},
	TecXCHAIN_CLAIM_NO_QUORUM:             {"tecXCHAIN_CLAIM_NO_QUORUM", "Quorum was not reached on the xchain claim. "},
	TecXCHAIN_PROOF_UNKNOWN_KEY:           {"tecXCHAIN_PROOF_UNKNOWN_KEY", "Unknown key for the xchain proof. "},
	TecXCHAIN_CREATE_ACCOUNT_NONXRP_ISSUE: {"tecXCHAIN_CREATE_ACCOUNT_NONXRP_ISSUE", "Only XRP may be used for xchain create account. "},
	TecXCHAIN_WRONG_CHAIN:                 {"tecXCHAIN_WRONG_CHAIN", "XChain Transaction was submitted to the wrong chain. "},
	TecXCHAIN_REWARD_MISMATCH:             {"tecXCHAIN_REWARD_MISMATCH", "The reward amount must match the reward specified in the xchain bridge. "},
	TecXCHAIN_NO_SIGNERS_LIST:             {"tecXCHAIN_NO_SIGNERS_LIST", "The account did not have a signers list. "},
	TecXCHAIN_SENDING_ACCOUNT_MISMATCH:    {"tecXCHAIN_SENDING_ACCOUNT_MISMATCH,", "The sending account did not match the expected sending account."},
	TecXCHAIN_INSUFF_CREATE_AMOUNT:        {"tecXCHAIN_INSUFF_CREATE_AMOUNT", "Insufficient amount to create an account. "},
	TecXCHAIN_ACCOUNT_CREATE_PAST:         {"tecXCHAIN_ACCOUNT_CREATE_PAST", "The account create count has already passed. "},
	TecXCHAIN_ACCOUNT_CREATE_TOO_MANY:     {"tecXCHAIN_ACCOUNT_CREATE_TOO_MANY", "There are too many pending account create transactions to submit a new one. "},
	TecXCHAIN_PAYMENT_FAILED:              {"tecXCHAIN_PAYMENT_FAILED", "Failed to transfer funds in a xchain transaction. "},
	TecXCHAIN_SELF_COMMIT:                 {"tecXCHAIN_SELF_COMMIT", "Account cannot commit funds to itself. "},
	TecXCHAIN_BAD_PUBLIC_KEY_ACCOUNT_PAIR: {"tecXCHAIN_BAD_PUBLIC_KEY_ACCOUNT_PAIR", "Bad public key account pair in an xchain transaction. "},
	TecXCHAIN_CREATE_ACCOUNT_DISABLED:     {"tecXCHAIN_CREATE_ACCOUNT_DISABLED", "This bridge does not support account creation. "},
	TecEMPTY_DID:                          {"tecEMPTY_DID", "The DID object did not have a URI or DIDDocument field. "},

	TefALREADY:                     {"tefALREADY", "The exact transaction was already in this ledger."},
	TefBAD_ADD_AUTH:                {"tefBAD_ADD_AUTH", "Not authorized to add account."},
	TefBAD_AUTH:                    {"tefBAD_AUTH", "Transaction's public key is not authorized."},
	TefBAD_LEDGER:                  {"tefBAD_LEDGER", "Ledger in unexpected state."},
	TefBAD_QUORUM:                  {"tefBAD_QUORUM", "Signatures provided do not meet the quorum."},
	TefBAD_SIGNATURE:               {"tefBAD_SIGNATURE", "A signature is provided for a non-signer."},
	TefCREATED:                     {"tefCREATED", "Can't add an already created account."},
	TefEXCEPTION:                   {"tefEXCEPTION", "Unexpected program state."},
	TefFAILURE:                     {"tefFAILURE", "Failed to apply."},
	TefINTERNAL:                    {"tefINTERNAL", "Internal error."},
	TefMASTER_DISABLED:             {"tefMASTER_DISABLED", "Master key is disabled."},
	TefMAX_LEDGER:                  {"tefMAX_LEDGER", "Ledger sequence too high."},
	TefNO_AUTH_REQUIRED:            {"tefNO_AUTH_REQUIRED", "Auth is not required."},
	TefNOT_MULTI_SIGNING:           {"tefNOT_MULTI_SIGNING", "Account has no appropriate list of multi-signers."},
	TefPAST_SEQ:                    {"tefPAST_SEQ", "This sequence number has already passed."},
	TefWRONG_PRIOR:                 {"tefWRONG_PRIOR", "This previous transaction does not match."},
	TefBAD_AUTH_MASTER:             {"tefBAD_AUTH_MASTER", "Auth for unclaimed account needs correct master key."},
	TefINVARIANT_FAILED:            {"tefINVARIANT_FAILED", "Fee claim violated invariants for the transaction."},
	TefTOO_BIG:                     {"tefTOO_BIG", "Transaction affects too many items."},
	TefNO_TICKET:                   {"tefNO_TICKET", "Ticket is not in ledger."},
	TefNFTOKEN_IS_NOT_TRANSFERABLE: {"tefNFTOKEN_IS_NOT_TRANSFERABLE", "The specified NFToken is not transferable."},

	TelLOCAL_ERROR:                       {"telLOCAL_ERROR", "Local failure."},
	TelBAD_DOMAIN:                        {"telBAD_DOMAIN", "Domain too long."},
	TelBAD_PATH_COUNT:                    {"telBAD_PATH_COUNT", "Malformed: Too many paths."},
	TelBAD_PUBLIC_KEY:                    {"telBAD_PUBLIC_KEY", "Public key is not valid."},
	TelFAILED_PROCESSING:                 {"telFAILED_PROCESSING", "Failed to correctly process transaction."},
	TelINSUF_FEE_P:                       {"telINSUF_FEE_P", "Fee insufficient."},
	TelNO_DST_PARTIAL:                    {"telNO_DST_PARTIAL", "Partial payment to create account not allowed."},
	TelCAN_NOT_QUEUE:                     {"telCAN_NOT_QUEUE", "Can not queue at this time."},
	TelCAN_NOT_QUEUE_BALANCE:             {"telCAN_NOT_QUEUE_BALANCE", "Can not queue at this time: insufficient balance to pay all queued fees."},
	TelCAN_NOT_QUEUE_BLOCKS:              {"telCAN_NOT_QUEUE_BLOCKS", "Can not queue at this time: would block later queued transaction: {s)."},
	TelCAN_NOT_QUEUE_BLOCKED:             {"telCAN_NOT_QUEUE_BLOCKED", "Can not queue at this time: blocking transaction in queue."},
	TelCAN_NOT_QUEUE_FEE:                 {"telCAN_NOT_QUEUE_FEE", "Can not queue at this time: fee insufficient to replace queued transaction."},
	TelCAN_NOT_QUEUE_FULL:                {"telCAN_NOT_QUEUE_FULL", "Can not queue at this time: queue is full."},
	TelWRONG_NETWORK:                     {"telWRONG_NETWORK", "Transaction specifies a Network ID that differs from that of the local node."},
	TelREQUIRES_NETWORK_ID:               {"telREQUIRES_NETWORK_ID", "Transactions submitted to this node/network must include a correct NetworkID field."},
	TelNETWORK_ID_MAKES_TX_NON_CANONICAL: {"telNETWORK_ID_MAKES_TX_NON_CANONICAL", "Transactions submitted to this node/network must NOT include a NetworkID field."},

	TemMALFORMED:                   {"temMALFORMED", "Malformed transaction."},
	TemBAD_AMM_TOKENS:              {"temBAD_AMM_TOKENS", "Malformed: Invalid LPTokens."},
	TemBAD_AMOUNT:                  {"temBAD_AMOUNT", "Can only send positive amounts."},
	TemBAD_CURRENCY:                {"temBAD_CURRENCY", "Malformed: Bad currency."},
	TemBAD_EXPIRATION:              {"temBAD_EXPIRATION", "Malformed: Bad expiration."},
	TemBAD_FEE:                     {"temBAD_FEE", "Invalid fee, negative or not XRP."},
	TemBAD_ISSUER:                  {"temBAD_ISSUER", "Malformed: Bad issuer."},
	TemBAD_LIMIT:                   {"temBAD_LIMIT", "Limits must be non-negative."},
	TemBAD_OFFER:                   {"temBAD_OFFER", "Malformed: Bad offer."},
	TemBAD_PATH:                    {"temBAD_PATH", "Malformed: Bad path."},
	TemBAD_PATH_LOOP:               {"temBAD_PATH_LOOP", "Malformed: Loop in path."},
	TemBAD_QUORUM:                  {"temBAD_QUORUM", "Malformed: Quorum is unreachable."},
	TemBAD_REGKEY:                  {"temBAD_REGKEY", "Malformed: Regular key cannot be same as master key."},
	TemBAD_SEND_XRP_LIMIT:          {"temBAD_SEND_XRP_LIMIT", "Malformed: Limit quality is not allowed for XRP to XRP."},
	TemBAD_SEND_XRP_MAX:            {"temBAD_SEND_XRP_MAX", "Malformed: Send max is not allowed for XRP to XRP."},
	TemBAD_SEND_XRP_NO_DIRECT:      {"temBAD_SEND_XRP_NO_DIRECT", "Malformed: No Ripple direct is not allowed for XRP to XRP."},
	TemBAD_SEND_XRP_PARTIAL:        {"temBAD_SEND_XRP_PARTIAL", "Malformed: Partial payment is not allowed for XRP to XRP."},
	TemBAD_SEND_XRP_PATHS:          {"temBAD_SEND_XRP_PATHS", "Malformed: Paths are not allowed for XRP to XRP."},
	TemBAD_SEQUENCE:                {"temBAD_SEQUENCE", "Malformed: Sequence is not in the past."},
	TemBAD_SIGNATURE:               {"temBAD_SIGNATURE", "Malformed: Bad signature."},
	TemBAD_SIGNER:                  {"temBAD_SIGNER", "Malformed: No signer may duplicate account or other signers."},
	TemBAD_SRC_ACCOUNT:             {"temBAD_SRC_ACCOUNT", "Malformed: Bad source account."},
	TemBAD_TRANSFER_RATE:           {"temBAD_TRANSFER_RATE", "Malformed: Transfer rate must be >= 1.0 and <= 2.0"},
	TemBAD_WEIGHT:                  {"temBAD_WEIGHT", "Malformed: Weight must be a positive value."},
	TemDST_IS_SRC:                  {"temDST_IS_SRC", "Destination may not be source."},
	TemDST_NEEDED:                  {"temDST_NEEDED", "Destination not specified."},
	TemEMPTY_DID:                   {"temEMPTY_DID", "Malformed: No DID data provided."},
	TemINVALID:                     {"temINVALID", "The transaction is ill-formed."},
	TemINVALID_FLAG:                {"temINVALID_FLAG", "The transaction has an invalid flag."},
	TemREDUNDANT:                   {"temREDUNDANT", "The transaction is redundant."},
	TemRIPPLE_EMPTY:                {"temRIPPLE_EMPTY", "PathSet with no paths."},
	TemUNCERTAIN:                   {"temUNCERTAIN", "In process of determining result. Never returned."},
	TemUNKNOWN:                     {"temUNKNOWN", "The transaction requires logic that is not implemented yet."},
	TemDISABLED:                    {"temDISABLED", "The transaction requires logic that is currently disabled."},
	TemBAD_TICK_SIZE:               {"temBAD_TICK_SIZE", "Malformed: Tick size out of range."},
	TemINVALID_ACCOUNT_ID:          {"temINVALID_ACCOUNT_ID", "Malformed: A field contains an invalid account ID."},
	TemCANNOT_PREAUTH_SELF:         {"temCANNOT_PREAUTH_SELF", "Malformed: An account may not preauthorize itself."},
	TemINVALID_COUNT:               {"temINVALID_COUNT", "Malformed: Count field outside valid range."},
	TemSEQ_AND_TICKET:              {"temSEQ_AND_TICKET", "Transaction contains a TicketSequence and a non-zero Sequence."},
	TemBAD_NFTOKEN_TRANSFER_FEE:    {"temBAD_NFTOKEN_TRANSFER_FEE", "Malformed: The NFToken transfer fee must be between 1 and 5000, inclusive."},
	TemXCHAIN_EQUAL_DOOR_ACCOUNTS:  {"temXCHAIN_EQUAL_DOOR_ACCOUNTS", "Malformed: Bridge must have unique door accounts."},
	TemXCHAIN_BAD_PROOF:            {"temXCHAIN_BAD_PROOF", "Malformed: Bad cross-chain claim proof."},
	TemXCHAIN_BRIDGE_BAD_ISSUES:    {"temXCHAIN_BRIDGE_BAD_ISSUES", "Malformed: Bad bridge issues."},
	TemXCHAIN_BRIDGE_NONDOOR_OWNER: {"temXCHAIN_BRIDGE_NONDOOR_OWNER", "Malformed: Bridge owner must be one of the door accounts."},
	TemXCHAIN_BRIDGE_BAD_MIN_ACCOUNT_CREATE_AMOUNT: {"temXCHAIN_BRIDGE_BAD_MIN_ACCOUNT_CREATE_AMOUNT", "Malformed: Bad min account create amount."},
	TemXCHAIN_BRIDGE_BAD_REWARD_AMOUNT:             {"temXCHAIN_BRIDGE_BAD_REWARD_AMOUNT", "Malformed: Bad reward amount."},

	TerRETRY:       {"terRETRY", "Retry transaction."},
	TerFUNDS_SPENT: {"terFUNDS_SPENT", "DEPRECATED."},
	TerINSUF_FEE_B: {"terINSUF_FEE_B", "Account balance can't pay fee."},
	TerLAST:        {"terLAST", "DEPRECATED."},
	TerNO_RIPPLE:   {"terNO_RIPPLE", "Path does not permit rippling."},
	TerNO_ACCOUNT:  {"terNO_ACCOUNT", "The source account does not exist."},
	TerNO_AUTH:     {"terNO_AUTH", "Not authorized to hold IOUs."},
	TerNO_LINE:     {"terNO_LINE", "No such line."},
	TerPRE_SEQ:     {"terPRE_SEQ", "Missing/inapplicable prior transaction."},
	TerOWNERS:      {"terOWNERS", "Non-zero owner count."},
	TerQUEUED:      {"terQUEUED", "Held until escalated fee drops."},
	TerPRE_TICKET:  {"terPRE_TICKET", "Ticket is not yet in ledger."},
	TerNO_AMM:      {"terNO_AMM", "AMM doesn't exist for the asset pair."},
}

var reverseResults map[string]TransactionResult

func init() {
	reverseResults = make(map[string]TransactionResult)
	for result, name := range resultNames {
		reverseResults[name.Token] = result
	}
}

func (r TransactionResult) String() string {
	return resultNames[r].Token
}

func (r TransactionResult) Human() string {
	return resultNames[r].Human
}

func (r TransactionResult) Success() bool {
	return r == TesSUCCESS
}

func (r TransactionResult) Queued() bool {
	return r == TerQUEUED
}

func (r TransactionResult) Symbol() string {
	switch r {
	case TesSUCCESS, TecCLAIM:
		return "✓"
	case TecPATH_PARTIAL, TecPATH_DRY:
		return "½"
	case TecUNFUNDED, TecUNFUNDED_ADD, TecUNFUNDED_OFFER, TecUNFUNDED_PAYMENT:
		return "$"
	default:
		return "✗"
	}
}
