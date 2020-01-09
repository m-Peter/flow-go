package hotstuff

type ViewState interface {
	IsSelf(view uint64, identityIdx uint32) bool
	IsSelfLeaderForView(view uint64) bool
	GetSelfIdxForView(view uint64) uint32
	GetIdxOfPubKeyForView(view uint64) uint32
}
