package include

/*
// A gateway can be set to one of 3 modes, dataonly, light and full
// each mode then has a bitmask defined via a chain var which
// determines the range of capabilities applicable for the given mode
*/

// The current full range of possible capabilities supported across the various gateway modes
const (
	// determines if a GW can route packets
	GW_CAPABILITY_ROUTE_PACKETS 	=    0x01
	// determines if a GW can issue POC Challenges
	GW_CAPABILITY_POC_CHALLENGER 	=  0x02
	// determines if a GW can accept POC Challenges
 	GW_CAPABILITY_POC_CHALLENGEE 	=   0x04
	// determines if a GW can witness challenges
 	GW_CAPABILITY_POC_WITNESS 		=      0x08
	// determines if a GW can issue receipts
 	GW_CAPABILITY_POC_RECEIPT 		=      0x10
	// determines if a GW can participate in consensus group
 	GW_CAPABILITY_CONSENSUS_GROUP 	=  0x20
)
//
// V1 capabilities for each gateway type defined below
// In practise we should never need to use these but they cover a case
// in blockchain_ledger_gateway_v2:mask_for_gateway_mode/2 whereby we have a GW with its mode field set
// but the corresponding bitmask chain var is not found....which should never really happen but
// I feel better for covering the scenario
//
const (
	 GW_CAPABILITIES_DATAONLY_GATEWAY_V1 = 0 ^ GW_CAPABILITY_ROUTE_PACKETS

	 GW_CAPABILITIES_LIGHT_GATEWAY_V1_PART1 = 0 ^ GW_CAPABILITY_ROUTE_PACKETS ^ GW_CAPABILITY_POC_CHALLENGER
	 GW_CAPABILITIES_LIGHT_GATEWAY_V1_PART2 = GW_CAPABILITY_POC_CHALLENGEE ^ GW_CAPABILITY_POC_WITNESS ^ GW_CAPABILITY_POC_RECEIPT
	 GW_CAPABILITIES_LIGHT_GATEWAY_V1 = GW_CAPABILITIES_LIGHT_GATEWAY_V1_PART1 ^ GW_CAPABILITIES_LIGHT_GATEWAY_V1_PART2
)

const (
	GW_CAPABILITIES_FULL_GATEWAY_V1_PART1 = 0 ^ GW_CAPABILITY_ROUTE_PACKETS ^ GW_CAPABILITY_POC_CHALLENGER ^ GW_CAPABILITY_POC_CHALLENGEE
	GW_CAPABILITIES_FULL_GATEWAY_V1_PART2 = GW_CAPABILITY_POC_WITNESS ^ GW_CAPABILITY_POC_RECEIPT ^ GW_CAPABILITY_CONSENSUS_GROUP
	GW_CAPABILITIES_FULL_GATEWAY_V1 = GW_CAPABILITIES_FULL_GATEWAY_V1_PART1 ^ GW_CAPABILITIES_FULL_GATEWAY_V1_PART2
)

// helper macros
func GW_CAPABILITIES_SET(Capabilities []byte) byte {
	var ret byte = 0x0
	for _, Capability := range Capabilities {
		ret = ret ^ Capability
	}
	return ret
}

func GW_CAPABILITY_QUERY(Mask byte, Capability byte) bool {
	return Mask & Capability == Capability
}

