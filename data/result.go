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
	tesSUCCESS TransactionResult = 0
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
	tecCLAIM                              = TransactionResult(100)
	tecPATH_PARTIAL                       = TransactionResult(101)
	tecUNFUNDED_ADD                       = TransactionResult(102) // Unused legacy code
	tecUNFUNDED_OFFER                     = TransactionResult(103)
	tecUNFUNDED_PAYMENT                   = TransactionResult(104)
	tecFAILED_PROCESSING                  = TransactionResult(105)
	tecDIR_FULL                           = TransactionResult(121)
	tecINSUF_RESERVE_LINE                 = TransactionResult(122)
	tecINSUF_RESERVE_OFFER                = TransactionResult(123)
	tecNO_DST                             = TransactionResult(124)
	tecNO_DST_INSUF_XRP                   = TransactionResult(125)
	tecNO_LINE_INSUF_RESERVE              = TransactionResult(126)
	tecNO_LINE_REDUNDANT                  = TransactionResult(127)
	tecPATH_DRY                           = TransactionResult(128)
	tecUNFUNDED                           = TransactionResult(129)
	tecNO_ALTERNATIVE_KEY                 = TransactionResult(130)
	tecNO_REGULAR_KEY                     = TransactionResult(131)
	tecOWNERS                             = TransactionResult(132)
	tecNO_ISSUER                          = TransactionResult(133)
	tecNO_AUTH                            = TransactionResult(134)
	tecNO_LINE                            = TransactionResult(135)
	tecINSUFF_FEE                         = TransactionResult(136)
	tecFROZEN                             = TransactionResult(137)
	tecNO_TARGET                          = TransactionResult(138)
	tecNO_PERMISSION                      = TransactionResult(139)
	tecNO_ENTRY                           = TransactionResult(140)
	tecINSUFFICIENT_RESERVE               = TransactionResult(141)
	tecNEED_MASTER_KEY                    = TransactionResult(142)
	tecDST_TAG_NEEDED                     = TransactionResult(143)
	tecINTERNAL                           = TransactionResult(144)
	tecOVERSIZE                           = TransactionResult(145)
	tecCRYPTOCONDITION_ERROR              = TransactionResult(146)
	tecINVARIANT_FAILED                   = TransactionResult(147)
	tecEXPIRED                            = TransactionResult(148)
	tecDUPLICATE                          = TransactionResult(149)
	tecKILLED                             = TransactionResult(150)
	tecHAS_OBLIGATIONS                    = TransactionResult(151)
	tecTOO_SOON                           = TransactionResult(152)
	tecHOOK_REJECTED                      = TransactionResult(153)
	tecMAX_SEQUENCE_REACHED               = TransactionResult(154)
	tecNO_SUITABLE_NFTOKEN_PAGE           = TransactionResult(155)
	tecNFTOKEN_BUY_SELL_MISMATCH          = TransactionResult(156)
	tecNFTOKEN_OFFER_TYPE_MISMATCH        = TransactionResult(157)
	tecCANT_ACCEPT_OWN_NFTOKEN_OFFER      = TransactionResult(158)
	tecINSUFFICIENT_FUNDS                 = TransactionResult(159)
	tecOBJECT_NOT_FOUND                   = TransactionResult(160)
	tecINSUFFICIENT_PAYMENT               = TransactionResult(161)
	tecUNFUNDED_AMM                       = TransactionResult(162)
	tecAMM_BALANCE                        = TransactionResult(163)
	tecAMM_FAILED                         = TransactionResult(164)
	tecAMM_INVALID_TOKENS                 = TransactionResult(165)
	tecAMM_EMPTY                          = TransactionResult(166)
	tecAMM_NOT_EMPTY                      = TransactionResult(167)
	tecAMM_ACCOUNT                        = TransactionResult(168)
	tecINCOMPLETE                         = TransactionResult(169)
	tecXCHAIN_BAD_TRANSFER_ISSUE          = TransactionResult(170)
	tecXCHAIN_NO_CLAIM_ID                 = TransactionResult(171)
	tecXCHAIN_BAD_CLAIM_ID                = TransactionResult(172)
	tecXCHAIN_CLAIM_NO_QUORUM             = TransactionResult(173)
	tecXCHAIN_PROOF_UNKNOWN_KEY           = TransactionResult(174)
	tecXCHAIN_CREATE_ACCOUNT_NONXRP_ISSUE = TransactionResult(175)
	tecXCHAIN_WRONG_CHAIN                 = TransactionResult(176)
	tecXCHAIN_REWARD_MISMATCH             = TransactionResult(177)
	tecXCHAIN_NO_SIGNERS_LIST             = TransactionResult(178)
	tecXCHAIN_SENDING_ACCOUNT_MISMATCH    = TransactionResult(179)
	tecXCHAIN_INSUFF_CREATE_AMOUNT        = TransactionResult(180)
	tecXCHAIN_ACCOUNT_CREATE_PAST         = TransactionResult(181)
	tecXCHAIN_ACCOUNT_CREATE_TOO_MANY     = TransactionResult(182)
	tecXCHAIN_PAYMENT_FAILED              = TransactionResult(183)
	tecXCHAIN_SELF_COMMIT                 = TransactionResult(184)
	tecXCHAIN_BAD_PUBLIC_KEY_ACCOUNT_PAIR = TransactionResult(185)
	tecXCHAIN_CREATE_ACCOUNT_DISABLED     = TransactionResult(186)
	tecEMPTY_DID                          = TransactionResult(187)
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
	telLOCAL_ERROR TransactionResult = iota - 399
	telBAD_DOMAIN
	telBAD_PATH_COUNT
	telBAD_PUBLIC_KEY
	telFAILED_PROCESSING
	telINSUF_FEE_P
	telNO_DST_PARTIAL
	telCAN_NOT_QUEUE
	telCAN_NOT_QUEUE_BALANCE
	telCAN_NOT_QUEUE_BLOCKS
	telCAN_NOT_QUEUE_BLOCKED
	telCAN_NOT_QUEUE_FEE
	telCAN_NOT_QUEUE_FULL
	telWRONG_NETWORK
	telREQUIRES_NETWORK_ID
	telNETWORK_ID_MAKES_TX_NON_CANONICAL
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
	temMALFORMED TransactionResult = iota - 299
	temBAD_AMOUNT
	temBAD_CURRENCY
	temBAD_EXPIRATION
	temBAD_FEE
	temBAD_ISSUER
	temBAD_LIMIT
	temBAD_OFFER
	temBAD_PATH
	temBAD_PATH_LOOP
	temBAD_REGKEY
	temBAD_SEND_XRP_LIMIT
	temBAD_SEND_XRP_MAX
	temBAD_SEND_XRP_NO_DIRECT
	temBAD_SEND_XRP_PARTIAL
	temBAD_SEND_XRP_PATHS
	temBAD_SEQUENCE
	temBAD_SIGNATURE
	temBAD_SRC_ACCOUNT
	temBAD_TRANSFER_RATE
	temDST_IS_SRC
	temDST_NEEDED
	temINVALID
	temINVALID_FLAG
	temREDUNDANT
	temRIPPLE_EMPTY
	temDISABLED
	temBAD_SIGNER
	temBAD_QUORUM
	temBAD_WEIGHT
	temBAD_TICK_SIZE
	temINVALID_ACCOUNT_ID
	temCANNOT_PREAUTH_SELF
	temINVALID_COUNT
	temUNCERTAIN // An internal intermediate result; should never be returned.
	temUNKNOWN   // An internal intermediate result; should never be returned.
	temSEQ_AND_TICKET
	temBAD_NFTOKEN_TRANSFER_FEE
	temBAD_AMM_TOKENS
	temXCHAIN_EQUAL_DOOR_ACCOUNTS
	temXCHAIN_BAD_PROOF
	temXCHAIN_BRIDGE_BAD_ISSUES
	temXCHAIN_BRIDGE_NONDOOR_OWNER
	temXCHAIN_BRIDGE_BAD_MIN_ACCOUNT_CREATE_AMOUNT
	temXCHAIN_BRIDGE_BAD_REWARD_AMOUNT
	temEMPTY_DID
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
	tefFAILURE TransactionResult = iota - 199
	tefALREADY
	tefBAD_ADD_AUTH
	tefBAD_AUTH
	tefBAD_LEDGER
	tefCREATED
	tefEXCEPTION
	tefINTERNAL
	tefNO_AUTH_REQUIRED // Can't set auth if auth is not required.
	tefPAST_SEQ
	tefWRONG_PRIOR
	tefMASTER_DISABLED
	tefMAX_LEDGER
	tefBAD_SIGNATURE
	tefBAD_QUORUM
	tefNOT_MULTI_SIGNING
	tefBAD_AUTH_MASTER
	tefINVARIANT_FAILED
	tefTOO_BIG
	tefNO_TICKET
	tefNFTOKEN_IS_NOT_TRANSFERABLE
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
	terRETRY       TransactionResult = iota - 99
	terFUNDS_SPENT                   // DEPRECATED.
	terINSUF_FEE_B                   // Can't pay fee, therefore don't burden network.
	terNO_ACCOUNT                    // Can't pay fee, therefore don't burden network.
	terNO_AUTH                       // Not authorized to hold IOUs.
	terNO_LINE                       // Internal flag.
	terOWNERS                        // Can't succeed with non-zero owner count.
	terPRE_SEQ                       // Can't pay fee, no point in forwarding, so don't
	// burden network.
	terLAST       // DEPRECATED.
	terNO_RIPPLE  // Rippling not allowed
	terQUEUED     // Transaction is being held in TxQ until fee drops
	terPRE_TICKET // Ticket is not yet in ledger but might be on its way
	terNO_AMM     // AMM doesn't exist for the asset pair
)

