package cidlink

import (
	"bytes"
	"context"
	"fmt"
	"io"

	cid "github.com/ipfs/go-cid"
	ipld "github.com/ipld/go-ipld-prime"
	multihash "github.com/multiformats/go-multihash"
)

var (
	_ ipld.Link        = Link{}
	_ ipld.LinkBuilder = LinkBuilder{}
)

type Link struct {
	cid.Cid
}

func (lnk Link) Load(ctx context.Context, lnkCtx ipld.LinkContext, nb ipld.NodeBuilder, loader ipld.Loader) (ipld.Node, error) {
	// Open the byte reader.
	r, err := loader(lnk, lnkCtx)
	if err != nil {
		return nil, err
	}
	// Tee into hash checking and unmarshalling.
	mcDecoder, exists := multicodecDecodeTable[lnk.Prefix().Codec]
	if !exists {
		return nil, fmt.Errorf("no decoder registered for multicodec %d", lnk.Prefix().Codec)
	}
	var hasher bytes.Buffer // multihash only exports bulk use, which is... really inefficient and should be fixed.
	node, decodeErr := mcDecoder(nb, io.TeeReader(r, &hasher))
	// Error checking order here is tricky.
	//  If decoding errored out, we should still run the reader to the end, to check the hash.
	//  (We still don't implement this by running the hash to the end first, because that would increase the high-water memory requirement.)
	//   ((Which we experience right now anyway because multihash's interface is silly, but we're acting as if that's fixed or will be soon.))
	//  If the hash is rejected, we should return that error (and even if there was a decodeErr, it becomes irrelevant).
	if decodeErr != nil {
		_, err := io.Copy(&hasher, r)
		if err != nil {
			return nil, err
		}
	}
	hash, err := multihash.Sum(hasher.Bytes(), lnk.Prefix().MhType, lnk.Prefix().MhLength)
	if err != nil {
		return nil, err
	}
	if hash.B58String() != lnk.Hash().B58String() {
		return nil, fmt.Errorf("hash mismatch!")
	}
	if decodeErr != nil {
		return nil, decodeErr
	}
	return node, nil
}
func (lnk Link) LinkBuilder() ipld.LinkBuilder {
	return LinkBuilder{lnk.Cid.Prefix()}
}
func (lnk Link) String() string {
	return lnk.Cid.String()
}

type LinkBuilder struct {
	cid.Prefix
}

func (lb LinkBuilder) Build(ctx context.Context, lnkCtx ipld.LinkContext, node ipld.Node, storer ipld.Storer) (ipld.Link, error) {
	// Open the byte writer.
	w, commit, err := storer(lnkCtx)
	if err != nil {
		return nil, err
	}
	// Marshal, teeing into the storage writer and the hasher.
	mcEncoder, exists := multicodecEncodeTable[lb.Prefix.MhType]
	if !exists {
		return nil, fmt.Errorf("no encoder registered for multicodec %d", lb.Prefix.MhType)
	}
	var hasher bytes.Buffer // multihash only exports bulk use, which is... really inefficient and should be fixed.
	w = io.MultiWriter(&hasher, w)
	err = mcEncoder(node, w)
	if err != nil {
		return nil, err
	}
	hash, err := multihash.Sum(hasher.Bytes(), lb.Prefix.MhType, lb.Prefix.MhLength)
	// FIXME finish making a CID out of this.
	// the cid package is a maze of twisty little passages all alike and I don't honestly know what's up where why.
	_ = hash
	if err := commit(nil); err != nil {
		return nil, err
	}
	panic("TODO")
}
