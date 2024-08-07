package repository

import (
	"artion-api-graphql/internal/types"
	"encoding/json"
	"fmt"
	neturl "net/url"
	"strings"
)

// GetTokenJsonMetadata provides decoded JSON metadata for the given token metadata URI.
func (p *Proxy) GetTokenJsonMetadata(uri string) (*types.JsonMetadata, error) {
	data, err := p.GetRawTokenJsonMetadata(uri)
	if err != nil {
		return nil, err
	}
	jsonMeta, err := types.DecodeJsonMetadata(data)
	if err != nil {
		return nil, fmt.Errorf("unable to decode json; %s", err)
	}
	return jsonMeta, err
}

// GetRawTokenJsonMetadata provides decoded JSON metadata for the given token metadata URI.
func (p *Proxy) GetRawTokenJsonMetadata(uri string) ([]byte, error) {
	// make sure to do this only once, if parallel requests were fired
	var key strings.Builder
	key.WriteString("GetRawTokenJsonMetadata")
	key.WriteString(uri)

	data, err, _ := p.callGroup.Do(key.String(), func() (interface{}, error) {
		data, _, err := p.getFileFromUri(uri)
		if err != nil {
			return nil, fmt.Errorf("unable to download json; %s", err)
		}
		return data, err
	})
	if err != nil {
		return nil, err
	}
	return data.([]byte), err
}

// GetImage downloads an image expected on the given URI.
func (p *Proxy) GetImage(imgUri string) (*types.Image, error) {
	var key strings.Builder
	key.WriteString("GetImage")
	key.WriteString(imgUri)

	data, err, _ := p.callGroup.Do(key.String(), func() (interface{}, error) {
		return p.getImageFromUri(imgUri)
	})
	if err != nil {
		log.Warningf("image can not be loaded from %s; %s", imgUri, err.Error())
		return nil, err
	}
	if data == nil {
		log.Warningf("image not found at %s", imgUri)
		return nil, fmt.Errorf("image not found at given URI")
	}
	return data.(*types.Image), err
}

// GetImageThumbnail generates a thumbnail for an image expected to be downloadable from the given URI.
func (p *Proxy) GetImageThumbnail(imgUri string) (*types.Image, error) {
	var key strings.Builder
	key.WriteString("GetImageThumbnail")
	key.WriteString(imgUri)

	data, err, _ := p.callGroup.Do(key.String(), func() (interface{}, error) {
		image, err := p.GetImage(imgUri)
		if err != nil {
			return nil, fmt.Errorf("image loading failed for %s; %s", imgUri, err)
		}
		if nil == image {
			return nil, fmt.Errorf("image %s not found", imgUri)
		}

		log.Infof("loaded %s of type %s", imgUri, image.Type.Mimetype())
		thumb, err := createThumbnail(*image)
		if err != nil {
			return nil, fmt.Errorf("thumbnail creation failed; %s", err)
		}
		return &thumb, nil
	})
	if data == nil {
		return nil, err
	}
	return data.(*types.Image), err
}

// UploadTokenData stores JSON metadata along with the token image into the IPFS storage.
func (p *Proxy) UploadTokenData(metadata types.JsonMetadata, image types.Image) (uri string, err error) {
	cid, err := p.pinFile("token-image", image.Data)
	if err != nil {
		return "", fmt.Errorf("uploading token image failed; %s", err)
	}

	// MM should use ipfs/gateway config (or relative api links when emulate_on_shared_db is enabled)
	//imageUri := "https://artion.mypinata.cloud/ipfs/" + cid
	baseUri := "https://artion.mypinata.cloud/ipfs/"
	if p.pinner.EmulateOnSharedDb() {
		baseUri = "/volcano-api/metadata/" // ?!..
	}
	imageUri := baseUri + cid

	metadata.Image = &imageUri

	data, err := json.Marshal(metadata)
	if err != nil {
		return "", fmt.Errorf("marshaling json meta failed; %s", err)
	}

	cid, err = p.pinFile("token-meta", data)
	if err != nil {
		return "", fmt.Errorf("uploading token meta failed; %s", err)
	}
	// MM
	// return "https://artion.mypinata.cloud/ipfs/" + cid, nil
	return baseUri + cid, nil
}