var resultNames = map[TransactionResult]struct {
	Token string
	Human string
}{
	tesSUCCESS: {"tesSUCCESS", "The transaction was applied."},

	tecAMM_BALANCE:                        {"tecAMM_BALANCE", "AMM has invalid balance. "},
	tecAMM_INVALID_TOKENS:                 {"tecAMM_INVALID_TOKENS", "AMM invalid LP tokens. "},
	tecAMM_FAILED:                         {"tecAMM_FAILED", "AMM transaction failed. "},
	tecAMM_EMPTY:                          {"tecAMM_EMPTY", "AMM is in empty state. "},
	tecAMM_NOT_EMPTY:                      {"tecAMM_NOT_EMPTY", "AMM is not in empty state. "},
	tecAMM_ACCOUNT:                        {"tecAMM_ACCOUNT", "This operation is not allowed on an AMM Account. "},
	tecCLAIM:                              {"tecCLAIM", "Fee claimed. Sequence used. No action. "},
	tecDIR_FULL:                           {"tecDIR_FULL", "Can not add entry to full directory. "},
	tecFAILED_PROCESSING:                  {"tecFAILED_PROCESSING", "Failed to correctly process transaction. "},
	tecINSUF_RESERVE_LINE:                 {"tecINSUF_RESERVE_LINE", "Insufficient reserve to add trust line. "},
	tecINSUF_RESERVE_OFFER:                {"tecINSUF_RESERVE_OFFER", "Insufficient reserve to create offer. "},
	tecNO_DST:                             {"tecNO_DST", "Destination does not exist. Send XRP to create it. "},
	tecNO_DST_INSUF_XRP:                   {"tecNO_DST_INSUF_XRP", "Destination does not exist. Too little XRP sent to create it. "},
	tecNO_LINE_INSUF_RESERVE:              {"tecNO_LINE_INSUF_RESERVE", "No such line. Too little reserve to create it. "},
	tecNO_LINE_REDUNDANT:                  {"tecNO_LINE_REDUNDANT", "Can't set non-existent line to default. "},
	tecPATH_DRY:                           {"tecPATH_DRY", "Path could not send partial amount. "},
	tecPATH_PARTIAL:                       {"tecPATH_PARTIAL", "Path could not send full amount. "},
	tecNO_ALTERNATIVE_KEY:                 {"tecNO_ALTERNATIVE_KEY", "The operation would remove the ability to sign transactions with the account. "},
	tecNO_REGULAR_KEY:                     {"tecNO_REGULAR_KEY", "Regular key is not set. "},
	tecOVERSIZE:                           {"tecOVERSIZE", "Object exceeded serialization limits. "},
	tecUNFUNDED:                           {"tecUNFUNDED", "Not enough XRP to satisfy the reserve requirement. "},
	tecUNFUNDED_ADD:                       {"tecUNFUNDED_ADD", "DEPRECATED. "},
	tecUNFUNDED_AMM:                       {"tecUNFUNDED_AMM", "Insufficient balance to fund AMM. "},
	tecUNFUNDED_OFFER:                     {"tecUNFUNDED_OFFER", "Insufficient balance to fund created offer. "},
	tecUNFUNDED_PAYMENT:                   {"tecUNFUNDED_PAYMENT", "Insufficient XRP balance to send. "},
	tecOWNERS:                             {"tecOWNERS", "Non-zero owner count. "},
	tecNO_ISSUER:                          {"tecNO_ISSUER", "Issuer account does not exist. "},
	tecNO_AUTH:                            {"tecNO_AUTH", "Not authorized to hold asset. "},
	tecNO_LINE:                            {"tecNO_LINE", "No such line. "},
	tecINSUFF_FEE:                         {"tecINSUFF_FEE", "Insufficient balance to pay fee. "},
	tecFROZEN:                             {"tecFROZEN", "Asset is frozen. "},
	tecNO_TARGET:                          {"tecNO_TARGET", "Target account does not exist. "},
	tecNO_PERMISSION:                      {"tecNO_PERMISSION", "No permission to perform requested operation. "},
	tecNO_ENTRY:                           {"tecNO_ENTRY", "No matching entry found. "},
	tecINSUFFICIENT_RESERVE:               {"tecINSUFFICIENT_RESERVE", "Insufficient reserve to complete requested operation. "},
	tecNEED_MASTER_KEY:                    {"tecNEED_MASTER_KEY", "The operation requires the use of the Master Key. "},
	tecDST_TAG_NEEDED:                     {"tecDST_TAG_NEEDED", "A destination tag is required. "},
	tecINTERNAL:                           {"tecINTERNAL", "An internal error has occurred during processing. "},
	tecCRYPTOCONDITION_ERROR:              {"tecCRYPTOCONDITION_ERROR", "Malformed, invalid, or mismatched conditional or fulfillment. "},
	tecINVARIANT_FAILED:                   {"tecINVARIANT_FAILED", "One or more invariants for the transaction were not satisfied. "},
	tecEXPIRED:                            {"tecEXPIRED", "Expiration time is passed. "},
	tecDUPLICATE:                          {"tecDUPLICATE", "Ledger object already exists. "},
	tecKILLED:                             {"tecKILLED", "No funds transferred and no offer created. "},
	tecHAS_OBLIGATIONS:                    {"tecHAS_OBLIGATIONS", "The account cannot be deleted since it has obligations. "},
	tecTOO_SOON:                           {"tecTOO_SOON", "It is too early to attempt the requested operation. Please wait. "},
	tecMAX_SEQUENCE_REACHED:               {"tecMAX_SEQUENCE_REACHED", "The maximum sequence number was reached. "},
	tecNO_SUITABLE_NFTOKEN_PAGE:           {"tecNO_SUITABLE_NFTOKEN_PAGE", "A suitable NFToken page could not be located. "},
	tecNFTOKEN_BUY_SELL_MISMATCH:          {"tecNFTOKEN_BUY_SELL_MISMATCH", "The 'Buy' and 'Sell' NFToken offers are mismatched. "},
	tecNFTOKEN_OFFER_TYPE_MISMATCH:        {"tecNFTOKEN_OFFER_TYPE_MISMATCH", "The type of NFToken offer is incorrect. "},
	tecCANT_ACCEPT_OWN_NFTOKEN_OFFER:      {"tecCANT_ACCEPT_OWN_NFTOKEN_OFFER", "An NFToken offer cannot be claimed by its owner. "},
	tecINSUFFICIENT_FUNDS:                 {"tecINSUFFICIENT_FUNDS", "Not enough funds available to complete requested transaction. "},
	tecOBJECT_NOT_FOUND:                   {"tecOBJECT_NOT_FOUND", "A requested object could not be located. "},
	tecINSUFFICIENT_PAYMENT:               {"tecINSUFFICIENT_PAYMENT", "The payment is not sufficient. "},
	tecINCOMPLETE:                         {"tecINCOMPLETE", "Some work was completed, but more submissions required to finish. "},
	tecXCHAIN_BAD_TRANSFER_ISSUE:          {"tecXCHAIN_BAD_TRANSFER_ISSUE", "Bad xchain transfer issue. "},
	tecXCHAIN_NO_CLAIM_ID:                 {"tecXCHAIN_NO_CLAIM_ID", "No such xchain claim id. "},
	tecXCHAIN_BAD_CLAIM_ID:                {"tecXCHAIN_BAD_CLAIM_ID", "Bad xchain claim id. "},
	tecXCHAIN_CLAIM_NO_QUORUM:             {"tecXCHAIN_CLAIM_NO_QUORUM", "Quorum was not reached on the xchain claim. "},
	tecXCHAIN_PROOF_UNKNOWN_KEY:           {"tecXCHAIN_PROOF_UNKNOWN_KEY", "Unknown key for the xchain proof. "},
	tecXCHAIN_CREATE_ACCOUNT_NONXRP_ISSUE: {"tecXCHAIN_CREATE_ACCOUNT_NONXRP_ISSUE", "Only XRP may be used for xchain create account. "},
	tecXCHAIN_WRONG_CHAIN:                 {"tecXCHAIN_WRONG_CHAIN", "XChain Transaction was submitted to the wrong chain. "},
	tecXCHAIN_REWARD_MISMATCH:             {"tecXCHAIN_REWARD_MISMATCH", "The reward amount must match the reward specified in the xchain bridge. "},
	tecXCHAIN_NO_SIGNERS_LIST:             {"tecXCHAIN_NO_SIGNERS_LIST", "The account did not have a signers list. "},
	tecXCHAIN_SENDING_ACCOUNT_MISMATCH:    {"tecXCHAIN_SENDING_ACCOUNT_MISMATCH,", "The sending account did not match the expected sending account."},
	tecXCHAIN_INSUFF_CREATE_AMOUNT:        {"tecXCHAIN_INSUFF_CREATE_AMOUNT", "Insufficient amount to create an account. "},
	tecXCHAIN_ACCOUNT_CREATE_PAST:         {"tecXCHAIN_ACCOUNT_CREATE_PAST", "The account create count has already passed. "},
	tecXCHAIN_ACCOUNT_CREATE_TOO_MANY:     {"tecXCHAIN_ACCOUNT_CREATE_TOO_MANY", "There are too many pending account create transactions to submit a new one. "},
	tecXCHAIN_PAYMENT_FAILED:              {"tecXCHAIN_PAYMENT_FAILED", "Failed to transfer funds in a xchain transaction. "},
	tecXCHAIN_SELF_COMMIT:                 {"tecXCHAIN_SELF_COMMIT", "Account cannot commit funds to itself. "},
	tecXCHAIN_BAD_PUBLIC_KEY_ACCOUNT_PAIR: {"tecXCHAIN_BAD_PUBLIC_KEY_ACCOUNT_PAIR", "Bad public key account pair in an xchain transaction. "},
	tecXCHAIN_CREATE_ACCOUNT_DISABLED:     {"tecXCHAIN_CREATE_ACCOUNT_DISABLED", "This bridge does not support account creation. "},
	tecEMPTY_DID:                          {"tecEMPTY_DID", "The DID object did not have a URI or DIDDocument field. "},

	tefALREADY:                     {"tefALREADY", "The exact transaction was already in this ledger."},
	tefBAD_ADD_AUTH:                {"tefBAD_ADD_AUTH", "Not authorized to add account."},
	tefBAD_AUTH:                    {"tefBAD_AUTH", "Transaction's public key is not authorized."},
	tefBAD_LEDGER:                  {"tefBAD_LEDGER", "Ledger in unexpected state."},
	tefBAD_QUORUM:                  {"tefBAD_QUORUM", "Signatures provided do not meet the quorum."},
	tefBAD_SIGNATURE:               {"tefBAD_SIGNATURE", "A signature is provided for a non-signer."},
	tefCREATED:                     {"tefCREATED", "Can't add an already created account."},
	tefEXCEPTION:                   {"tefEXCEPTION", "Unexpected program state."},
	tefFAILURE:                     {"tefFAILURE", "Failed to apply."},
	tefINTERNAL:                    {"tefINTERNAL", "Internal error."},
	tefMASTER_DISABLED:             {"tefMASTER_DISABLED", "Master key is disabled."},
	tefMAX_LEDGER:                  {"tefMAX_LEDGER", "Ledger sequence too high."},
	tefNO_AUTH_REQUIRED:            {"tefNO_AUTH_REQUIRED", "Auth is not required."},
	tefNOT_MULTI_SIGNING:           {"tefNOT_MULTI_SIGNING", "Account has no appropriate list of multi-signers."},
	tefPAST_SEQ:                    {"tefPAST_SEQ", "This sequence number has already passed."},
	tefWRONG_PRIOR:                 {"tefWRONG_PRIOR", "This previous transaction does not match."},
	tefBAD_AUTH_MASTER:             {"tefBAD_AUTH_MASTER", "Auth for unclaimed account needs correct master key."},
	tefINVARIANT_FAILED:            {"tefINVARIANT_FAILED", "Fee claim violated invariants for the transaction."},
	tefTOO_BIG:                     {"tefTOO_BIG", "Transaction affects too many items."},
	tefNO_TICKET:                   {"tefNO_TICKET", "Ticket is not in ledger."},
	tefNFTOKEN_IS_NOT_TRANSFERABLE: {"tefNFTOKEN_IS_NOT_TRANSFERABLE", "The specified NFToken is not transferable."},

	telLOCAL_ERROR:                       {"telLOCAL_ERROR", "Local failure."},
	telBAD_DOMAIN:                        {"telBAD_DOMAIN", "Domain too long."},
	telBAD_PATH_COUNT:                    {"telBAD_PATH_COUNT", "Malformed: Too many paths."},
	telBAD_PUBLIC_KEY:                    {"telBAD_PUBLIC_KEY", "Public key is not valid."},
	telFAILED_PROCESSING:                 {"telFAILED_PROCESSING", "Failed to correctly process transaction."},
	telINSUF_FEE_P:                       {"telINSUF_FEE_P", "Fee insufficient."},
	telNO_DST_PARTIAL:                    {"telNO_DST_PARTIAL", "Partial payment to create account not allowed."},
	telCAN_NOT_QUEUE:                     {"telCAN_NOT_QUEUE", "Can not queue at this time."},
	telCAN_NOT_QUEUE_BALANCE:             {"telCAN_NOT_QUEUE_BALANCE", "Can not queue at this time: insufficient balance to pay all queued fees."},
	telCAN_NOT_QUEUE_BLOCKS:              {"telCAN_NOT_QUEUE_BLOCKS", "Can not queue at this time: would block later queued transaction: {s)."},
	telCAN_NOT_QUEUE_BLOCKED:             {"telCAN_NOT_QUEUE_BLOCKED", "Can not queue at this time: blocking transaction in queue."},
	telCAN_NOT_QUEUE_FEE:                 {"telCAN_NOT_QUEUE_FEE", "Can not queue at this time: fee insufficient to replace queued transaction."},
	telCAN_NOT_QUEUE_FULL:                {"telCAN_NOT_QUEUE_FULL", "Can not queue at this time: queue is full."},
	telWRONG_NETWORK:                     {"telWRONG_NETWORK", "Transaction specifies a Network ID that differs from that of the local node."},
	telREQUIRES_NETWORK_ID:               {"telREQUIRES_NETWORK_ID", "Transactions submitted to this node/network must include a correct NetworkID field."},
	telNETWORK_ID_MAKES_TX_NON_CANONICAL: {"telNETWORK_ID_MAKES_TX_NON_CANONICAL", "Transactions submitted to this node/network must NOT include a NetworkID field."},

	temMALFORMED:                   {"temMALFORMED", "Malformed transaction."},
	temBAD_AMM_TOKENS:              {"temBAD_AMM_TOKENS", "Malformed: Invalid LPTokens."},
	temBAD_AMOUNT:                  {"temBAD_AMOUNT", "Can only send positive amounts."},
	temBAD_CURRENCY:                {"temBAD_CURRENCY", "Malformed: Bad currency."},
	temBAD_EXPIRATION:              {"temBAD_EXPIRATION", "Malformed: Bad expiration."},
	temBAD_FEE:                     {"temBAD_FEE", "Invalid fee, negative or not XRP."},
	temBAD_ISSUER:                  {"temBAD_ISSUER", "Malformed: Bad issuer."},
	temBAD_LIMIT:                   {"temBAD_LIMIT", "Limits must be non-negative."},
	temBAD_OFFER:                   {"temBAD_OFFER", "Malformed: Bad offer."},
	temBAD_PATH:                    {"temBAD_PATH", "Malformed: Bad path."},
	temBAD_PATH_LOOP:               {"temBAD_PATH_LOOP", "Malformed: Loop in path."},
	temBAD_QUORUM:                  {"temBAD_QUORUM", "Malformed: Quorum is unreachable."},
	temBAD_REGKEY:                  {"temBAD_REGKEY", "Malformed: Regular key cannot be same as master key."},
	temBAD_SEND_XRP_LIMIT:          {"temBAD_SEND_XRP_LIMIT", "Malformed: Limit quality is not allowed for XRP to XRP."},
	temBAD_SEND_XRP_MAX:            {"temBAD_SEND_XRP_MAX", "Malformed: Send max is not allowed for XRP to XRP."},
	temBAD_SEND_XRP_NO_DIRECT:      {"temBAD_SEND_XRP_NO_DIRECT", "Malformed: No Ripple direct is not allowed for XRP to XRP."},
	temBAD_SEND_XRP_PARTIAL:        {"temBAD_SEND_XRP_PARTIAL", "Malformed: Partial payment is not allowed for XRP to XRP."},
	temBAD_SEND_XRP_PATHS:          {"temBAD_SEND_XRP_PATHS", "Malformed: Paths are not allowed for XRP to XRP."},
	temBAD_SEQUENCE:                {"temBAD_SEQUENCE", "Malformed: Sequence is not in the past."},
	temBAD_SIGNATURE:               {"temBAD_SIGNATURE", "Malformed: Bad signature."},
	temBAD_SIGNER:                  {"temBAD_SIGNER", "Malformed: No signer may duplicate account or other signers."},
	temBAD_SRC_ACCOUNT:             {"temBAD_SRC_ACCOUNT", "Malformed: Bad source account."},
	temBAD_TRANSFER_RATE:           {"temBAD_TRANSFER_RATE", "Malformed: Transfer rate must be >= 1.0 and <= 2.0"},
	temBAD_WEIGHT:                  {"temBAD_WEIGHT", "Malformed: Weight must be a positive value."},
	temDST_IS_SRC:                  {"temDST_IS_SRC", "Destination may not be source."},
	temDST_NEEDED:                  {"temDST_NEEDED", "Destination not specified."},
	temEMPTY_DID:                   {"temEMPTY_DID", "Malformed: No DID data provided."},
	temINVALID:                     {"temINVALID", "The transaction is ill-formed."},
	temINVALID_FLAG:                {"temINVALID_FLAG", "The transaction has an invalid flag."},
	temREDUNDANT:                   {"temREDUNDANT", "The transaction is redundant."},
	temRIPPLE_EMPTY:                {"temRIPPLE_EMPTY", "PathSet with no paths."},
	temUNCERTAIN:                   {"temUNCERTAIN", "In process of determining result. Never returned."},
	temUNKNOWN:                     {"temUNKNOWN", "The transaction requires logic that is not implemented yet."},
	temDISABLED:                    {"temDISABLED", "The transaction requires logic that is currently disabled."},
	temBAD_TICK_SIZE:               {"temBAD_TICK_SIZE", "Malformed: Tick size out of range."},
	temINVALID_ACCOUNT_ID:          {"temINVALID_ACCOUNT_ID", "Malformed: A field contains an invalid account ID."},
	temCANNOT_PREAUTH_SELF:         {"temCANNOT_PREAUTH_SELF", "Malformed: An account may not preauthorize itself."},
	temINVALID_COUNT:               {"temINVALID_COUNT", "Malformed: Count field outside valid range."},
	temSEQ_AND_TICKET:              {"temSEQ_AND_TICKET", "Transaction contains a TicketSequence and a non-zero Sequence."},
	temBAD_NFTOKEN_TRANSFER_FEE:    {"temBAD_NFTOKEN_TRANSFER_FEE", "Malformed: The NFToken transfer fee must be between 1 and 5000, inclusive."},
	temXCHAIN_EQUAL_DOOR_ACCOUNTS:  {"temXCHAIN_EQUAL_DOOR_ACCOUNTS", "Malformed: Bridge must have unique door accounts."},
	temXCHAIN_BAD_PROOF:            {"temXCHAIN_BAD_PROOF", "Malformed: Bad cross-chain claim proof."},
	temXCHAIN_BRIDGE_BAD_ISSUES:    {"temXCHAIN_BRIDGE_BAD_ISSUES", "Malformed: Bad bridge issues."},
	temXCHAIN_BRIDGE_NONDOOR_OWNER: {"temXCHAIN_BRIDGE_NONDOOR_OWNER", "Malformed: Bridge owner must be one of the door accounts."},
	temXCHAIN_BRIDGE_BAD_MIN_ACCOUNT_CREATE_AMOUNT: {"temXCHAIN_BRIDGE_BAD_MIN_ACCOUNT_CREATE_AMOUNT", "Malformed: Bad min account create amount."},
	temXCHAIN_BRIDGE_BAD_REWARD_AMOUNT:             {"temXCHAIN_BRIDGE_BAD_REWARD_AMOUNT", "Malformed: Bad reward amount."},

	terRETRY:       {"terRETRY", "Retry transaction."},
	terFUNDS_SPENT: {"terFUNDS_SPENT", "DEPRECATED."},
	terINSUF_FEE_B: {"terINSUF_FEE_B", "Account balance can't pay fee."},
	terLAST:        {"terLAST", "DEPRECATED."},
	terNO_RIPPLE:   {"terNO_RIPPLE", "Path does not permit rippling."},
	terNO_ACCOUNT:  {"terNO_ACCOUNT", "The source account does not exist."},
	terNO_AUTH:     {"terNO_AUTH", "Not authorized to hold IOUs."},
	terNO_LINE:     {"terNO_LINE", "No such line."},
	terPRE_SEQ:     {"terPRE_SEQ", "Missing/inapplicable prior transaction."},
	terOWNERS:      {"terOWNERS", "Non-zero owner count."},
	terQUEUED:      {"terQUEUED", "Held until escalated fee drops."},
	terPRE_TICKET:  {"terPRE_TICKET", "Ticket is not yet in ledger."},
	terNO_AMM:      {"terNO_AMM", "AMM doesn't exist for the asset pair."},
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
	return r == tesSUCCESS
}

func (r TransactionResult) Queued() bool {
	return r == terQUEUED
}

func (r TransactionResult) Symbol() string {
	switch r {
	case tesSUCCESS, tecCLAIM:
		return "✓"
	case tecPATH_PARTIAL, tecPATH_DRY:
		return "½"
	case tecUNFUNDED, tecUNFUNDED_ADD, tecUNFUNDED_OFFER, tecUNFUNDED_PAYMENT:
		return "$"
	default:
		return "✗"
	}
}
