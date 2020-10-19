package main

import (
  "crypto/md5"
  "encoding/hex"
  "golang.org/x/crypto/bcrypt"
)

func GenerateToken(pass string) string {
  hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
  if err != nil {
    return ""
  }
  hasher := md5.New()
  hasher.Write(hash)
  return hex.EncodeToString(hasher.Sum(nil))
}

func Hash(input string) string {
  h := md5.New()
  h.Write([]byte(input))
  return string(h.Sum(nil))
}
