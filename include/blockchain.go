package include

const (
	GOSSIP_PROTOCOL_V1         	= "blockchain_gossip/1.0.0"
	SYNC_PROTOCOL_V1   			= "blockchain_sync/1.1.0"
	FASTFORWARD_PROTOCOL_V1		= "blockchain_fastforward/1.0.0"

	SYNC_PROTOCOL_V2       		= "blockchain_sync/1.2.0"
	FASTFORWARD_PROTOCOL_V2    	= "blockchain_fastforward/1.2.0"

	SUPPORTED_GOSSIP_PROTOCOLS	= GOSSIP_PROTOCOL_V1

	SNAPSHOT_PROTOCOL 			= "blockchain_snapshot/1.0.0"

	TX_PROTOCOL 				= "blockchain_txn/1.0.0"
	LOC_ASSERTION_PROTOCOL		= "loc_assertion/1.0.0"
	STATE_CHANNEL_PROTOCOL_V1	= "state_channel/1.0.0"

	// B58 Address Versions
	B58_MAINNET_VER	= 0
	B58_TESTNET_VER	= 2
	B58_HTLC_VER	= 24

	EVT_MGR 	= "blockchain_event_mgr"
)

// this is meant to be constant! Please don't mutate it!
var SUPPORTED_SYNC_PROTOCOLS		= []string{SYNC_PROTOCOL_V2, SYNC_PROTOCOL_V1}
var SUPPORTED_FASTFORWARD_PROTOCOLS = []string{FASTFORWARD_PROTOCOL_V2, FASTFORWARD_PROTOCOL_V1}

var BC_UPGRADE_NAMES = []string{
	"gateway_v2",
	"hex_targets",
	"gateway_oui",
	"h3dex",
	"h3dex2",
	"gateway_lg3",
	"clear_witnesses"}

func bones(HNT int) int {
	return HNT * 100000000
}