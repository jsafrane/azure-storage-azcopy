//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package generated_blob

import (
	"context"
	"encoding/base64"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"io"
	"net/http"
	"strconv"
	"time"
)

// AppendBlobClient contains the methods for the AppendBlob group.
// Don't use this type directly, use a constructor function instead.
type AppendBlobClient struct {
	internal *azcore.Client
	endpoint string
}

// AppendBlock - The Append Block operation commits a new block of data to the end of an existing append blob. The Append
// Block operation is permitted only if the blob was created with x-ms-blob-type set to
// AppendBlob. Append Block is supported only on version 2015-02-21 version or later.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-12-02
//   - contentLength - The length of the request.
//   - body - Initial data
//   - options - AppendBlobClientAppendBlockOptions contains the optional parameters for the AppendBlobClient.AppendBlock method.
//   - LeaseAccessConditions - LeaseAccessConditions contains a group of parameters for the ContainerClient.GetProperties method.
//   - AppendPositionAccessConditions - AppendPositionAccessConditions contains a group of parameters for the AppendBlobClient.AppendBlock
//     method.
//   - CPKInfo - CPKInfo contains a group of parameters for the BlobClient.Download method.
//   - CPKScopeInfo - CPKScopeInfo contains a group of parameters for the BlobClient.SetMetadata method.
//   - ModifiedAccessConditions - ModifiedAccessConditions contains a group of parameters for the ContainerClient.Delete method.
func (client *AppendBlobClient) AppendBlock(ctx context.Context, contentLength int64, body io.ReadSeekCloser, options *AppendBlobClientAppendBlockOptions, leaseAccessConditions *LeaseAccessConditions, appendPositionAccessConditions *AppendPositionAccessConditions, cpkInfo *CPKInfo, cpkScopeInfo *CPKScopeInfo, modifiedAccessConditions *ModifiedAccessConditions) (AppendBlobClientAppendBlockResponse, error) {
	req, err := client.appendBlockCreateRequest(ctx, contentLength, body, options, leaseAccessConditions, appendPositionAccessConditions, cpkInfo, cpkScopeInfo, modifiedAccessConditions)
	if err != nil {
		return AppendBlobClientAppendBlockResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AppendBlobClientAppendBlockResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusCreated) {
		return AppendBlobClientAppendBlockResponse{}, runtime.NewResponseError(resp)
	}
	return client.appendBlockHandleResponse(resp)
}

