package encrypt

import (
   "bytes"
   "crypto/aes"
   "io"
   "crypto/rand"
   "crypto/cipher"
   // "encoding/base64"
   "fmt"
)


// padding
func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
   padding := blockSize - len(ciphertext) % blockSize
   padtext := bytes.Repeat([]byte{byte(padding)}, padding)
   return append(ciphertext, padtext...)
}

func PKCS7UnPadding(origData []byte) []byte {
   length := len(origData)
   unpadding := int(origData[length-1])
   return origData[:(length - unpadding)]
}

func AesCBCEncrypt(rawData,key []byte) ([]byte, error) {
   block, err := aes.NewCipher(key)
   if err != nil {
       fmt.Println( err ) 
       return nil, err
   }

   // filling
   blockSize := block.BlockSize()
   rawData = PKCS7Padding(rawData, blockSize)
   // initial vector IV should be unique, be not necessarily to keep safe
   cipherText := make([]byte,blockSize+len(rawData))
   //block size : 16
   iv := cipherText[:blockSize]
   if _, err := io.ReadFull(rand.Reader,iv); err != nil {
       fmt.Println( err ) 
       return nil, err
   }

   // block size should be the same as IV
   mode := cipher.NewCBCEncrypter(block,iv)
   mode.CryptBlocks(cipherText[blockSize:],rawData)

   return cipherText, nil
}



func AesCBCDncrypt(encryptData, key []byte) ([]byte,error) {
   block, err := aes.NewCipher(key)
   if err != nil {
       fmt.Println( err ) 
       return nil, err
   }

   blockSize := block.BlockSize()

   if len(encryptData) < blockSize {
       fmt.Println( "ciphertext too short" ) 
       return nil, err
   }
   iv := encryptData[:blockSize]
   encryptData = encryptData[blockSize:]

   // CBC mode always works in whole blocks.
   if len(encryptData)%blockSize != 0 {
       fmt.Println( "ciphertext is not a multiple of the block size" ) 
       return nil, err
   }

   mode := cipher.NewCBCDecrypter(block, iv)

   // CryptBlocks can work in-place if the two arguments are the same.
   mode.CryptBlocks(encryptData, encryptData)
   // decrypt filling
   encryptData = PKCS7UnPadding(encryptData)
   return encryptData,nil
}



