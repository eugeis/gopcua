// Copyright 2018 gopcua authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package securitypolicy

import "crypto/rsa"

/*
"SecurityPolicy – None" Profile
http://opcfoundation.org/UA/SecurityPolicy#None

Security None CreateSession ActivateSession
When SecurityPolicy=None, the CreateSession and ActivateSession service allow for a NULL/empty signature and do not require Application Certificates or a Nonce.
(OPT) Security None CreateSession ActivateSession 1.0
	The Client can connect to Servers that require a certificate being passed on Session establishment.
	The Client in this case will first try without a certificate and if this fails present a certificate.
SymmetricSignatureAlgorithm_None 		This algorithm does not apply.
SymmetricEncryptionAlgorithm_None 		This algorithm does not apply.
AsymmetricSignatureAlgorithm_None 		This algorithm does not apply.
AsymmetricEncryptionAlgorithm_None 		This algorithm does not apply.
KeyDerivationAlgorithm_None 		This algorithm does not apply.
SecurityPolicy_None_Limits 		DerivedSignatureKeyLength: 0

*/
func newNoneAsymmetric(*rsa.PrivateKey, *rsa.PublicKey) (*EncryptionAlgorithm, error) {
	e := new(EncryptionAlgorithm)

	e.blockSize = blockSizeNone
	e.minPadding = minPaddingNone
	e.encrypt = encryptNone
	e.decrypt = decryptNone
	e.signature = signatureNone
	e.verifySignature = verifySignatureNone

	return e, nil
}

func newNoneSymmetric([]byte, []byte) (*EncryptionAlgorithm, error) {
	e := new(EncryptionAlgorithm)

	e.blockSize = blockSizeNone
	e.minPadding = minPaddingNone
	e.encrypt = encryptNone
	e.decrypt = decryptNone
	e.signature = signatureNone
	e.verifySignature = verifySignatureNone

	return e, nil
}
