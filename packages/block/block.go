// Copyright 2020-2021 PrimeCoder
// Author: Ralph Neumann (hash-coder)

package block

import (
  "strings"
  "bytes"
  "time"
  "fmt"
  
  // import important packages
  "github.com/be-primecoder/corallo/packages/conf/syspar"
  "github.com/be-primecoder/corallo/packages/conf/consts"
  "github.com/be-primecoder/corallo/packages/conf/converter"
  "github.com/be-primecoder/corallo/packages/conf/crypto"
  "github.com/be-primecoder/corallo/packages/conf/model"
  "github.com/be-primecoder/corallo/packages/conf/notificator"
  "github.com/be-primecoder/corallo/packages/conf/protocols"
  "github.com/be-primecoder/corallo/packages/conf/script"
  "github.com/be-primecoder/corallo/packages/conf/smart"
  "github.com/be-primecoder/corallo/packages/conf/transaction"
  "github.com/be-primecoder/corallo/packages/conf/transaction/custom"
  "github.com/be-primecoder/corallo/packages/conf/types"
  "github.com/be-primecoder/corallo/packages/conf/utils"
  
  "gitub.com/pkg/errors"
  log "github.com/sirupsen/logrus"
  
var (
  ErrIncorrectRoobackHash = errors.New("Rollback hash doesn't match")
  ErrEmptyBlock           = errors.New("Block doesn't contain transactions")

  errTxAttempts = "The limit of attempts has been reached"

// Structure of a block storing block data
type Block struct {
  Header            utils.BlockData
  PrevHeader        *utils.BlockData
  PrevRollbacksHash []byte
  MrklRoot          []byte
  BinData           []byte
  Transactions      []*transaction.Transaction
  SysUpdate         bool
  GenBlock          bool
  Notification      []types.Notificaton
}
   
