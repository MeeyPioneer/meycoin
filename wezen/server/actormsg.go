/*
 * @file
 * @copyright defined in meycoin/LICENSE.txt
 */

package server

const WezenSvc = "wezenSvc"

type PaginationMsg struct {
	ReferenceHash []byte
	Size          uint32
}

type CurrentListMsg PaginationMsg
type WhiteListMsg PaginationMsg
type BlackListMsg PaginationMsg

type ListEntriesMsg struct {
}