// pinFile requests pinning of the given file generating IPFS CID of the stored file.
func (p *Proxy) pinFile(filename string, content []byte) (cid string, err error) {
	// MM
	if p.pinner.EmulateOnSharedDb() {
		image := types.NewNamedImage(content, filename)
		err = p.shared.StoreImage(image)
		cid = image.ID().Hex()
	} else {
		cid, err = p.pinner.PinFile(filename, content)
	}
	if err == nil {
		err = p.filecache.PushIpfsFile(cid, content)
	}
	return
}

// getImageFromUri downloads image from given URI and detect its mimetype
func (p *Proxy) getImageFromUri(uri string) (*types.Image, error) {
	data, _, err := p.getFileFromUri(uri)
	if err != nil {
		return nil, fmt.Errorf("unable to download image; %s", err)
	}

	imgType, err := types.ImageTypeFromMimetype(data)
	if imgType == types.ImageTypeUnknown {
		imgType = types.ImageTypeFromExtension(uri)
	}

	out := types.Image{
		Data: data,
		Type: imgType,
	}
	return &out, nil
}

// getCidFromIpfsUri extracts IPFS CID from the given URI.
func getCidFromIpfsUri(uri string) string {
	cid := uri[6:] // skip /ipfs/
	slashIdx := strings.Index(cid, "/")
	if slashIdx != -1 {
		cid = cid[0:slashIdx]
	}
	return cid
}

// getFileFromUri resolves the URI and download file from the URI using appropriate protocol
func (p *Proxy) getFileFromUri(uri string) (data []byte, mimetype string, err error) {
	// the URI contains the data directly as BASE64 encoded data stream
	if strings.HasPrefix(uri, "data:") {
		return p.uri.GetFromDataUri(uri)
	}

	// serve IPFS (or IPFS gateway) URIs
	if ipfsUri := p.uri.GetIpfsUri(uri); ipfsUri != "" {
		// try the local cache first
		cachedContent := p.filecache.PullIpfsFile(getCidFromIpfsUri(ipfsUri))
		if cachedContent != nil {
			return cachedContent, "", nil
		}

		// MM
		if p.pinner.EmulateOnSharedDb() {
			image, err := p.shared.GetImage(getCidFromIpfsUri(ipfsUri))
			if err == nil {
				if image == nil {
					return nil, "", fmt.Errorf("not found image file %s on shared_db", uri)
				} else {
					return image.Data, image.Type.Mimetype(), nil
				}
			} else {
				return nil, "", fmt.Errorf("failed to retrieve image file %s from shared_db; %s", uri, err.Error())
			}
		}

		// extract data from IPFS, if possible
		data, mimetype, err = p.uri.GetFromIpfs(ipfsUri)
		if err == nil || strings.HasPrefix(uri, "ipfs:") || strings.HasPrefix(uri, "/ipfs/") {
			return data, mimetype, err
		} else {
			// if err, but it is http ipfs proxy uri, only log and continue to process as a http uri
			log.Warningf("failed to download file %s from IPFS, will try http; %s", ipfsUri, err)
		}
	}

	// try to decode the URI to validate the format
	url, err := neturl.Parse(uri)
	if err != nil {
		return nil, "", fmt.Errorf("invalid URI format at %s; %s", uri, err.Error())
	}

	// check the schema for known access methods
	if url.Scheme == "http" || url.Scheme == "https" {
		return p.uri.GetFromHttp(uri)
	}

	return nil, "", fmt.Errorf("invalid URI scheme; file not available at %s", uri)
}