// appendBlockCreateRequest creates the AppendBlock request.
func (client *AppendBlobClient) appendBlockCreateRequest(ctx context.Context, contentLength int64, body io.ReadSeekCloser, options *AppendBlobClientAppendBlockOptions, leaseAccessConditions *LeaseAccessConditions, appendPositionAccessConditions *AppendPositionAccessConditions, cpkInfo *CPKInfo, cpkScopeInfo *CPKScopeInfo, modifiedAccessConditions *ModifiedAccessConditions) (*policy.Request, error) {
	req, err := runtime.NewRequest(ctx, http.MethodPut, client.endpoint)
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("comp", "appendblock")
	if options != nil && options.Timeout != nil {
		reqQP.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Content-Length"] = []string{strconv.FormatInt(contentLength, 10)}
	if options != nil && options.TransactionalContentMD5 != nil {
		req.Raw().Header["Content-MD5"] = []string{base64.StdEncoding.EncodeToString(options.TransactionalContentMD5)}
	}
	if options != nil && options.TransactionalContentCRC64 != nil {
		req.Raw().Header["x-ms-content-crc64"] = []string{base64.StdEncoding.EncodeToString(options.TransactionalContentCRC64)}
	}
	if leaseAccessConditions != nil && leaseAccessConditions.LeaseID != nil {
		req.Raw().Header["x-ms-lease-id"] = []string{*leaseAccessConditions.LeaseID}
	}
	if appendPositionAccessConditions != nil && appendPositionAccessConditions.MaxSize != nil {
		req.Raw().Header["x-ms-blob-condition-maxsize"] = []string{strconv.FormatInt(*appendPositionAccessConditions.MaxSize, 10)}
	}
	if appendPositionAccessConditions != nil && appendPositionAccessConditions.AppendPosition != nil {
		req.Raw().Header["x-ms-blob-condition-appendpos"] = []string{strconv.FormatInt(*appendPositionAccessConditions.AppendPosition, 10)}
	}
	if cpkInfo != nil && cpkInfo.EncryptionKey != nil {
		req.Raw().Header["x-ms-encryption-key"] = []string{*cpkInfo.EncryptionKey}
	}
	if cpkInfo != nil && cpkInfo.EncryptionKeySHA256 != nil {
		req.Raw().Header["x-ms-encryption-key-sha256"] = []string{*cpkInfo.EncryptionKeySHA256}
	}
	if cpkInfo != nil && cpkInfo.EncryptionAlgorithm != nil {
		req.Raw().Header["x-ms-encryption-algorithm"] = []string{string(*cpkInfo.EncryptionAlgorithm)}
	}
	if cpkScopeInfo != nil && cpkScopeInfo.EncryptionScope != nil {
		req.Raw().Header["x-ms-encryption-scope"] = []string{*cpkScopeInfo.EncryptionScope}
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfModifiedSince != nil {
		req.Raw().Header["If-Modified-Since"] = []string{(*modifiedAccessConditions.IfModifiedSince).In(gmt).Format(time.RFC1123)}
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfUnmodifiedSince != nil {
		req.Raw().Header["If-Unmodified-Since"] = []string{(*modifiedAccessConditions.IfUnmodifiedSince).In(gmt).Format(time.RFC1123)}
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{string(*modifiedAccessConditions.IfMatch)}
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfNoneMatch != nil {
		req.Raw().Header["If-None-Match"] = []string{string(*modifiedAccessConditions.IfNoneMatch)}
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfTags != nil {
		req.Raw().Header["x-ms-if-tags"] = []string{*modifiedAccessConditions.IfTags}
	}
	req.Raw().Header["x-ms-version"] = []string{"2021-12-02"}
	if options != nil && options.RequestID != nil {
		req.Raw().Header["x-ms-client-request-id"] = []string{*options.RequestID}
	}
	req.Raw().Header["Accept"] = []string{"application/xml"}
	if err := req.SetBody(body, "application/octet-stream"); err != nil {
		return nil, err
	}
	return req, nil
}

// appendBlockHandleResponse handles the AppendBlock response.
func (client *AppendBlobClient) appendBlockHandleResponse(resp *http.Response) (AppendBlobClientAppendBlockResponse, error) {
	result := AppendBlobClientAppendBlockResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = (*azcore.ETag)(&val)
	}
	if val := resp.Header.Get("Last-Modified"); val != "" {
		lastModified, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return AppendBlobClientAppendBlockResponse{}, err
		}
		result.LastModified = &lastModified
	}
	if val := resp.Header.Get("Content-MD5"); val != "" {
		contentMD5, err := base64.StdEncoding.DecodeString(val)
		if err != nil {
			return AppendBlobClientAppendBlockResponse{}, err
		}
		result.ContentMD5 = contentMD5
	}
	if val := resp.Header.Get("x-ms-content-crc64"); val != "" {
		contentCRC64, err := base64.StdEncoding.DecodeString(val)
		if err != nil {
			return AppendBlobClientAppendBlockResponse{}, err
		}
		result.ContentCRC64 = contentCRC64
	}
	if val := resp.Header.Get("x-ms-client-request-id"); val != "" {
		result.ClientRequestID = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if val := resp.Header.Get("Date"); val != "" {
		date, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return AppendBlobClientAppendBlockResponse{}, err
		}
		result.Date = &date
	}
	if val := resp.Header.Get("x-ms-blob-append-offset"); val != "" {
		result.BlobAppendOffset = &val
	}
	if val := resp.Header.Get("x-ms-blob-committed-block-count"); val != "" {
		blobCommittedBlockCount32, err := strconv.ParseInt(val, 10, 32)
		blobCommittedBlockCount := int32(blobCommittedBlockCount32)
		if err != nil {
			return AppendBlobClientAppendBlockResponse{}, err
		}
		result.BlobCommittedBlockCount = &blobCommittedBlockCount
	}
	if val := resp.Header.Get("x-ms-request-server-encrypted"); val != "" {
		isServerEncrypted, err := strconv.ParseBool(val)
		if err != nil {
			return AppendBlobClientAppendBlockResponse{}, err
		}
		result.IsServerEncrypted = &isServerEncrypted
	}
	if val := resp.Header.Get("x-ms-encryption-key-sha256"); val != "" {
		result.EncryptionKeySHA256 = &val
	}
	if val := resp.Header.Get("x-ms-encryption-scope"); val != "" {
		result.EncryptionScope = &val
	}
	return result, nil
}

// AppendBlockFromURL - The Append Block operation commits a new block of data to the end of an existing append blob where
// the contents are read from a source url. The Append Block operation is permitted only if the blob was
// created with x-ms-blob-type set to AppendBlob. Append Block is supported only on version 2015-02-21 version or later.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-12-02
//   - sourceURL - Specify a URL to the copy source.
//   - contentLength - The length of the request.
//   - options - AppendBlobClientAppendBlockFromURLOptions contains the optional parameters for the AppendBlobClient.AppendBlockFromURL
//     method.
//   - CPKInfo - CPKInfo contains a group of parameters for the BlobClient.Download method.
//   - CPKScopeInfo - CPKScopeInfo contains a group of parameters for the BlobClient.SetMetadata method.
//   - LeaseAccessConditions - LeaseAccessConditions contains a group of parameters for the ContainerClient.GetProperties method.
//   - AppendPositionAccessConditions - AppendPositionAccessConditions contains a group of parameters for the AppendBlobClient.AppendBlock
//     method.
//   - ModifiedAccessConditions - ModifiedAccessConditions contains a group of parameters for the ContainerClient.Delete method.
//   - SourceModifiedAccessConditions - SourceModifiedAccessConditions contains a group of parameters for the BlobClient.StartCopyFromURL
//     method.
func (client *AppendBlobClient) AppendBlockFromURL(ctx context.Context, sourceURL string, contentLength int64, options *AppendBlobClientAppendBlockFromURLOptions, cpkInfo *CPKInfo, cpkScopeInfo *CPKScopeInfo, leaseAccessConditions *LeaseAccessConditions, appendPositionAccessConditions *AppendPositionAccessConditions, modifiedAccessConditions *ModifiedAccessConditions, sourceModifiedAccessConditions *SourceModifiedAccessConditions) (AppendBlobClientAppendBlockFromURLResponse, error) {
	req, err := client.appendBlockFromURLCreateRequest(ctx, sourceURL, contentLength, options, cpkInfo, cpkScopeInfo, leaseAccessConditions, appendPositionAccessConditions, modifiedAccessConditions, sourceModifiedAccessConditions)
	if err != nil {
		return AppendBlobClientAppendBlockFromURLResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AppendBlobClientAppendBlockFromURLResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusCreated) {
		return AppendBlobClientAppendBlockFromURLResponse{}, runtime.NewResponseError(resp)
	}
	return client.appendBlockFromURLHandleResponse(resp)
}

// appendBlockFromURLCreateRequest creates the AppendBlockFromURL request.
func (client *AppendBlobClient) appendBlockFromURLCreateRequest(ctx context.Context, sourceURL string, contentLength int64, options *AppendBlobClientAppendBlockFromURLOptions, cpkInfo *CPKInfo, cpkScopeInfo *CPKScopeInfo, leaseAccessConditions *LeaseAccessConditions, appendPositionAccessConditions *AppendPositionAccessConditions, modifiedAccessConditions *ModifiedAccessConditions, sourceModifiedAccessConditions *SourceModifiedAccessConditions) (*policy.Request, error) {
	req, err := runtime.NewRequest(ctx, http.MethodPut, client.endpoint)
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("comp", "appendblock")
	if options != nil && options.Timeout != nil {
		reqQP.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["x-ms-copy-source"] = []string{sourceURL}
	if options != nil && options.SourceRange != nil {
		req.Raw().Header["x-ms-source-range"] = []string{*options.SourceRange}
	}
	if options != nil && options.SourceContentMD5 != nil {
		req.Raw().Header["x-ms-source-content-md5"] = []string{base64.StdEncoding.EncodeToString(options.SourceContentMD5)}
	}
	if options != nil && options.SourceContentcrc64 != nil {
		req.Raw().Header["x-ms-source-content-crc64"] = []string{base64.StdEncoding.EncodeToString(options.SourceContentcrc64)}
	}
	req.Raw().Header["Content-Length"] = []string{strconv.FormatInt(contentLength, 10)}
	if options != nil && options.TransactionalContentMD5 != nil {
		req.Raw().Header["Content-MD5"] = []string{base64.StdEncoding.EncodeToString(options.TransactionalContentMD5)}
	}
	if cpkInfo != nil && cpkInfo.EncryptionKey != nil {
		req.Raw().Header["x-ms-encryption-key"] = []string{*cpkInfo.EncryptionKey}
	}
	if cpkInfo != nil && cpkInfo.EncryptionKeySHA256 != nil {
		req.Raw().Header["x-ms-encryption-key-sha256"] = []string{*cpkInfo.EncryptionKeySHA256}
	}
	if cpkInfo != nil && cpkInfo.EncryptionAlgorithm != nil {
		req.Raw().Header["x-ms-encryption-algorithm"] = []string{string(*cpkInfo.EncryptionAlgorithm)}
	}
	if cpkScopeInfo != nil && cpkScopeInfo.EncryptionScope != nil {
		req.Raw().Header["x-ms-encryption-scope"] = []string{*cpkScopeInfo.EncryptionScope}
	}
	if leaseAccessConditions != nil && leaseAccessConditions.LeaseID != nil {
		req.Raw().Header["x-ms-lease-id"] = []string{*leaseAccessConditions.LeaseID}
	}
	if appendPositionAccessConditions != nil && appendPositionAccessConditions.MaxSize != nil {
		req.Raw().Header["x-ms-blob-condition-maxsize"] = []string{strconv.FormatInt(*appendPositionAccessConditions.MaxSize, 10)}
	}
	if appendPositionAccessConditions != nil && appendPositionAccessConditions.AppendPosition != nil {
		req.Raw().Header["x-ms-blob-condition-appendpos"] = []string{strconv.FormatInt(*appendPositionAccessConditions.AppendPosition, 10)}
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfModifiedSince != nil {
		req.Raw().Header["If-Modified-Since"] = []string{(*modifiedAccessConditions.IfModifiedSince).In(gmt).Format(time.RFC1123)}
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfUnmodifiedSince != nil {
		req.Raw().Header["If-Unmodified-Since"] = []string{(*modifiedAccessConditions.IfUnmodifiedSince).In(gmt).Format(time.RFC1123)}
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{string(*modifiedAccessConditions.IfMatch)}
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfNoneMatch != nil {
		req.Raw().Header["If-None-Match"] = []string{string(*modifiedAccessConditions.IfNoneMatch)}
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfTags != nil {
		req.Raw().Header["x-ms-if-tags"] = []string{*modifiedAccessConditions.IfTags}
	}
	if sourceModifiedAccessConditions != nil && sourceModifiedAccessConditions.SourceIfModifiedSince != nil {
		req.Raw().Header["x-ms-source-if-modified-since"] = []string{(*sourceModifiedAccessConditions.SourceIfModifiedSince).In(gmt).Format(time.RFC1123)}
	}
	if sourceModifiedAccessConditions != nil && sourceModifiedAccessConditions.SourceIfUnmodifiedSince != nil {
		req.Raw().Header["x-ms-source-if-unmodified-since"] = []string{(*sourceModifiedAccessConditions.SourceIfUnmodifiedSince).In(gmt).Format(time.RFC1123)}
	}
	if sourceModifiedAccessConditions != nil && sourceModifiedAccessConditions.SourceIfMatch != nil {
		req.Raw().Header["x-ms-source-if-match"] = []string{string(*sourceModifiedAccessConditions.SourceIfMatch)}
	}
	if sourceModifiedAccessConditions != nil && sourceModifiedAccessConditions.SourceIfNoneMatch != nil {
		req.Raw().Header["x-ms-source-if-none-match"] = []string{string(*sourceModifiedAccessConditions.SourceIfNoneMatch)}
	}
	req.Raw().Header["x-ms-version"] = []string{"2021-12-02"}
	if options != nil && options.RequestID != nil {
		req.Raw().Header["x-ms-client-request-id"] = []string{*options.RequestID}
	}
	if options != nil && options.CopySourceAuthorization != nil {
		req.Raw().Header["x-ms-copy-source-authorization"] = []string{*options.CopySourceAuthorization}
	}
	req.Raw().Header["Accept"] = []string{"application/xml"}
	return req, nil
}

// appendBlockFromURLHandleResponse handles the AppendBlockFromURL response.
func (client *AppendBlobClient) appendBlockFromURLHandleResponse(resp *http.Response) (AppendBlobClientAppendBlockFromURLResponse, error) {
	result := AppendBlobClientAppendBlockFromURLResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = (*azcore.ETag)(&val)
	}
	if val := resp.Header.Get("Last-Modified"); val != "" {
		lastModified, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return AppendBlobClientAppendBlockFromURLResponse{}, err
		}
		result.LastModified = &lastModified
	}
	if val := resp.Header.Get("Content-MD5"); val != "" {
		contentMD5, err := base64.StdEncoding.DecodeString(val)
		if err != nil {
			return AppendBlobClientAppendBlockFromURLResponse{}, err
		}
		result.ContentMD5 = contentMD5
	}
	if val := resp.Header.Get("x-ms-content-crc64"); val != "" {
		contentCRC64, err := base64.StdEncoding.DecodeString(val)
		if err != nil {
			return AppendBlobClientAppendBlockFromURLResponse{}, err
		}
		result.ContentCRC64 = contentCRC64
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if val := resp.Header.Get("Date"); val != "" {
		date, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return AppendBlobClientAppendBlockFromURLResponse{}, err
		}
		result.Date = &date
	}
	if val := resp.Header.Get("x-ms-blob-append-offset"); val != "" {
		result.BlobAppendOffset = &val
	}
	if val := resp.Header.Get("x-ms-blob-committed-block-count"); val != "" {
		blobCommittedBlockCount32, err := strconv.ParseInt(val, 10, 32)
		blobCommittedBlockCount := int32(blobCommittedBlockCount32)
		if err != nil {
			return AppendBlobClientAppendBlockFromURLResponse{}, err
		}
		result.BlobCommittedBlockCount = &blobCommittedBlockCount
	}
	if val := resp.Header.Get("x-ms-encryption-key-sha256"); val != "" {
		result.EncryptionKeySHA256 = &val
	}
	if val := resp.Header.Get("x-ms-encryption-scope"); val != "" {
		result.EncryptionScope = &val
	}
	if val := resp.Header.Get("x-ms-request-server-encrypted"); val != "" {
		isServerEncrypted, err := strconv.ParseBool(val)
		if err != nil {
			return AppendBlobClientAppendBlockFromURLResponse{}, err
		}
		result.IsServerEncrypted = &isServerEncrypted
	}
	return result, nil
}

// Create - The Create Append Blob operation creates a new append blob.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-12-02
//   - contentLength - The length of the request.
//   - options - AppendBlobClientCreateOptions contains the optional parameters for the AppendBlobClient.Create method.
//   - BlobHTTPHeaders - BlobHTTPHeaders contains a group of parameters for the BlobClient.SetHTTPHeaders method.
//   - LeaseAccessConditions - LeaseAccessConditions contains a group of parameters for the ContainerClient.GetProperties method.
//   - CPKInfo - CPKInfo contains a group of parameters for the BlobClient.Download method.
//   - CPKScopeInfo - CPKScopeInfo contains a group of parameters for the BlobClient.SetMetadata method.
//   - ModifiedAccessConditions - ModifiedAccessConditions contains a group of parameters for the ContainerClient.Delete method.
func (client *AppendBlobClient) Create(ctx context.Context, contentLength int64, options *AppendBlobClientCreateOptions, blobHTTPHeaders *BlobHTTPHeaders, leaseAccessConditions *LeaseAccessConditions, cpkInfo *CPKInfo, cpkScopeInfo *CPKScopeInfo, modifiedAccessConditions *ModifiedAccessConditions) (AppendBlobClientCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, contentLength, options, blobHTTPHeaders, leaseAccessConditions, cpkInfo, cpkScopeInfo, modifiedAccessConditions)
	if err != nil {
		return AppendBlobClientCreateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AppendBlobClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusCreated) {
		return AppendBlobClientCreateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *AppendBlobClient) createCreateRequest(ctx context.Context, contentLength int64, options *AppendBlobClientCreateOptions, blobHTTPHeaders *BlobHTTPHeaders, leaseAccessConditions *LeaseAccessConditions, cpkInfo *CPKInfo, cpkScopeInfo *CPKScopeInfo, modifiedAccessConditions *ModifiedAccessConditions) (*policy.Request, error) {
	req, err := runtime.NewRequest(ctx, http.MethodPut, client.endpoint)
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Timeout != nil {
		reqQP.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["x-ms-blob-type"] = []string{"AppendBlob"}
	req.Raw().Header["Content-Length"] = []string{strconv.FormatInt(contentLength, 10)}
	if blobHTTPHeaders != nil && blobHTTPHeaders.BlobContentType != nil {
		req.Raw().Header["x-ms-blob-content-type"] = []string{*blobHTTPHeaders.BlobContentType}
	}
	if blobHTTPHeaders != nil && blobHTTPHeaders.BlobContentEncoding != nil {
		req.Raw().Header["x-ms-blob-content-encoding"] = []string{*blobHTTPHeaders.BlobContentEncoding}
	}
	if blobHTTPHeaders != nil && blobHTTPHeaders.BlobContentLanguage != nil {
		req.Raw().Header["x-ms-blob-content-language"] = []string{*blobHTTPHeaders.BlobContentLanguage}
	}
	if blobHTTPHeaders != nil && blobHTTPHeaders.BlobContentMD5 != nil {
		req.Raw().Header["x-ms-blob-content-md5"] = []string{base64.StdEncoding.EncodeToString(blobHTTPHeaders.BlobContentMD5)}
	}
	if blobHTTPHeaders != nil && blobHTTPHeaders.BlobCacheControl != nil {
		req.Raw().Header["x-ms-blob-cache-control"] = []string{*blobHTTPHeaders.BlobCacheControl}
	}
	if options != nil && options.Metadata != nil {
		for k, v := range options.Metadata {
			if v != nil {
				req.Raw().Header["x-ms-meta-"+k] = []string{*v}
			}
		}
	}
	if leaseAccessConditions != nil && leaseAccessConditions.LeaseID != nil {
		req.Raw().Header["x-ms-lease-id"] = []string{*leaseAccessConditions.LeaseID}
	}
	if blobHTTPHeaders != nil && blobHTTPHeaders.BlobContentDisposition != nil {
		req.Raw().Header["x-ms-blob-content-disposition"] = []string{*blobHTTPHeaders.BlobContentDisposition}
	}
	if cpkInfo != nil && cpkInfo.EncryptionKey != nil {
		req.Raw().Header["x-ms-encryption-key"] = []string{*cpkInfo.EncryptionKey}
	}
	if cpkInfo != nil && cpkInfo.EncryptionKeySHA256 != nil {
		req.Raw().Header["x-ms-encryption-key-sha256"] = []string{*cpkInfo.EncryptionKeySHA256}
	}
	if cpkInfo != nil && cpkInfo.EncryptionAlgorithm != nil {
		req.Raw().Header["x-ms-encryption-algorithm"] = []string{string(*cpkInfo.EncryptionAlgorithm)}
	}
	if cpkScopeInfo != nil && cpkScopeInfo.EncryptionScope != nil {
		req.Raw().Header["x-ms-encryption-scope"] = []string{*cpkScopeInfo.EncryptionScope}
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfModifiedSince != nil {
		req.Raw().Header["If-Modified-Since"] = []string{(*modifiedAccessConditions.IfModifiedSince).In(gmt).Format(time.RFC1123)}
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfUnmodifiedSince != nil {
		req.Raw().Header["If-Unmodified-Since"] = []string{(*modifiedAccessConditions.IfUnmodifiedSince).In(gmt).Format(time.RFC1123)}
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{string(*modifiedAccessConditions.IfMatch)}
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfNoneMatch != nil {
		req.Raw().Header["If-None-Match"] = []string{string(*modifiedAccessConditions.IfNoneMatch)}
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfTags != nil {
		req.Raw().Header["x-ms-if-tags"] = []string{*modifiedAccessConditions.IfTags}
	}
	req.Raw().Header["x-ms-version"] = []string{"2021-12-02"}
	if options != nil && options.RequestID != nil {
		req.Raw().Header["x-ms-client-request-id"] = []string{*options.RequestID}
	}
	if options != nil && options.BlobTagsString != nil {
		req.Raw().Header["x-ms-tags"] = []string{*options.BlobTagsString}
	}
	if options != nil && options.ImmutabilityPolicyExpiry != nil {
		req.Raw().Header["x-ms-immutability-policy-until-date"] = []string{(*options.ImmutabilityPolicyExpiry).In(gmt).Format(time.RFC1123)}
	}
	if options != nil && options.ImmutabilityPolicyMode != nil {
		req.Raw().Header["x-ms-immutability-policy-mode"] = []string{string(*options.ImmutabilityPolicyMode)}
	}
	if options != nil && options.LegalHold != nil {
		req.Raw().Header["x-ms-legal-hold"] = []string{strconv.FormatBool(*options.LegalHold)}
	}
	req.Raw().Header["Accept"] = []string{"application/xml"}
	return req, nil
}

// createHandleResponse handles the Create response.
func (client *AppendBlobClient) createHandleResponse(resp *http.Response) (AppendBlobClientCreateResponse, error) {
	result := AppendBlobClientCreateResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = (*azcore.ETag)(&val)
	}
	if val := resp.Header.Get("Last-Modified"); val != "" {
		lastModified, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return AppendBlobClientCreateResponse{}, err
		}
		result.LastModified = &lastModified
	}
	if val := resp.Header.Get("Content-MD5"); val != "" {
		contentMD5, err := base64.StdEncoding.DecodeString(val)
		if err != nil {
			return AppendBlobClientCreateResponse{}, err
		}
		result.ContentMD5 = contentMD5
	}
	if val := resp.Header.Get("x-ms-client-request-id"); val != "" {
		result.ClientRequestID = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if val := resp.Header.Get("x-ms-version-id"); val != "" {
		result.VersionID = &val
	}
	if val := resp.Header.Get("Date"); val != "" {
		date, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return AppendBlobClientCreateResponse{}, err
		}
		result.Date = &date
	}
	if val := resp.Header.Get("x-ms-request-server-encrypted"); val != "" {
		isServerEncrypted, err := strconv.ParseBool(val)
		if err != nil {
			return AppendBlobClientCreateResponse{}, err
		}
		result.IsServerEncrypted = &isServerEncrypted
	}
	if val := resp.Header.Get("x-ms-encryption-key-sha256"); val != "" {
		result.EncryptionKeySHA256 = &val
	}
	if val := resp.Header.Get("x-ms-encryption-scope"); val != "" {
		result.EncryptionScope = &val
	}
	return result, nil
}

// Seal - The Seal operation seals the Append Blob to make it read-only. Seal is supported only on version 2019-12-12 version
// or later.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-12-02
//   - options - AppendBlobClientSealOptions contains the optional parameters for the AppendBlobClient.Seal method.
//   - LeaseAccessConditions - LeaseAccessConditions contains a group of parameters for the ContainerClient.GetProperties method.
//   - ModifiedAccessConditions - ModifiedAccessConditions contains a group of parameters for the ContainerClient.Delete method.
//   - AppendPositionAccessConditions - AppendPositionAccessConditions contains a group of parameters for the AppendBlobClient.AppendBlock
//     method.
func (client *AppendBlobClient) Seal(ctx context.Context, options *AppendBlobClientSealOptions, leaseAccessConditions *LeaseAccessConditions, modifiedAccessConditions *ModifiedAccessConditions, appendPositionAccessConditions *AppendPositionAccessConditions) (AppendBlobClientSealResponse, error) {
	req, err := client.sealCreateRequest(ctx, options, leaseAccessConditions, modifiedAccessConditions, appendPositionAccessConditions)
	if err != nil {
		return AppendBlobClientSealResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AppendBlobClientSealResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AppendBlobClientSealResponse{}, runtime.NewResponseError(resp)
	}
	return client.sealHandleResponse(resp)
}

// sealCreateRequest creates the Seal request.
func (client *AppendBlobClient) sealCreateRequest(ctx context.Context, options *AppendBlobClientSealOptions, leaseAccessConditions *LeaseAccessConditions, modifiedAccessConditions *ModifiedAccessConditions, appendPositionAccessConditions *AppendPositionAccessConditions) (*policy.Request, error) {
	req, err := runtime.NewRequest(ctx, http.MethodPut, client.endpoint)
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("comp", "seal")
	if options != nil && options.Timeout != nil {
		reqQP.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["x-ms-version"] = []string{"2021-12-02"}
	if options != nil && options.RequestID != nil {
		req.Raw().Header["x-ms-client-request-id"] = []string{*options.RequestID}
	}
	if leaseAccessConditions != nil && leaseAccessConditions.LeaseID != nil {
		req.Raw().Header["x-ms-lease-id"] = []string{*leaseAccessConditions.LeaseID}
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfModifiedSince != nil {
		req.Raw().Header["If-Modified-Since"] = []string{(*modifiedAccessConditions.IfModifiedSince).In(gmt).Format(time.RFC1123)}
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfUnmodifiedSince != nil {
		req.Raw().Header["If-Unmodified-Since"] = []string{(*modifiedAccessConditions.IfUnmodifiedSince).In(gmt).Format(time.RFC1123)}
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{string(*modifiedAccessConditions.IfMatch)}
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfNoneMatch != nil {
		req.Raw().Header["If-None-Match"] = []string{string(*modifiedAccessConditions.IfNoneMatch)}
	}
	if appendPositionAccessConditions != nil && appendPositionAccessConditions.AppendPosition != nil {
		req.Raw().Header["x-ms-blob-condition-appendpos"] = []string{strconv.FormatInt(*appendPositionAccessConditions.AppendPosition, 10)}
	}
	req.Raw().Header["Accept"] = []string{"application/xml"}
	return req, nil
}

// sealHandleResponse handles the Seal response.
func (client *AppendBlobClient) sealHandleResponse(resp *http.Response) (AppendBlobClientSealResponse, error) {
	result := AppendBlobClientSealResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = (*azcore.ETag)(&val)
	}
	if val := resp.Header.Get("Last-Modified"); val != "" {
		lastModified, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return AppendBlobClientSealResponse{}, err
		}
		result.LastModified = &lastModified
	}
	if val := resp.Header.Get("x-ms-client-request-id"); val != "" {
		result.ClientRequestID = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if val := resp.Header.Get("Date"); val != "" {
		date, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return AppendBlobClientSealResponse{}, err
		}
		result.Date = &date
	}
	if val := resp.Header.Get("x-ms-blob-sealed"); val != "" {
		isSealed, err := strconv.ParseBool(val)
		if err != nil {
			return AppendBlobClientSealResponse{}, err
		}
		result.IsSealed = &isSealed
	}
	return result, nil
}
