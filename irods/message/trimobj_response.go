package message

import (
	"fmt"

	"github.com/cyverse/go-irodsclient/irods/common"
)

// IRODSMessageTrimobjResponse stores data object trim response
type IRODSMessageTrimobjResponse struct {
	// empty structure
	Result int
}

// CheckError returns error if server returned an error
func (msg *IRODSMessageTrimobjResponse) CheckError() error {
	if msg.Result < 0 {
		return common.MakeIRODSError(common.ErrorCode(msg.Result))
	}
	return nil
}

// FromMessage returns struct from IRODSMessage
func (msg *IRODSMessageTrimobjResponse) FromMessage(msgIn *IRODSMessage) error {
	if msgIn.Body == nil {
		return fmt.Errorf("Cannot create a struct from an empty body")
	}

	msg.Result = int(msgIn.Body.IntInfo)
	return nil
}
